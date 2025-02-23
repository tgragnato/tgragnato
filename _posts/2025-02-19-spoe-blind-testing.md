---
title: HAProxy Stream Processing Offload Engine
description: Blind testing & Performance Benchmark
layout: default
lang: it
images:
  - loc: /images/2025-02-19-spoe.webp
    caption: HAProxy Stream Processing Offload Engine
  - loc: /images/2025-02-19-spoe-blind-testing.webp
    caption: HAProxy SPOE blind testing results
prefetch:
  - www.haproxy.com
---

Lo Stream Processing Offload Engine di HAProxy è un componente che permette di delegare il processamento dei flussi dati a entità esterne. Consente di eseguire logiche di elaborazione personalizzate (come controlli di sicurezza o manipolazioni specifiche del traffico) fuori dal contesto principale di HAProxy, migliorando così la flessibilità e la modularità del proxy.

[![HAProxy Stream Processing Offload Engine](/images/2025-02-19-spoe.webp)](https://www.haproxy.com/blog/extending-haproxy-with-the-stream-processing-offload-engine)

> [Imagine you’re living in the distant future and you’re stomping around town in the latest 15-foot tall mech robot. You’ve outfitted it in the latest gadgetry. It’s fast, agile and responsive. When the new arm cannon upgrade becomes available next month, you’ll be able to augment your robot by simply heading to the nearest shop and having the cannons snapped into place.](https://www.haproxy.com/blog/extending-haproxy-with-the-stream-processing-offload-engine)

Viene spesso utilizzato per implementare i moduli di sicurezza, come i firewall per le applicazioni web e le protezioni dai bot. Durante la messa in opera le sue performance possono essere monitorate attraverso l'exporter prometheus: questo permette di misurare continuamente la reale latenza media introdotta nello spoe backend dal meccanismo di protezione.

Il protocollo di stream processing offload è molto efficiente e i risultati ottenuti da questa misurazione possono sembrare impossibili, per alcuni aspetti sono contro-intuitivi.

## Blind Testing

Il blind testing è una metodologia di valutazione delle prestazioni in cui i soggetti coinvolti nel test non sono a conoscenza di alcuni dettagli critici dell'esperimento. Nel contesto dei benchmark di performance, questo approccio elimina i pregiudizi consci e inconsci, garantisce risultati più oggettivi e riduce l'influenza delle aspettative sui risultati.

Il test viene condotto senza che i partecipanti sappiano quali sono i funzionamenti interni del sistema che stanno effettivamente testando, assicurando così una valutazione imparziale delle performance.

In questo caso specifico i risultati possono essere utili per confermare o smentire i valori recuperati di latenza recuperati tramite le misurazioni effettuate con l'exporter Prometheus.

haproxy.cfg
```
frontend myproxy
    mode http
    bind 127.0.0.1:12346

    # dichiarazione del filtro e del suo file di configurazione
    filter spoe engine ip-reputation config iprep.conf

    # rifiuta la connessione se la reputazione dell'IP è inferiore a 20
    tcp-request content reject if { var(sess.iprep.ip_score) -m int lt 20 }
    default_backend webservers

# usuale backend per i server web
backend webservers
    mode http
    stats uri /stats
    stats refresh 10s
    server web1 127.0.0.1:12347 check

# backend utilizzato dallo SPOE ip-reputation
backend agents
    mode tcp
    balance roundrobin
    timeout connect 10s
    timeout server  10s
    option spop-check
    server agent 127.0.0.1:12345 check
```

iprep.conf
```
[ip-reputation]
spoe-agent iprep-agent
    messages check-client-ip
    option var-prefix iprep
    timeout hello 2m
    timeout idle  2m
    timeout processing 1m
    use-backend agents
    log global
    
spoe-message check-client-ip
    args ip=src
    event on-client-session
```

main.go
```go
package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/negasus/haproxy-spoe-go/action"
	"github.com/negasus/haproxy-spoe-go/agent"
	"github.com/negasus/haproxy-spoe-go/logger"
	"github.com/negasus/haproxy-spoe-go/request"
)

var (
	latency       int64 // artificially add latency to the spoa handler
	testLatencies = []int64{0, 1, 10, 100, 1000}
	testRequests  = []uint{10, 100, 1000}
)

func main() {
	log.SetFlags(0)

	go func() {
		listener, err := net.Listen("tcp4", "127.0.0.1:12345")
		if err != nil {
			log.Fatalf("error listen: %s\n", err.Error())
		}

		a := agent.New(handler, logger.NewDefaultLog())
		if err := a.Serve(listener); err != nil {
			log.Printf("error agent serve: %+v\n", err)
		}
	}()

	go func() {
		srv := &http.Server{
			Addr: "127.0.0.1:12347",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("ok"))
			}),
		}
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("error starting the webserver: %s\n", err.Error())
		}
	}()

	time.Sleep(time.Minute)

	for _, testLatency := range testLatencies {
		latency = testLatency
		avgLatency := float64(0)

		for _, testRequest := range testRequests {

			for i := uint(0); i < testRequest; i++ {
				start := time.Now()
				resp, err := http.Get("http://127.0.0.1:12346")
				if err != nil {
					log.Fatalf("error making request: %s\n", err.Error())
				} else {
					defer resp.Body.Close()
					elapsed := time.Since(start)
					avgLatency += float64(elapsed.Milliseconds())
				}
			}

			avgLatency /= float64(testRequest)
			log.Printf("Average latency for %d requests with latency %d ms: %f ms\n", testRequest, latency, avgLatency)

			time.Sleep(10 * time.Second)
		}
	}
}

func handler(req *request.Request) {
	mes, err := req.Messages.GetByIndex(0)
	if err != nil {
		log.Fatalf("no message was found: %s\n", err.Error())
		return
	}

	ipValue, ok := mes.KV.Get("ip")
	if !ok {
		log.Fatalf("var 'ip' not found in message")
		return
	}

	_, ok = ipValue.(net.IP)
	if !ok {
		log.Fatalf("var 'ip' has wrong type. expect IP addr")
		return
	}

	time.Sleep(time.Duration(latency) * time.Millisecond)
	req.Actions.SetVar(action.ScopeSession, "ip_score", 50)
}
```

Lo scopo di questo codice è valutare le prestazioni di un agente HAProxy SPOE misurando la latenza media introdotta dall'agente in diverse condizioni di tempi di risposta dell'agente e carico di richieste.

Ogni aspetto della struttura del codice è importante:

- **Benchmark delle Performance dell'Agente SPOE**: L'obiettivo principale è quantificare l'impatto dell'introduzione di SPOE sulla base dei tempi di risposta dell'agente. Questo viene ottenuto simulando diversi scenari di latenza all'interno della funzione `handler` dell'agente e misurando i tempi di risposta end-to-end.
- **Latenza Artificiale**: La variabile `latency` e la funzione `time.Sleep` nell'`handler` vengono utilizzate per simulare i diversi tempi di risposta che un agente SPOE reale potrebbe impiegare per eseguire attività come controlli di sicurezza, trasformazioni dei dati, altra logica personalizzata, trasporto su rete.
- **Carichi di Richieste Variabili**: Il codice testa le prestazioni dell'agente sotto diversi carichi di richieste (10, 100 e 1000 richieste) per valutarne la scalabilità.
- **Media della Latenza**: Mediando la latenza su più richieste, il codice mira a ottenere una misura più stabile e rappresentativa delle prestazioni dell'agente.
- **Registrazione dei Risultati**: Le istruzioni `log.Printf` producono la latenza media per ogni combinazione di delay artificiale introdotto e carico di richieste, permettendo l'analisi e il confronto del risultato.

L'esecuzione del test sull'**interfaccia di loopback** (`127.0.0.1`) è cruciale per isolare l'ambiente di test, per minimizzare l'impatto di fattori di rete esterni sulle misurazioni della latenza, per garantire che i risultati del test siano più coerenti e riproducibili.


## Risultati

- Average latency for 10 requests with latency 0 ms: 0.400000 ms
- Average latency for 100 requests with latency 0 ms: 0.024000 ms
- Average latency for 1000 requests with latency 0 ms: 0.003024 ms
- Average latency for 10 requests with latency 1 ms: 1.400000 ms
- Average latency for 100 requests with latency 1 ms: 1.804000 ms
- Average latency for 1000 requests with latency 1 ms: 1.889804 ms
- Average latency for 10 requests with latency 10 ms: 11.900000 ms
- Average latency for 100 requests with latency 10 ms: 11.869000 ms
- Average latency for 1000 requests with latency 10 ms: 11.936869 ms
- Average latency for 10 requests with latency 100 ms: 102.900000 ms
- Average latency for 100 requests with latency 100 ms: 103.599000 ms
- Average latency for 1000 requests with latency 100 ms: 102.555599 ms
- Average latency for 10 requests with latency 1000 ms: 1004.300000 ms
- Average latency for 100 requests with latency 1000 ms: 1013.343000 ms
- Average latency for 1000 requests with latency 1000 ms: 1003.964343 ms

> set xlabel 'Response Time' / set ylabel 'Average Latency' / set logscale x

> plot 'plot.csv' using 1:2 with linespoints title '10 Requests', 'plot.csv' using 1:3 with linespoints title '100 Requests', 'plot.csv' using 1:4 with linespoints title '1000 Requests'

![HAProxy SPOE blind testing results](/images/2025-02-19-spoe-blind-testing.webp){:loading="lazy"}

```
0, 0.400000, 0.024000, 0.003024
1, 1.400000, 1.804000, 1.889804
10, 11.900000, 11.936869, 11.936869
100, 102.900000, 103.599000, 102.555599
1000, 1004.300000, 1013.343000, 1003.964343
```
