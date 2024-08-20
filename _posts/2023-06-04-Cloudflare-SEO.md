---
title: Fondamenti SEO con Cloudflare
description: Una guida completa su come Cloudflare può migliorare la SEO del tuo sito web
layout: default
lang: it
images:
  - loc: /images/2023-06-04-Introduction.webp
  - loc: /images/2023-06-04-Estimated-Traffic.webp
    caption: Estimated Traffic
prefetch:
  - web.dev
---

L'obiettivo della SEO è posizionare le pagine il più in alto possibile nella SERP, il che alla fine porta più traffico al sito web.

Tra i vari fattori che influenzano il posizionamento nei motori di ricerca ci sono: qualità e dinamismo dei contenuti, struttura della pagina, numero di link in inbound/outbound, UX (attraverso i core web vitals), ...

![2023-06-04-Introduction](/images/2023-06-04-Introduction.webp){:loading="lazy"}
![2023-06-04-Estimated-Traffic](/images/2023-06-04-Estimated-Traffic.webp){:loading="lazy"}

## Verifica del corretto setup dei redirect

Ricorda sempre di reindirizzare il traffico da `apex` a `www` (o viceversa).

Utilizza le regole di reindirizzamento per **precaricare nella CDN le rotte per cui vuoi ritornare 301 o 302**, è possibile faro in tre modi: reindirizzamenti statici, reindirizzamenti dinamici, reindirizzamenti bulk.

## Controllo degli errori di scansione e degli errori del server

La search console riporta gli errori di scansione con una certa latenza. Durante una migrazione è facile sbagliare nello scrivere qualche regola WAF.
Per evitare di bloccare i crawler è importante **imparare ad utilizzare i filtri di selezione nell'analisi del traffico**.

La lista dei "bot noti" è una lista di signature, mantenuta da Cloudflare, che identifica crawler e spider "utili", che non abusano delle risorse dei siti.
Escludere i "known bots" dalle challenge di verifica è molto utile per evitare errori di scansione dovuti al WAF.

L'ASN di Google è `15169`: da qui vedremo provenire il crawler del search engine, il fetcher di ggpht, ed il privacy preserving prefetcher.
Per evitare di bloccare in qualsiasi modo uno degli innumerevoli robot di Google, è potenzialmente utile evitare di utilizzare challenge per l'intero sistema autonomo, lasciando che sia la protezione antibot ad occuparsi di eventuali casi di abuso.

Semrush e Criteo non sono nella lista dei "bot noti", se vogliamo rendere possibili le loro scansioni sarà necessario creare delle apposite regole di bypass.
Semrush ha due ASN: `30161` e `209366`, lo user agent contiente proprio la stringa "Semrush". 
Anche Criteo ha due ASN: `55569` e `44788`, lo UA nuovamente contiene la stringa "Criteo".

## Caching, caching e ancora caching

Se la tua inifrastruttura fosse geograficamente distribuita, avessi `anycast` a dispozione ed il tuo cluster utilizzasse algoritmi e funzionalità di `edge computing`, allora Cloudflare potrebbe non servirti.

Con molta probabilità non è così, quindi per utilizzare il più possibile le performance dell'**addizionale layer di networking** che stai introducendo è necessario sfruttare al meglio le cache della CDN.

Le analitiche sono utili e importanti, tanto quanto gli interventi correttivi sull'applicativo.

Per ottenere la massima cacheability è necessario infatti:

1. definire correttamente le regole custom nella sezione deputata al caching, distinguendo con attenzione le varie frequenze di aggiornamento dei contenuti;
2. verificare che l'applicazione utilizzi i `metodi di richiesta HTTP` giusti, applicando eventuali correzioni;
3. attivare i `crawler hints` ed i `signed exchanges`.

**Gli scambi firmati** ([https://web.dev/signed-exchanges/](https://web.dev/signed-exchanges/)) sono un meccanismo introdotto dal progetto Web Packaging, che consente ai siti web di fornire contenuti firmati digitalmente ai browser. Ciò permette ai browser di verificare l'autenticità dei contenuti scaricati da cache di terze parti, come l'accesso offline. Sono utili per migliorare la velocità di caricamento delle pagine web e offrire un'esperienza utente più fluida.

## Migliorare i Web Vitals

Le **pipeline di compilazione** di JS e CSS moderne sono già in grado di minificare le risorse statiche di questo tipo.
Per migliorare ulteriormente le performance possiamo scegliere di attivare la compressione dell'HTML, attivare `brotli` e l'ottimizzazione delle immagini.

**Il rocket loader è uno script iniettato nelle pagine che ritarda il caricamento di tutto il JS fino al termine del rendering**.
Questo implica una renderizzazione anticipata dei contenuti della pagina, e una miglioria di `TTFP`, `TTFCP`, `TTFMP` e `document load`.

Gli script che non tengono in considerazione eventuali asincronicità di esecuzione nella loro logica potrebbero rompersi.
Diventa fondamentale verificare attraverso le analitiche che impatto migliorativo apporta la funzionalità, quali script si rompono, ed apportare le dovute correzioni.