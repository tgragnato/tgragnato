---
title: Mitigations for the NoName057(16) DDoSia Project
description: Understanding the attack methodology and the source infrastructure to implement cost-effective mitigations
layout: default
lang: en
prefetch:
  - cs.opensource.google
  - ton.tg
  - www.youtube-nocookie.com
---

<iframe src="https://www.youtube-nocookie.com/embed/ROf4oNqGEUc" width="100%" height="560" title="Inside The World Of Russian Hackers" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" crossorigin="anonymous" frameborder="0" allowfullscreen></iframe>

> If you know the enemy and know yourself, you need not fear the result of a hundred battles. `Sun Tzu`

## Super Nodes

To easily support multiple operating systems and offer cryptocurrency payments to partisan volunteers, the nodes making up the botnet are written in Golang and cross-compiled for multiple architectures.

Upon startup, the client connects to the NoName057(16) infrastructure and registers by providing its credentials. Bot credentials are tied to a user account.

An account is often managed by a **volunteer hacktivist** who will run the binaries on their own infrastructure and will be rewarded in [cryptocurrency](https://ton.tg) based on the number of successful attacks.
However, credentials may be linked to criminal organizations that deploy the binaries through traditional **malware droppers**. This provides a form of plausible deniability to volunteers.
It's also clear the operators are trying to improve the efficacy of the operation with **infrastructure** they have **procured firsthand** and are abusing.

Once registered, the client node periodically downloads from super-nodes a _JSON configuration file_ that specifies the behavior to adopt until the next update.
Many publications about DDoSia refer to botnet super-nodes as C2s. This is inaccurate.
In reality, the botnet distinguishes between C2s and super-nodes: super-nodes are meant to be sacrificial and reachable by common nodes, providing a bridgehead to Western internet infrastructure, while command and control functions are hosted on Russian soil.

Super-nodes store information reported by client nodes during each operational interval in a `MongoDB` datastore. A `Prometheus` exporter exposes data about the health status of the system. `RabbitMQ` is used as a message bus to transmit control events.

Super-nodes are particularly interesting because operators don't invest much effort in protecting them, indicating they are considered expendable. As a result, **it's possible to infer the whole botnet's operational status through the netflow telemetry of a very limited number of hosts**.

## Bots Behaviour

The voluntary organization of "contributions", the centralized architecture of botnet super-nodes, the absence of P2P communications between client nodes, the operators' inability to implement an effective and robust communication encryption mechanisms, necessarily imply the **impossibility of protecting the current clients configuration**.

Visibility can be gained by introducing fake nodes: creating a **_rogue identity_** to register an account and modifying the client node so that it doesn't actually perform attacks, but still reports plausible results to super-nodes. If desired, it's also possible to deceive operators and get compensated for the "achievements".

<details>
	<summary>Click to show the configuration of yesterday evening.</summary>
	<pre>
{
    "targets": [
        {
            "target_id": "642fe45d318a95e391426659",
            "request_id": "6419d35236673e75cff253c2",
            "host": "www.mit.gov.it",
            "ip": "2.42.233.153",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/",
            "body": {
                "type": "str",
                "value": "keys=$_1&form_build_id=form-HeYtZOikPZbAPE3XiXsWt5-ODPq4JFbcnI5CHONvCm0&form_id=search_api_page_block_form_mioindice&op=Cerca"
            },
            "headers": null
        },
        {
            "target_id": "642fe45d318a95e391426659",
            "request_id": "6419d35236673e75cff253c3",
            "host": "www.mit.gov.it",
            "ip": "2.42.233.153",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/ricerca/$_1",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe45d318a95e391426659",
            "request_id": "6419d35336673e75cff253c4",
            "host": "www.mit.gov.it",
            "ip": "2.42.233.153",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/comunicazione/news?page=$_2",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe45d318a95e391426659",
            "request_id": "6419d35336673e75cff253c5",
            "host": "www.mit.gov.it",
            "ip": "2.42.233.153",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/comunicazione/archivio/news?page=$_2",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe45d318a95e391426659",
            "request_id": "6419d35336673e75cff253c6",
            "host": "www.mit.gov.it",
            "ip": "2.42.233.153",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/comunicazione/pubblicazioni?page=$_2",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe45d318a95e391426659",
            "request_id": "6419d35436673e75cff253c8",
            "host": "www.mit.gov.it",
            "ip": "2.42.233.153",
            "type": "tcp",
            "method": "syn",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe45d318a95e391426659",
            "request_id": "6419d35436673e75cff253c9",
            "host": "www.mit.gov.it",
            "ip": "2.42.233.153",
            "type": "tcp",
            "method": "syn",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe45d318a95e391426659",
            "request_id": "6419d35436673e75cff253ca",
            "host": "www.mit.gov.it",
            "ip": "2.42.233.153",
            "type": "nginx_loris",
            "method": "",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe45d318a95e391426659",
            "request_id": "646be47d5068723b6915d7eb",
            "host": "www.mit.gov.it",
            "ip": "2.42.233.153",
            "type": "tcp",
            "method": "PING",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe45d318a95e391426659",
            "request_id": "646be48c5068723b6915d7ed",
            "host": "www.mit.gov.it",
            "ip": "2.42.233.153",
            "type": "tcp",
            "method": "PING",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe45d318a95e391426659",
            "request_id": "6638d2880560c5f1ff54c54c",
            "host": "www.mit.gov.it",
            "ip": "2.42.233.153",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/comunicazione/archivio/news?page=1$_2",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe45d318a95e391426659",
            "request_id": "6638d2880560c5f1ff54c54d",
            "host": "www.mit.gov.it",
            "ip": "2.42.233.153",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/comunicazione/archivio/news?page=1$_2",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe45d318a95e391426659",
            "request_id": "6638d2880560c5f1ff54c54e",
            "host": "www.mit.gov.it",
            "ip": "2.42.233.153",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/comunicazione/archivio/news?page=1$_2",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe45d318a95e391426659",
            "request_id": "6638d28a0560c5f1ff54c550",
            "host": "www.mit.gov.it",
            "ip": "2.42.233.153",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/comunicazione/archivio/news?page=1$_2",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe45d318a95e391426659",
            "request_id": "6638d28a0560c5f1ff54c54f",
            "host": "www.mit.gov.it",
            "ip": "2.42.233.153",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/comunicazione/archivio/news?page=1$_2",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe45d318a95e391426659",
            "request_id": "6638d2980560c5f1ff54c557",
            "host": "www.mit.gov.it",
            "ip": "2.42.233.153",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/comunicazione/archivio/interviste?page=$_3",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe45d318a95e391426659",
            "request_id": "6638d2ab0560c5f1ff54c55f",
            "host": "www.mit.gov.it",
            "ip": "2.42.233.153",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/node/$_5",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe45d318a95e391426659",
            "request_id": "6638d2b50560c5f1ff54c560",
            "host": "www.mit.gov.it",
            "ip": "2.42.233.153",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/documentazione?keys=$_1&field_temi_argomento_target_id=All",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe45d318a95e391426659",
            "request_id": "6638d3230560c5f1ff54c5de",
            "host": "www.mit.gov.it",
            "ip": "2.42.233.153",
            "type": "tcp",
            "method": "syn_ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe45d318a95e391426659",
            "request_id": "6638d32c0560c5f1ff54c5e3",
            "host": "www.mit.gov.it",
            "ip": "2.42.233.153",
            "type": "tcp",
            "method": "syn_ack",
            "port": 80,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe4cc318a95e3914267ac",
            "request_id": "6419cf7e36673e75cff252eb",
            "host": "www.autorita-trasporti.it",
            "ip": "84.240.191.112",
            "type": "http2",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/ricerca-avanzata/?level=1&q=$_1&date_from=&date_to=",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe4cc318a95e3914267ac",
            "request_id": "6419cf7f36673e75cff252ec",
            "host": "www.autorita-trasporti.it",
            "ip": "84.240.191.112",
            "type": "http2",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/indice-delibere/page/$_5/",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe4cc318a95e3914267ac",
            "request_id": "6419cf7f36673e75cff252ed",
            "host": "www.autorita-trasporti.it",
            "ip": "84.240.191.112",
            "type": "http2",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/indice-comunicati-stampa/?lang=it&q=$_0",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe4cc318a95e3914267ac",
            "request_id": "6419cf7f36673e75cff252ee",
            "host": "www.autorita-trasporti.it",
            "ip": "84.240.191.112",
            "type": "http2",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/protocolli-dintesa/?lang=it&q=$_1",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe4cc318a95e3914267ac",
            "request_id": "6419cf8036673e75cff252ef",
            "host": "www.autorita-trasporti.it",
            "ip": "84.240.191.112",
            "type": "http2",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/indice-comunicati-stampa/?lang=it&q=$_1",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe4cc318a95e3914267ac",
            "request_id": "6419cf8036673e75cff252f0",
            "host": "www.autorita-trasporti.it",
            "ip": "84.240.191.112",
            "type": "http2",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/indice-comunicati-stampa/page/$_3/",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe4cc318a95e3914267ac",
            "request_id": "6419cf8036673e75cff252f1",
            "host": "www.autorita-trasporti.it",
            "ip": "84.240.191.112",
            "type": "http2",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/archivio-news/page/$_5/",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe4cc318a95e3914267ac",
            "request_id": "6419cf8136673e75cff252f2",
            "host": "www.autorita-trasporti.it",
            "ip": "84.240.191.112",
            "type": "http2",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/xmlrpc.php",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe4cc318a95e3914267ac",
            "request_id": "6419cf8236673e75cff252f3",
            "host": "www.autorita-trasporti.it",
            "ip": "84.240.191.112",
            "type": "tcp",
            "method": "syn",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe4cc318a95e3914267ac",
            "request_id": "6419cf8236673e75cff252f4",
            "host": "www.autorita-trasporti.it",
            "ip": "84.240.191.112",
            "type": "tcp",
            "method": "syn",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe4cc318a95e3914267ac",
            "request_id": "6419cf8336673e75cff252f5",
            "host": "www.autorita-trasporti.it",
            "ip": "84.240.191.112",
            "type": "nginx_loris",
            "method": "",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe4cc318a95e3914267ac",
            "request_id": "646aa0314c440e1fe2ce5a2f",
            "host": "www.autorita-trasporti.it",
            "ip": "84.240.191.112",
            "type": "tcp",
            "method": "PING",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe4cc318a95e3914267ac",
            "request_id": "646aa0474c440e1fe2ce5a30",
            "host": "www.autorita-trasporti.it",
            "ip": "84.240.191.112",
            "type": "tcp",
            "method": "PING",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe4cc318a95e3914267ac",
            "request_id": "6641b8f8f361d0321ca9d723",
            "host": "www.autorita-trasporti.it",
            "ip": "84.240.191.112",
            "type": "tcp",
            "method": "ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "642fe4cc318a95e3914267ac",
            "request_id": "6641b902f361d0321ca9d724",
            "host": "www.autorita-trasporti.it",
            "ip": "84.240.191.112",
            "type": "tcp",
            "method": "ack",
            "port": 80,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e76c",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/ita/modulistica/permessi-accesso-porto",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e76d",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/ita/il-porto/descrizione",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e76e",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/xmlrpc.php",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e76f",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/ita/il-porto/storia",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e770",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/ita/il-porto/sviluppo-possibilita-investimento",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e771",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/ita/il-porto/brochure-promozionale",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e772",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/ita/operatori/imprese-ex-art-17-lex-84-94",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e773",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/ita/operatori/veterinario",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e774",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/ita/operatori/rimorchiatori-tripmare-tripnavi",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e775",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/ita/statistiche/stat-anno-20$_5",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e776",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/category/ita/informazione/comunicati-stampa",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e777",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/ita/autorita-di-sistema-portuale-del-mare-adriatico-orientale/mission",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e778",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/ita/autorita-di-sistema-portuale-del-mare-adriatico-orientale/progetti-europei",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e779",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/category/ita/comunicazione-istituzionale/avvisi",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e77a",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/category/avvisi/year/20$_5",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e77b",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/category/decreti/year/20$_5",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e77c",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "tcp",
            "method": "syn",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e77d",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "tcp",
            "method": "PING",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e77e",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "nginx_loris",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e77f",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "tcp",
            "method": "syn",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e780",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "tcp",
            "method": "PING",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e781",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "tcp",
            "method": "syn",
            "port": 8010,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e782",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "tcp",
            "method": "PING",
            "port": 8010,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e783",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "tcp",
            "method": "syn",
            "port": 8008,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481ca4eb9cf72aa31d0e76b",
            "request_id": "6481ca4fb9cf72aa31d0e784",
            "host": "www.porto.trieste.it",
            "ip": "151.0.149.65",
            "type": "tcp",
            "method": "PING",
            "port": 8008,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481d24bb9cf72aa31d0ea72",
            "request_id": "6481d24cb9cf72aa31d0ea73",
            "host": "port.taranto.it",
            "ip": "94.188.237.164",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/index.php/en/",
            "body": {
                "type": "str",
                "value": "searchword=$_1&task=search&option=com_search&Itemid=198"
            },
            "headers": null
        },
        {
            "target_id": "6481d24bb9cf72aa31d0ea72",
            "request_id": "6481d24cb9cf72aa31d0ea74",
            "host": "port.taranto.it",
            "ip": "94.188.237.164",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/index.php/en/component/search/?searchword=$_1&searchphrase=all&Itemid=198",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481d24bb9cf72aa31d0ea72",
            "request_id": "6481d24cb9cf72aa31d0ea75",
            "host": "port.taranto.it",
            "ip": "94.188.237.164",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/index.php/en/user-services/336-supplier-register/2222-avviso-pubblico-di-chiusura-presentazione-delle-domande-di-iscrizione-nell-elenco-degli-operatori-economici-dell-ente",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481d24bb9cf72aa31d0ea72",
            "request_id": "6481d24cb9cf72aa31d0ea76",
            "host": "port.taranto.it",
            "ip": "94.188.237.164",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/albopretorio/index.php?lang=it",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481d24bb9cf72aa31d0ea72",
            "request_id": "6481d24cb9cf72aa31d0ea77",
            "host": "port.taranto.it",
            "ip": "94.188.237.164",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/albopretorio/index.php?option=com_content&view=article&id=7&Itemid=126&lang=it",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481d24bb9cf72aa31d0ea72",
            "request_id": "6481d24cb9cf72aa31d0ea78",
            "host": "port.taranto.it",
            "ip": "94.188.237.164",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/albopretorio/index.php?option=com_content&view=article&id=7&Itemid=126&lang=it",
            "body": {
                "type": "str",
                "value": "username=$_1&password=$_1&option=com_users&task=user.login&return=aHR0cHM6Ly9wb3J0LnRhcmFudG8uaXQvYWxib3ByZXRvcmlvL2luZGV4LnBocD9vcHRpb249Y29tX2NvbnRlbnQmdmlldz1hcnRpY2xlJmlkPTcmSXRlbWlkPTEyNiZsYW5nPWl0&60623da9976804b209800a20ee9461a9=1"
            },
            "headers": null
        },
        {
            "target_id": "6481d24bb9cf72aa31d0ea72",
            "request_id": "6481d24cb9cf72aa31d0ea79",
            "host": "port.taranto.it",
            "ip": "94.188.237.164",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/albopretorio/index.php?option=com_users&task=remind.remind&lang=it&Itemid=101",
            "body": {
                "type": "str",
                "value": "jform%5Bemail%5D=$_1%40gmail.com&60623da9976804b209800a20ee9461a9=1"
            },
            "headers": null
        },
        {
            "target_id": "6481d24bb9cf72aa31d0ea72",
            "request_id": "6481d24cb9cf72aa31d0ea7a",
            "host": "port.taranto.it",
            "ip": "94.188.237.164",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/albopretorio/index.php?option=com_users&task=reset.request&lang=it&Itemid=101",
            "body": {
                "type": "str",
                "value": "jform%5Bemail%5D=$_1%40gmail.com&60623da9976804b209800a20ee9461a9=1"
            },
            "headers": null
        },
        {
            "target_id": "6481d24bb9cf72aa31d0ea72",
            "request_id": "6481d24cb9cf72aa31d0ea7b",
            "host": "port.taranto.it",
            "ip": "94.188.237.164",
            "type": "tcp",
            "method": "syn",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481d24bb9cf72aa31d0ea72",
            "request_id": "6481d24cb9cf72aa31d0ea7c",
            "host": "port.taranto.it",
            "ip": "94.188.237.164",
            "type": "tcp",
            "method": "PING",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481d24bb9cf72aa31d0ea72",
            "request_id": "6481d24cb9cf72aa31d0ea7d",
            "host": "port.taranto.it",
            "ip": "94.188.237.164",
            "type": "nginx_loris",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481d24bb9cf72aa31d0ea72",
            "request_id": "6481d24cb9cf72aa31d0ea7e",
            "host": "port.taranto.it",
            "ip": "94.188.237.164",
            "type": "tcp",
            "method": "syn",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481d24bb9cf72aa31d0ea72",
            "request_id": "6481d24cb9cf72aa31d0ea7f",
            "host": "port.taranto.it",
            "ip": "94.188.237.164",
            "type": "tcp",
            "method": "PING",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481d24bb9cf72aa31d0ea72",
            "request_id": "6481d24cb9cf72aa31d0ea80",
            "host": "port.taranto.it",
            "ip": "94.188.237.164",
            "type": "nginx_loris",
            "method": "GET",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481d24bb9cf72aa31d0ea72",
            "request_id": "6481d24cb9cf72aa31d0ea81",
            "host": "port.taranto.it",
            "ip": "94.188.237.164",
            "type": "tcp",
            "method": "syn",
            "port": 5060,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481d24bb9cf72aa31d0ea72",
            "request_id": "6481d24cb9cf72aa31d0ea82",
            "host": "port.taranto.it",
            "ip": "94.188.237.164",
            "type": "tcp",
            "method": "syn",
            "port": 2000,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481d24bb9cf72aa31d0ea72",
            "request_id": "6481d24cb9cf72aa31d0ea83",
            "host": "port.taranto.it",
            "ip": "94.188.237.164",
            "type": "tcp",
            "method": "syn",
            "port": 8008,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481e025b9cf72aa31d0ee53",
            "request_id": "6481e025b9cf72aa31d0ee55",
            "host": "www.sinfomar.it",
            "ip": "89.46.225.112",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/?s=$_1",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481e025b9cf72aa31d0ee53",
            "request_id": "6481e025b9cf72aa31d0ee56",
            "host": "www.sinfomar.it",
            "ip": "89.46.225.112",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/20$_5/0$_3/",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481e025b9cf72aa31d0ee53",
            "request_id": "6481e025b9cf72aa31d0ee58",
            "host": "www.sinfomar.it",
            "ip": "89.46.225.112",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/news/page/$_2/",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481e025b9cf72aa31d0ee53",
            "request_id": "6481e025b9cf72aa31d0ee5a",
            "host": "www.sinfomar.it",
            "ip": "89.46.225.112",
            "type": "tcp",
            "method": "syn",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481e025b9cf72aa31d0ee53",
            "request_id": "6481e025b9cf72aa31d0ee5b",
            "host": "www.sinfomar.it",
            "ip": "89.46.225.112",
            "type": "tcp",
            "method": "PING",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481e025b9cf72aa31d0ee53",
            "request_id": "6481e025b9cf72aa31d0ee5c",
            "host": "www.sinfomar.it",
            "ip": "89.46.225.112",
            "type": "nginx_loris",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481e025b9cf72aa31d0ee53",
            "request_id": "6481e025b9cf72aa31d0ee5d",
            "host": "www.sinfomar.it",
            "ip": "89.46.225.112",
            "type": "tcp",
            "method": "syn",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481e025b9cf72aa31d0ee53",
            "request_id": "6481e025b9cf72aa31d0ee5e",
            "host": "www.sinfomar.it",
            "ip": "89.46.225.112",
            "type": "tcp",
            "method": "PING",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481e025b9cf72aa31d0ee53",
            "request_id": "6481e025b9cf72aa31d0ee5f",
            "host": "www.sinfomar.it",
            "ip": "85.18.56.93",
            "type": "tcp",
            "method": "syn",
            "port": 21,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481e025b9cf72aa31d0ee53",
            "request_id": "6481e025b9cf72aa31d0ee60",
            "host": "www.sinfomar.it",
            "ip": "85.18.56.93",
            "type": "tcp",
            "method": "syn",
            "port": 8008,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6481e025b9cf72aa31d0ee53",
            "request_id": "6481e025b9cf72aa31d0ee61",
            "host": "www.sinfomar.it",
            "ip": "85.18.56.93",
            "type": "tcp",
            "method": "syn",
            "port": 888,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3cf4c4192c7e2e6fc358b",
            "request_id": "64c3cf4c4192c7e2e6fc358c",
            "host": "www.sienamobilita.it",
            "ip": "81.88.52.184",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/?s=$_1",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3cf4c4192c7e2e6fc358b",
            "request_id": "64c3cf4c4192c7e2e6fc3595",
            "host": "www.sienamobilita.it",
            "ip": "81.88.52.184",
            "type": "tcp",
            "method": "syn",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3cf4c4192c7e2e6fc358b",
            "request_id": "64c3cf4c4192c7e2e6fc3596",
            "host": "www.sienamobilita.it",
            "ip": "81.88.52.184",
            "type": "tcp",
            "method": "PING",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3cf4c4192c7e2e6fc358b",
            "request_id": "64c3cf4c4192c7e2e6fc3597",
            "host": "www.sienamobilita.it",
            "ip": "81.88.52.184",
            "type": "nginx_loris",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3cf4c4192c7e2e6fc358b",
            "request_id": "64c3cf4c4192c7e2e6fc3598",
            "host": "www.sienamobilita.it",
            "ip": "81.88.52.184",
            "type": "tcp",
            "method": "syn",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3cf4c4192c7e2e6fc358b",
            "request_id": "64c3cf4c4192c7e2e6fc3599",
            "host": "www.sienamobilita.it",
            "ip": "81.88.52.184",
            "type": "tcp",
            "method": "PING",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3cf4c4192c7e2e6fc358b",
            "request_id": "64c3cf4c4192c7e2e6fc359a",
            "host": "www.sienamobilita.it",
            "ip": "81.88.52.184",
            "type": "nginx_loris",
            "method": "GET",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3cf4c4192c7e2e6fc358b",
            "request_id": "64c3cf4c4192c7e2e6fc359b",
            "host": "www.sienamobilita.it",
            "ip": "81.88.52.184",
            "type": "tcp",
            "method": "syn",
            "port": 3306,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3cf4c4192c7e2e6fc358b",
            "request_id": "64c3cf4c4192c7e2e6fc359c",
            "host": "www.sienamobilita.it",
            "ip": "81.88.52.184",
            "type": "tcp",
            "method": "PING",
            "port": 3306,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3cf4c4192c7e2e6fc358b",
            "request_id": "64c3cf4c4192c7e2e6fc359d",
            "host": "www.sienamobilita.it",
            "ip": "81.88.52.184",
            "type": "tcp",
            "method": "syn",
            "port": 21,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3cf4c4192c7e2e6fc358b",
            "request_id": "64c3cf4c4192c7e2e6fc359e",
            "host": "www.sienamobilita.it",
            "ip": "81.88.52.184",
            "type": "tcp",
            "method": "PING",
            "port": 21,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3cf4c4192c7e2e6fc358b",
            "request_id": "64d4d5dad970c324f1163ffb",
            "host": "www.sienamobilita.it",
            "ip": "81.88.52.184",
            "type": "nginx_loris",
            "method": "",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3cf4c4192c7e2e6fc358b",
            "request_id": "676ecfe88823a512f27dd313",
            "host": "www.sienamobilita.it",
            "ip": "81.88.52.184",
            "type": "http2",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/wp-cron.php",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3cf4c4192c7e2e6fc358b",
            "request_id": "676ecff58823a512f27dd316",
            "host": "www.sienamobilita.it",
            "ip": "81.88.52.184",
            "type": "http2",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/wp-admin/load-scripts.php?c=1&load%5B%5D=eutil,common,wp-a11y,sack,quicktag,colorpicker,editor,wp-fullscreen-stu,wp-ajax-response,wp-api-request,wp-pointer,autosave,heartbeat,wp-auth-check,wp-lists,prototype,scriptaculous-root,scriptaculous-builder,scriptaculous-dragdrop,scriptaculous-effects,scriptaculous-slider,scriptaculous-sound,scriptaculous-controls,scriptaculous,cropper,jquery,jquery-core,jquery-migrate,jquery-ui-core,jquery-effects-core,jquery-effects-blind,jquery-effects-bounce,jquery-effects-clip,jquery-effects-drop,jquery-effects-explode,jquery-effects-fade,jquery-effects-fold,jquery-effects-highlight,jquery-effects-puff,jquery-effects-pulsate,jquery-effects-scale,jquery-effects-shake,jquery-effects-size,jquery-effects-slide,jquery-effects-transfer,jquery-ui-accordion,jquery-ui-autocomplete,jquery-ui-button,jquery-ui-datepicker,jquery-ui-dialog,jquery-ui-draggable,jquery-ui-droppable,jquery-ui-menu,jquery-ui-mouse,jquery-ui-position,jquery-ui-progressbar,jquery-ui-resizable,jquery-ui-selectable,jquery-ui-selectmenu,jquery-ui-slider,jquery-ui-sortable,jquery-ui-spinner,jquery-ui-tabs,jquery-ui-tooltip,jquery-ui-widget,jquery-form,jquery-color,schedule,jquery-query,jquery-serialize-object,jquery-hotkeys,jquery-table-hotkeys,jquery-touch-punch,suggest,imagesloaded,masonry,jquery-masonry,thickbox,jcrop,swfobject,moxiejs,plupload,plupload-handlers,wp-plupload,swfupload,swfupload-all,swfupload-handlers,comment-repl,json2,underscore,backbone,wp-util,wp-sanitize,wp-backbone,revisions,imgareaselect,mediaelement,mediaelement-core,mediaelement-migrat,mediaelement-vimeo,wp-mediaelement,wp-codemirror,csslint,jshint,esprima,jsonlint,htmlhint,htmlhint-kses,code-editor,wp-theme-plugin-editor,wp-playlist,zxcvbn-async,password-strength-meter,user-profile,language-chooser,user-suggest,admin-ba,wplink,wpdialogs,word-coun,media-upload,hoverIntent,customize-base,customize-loader,customize-preview,customize-models,customize-views,customize-controls,customize-selective-refresh,customize-widgets,customize-preview-widgets,customize-nav-menus,customize-preview-nav-menus,wp-custom-header,accordion,shortcode,media-models,wp-embe,media-views,media-editor,media-audiovideo,mce-view,wp-api,admin-tags,admin-comments,xfn,postbox,tags-box,tags-suggest,post,editor-expand,link,comment,admin-gallery,admin-widgets,media-widgets,media-audio-widget,media-image-widget,media-gallery-widget,media-video-widget,text-widgets,custom-html-widgets,theme,inline-edit-post,inline-edit-tax,plugin-install,updates,farbtastic,iris,wp-color-picker,dashboard,list-revision,media-grid,media,image-edit,set-post-thumbnail,nav-menu,custom-header,custom-background,media-gallery,svg-painter&ver=4.9",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3cf4c4192c7e2e6fc358b",
            "request_id": "676ecffe8823a512f27dd318",
            "host": "www.sienamobilita.it",
            "ip": "81.88.52.184",
            "type": "http2",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/wp-admin/load-styles.php?c=1&load%5B%5D=common,forms,admin-menu,dashboard,list-tables,edit,revisions,media,themes,about,nav-menus,widgets,site-icon,l10n,install,updates,farbtastic,wp-jquery-ui-dialog,wp-pointer,customize-controls,customize-widgets,customize-preview,customize-nav-menus,customize-preview-nav-menus,wp-color-picker,colors,ie,buttons,wp-auth-check,media-views,wp-mediaelement,imgareaselect,thickbox,editor-buttons,wp-editor-buttons,wp-editor,wp-editor-autop,wp-editor-rtl,wp-codemirror,wp-theme-plugin-editor,mediaelement,jquery-ui-core,jquery-ui-autocomplete,jquery-ui-button,jquery-ui-dialog,jquery-ui-draggable,jquery-ui-droppable,jquery-ui-menu,jquery-ui-position,jquery-ui-progressbar,jquery-ui-resizable,jquery-ui-selectable,jquery-ui-selectmenu,jquery-ui-slider,jquery-ui-sortable,jquery-ui-spinner,jquery-ui-tabs,jquery-ui-tooltip,jquery-ui-widget,jquery-ui-accordion,jquery-ui-datepicker,wp-plupload,wp-plupload-handlers,plupload,plupload-handlers,swfupload,swfupload-handlers,swfupload-all,wp-jquery-ui-dialog,wp-pointer,customize-controls,customize-widgets,customize-preview,customize-nav-menus,customize-preview-nav-menus,wp-color-picker,colors,ie,buttons,wp-auth-check,media-views,wp-mediaelement,imgareaselect,thickbox,editor-buttons,wp-editor-buttons,wp-editor,wp-editor-autop,wp-editor-rtl,wp-codemirror,wp-theme-plugin-editor,mediaelement&ver=4.9",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3cf4c4192c7e2e6fc358b",
            "request_id": "676ed02f8823a512f27dd324",
            "host": "www.sienamobilita.it",
            "ip": "81.88.52.184",
            "type": "http2",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/xmlrpc.php",
            "body": {
                "type": "str",
                "value": "<methodResponse>\n  <params>\n    <param>\n      <value>\n      <array><data>\n  <value><string>system.multicall</string></value>\n  <value><string>system.listMethods</string></value>\n  <value><string>system.getCapabilities</string></value>\n  <value><string>translationproxy.updated_job_status</string></value>\n  <value><string>translationproxy.test_xmlrpc</string></value>\n  <value><string>translationproxy.get_languages_list</string></value>\n  <value><string>wpml.get_languages</string></value>\n  <value><string>wpml.get_post_trid</string></value>\n  <value><string>demo.addTwoNumbers</string></value>\n  <value><string>demo.sayHello</string></value>\n  <value><string>pingback.extensions.getPingbacks</string></value>\n  <value><string>pingback.ping</string></value>\n  <value><string>mt.publishPost</string></value>\n  <value><string>mt.getTrackbackPings</string></value>\n  <value><string>mt.supportedTextFilters</string></value>\n  <value><string>mt.supportedMethods</string></value>\n  <value><string>mt.setPostCategories</string></value>\n  <value><string>mt.getPostCategories</string></value>\n  <value><string>mt.getRecentPostTitles</string></value>\n  <value><string>mt.getCategoryList</string></value>\n  <value><string>metaWeblog.getUsersBlogs</string></value>\n  <value><string>metaWeblog.deletePost</string></value>\n  <value><string>metaWeblog.newMediaObject</string></value>\n  <value><string>metaWeblog.getCategories</string></value>\n  <value><string>metaWeblog.getRecentPosts</string></value>\n  <value><string>metaWeblog.getPost</string></value>\n  <value><string>metaWeblog.editPost</string></value>\n  <value><string>metaWeblog.newPost</string></value>\n  <value><string>blogger.deletePost</string></value>\n  <value><string>blogger.editPost</string></value>\n  <value><string>blogger.newPost</string></value>\n  <value><string>blogger.getRecentPosts</string></value>\n  <value><string>blogger.getPost</string></value>\n  <value><string>blogger.getUserInfo</string></value>\n  <value><string>blogger.getUsersBlogs</string></value>\n  <value><string>wp.restoreRevision</string></value>\n  <value><string>wp.getRevisions</string></value>\n  <value><string>wp.getPostTypes</string></value>\n  <value><string>wp.getPostType</string></value>\n  <value><string>wp.getPostFormats</string></value>\n  <value><string>wp.getMediaLibrary</string></value>\n  <value><string>wp.getMediaItem</string></value>\n  <value><string>wp.getCommentStatusList</string></value>\n  <value><string>wp.newComment</string></value>\n  <value><string>wp.editComment</string></value>\n  <value><string>wp.deleteComment</string></value>\n  <value><string>wp.getComments</string></value>\n  <value><string>wp.getComment</string></value>\n  <value><string>wp.setOptions</string></value>\n  <value><string>wp.getOptions</string></value>\n  <value><string>wp.getPageTemplates</string></value>\n  <value><string>wp.getPageStatusList</string></value>\n  <value><string>wp.getPostStatusList</string></value>\n  <value><string>wp.getCommentCount</string></value>\n  <value><string>wp.deleteFile</string></value>\n  <value><string>wp.uploadFile</string></value>\n  <value><string>wp.suggestCategories</string></value>\n  <value><string>wp.deleteCategory</string></value>\n  <value><string>wp.newCategory</string></value>\n  <value><string>wp.getTags</string></value>\n  <value><string>wp.getCategories</string></value>\n  <value><string>wp.getAuthors</string></value>\n  <value><string>wp.getPageList</string></value>\n  <value><string>wp.editPage</string></value>\n  <value><string>wp.deletePage</string></value>\n  <value><string>wp.newPage</string></value>\n  <value><string>wp.getPages</string></value>\n  <value><string>wp.getPage</string></value>\n  <value><string>wp.editProfile</string></value>\n  <value><string>wp.getProfile</string></value>\n  <value><string>wp.getUsers</string></value>\n  <value><string>wp.getUser</string></value>\n  <value><string>wp.getTaxonomies</string></value>\n  <value><string>wp.getTaxonomy</string></value>\n  <value><string>wp.getTerms</string></value>\n  <value><string>wp.getTerm</string></value>\n  <value><string>wp.deleteTerm</string></value>\n  <value><string>wp.editTerm</string></value>\n  <value><string>wp.newTerm</string></value>\n  <value><string>wp.getPosts</string></value>\n  <value><string>wp.getPost</string></value>\n  <value><string>wp.deletePost</string></value>\n  <value><string>wp.editPost</string></value>\n  <value><string>wp.newPost</string></value>\n  <value><string>wp.getUsersBlogs</string></value>\n</data></array>\n      </value>\n    </param>\n  </params>\n</methodResponse>"
            },
            "headers": null
        },
        {
            "target_id": "64c3d00f4192c7e2e6fc360f",
            "request_id": "64c3d00f4192c7e2e6fc3610",
            "host": "www.gtt.to.it",
            "ip": "158.102.161.11",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/cms/cerca?searchword=$_1&searchphrase=all",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3d00f4192c7e2e6fc360f",
            "request_id": "64c3d00f4192c7e2e6fc3611",
            "host": "www.gtt.to.it",
            "ip": "158.102.161.11",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/cms/gtt/presentazione-del-gruppo",
            "body": {
                "type": "str",
                "value": "searchword=$_1&task=search&option=com_search&Itemid=236"
            },
            "headers": null
        },
        {
            "target_id": "64c3d00f4192c7e2e6fc360f",
            "request_id": "64c3d00f4192c7e2e6fc3619",
            "host": "www.gtt.to.it",
            "ip": "158.102.161.11",
            "type": "tcp",
            "method": "syn",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3d00f4192c7e2e6fc360f",
            "request_id": "64c3d00f4192c7e2e6fc361a",
            "host": "www.gtt.to.it",
            "ip": "158.102.161.11",
            "type": "tcp",
            "method": "PING",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3d00f4192c7e2e6fc360f",
            "request_id": "64c3d00f4192c7e2e6fc361b",
            "host": "www.gtt.to.it",
            "ip": "158.102.161.11",
            "type": "nginx_loris",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3d00f4192c7e2e6fc360f",
            "request_id": "64c3d00f4192c7e2e6fc361c",
            "host": "www.gtt.to.it",
            "ip": "158.102.161.11",
            "type": "tcp",
            "method": "syn",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3d00f4192c7e2e6fc360f",
            "request_id": "64c3d00f4192c7e2e6fc361d",
            "host": "www.gtt.to.it",
            "ip": "158.102.161.11",
            "type": "tcp",
            "method": "PING",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3d00f4192c7e2e6fc360f",
            "request_id": "64c3d00f4192c7e2e6fc361e",
            "host": "www.gtt.to.it",
            "ip": "158.102.161.11",
            "type": "nginx_loris",
            "method": "GET",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "64c3d00f4192c7e2e6fc360f",
            "request_id": "6638ec3a0560c5f1ff54cf39",
            "host": "www.gtt.to.it",
            "ip": "158.102.161.11",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/assistenza/bridge/public/gttform/knowledge_bases/search",
            "body": {
                "type": "str",
                "value": "{\"knowledge_base_id\":1,\"locale\":\"it-it\",\"flavor\":\"agent\",\"query\":\"$_1\"}"
            },
            "headers": null
        },
        {
            "target_id": "6515862e6bdbcfa9997f1270",
            "request_id": "6515862e6bdbcfa9997f1271",
            "host": "concorsi.gdf.gov.it",
            "ip": "2.45.152.31",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/index.aspx",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6515862e6bdbcfa9997f1270",
            "request_id": "6515862e6bdbcfa9997f1272",
            "host": "concorsi.gdf.gov.it",
            "ip": "2.45.152.31",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/Portale-Concorsi/Ricerca-con-risultati.aspx",
            "body": {
                "type": "str",
                "value": "__VIEWSTATE=jhhf9m1f7LJGvxmi3POc%2BT1JGGecoxqSrxioEoXOXaQtsTXU5zeLC0iHntZCAyQneV%2BxFu8cXm63Ahbk%2FhTwlDmd%2FYThdbwTC8iVzMvg2PHqMYJilnU%2B97hbuYqdyYDqAr3rmN81VoyU6PwB%2BW3IjC%2B4jpdbMAcqFHCB0I2es%2FgIfMlJqKq4Y3rrTpSpVeLrZoA9ISXb%2FFUD00xW9NERPxfB1Q9GfCp38LRVDZcGA1NJm4VpaaEiCBABAIPYjqhihkpDe5%2FMQyu6%2B%2BBt2y6k%2BBMB2xdVMQd0gXOJN8JbpdeEqcwkXD8tktQBeMUcTeAgLQ7PBdrw15fjoz%2BKrE0BRJVVbUx4o7qJGINu7zX1iU9s0pXQuY6WsmCMcC5IzPXuo2%2FXeOt8%2BNEUFsvJxP23vNVBIU4QIRrbZ8opIeDzoo6k%2BtwvFVeygtW4ynuPOayQdMj3Z6D4F9UalKKm5ZGzwHDdG%2Ftq78EB1d7G%2BcUf6FHQ6XF%2BGNl44CKuXdryCezGkmodTcsFHC6MCBEnW5Nh5ek7BPNc%2FbR5PJvcV2wbhr%2BXpj0Pr1mQ2DzhMbMcNLbEagsW8CnGjkBnOOuN22zjg%2Fg5nFasQgvVNA8SRMNmvbbolAFXR40gHJJMcqOz4hOvJiU01F5mO4iiBzA6yYEtRBZPUoVWmfl2f9KqPR0PD8sGssy0lbp9sQlUY%2BqC5jRO4PTMBSPkJoYB5SW9EslVPBCfDYOHS%2BoDy4ppuvCknQq6xPOQamhB2jN3tPJVJbThotBCYtjPCNGvr7IWMJTMlnmBRQGsQL%2F4fkL2eNvC%2FvnbBFxdtw%2FX5QUqqw5jFT52WMGeP1aOAiSxW107wW1GVG%2Fwl1UgtJBuRN71GCxvF6b27AroLFn4Q8rc8cbdRlQzs0sY9VV9hImdR41Nwo%2BGHP%2FDIYs229pGPw1YoaHrU9S3G0vxfP4%2FUI1e12usneqjdH7paqWVLuCnjE2oaxr89F7TlGzfHDQNZUQnfznKp6Bixo1vY2IF6SmmSvyJ%2FOOHhK6VXpgmo%2FoDjEyLflKuxDqZTtBc8TdT3NdiFcGoJd%2FmaSEFTPTPnn1jaYvx3WlW0rT%2By0%2F%2BFOv%2BYoCrztOPpAyEkEoGb3LqOAHLcPWi5Tqj94Jz75%2ByQeT4uTvepi57HMjxYt%2FU%2FuAcW5nvtg%2BxJTGWzC3qvnzHHfb4gmSGknWfwYcNWMkhBiavtNpDgrdngVK2HNrfc76nLycmtNfJJpYYmE1f0EsUPU1ttJIEsM4V7uwppjPW6tAJYII7M4V3VPErWFMVTT4fWkq3nfsP%2FWjfAXT6lKiCYszWfx4QAmFAsl1euIHPQCcug62n8WqT2MUe5ptWuzJJL121ZkAF8nlngofjcwqKH6Ti3ZLwBofuIzXEwDEDRcKkgUvCal2PltzIpPBKxov1D%2Bn9Ejt%2BQ409uR1o0RsUa%2FxQzTCQ%2B%2B2QcxGT4kKTzzDWD186kQde7EBXGKVWJSnXqWFAUdFK43IyNKJoMNDR40NgXjg2FlEmo%2FNGy8qB2pd9o%2FUNT52gltBH%2Bug4uI8pHWt5PoguIJ5acVeJP9LANgQSpw0YtE0EYrg9X8%2Bdkn%2Bs7jcDVLUwyFaecsYofvxIiHNu8tr8stuB5ojoWeIZgeJzrX6cxxunWUgn9PsjDG%2FsUcMwX6DVPc47%2BTnM8EfTAqOB%2BDR%2FshWmHjd0L5EDGr7khWsfln5Mv6p0M%2FjucFL4VCTHRt4p9lJfLG38h4x9Urtq%2BA9CV%2FtjjqOdM7gWertMBwa9Hzgct%2Bi%2B24rlOVyAnMob2G3smxSFIa99yHbdhRcM6snT64ExpnYA03105tJc3wObUXWEuyGyaYCz0MFo93x1PMzsjVex9mvwwvBAH8atnyTEAw2DX3JRW0gRkuuruZOB%2F5GuEQ5croa0XCZtyODmZwoCV%2Bj5V%2Bst%2BvK7OEmtYvQuh1WsxEWuUlQ1R7olSlcSgYXWlZG4H6rOj8NowFMTkwjtcAPtpmQXxEX3fFrJXyAYmBlM9pTskUFcxmjKfgDqSMvcfcfYOr11xM0d7ie8KYF%2FXq%2BrbPR0L8z1i%2BW5AHLWF32qFbVSW5OHqnE23MtYudi0reVQJ4kZw3DzeNS26JXMAIcpcSBwXkhUjHSq%2FGupG4XGqvuIeD3d%2Fs0V3bUn7smrpJoGuMBPV2BcAmzXeVF6CVMIHRwwU8qrFJf%2Bw4vzDdycevf0obtI2948ushoQYyhkWDoQ34P0IitcVkIeTTl9JQnfg3pA9gI%2BJW3BKPVAGabjDs%2B9wUxRdoZOCo6kKdVMp2JLaDflHgqIBq1aNFT5SpnPF2kKa57B7kWTENgEyY6a%2B5icfcYwUNSZ98xte7tL28gTwKglFHwMYjQWj%2FBSsQAvy2OA8Y7jtC8eIbwO6lcnyy3aKlVWs5qupIsZ%2FtV0VdXQu5UxaGsiqtl%2B9sqCpcVBxaFSSvQyLOO49TRVsLP1w7iaHn71ySFm6a6Aidi0efIuBlVh73SGXu1R54wQMoKb9m4UGxGrWtCnCRD4VuPOZqfL6WBAa%2B5zot6Nu5CRZqv4ll2H5fuHwytgfpTvYJMFDFJiYMIYGQpvL2vprA0kO7loi2ROqtYefY3f8LezN52oPrP1u1IPm7jRQwho13kA2gMne%2BCd7WYiIfD%2BvXy2KYE4JAMdRSo2SPTWHqPtFplZAkKyw1W1okjRW5CTfP%2BT782jpzIUdE7Y%2BFZBjq8Pd8cwePuQWy0hBmdN7v6ROfgh9rD6mxsjc1HpslfQix8lsFInh%2BmnEYEFACzAB%2FFxyj4B3VlDcAjmK3SzpBLbzhPm4hkD1xZUwoD%2BQsCB3TBBEfuJsjg6thKoUd%2FcYQApy%2FOF%2F1baFvjhM%2Fw%2BI7wkQBUlSmzsjAjeWFzSI6LRq1m3weeA8zEvqVqRyGnT5fGMN4VxxTAaHlPw%2FWOxwqb9R8nEUjvHwAVgVfjlrAYpEtkDE%2FgmameIZ6MFSgKF%2BAxLLt5%2FekscURaGEGPfPXQcBHj4GbIsmMdM3IlQRS6JepTdv80gi2mNVhPFmA%2BuKIqJ6XuTRjqIuJBeoL3noyu%2BY1nZ%2BtLmyF3UdFQ4srdd0YPeDwJdUeGvaJ2Pe0GzD3whabyegdWgJhw25Uy3IFXFimXBIO35kL7HqQ32QRQPlDvsKf56MyCfC%2B%2B03QNiDrtE3So0EWG9q%2BlR%2FHUHyjkzQtWraZS4gL7ZJ4PJvRRjoV8vF3fa%2BBH3cKJEUzvIHIQe72aUSEWcvQKmJOXUJwu8k1rXJ0AUTgu%2B6mvJjncEvFbwLHNLxjFMzhmhHG571nkgH1802Zlq9Fh813Dz%2B0DzN5lOVbGXETEhgCcljP5qb9udydRYVPl%2BU%2BS9MCGgaV18Fdp3fWz6VaCpHu%2FEYvNwIwTtcsZvFEhlIdJsCYvZWP7pKEK1xXsxBSkmuy29e398ej1UAEtMf8iS%2B1FlSw4q3AaXvKF9oh1%2BvBW9H5PD%2FlgWIvSYIrBIK7mfCywHGuAr2xYnw3DCJpMPbwmwGdVWwSMgTMrNboXKaOHOuyVpP%2FqjCloSAuN8xRm2rg3e8oZ4HSH3EJ5mR5kmykZse8YS%2FbN6KMrPJaqiCF3syXzlMv3PW0GCUfq5Di8Rn10G0Jb9gUAwpAw9Xz3riumJe3c2jlOFlOUhSZrGJtuDfe8IMd8eV7AviBw2yYGgCJfYE2JwDpYozCd074i%2Fn1ph%2BcKITxt2if6HZxWcsw7w6gSphgQhseeSVf0IYl8QiMw4Tjmr2kzeDAINBg9%2ByHAtHGcvGlqQTviL20etZEjTBSgb6lDs36WeymFfGWE7fYqLKy7JHTKhUV6ziH0%2BfiaRdWwgO9G4EHJMyzsCjIcTJUdEa1E2HelqKsCZcdY7oeIOuXsnISyhTBldwZ37K%2F9cBd%2BILpTqqF5jCVTndo7gxRQHpazSKcwia3Y4%2BpkVM13%2Fy1ndERuRU1aTJigz5pDDwiLiZrk7dxegptea7xT6xEjSgUlTL5l23AbdNzwbWwzFFuMFtyru3uLtF3Ej8ILCTzMxEG6lC%2B3F%2BBqOXitpDlbsk9512YeBPY2XucVli6IaB6YM%2B6PsNUtB%2FGSZjTOeaCOs6UKd14hAVWqzB8SXSDPocOfuCNCtgfIDI8aKb6rAMCct6pBtft6ANkC%2BkOGICVGgvaBD2X%2FSyt2yHpWzG5M9mRRlVWAhcuEtyu81meXM1SeMvE74xgx72YETRyGtjH9aMX8rZCMjLaa8Uzjph1K9EGLDDnwdbsoNcvVV5yG4vjZ8hUFcUzJoBdeBVGhu15jwaCwEaGKlCH%2BUJSB92KkNgexVE4vT%2FGLc9bactzklukYcTEKQyFQ5Fm6AiHzwwneqZDTHbS01JNIHwaduFSb8knCAw60Tb%2B2bj9ohqsBM9dWWU0qNwNvBH5F35WJ7rbZPPrK6c75a%2FCnRaWPcfDka1e6%2FUCPuasJrWCJFRbpsLZxy1HmWIdHLCCc7nIYQmzdCsX49aWKv%2FixJ0wIMgK3KuhzM71MH4YmeM%2BBNhCnvDwHFHiMLN2H0P%2FwIMfi0i5hQX7b55CbLwC%2BDWrbE2bWCAr%2FfhplgrJCKoDD9S48aKjtenwdWt3BwX1pWqy%2FN6Sifb32HPnjCO8sRde9UWLCBl63ggxZJq%2F7rpJYaRm7U75xnjieUORTFd8miUucr2Eo3reelz5j3t6fWjJDmLOeWLuKHZO7qv55kG5GKYT%2FPA%2B2JyZD%2B3LO3TxiKF0KN2mcjqMWV1BlaYttqs4umEFzPfEgbBTr1QCTDi%2BBMEZMUPQ8DngFOMYmyRqUODzRNrOiYv4nVDFDTkfs01NLu1T%2F2ieVY1UPU%2Bne2KzV5xGZiQ0CMHHIKEZOvXLguplYh1keN8Cf74aIb372ukMWyg4LiD7MvwFvZ01N1oX%2Bjvw2Q%2FKLrqWVEhhurnXKXbi5egpemgUnBvJC6paoj3JDLQb30w7BK4I2e2lHwCN5p2kPj%2F6bm4GSFxJkULhXotIWOSKeEGBPWhAdB2Fe7FQXJqBIoko1HLoKRedt02U4QC3I4HYum04W5cuT1%2FLRJr%2BwPbkUOeVg9n5fVy0AgMqnVkcAgtdCuE9eE66fZ3Ad%2BbrxEvViwvZE%2FCODAGfLTu0vvaOJop4P8tc%2Ffn3%2Fh26tHQHw%2F7144IczC6mSDXSQj7PmsRRPI54Cg0zPxQdJdl0YRkIOKprqXU%2FALtJz0d%2BYDcodo%2BbQhcmzTBnF3YAllO4WAvDXq8VD6pb8wJqwgYBq90DVMcxz5tzw0huA1tteuhGANQx5zavuTID7FUCNDSGQh1dEdqmSdA9EUOAhePZwL0712U0x4IFD4XeWZJl3RdTGWosQr3yGgZGBfAUwoelF8ylKxgWiD38AEjBSJ3J%2B%2BVf1kpLJW7MSy6bnBM%2FQ%2F1uGC3BhZ0AhDy8OD%2FAG6jpUMD7jai%2F6gT%2FJAgcaHft9%2Fdrr97Berzn5mGNrWrjhFOuf4mECEcViFuSMWaTn9USeXMDeTPcykc1jgdwfuVs6ospjOeP42adNkiPtOG43GcMAVY11tpxjL3QnqynOZmwDHcxjdz%2FGVjsdu6yUeMaxvtzeD4PdIXw0sLHZzfPu6Hft18gP%2FecALPHYRretjsVyi0DbtMbszanSHayU4GwAq%2FE1silY3DX98uUFEFJkj0gE1LLUVKXrAJodZ1oLd4kXppC8aWLG5FQvxxyDyn2mpaxR4qCvMx90bJjKdGbWIOHnMvMETPsrBjoVxoKA8VgJYGFTnMksx7yBg9My4wIzIy6e5m%2FLWFF%2FtHnLtLSYB6LyLrg2mIPaEVJYKpw41Cmgyk2pN4JEdRGeuVI8%2F3a1kzORwmo9rVPyHsjteTe749R6iWjphaLNlHBikP%2Fl5oS41ZFshth4%2BgdEeymHf9vqJPqn3A%2F2cN84vBqXfq98RV1u8a%2F7O%2F4dczt1Afj9Gv85r6%2BINBHHXNvIQKI3WZqHplV%2FX%2FHo1DaK6yf9lQ4rn%2FbVBBs5Zi46CmNpKzY3Iwn6wDHj7p3P9ohkJ3AM6PP0uGGOcmGIKdzQe8%2BXHOSI8KOuBuihFjkW%2FTrdzqXVDyHJwB4tzrp80t7SPIwC9p%2B8ROuMOcKSHC3YSw%2FYAu92WLOoLY9WupIVfkL40oNXFR3fsM1QF4%2Ba%2FE9o8L2mR610sazPe6vMes9MaoHwBkoO4A38Vnh3MJokXGhwROSBRHxhgkYEvVZVAy3378V5EFOb0HdaYZEeLtkxjADCw40p657xY5GYoGiffgcIWCFo%2FhBLXcTXFaJK2ico5m6%2B8bNNHWLOwdhfHfCvX%2FB086KKDFRUVYCKx0%2BnGdNFyIH6jC5UhsaiOea1Ia9DGP6OK8qVyczeVmRfamQN3epJu7ta7zqXsKnoRINpcU4VeXkG%2FY8wkb2YER%2FvwICzukRGDqHQ2rtJppNceww5ePkqlOEKzKEhgtG6HxyOVa8u1REEGtYXZN1%2Brz2Yon%2BeQ%2BX1UKSU0n00yu%2BXRDObDSV8FIh9U5JF3a%2F6vpnzRq7Ci64QT5NDw1QOKucElHxeZ5r%2FaFCnTX2T4cIhYByuUg1VeLD1aV9Gm8TztTMxEeZynmfqaD3jhO9s%2B2GCqP3%2BBF6fftl2oxESIIoCfjLDwkBXdHW3bl0TC69lvMIRxYoh%2FmKhkHOnxiqRYcK6ww1dHjxo4XO7YAlhX1sD%2Ftm6vCqBqH1Qu%2BeP%2Bt2ZIZJzQghnFl3ib%2BQ3vZ%2FfJax6dY6CV7AEdCLkYzst9E%2FaGRQwwLT7q%2BMxBH5Tdtr4R8y63FxOJkug1qLgCMNNIhpoJoAZvA%2BvgrQZwrVY7VlIMuvVazDMSkSZ%2FM8XlZG7o3VmagBYULfIOx9GiTTRV4U0B%2FwrMUI5Ex%2FZlEjKvVW0Tfl8q62YOqhU4M4SXINnF0uS29yvAQHKKJAKR1oWxcWgLsKTU42M4ncNT4SxFbyGbnhpPdx3P8Bi2Xg2RYRWEDmdbxBIYgKWNs5hm3pui2RNKY%2FfzHbTYmlmBclsWHBmL9LdbbNuf8n1jUu4%2Fsnb1HjWvTGRWBbDxyPHFW09XY9Uq7uMFuf%2BqPo25gdyK5WAwbFoMyIQvOhSRLdWZmdiZsawR7%2BVaqM7b7qYhD9aIc4uzGW5iI1iebN%2F%2BSXIxWAGNmaNDZIBMDVJB59yw8HOcUYm2hJCUzM8mlVEPmoieX5A9EMIr6drEIYQgbm8m14pqlfXpXu227vPES5H4NxMN1aCbi%2B5Ul7V21qnlu9tLetnjuFPoN0oGjjcF41DcjPkC1i77eLxPYaveLZmXYK%2BymmTQ1GM9hfqqjzf5KXGaukuoCGKxRVxwMlfXy0eQvJzGJS6GDI1iRl4kr30HyK0asgaZl3bye%2BOPCLVGHuM5vhAgQ87P%2FlIHn3%2F3QYDMSXcfi4vVrPe9qZrM0Xaqwbvp3k5BfyH7UvsipsU%2BVSUU8iixGFFRRtfp0WQ%2BNER1cnpZcKasxEg0uT9Dx0j0Bc8qMuyujQdYXLEVVQ4MI7kT0dlnefv0alMTUpKoi8YjUkwURLiblE2SEgWy5ZFwbkBi0ZABt8unNy8OL7h3oUwadZGnaqZmIpSwnGRQ1AG2JTReBpOOoomx4G%2BL8%2FJbIaSfiLsShMIFdo69Zj%2BlFq%2Bwnw11QWRm1kww1Gt%2BMddjYhdQtsor2G5W1FuhddRjgrSjJidzojsSOgnRiM8%2FVwWROYZ5CnFeNMaOT%2Bau4mC8Voqj%2FPniiNZwnn0S9W5Lbr%2FT20XlOjarhYri%2FgZ2B5Xsm4NpGhQM6laW%2BvsWrngLJQeQZ2R6l3XdUoSwiWJXxSK08trZl37oDpE8xkwMK2dX%2BcbgO0gwoCK6XWI0AF798Hf2fIy85FfJlHGI9PvTJn1cBSLkBhh%2B8PtZQUFmPhI8dSq15%2F8KRpOo48RkXDHOZBRoJM30CgxVg3nII0RYY3g85X2vmPuZJYrSD8B5p0Ks3a1J8f5pK4PGZVouFsxmdJ0on6Rzbobrd2blvV0cGdCWpjNmTZynLoBfWvvCBpmLYx2nLRkyL8wzlDKMEcfTmCCXvlxk7sHH0VedlRmtvQb4IVch8D727SOD9ES6VDzaRFxZtRts4wc97O0wF%2BCe6PEgE%2BOqpsPFtxXap9T6ngPoh6PlaG7axPu3UUFNS2bd8A47BXGo7h20KL76d5%2BAp6NXrkz1%2BqphTj4MIHDVMx7m%2BjcXySdwTMGOARXrDK0OzmePPDZmNi7Y6Q3YrvTXUtFoZz2IdlgNwYmK%2BUKcma3ljp0vdutMTSFbZbIB3v1%2BKlbx49UMPWoNtHW2bb1gfKiSZHlREtI5wFKqrlhfgK0hP4oR2nlUvthqW1dv8VJFZ5t6yxDNl9tjYD5CQjx%2F0RRclhNRGVT73%2FY0cXso8zLI32D13DOGOYHK%2B9PiRxoEYRpZFx9TRqtdEMA3rbw7LrNDl%2F5p9Y0UK9WNDjCmJFSapG%2Bt6%2BQEwumUIwVuWiX5J9rWTjA86AND0qStsBiKc0pOsi1x7T6MeePsjrPEheYmysLb58%2F3ZCBbXyWJEaGTYi8y02DOrbZ%2F4HJPn0fxAzH3Yv0UBefX6z3S0tU%2FqZfFq9qE9rxCO%2BVoOo%2FtXyqVnS9f%2F687ocJ1KRsz7iqgpxwD8LRvTWg2ncb%2BSKL42PQez8CHCstUEpxunOUW1Vh4eD282K9RC8PIfMpXTwkGF3aKyqjj96dOAF8GlOrS6xiX2PycAb6VVB6OLAQDJjFOjUGvYc4ilpSBPPnolLsDOO9ahcz0zLqhM1hM8qIiwTgpzNXolEZf%2FEDw2BdrItHtM%2FYM4cJkP%2BGWd4Yd1g%2F1E2GuLv7VJcTdsFEkazacOKIPCwPcSB7SiBp%2B1O6rZpWdX%2FA%2ByXtml1mSTK1TrOYcjqlVYq3IFPsvGxaG5BMzkLH91Z4NmNWhHRSnQXYJiwPbZp3DtqsL7nx4XgPyGUQE5bsCXYRURUQ17ABBU1he9HT7RB1aTMZMk1AaC84xCxyIu%2F0m9z4K1BJVg9WLrGWSQ1ZnRUleKmIAwdUwwm%2FrNN4oH4Yl%2BKMnMFgRnAVQ58dv6lUV%2Bu9g22WQSvEZtBsJPyUtQruASO8299juqoW8xGCZ6Tn8dJLU%2BzbPFKy1e1uMNpw6rgWL42BZhmutmaqeE8vInWpnIisfoZoLdbibaQX6vjJvD8Z2uI8nRn88nq5G7Lzi4riVarnqP5KyVv8k%2BIu9J%2B0%2Bc5mLh0FvjOft8ovmm0dOqwyKSMeFhXQekFD5qbJntFSV9A3cMT1Px%2F6Hke7TbJTbhHtdyXoJJ5b8hrpWgenTJtG2z1Ceolr1Yxr8JaawKC%2FsvMVW0yIGTm3iyUYZuXzsA5oPDrr0GsPr8qXNCmeALLYYav%2FV%2BN6zUOGtBmDtPlVn4OpNpHyM8bzB%2FgOedzQTfoVDsQN6HthAi3hRRYYAetNWmeNwoHJ%2BfpayL%2FvDfaSyrA9ZAUyKyK%2FNGzJJCBHPXerk5YR1ZiH%2BFavrIRRfv4WVZen%2FDty7tGb2nVqL4iR1ICMCMaJML%2BNFaeZOYMx7MH92bqw4We4MELqWJhMQgpb8zppmWtNTp0hqaeHelC0TM0DvBICqm2SZjbpBi%2BGQ3mhiEmy3%2BAc2hUp5UqxdkQpSVKQHbbSFcGRJz12PbQAJHrUIXglHwc1pXJ2CayuAMjmGqZqJi9TLL2VCYPo3CTG%2Fu1MBpSUd5FuCCTyG%2FXj8otnS26YjcYSzYh9P8ryC7UrZOsrssnGZNV3JTdWTDqSxK5o3IWYYy1m1DE5NAwDx1ZETbFW%2FBEXU781sMkm%2B5Zay5MCTIAkfs3pCp%2Bl5Mhp1oWUIWLRHAYSSKepTghXU6ZNmaMZdJlCKujm0w6cTQTPhPWXkyiqxa%2FxFS1arfyMvaOy3M0klwdU25QrUZVGTh4dhJguEgZcNfY%2F2Dh5%2FiKF%2F5q%2BlOMn6BLWUpkp54o7owbfY%2Fdcsyll%2B%2FTXQS28vBrzSFFlz3qCRj0tVds44EEgFY6ckL3CEQCsFkoFtYcUwpmIMGF0zB25DcuZsaQZYF6yAjhou6V%2BRVtRRLY6Q4ZHWCtlUY%2FqT%2FQxjqC37Rh9B7ofsdQ%2Bn0D%2F0QyS7OeJqeaIbscT5LX3uQ3IEMKRS9tgOf%2BAptFz8KU%2F5%2FwBK3ZcZ51qBOuV7wZbDo7NMg2xT1Oa4eprRqg67LsSSs3wI0hYJgp1Hx3QirBrnol1kc9bpxd7B3kB4k%2Fk%2FDrnyz%2FsEIvlduG%2Bv9wa9zF2iTwHvdYgFpr3xNKUAh1PTmG92MgLmh8j6k2b%2B3ajaXT7ZJucB7emf%2FUZVZT%2BUsUwhM0UiX6seucJQba2hkmseXUph%2FNFQPVZoAMIgZqbXHBp0XGgtz5GRYFfPKgl6BkIi97cKHKStgSox5q67Dh2Gp600PF3tU0xquj%2BGwEsG9men1tx%2B%2BaM%2BEsq2dFHrjS2dhuPoAsUrk7XCW0MSMHf2YDVmX6Sz31goxhIbJJrq1ZRHwQtsZAxW7IjIA%2FwLnqUlazsiLw41o0hkBMNi0ecdXW5Vxe8LeLNzu0C%2FJtSdjTuhm68DQbYLdu5zysFCiKfqWFkc7WkHbh2Ju8Fv7%2F0fXiAZe2BLB0S7M%2FSDFdbdUHvS1xFvKgSzvvLwycdEuFSbU3ngPbcduMjAwIXWNfYwvOgVidEIY3A4ZB9ke0MDcP2q2R6ZYs3jwjts6Ct5EZ6Q8lp3zbo8dATDJYMSwH2WXEkBtN0jAU2hxFqWpWrGkbWDsA9snDkJlcaQW2tSKOzz7vrUKYfAw0CanTeusVGEi8e85ixnEtk2RYoYxnlnkE2BjYasEjRtGw2u8dfNN2PGkh%2FbEC4qqsFz2h7b1iKv%2FauBJzav8AxN70wvXAxnDxMpXWDdI0sEA5tiLGbVZiootIrWgKBhqtmly6OlMJBxxNSV6mnk1WkZPQoI8zR9kt7AdycW6crGi%2Fr%2F0XtIbMcDsdlQBLUYE5BXf0oOsWytRMM%2BOhKJ%2FuesXs1YjnYiOxck845XTDGBchn8j%2FLyz4Dtodix%2FE6ZzH%2BkRWvR%2BkJkEw0tODRHUoaznYzg7EL6Q2XE6Rkxdm4v3DhJdxk8zere2FMyYA1%2Bqvg%2BB67hdVCscYvMFLGMG5V6y5SlCHXcdn%2FeOL1fWGFyFseIkdd42WuvjsA2hknExCYRUALBKse83FERLHC4Mc%2FfsuonCj%2FCKn%2FAiYfXCnwvHNq%2Fhx2akEowY11F5tCTgp2b%2B9qTwKR6Y3DqPavDaQXRXiRcaNQFQUqSiY3zEhxkhiHPshDQuOD4jkrZM2bV8CezBwWcmCPy9byh7PjGhQ0YqO%2FuIQ3xIQF9smR4fL8q0jtthbpQXUkMcJca5ibkGqnBhNGPKdwLGNXhGlb0zQQEage1q9nndyzmosKq1%2BNr9ndHRKXBMYXg8brFrl7P4rTpJ7gqWVhD5EQFtblG4K7Nnw6xbZHjnSw95nYOLeVO6VK5NLZNPw5moV8AM9QoJRqIJJzkq1s7iPGrKA7eHBLXs6l56FFQvepD8AkgWZgq%2FCisfJL67VenjV9zChmGKmY5qbE9B3NrknTVyoRRfddmErk4PlsOEHqtaMzzTmqK3ieDx9gxjdXsjqWa1yIfaYb58VmgvcKuFR8Hw6Cj4CwXWERklyzMntB94Pv6AwCclYOMZjf5EZNwyD%2FMW2r%2FEsOTx%2BWpWAj3PMta4yrKcs12IF974lSxb3RF%2F73pHxWInR0wX9heWdazNqui8nJH1cnGrAF6EXUj%2BNihvaAxuNAZ5WXKBafHf64bvouYab1YVld8Ma5ufJQW5ZPqc2x6f5l%2BmB48SgeJcbzdDeJno6ZvoPnPuHNaraBnL6QisgBomd0vcEt9Sd0XgwztuGCgbSF%2B6If5KFhE%2BteTzd5kwYxAgenNmqyTG8muQHyolvEjsIEgvtNfGhYsiPfsmVz9lFC56cBYpsjQewXSii3KKExyklWYvUlf2TynhmzybUeoRuflC2hveEhh3Rzgc1XDP4Kflpb8%2FfyFmAoxWCDH761i7a3J%2Bbq8OuyxHtBcSYZx1kgD3o0wFbQUc9iydKBA34Uhu7Hb7Z%2BaoheivQU93h77W9%2BFtYKTs%2B1jTlxaC4InJ3oHAIHgEnV3pSKKtqGSvKvhMULQrP6fCFWi3%2BOi1RuCnLXyGIPfdiQKlS6XGnPSl03T4C51gRAJngY31WgPGVS2B99jBDTN3WQrZc1Q2OiApUTBjFcxKtg3DBt5I2R9fbKU56j2aAei6efa6zlBHRQ5dzhnKfPrv0G4t7bwYBMUa96qhoHdOYu39Ug0fr%2FlzsW937LmJnlgAXhEfpGk99wvp22Li5LgQEnUb%2F6oswF15d9MlaLHbCD%2Fz0s5AhiTKdmE54YyF%2Bu%2FRHQUB%2BTeVaffcuXqZ6qLCoiORQOjmw2QuzFobb%2BTvglCYcsHIsEM%2FJ8s3LUr%2BwV%2FwwoI9sfzygvPQwIsEBYKAyYGg%2FWWRDIXSlC290GGygHkEwvLwObJCeMennT5W7zH7hA4L9Q4PoE%2BDHKmLKvJ7F4xCHF7IK9GG1md3F4vym4%2FNAv%2Fw%2FYsJu7QqoysOQ1N7qtWkmH%2FGJxdKs5nTPHsUyPJWL5MjH3k0aDxLkNeonfvyXaKvkCGm0H36cG66YKY%2Bbvylc49Br3XXD33rX7%2FUaIRaaUeA5jDruQJudv2m5SXYY6KmfEabuLnvdBDqIlvrsB75seGs14fMmo%2Fh%2BSse0uV%2FXO%2F9tJPcs4EwYe7tCo0QD3%2FOP17x9cdfFFjoQrBhB%2B644X6AZeliQ8hn4RGochBSl8h9%2FtoufmpsEvZtqZ%2BVEHvjI2qBP5PzY%2BdKHG%2F84XV2A9Q7dspJuoZsRrZ8%2BGZ1XjlN%2B%2BIh0HvNR8DVS0CBAODtwh5b7BpgL7OXqSRzSVgXDxsnqg8uaJUaj8X9qrRHi8QJ%2FwNyrwyqumDDh2pGg6f%2F61bIF0sY449wn5BP269TDBjirfbb4kHHEHIYc8eMP0haA7bghqQ%2FNjM3655ZsDNQWKpT9CXy8VDs0ovivYNtvmoUkIcKodRxxad0RGBbEuQZekRj65Oe0VWBEXz4oA7nknGHs0IA4ZsHWwPRrddYKz8jx3swM%2Fk%2BCA4ImRCJkuxhCKvGQhmdGfRQIufmmdP88JKOyJ8bMmcHfvU9Adp4THdalpNkeMdR5%2FJzIGyYi5%2B2TKDf0xXlkbqExl6Nm9DQXIlPYlWyswaK%2BCtZrc1Ll4Uj6wnUyxBPIBu0lnrVV1YA617YJmX67GeuNrEtcoBQ4JEQLCSZufe1D9tpka8D2QrM5le8hnkbTyt%2BqRLuJhA7V5Bv1YN7WiJOP5wmgsoQb0ljOW8doLHxUHTxC9lYkHi9e5kzRZZKmPbx4lohjS97bKBfnhaB1I8g8Vz6hFc2Csx8AB2LY4dCjDJWG2sv8%2FbzNmjvV6zxuGxkiu7fbhG%2F4xfqyUB5T8DNffEEyGXG7%2BaoHS8FTYoy7%2Bq4%2BAR1mVmyTcAH7leW%2BKeDp2hgRn70TpOCB275h5C7y2to0prtqiPEa8J30n%2FKQpzq1IqoUuSBx%2FusJFWxX1e9iBWNLf141UmWrAKUUZ9coklWT5P23cMvmElG6KILbwpV6V%2BQwLka7MR%2BJc5bXg%2FclBZXP6lEL5oFvGWXCC4WpKBKNn7prT6r%2FzsK0yZCvHQ%2Bba%2FfZnnCgGOuLh4wcsnhJiJNBNVtSfki8kn5rVQzr6XnRnBN5t4M3ZgM4n1VMmhe1ns2vDxU3NIIJoPTg2bF3lPrWsWozSAStx57g39cRU5UGVqg%2BPhdPmAJ%2FWX0ZEacViDUa5cM95kgABE03wj3nugbeHvuMStiqPgmkhTPS3IRHdj04Vv%2BAiFGWt43O1nH0xkA7ecrpAWRwCCiUABE50EKW2DBZLVj8ASV72vt1xf6VgCggNpNFOz83I%2Bz8i1jcLWjbxw92bq6DnTePAYMLZAJJP%2FOvthRi8oDY7S47QEMwUELWaN3%2B%2FqMk0vCcOfbljBw6efCqwP3BFgUa6Zfo92w39cPAI61eNciWQZOxKa0Y%2Fs3Txngs79SeDmDJGRfb9Ytzfyresk5KchZALw040K1XANIA3WcQ8bHP%2FC9JJEdpzP0rz1dMsoBdq7edcO8U7n8L%2B6GO5DRh%2B7Kqgaff%2FGFkFZ%2FumqEsypADabYz4Q6sUc5pTPsuUTJnNvgWr57D8eBkp39GQYGG9255RigmYjVUO91wCdSaENUwZgbakijEtE6dkGQXmeqU0ZRN8jPqctZRM813RaMu3qDo1OP5tgBI01ZDi%2Fy7zYmTzKn%2BW8CJaCT1QwAouGo3aScW0%2F0eZiAbuDpVyjlNE%2BpiFfEeUPfry%2BV7DkfK%2Bu06L6s1koM21Xxs%2BEqw%2FS%2BBaX4A%2FWSKcwCzRh6CXN4Ue2yITtRMoJpiD6hfdLQtBNSCFZtKJPahews2miv%2Bd%2F3j78MiRD4peFmqvMtdfFks66pGcxeFiHX9FGmfT5%2BI6xLXrfarH5%2BwO2JURe7EuggGlwRhN9b5m4PQtALngPrEjjXsUm7VQ%2BYszUd9Do93vp6Mi%2BS5HVBfvh9YRXwISBSrX6fzJLxbcXlS9acYFknytx5WDMq9SVBL%2Bi8RgOGntM7U%2F1bqTEIzv6XMQwaTRMAM2CKD9KaP7E3OgBWCJ2uTGtLqlYWLKnEVywiiq6ChWTNXSRJjc2GA8nn7gDTssTvS4%2B5bydP%2F7TJ07dDWSq2TT%2FYEOpcRbjPzFPdCUMOke12ESwUQeJFjC2kw9VApp5vELXgxpcUfkf1tK55AXI45%2B%2BK9Qjz5I6zzZj7eomHatgh1rfGR9RyaJe4gnakyoNkMsKoTsYOw%2Fns38F4jP3D%2FSvj7tSXagRWTuQij7twu6ACmkDAi8DzbwU0Nv7%2FcUF27IOvyusaVq4TQDIDi3XfTJa%2Bv5obPRZESQLPDSSfR2s5bpOVauyPvpBQ%2BS4p11SKucxkOnPBquyVZyh1mUTLLLXe4WyjDsAN8%2FnNzgut%2BAMfX8JCJGm%2FfoEVYn79jTD93dUftYo5Ak3wBHPCggGG4VTeFuH62H%2BH%2FAkz7UKiGIwrfk3IU401TfI6Qa7oDVbU6hsEqjPfnTW2s1jaiT5ngemPHqH5dja7VB4CIrrLSq3vgNROxSy4OuIjyWVIaAVds5mgAli3UKDhGddR2U8ceYvXKVUyJ1jluFm7LkBe2M%2BFCDAWXY0V9KLdAe%2FAkYhIb%2BR5E25pXEDbJjadhRM7gwm%2BB7exOKbauNOQM5tY7vhsYzmSDlIGWUkExWsaRHpkA8uybr7DxkEszOQT%2BkDNjVS%2B8H%2BUGHRj5mLWEE8rW0zFmk9LtioIRj2Fqaen6HKbj4p2daQFXJ9i75t8tg4wrshVfkmLvt9HkGPY7YKbzBl5Jj9EstQq72h3xfoZXni%2BGZ4%2F4eYNGYj4AdpuyJdKCfeCAXls1A8UIItwxCYYoZwGsq238hdBF9uPVXa6E2iQ%2BfMjlVgC5flC5sB4jf3bE7%2B13wu7WAjiK%2BTywW%2B6Jn5tcwYfBN%2Bd6m8e4S7f39tVgNWCl7Dvb%2FRC4sdEYMnASGOq1vBXsC0XIPUDGNCJOu%2Bs%2FfP8wk3OGv9%2F2otpWzT3RdvRNKLTwndVPEVzcZnFmslMueC8GSGWW1cYdQJ6qMU8LAJbaQ%2BebdFJB2T0PFQLZRRPaYodi8QJa36CoRLXA6wf4irzZht9apn9wST0v3c8eQihzRRe%2FSSC9k81g%2FGny3XzqlToFHXLEMPaFuaq0sZDCQmLA0bmzzDd%2FRqiCuhfw7KCX7kuY6I%2FOuNsg8pVLH7C%2F4ATJUhZJNwJNEcYqBdSBY7%2F9O6U5%2BZlGxZLN9vuCyOZ0jAkUW2RtvtOd%2FlL0gp4Aeg1h67FHZ4MCHmiJyyh5uc%2Bq1t4WUu5yCYbmkofWqN4yVfxlI54Duf4bdkHuula92Ou1R%2Fxx%2BHYe%2BnrFx05jKeF%2FEZjL1o9T%2FgEAa6nJ%2FKXSukKS0nwkJCo0urhiHmQXdDSCpQbJ0pFXelIdqzGfU%2BjHP3fQ0vzrXFJ%2BRwlGDrdzetyneCMrn5nOEu%2BGytEF1DFXIGzWQPre%2FoH5U632A4q8MqIfaJkxNQfe%2FxLvFXEV176g1PU3F2HRxHJoPaMcHGGHRULqA5J42FM8Sxp2xR1kqapEeMuNTilMO63mb%2F1SxBYgxHvKQ5IZTHiX3hwu2kZde9mIrO%2BZizKnL8%2FkhFzgDP7D9S10sNgzaB8ksTMcVIdAJAv2Xpjko10iJDxthmkSgjhZ3XBQhwm3r%2BXFaLjoxtROoPySRr%2FwGeQaMThOFWauXAFANoawkoeQWQFCRCCIuw22B4QFIhUC8yJ2%2B2PqrCPm7gTlnA6Ccxi9W29dLbFB8%2F7VyaJPra4Y2O3BlQ2jrV7cMiipBtm8E5BwytEvAHBLBqPDvefEcHlj%2FnlMiMhvc2%2BQ7I3%2FfROW6aPa0%2FnU9ZWc%2FO%2FeTMoyplBfXclu4QWvEd4EbDr6SA5Ezvn09jijyt%2BHSRjRMOg9oATCNqXA%2Fhnd4izzR1mTpCHw%2Fk7LeohIkboxgys%2F4LD%2F3lu8aGUUiL%2FXY7kCd%2Fk89UVSbXnGIyCFlCOa4XXQeM05psl%2FDVKPLZz2cGohub%2FJfRukXvkP%2F3hKeklK2N3QOicNAoeiyageDjoggNva7Rc6FbU%2BKz%2BUxIO049OyrrLnOzekv7QR3jbZTbQ%2BOhLMb0n5pOeEj41q8Iw0pXOq3bMl7u8QCEoJtvLfYbWwRZZWcOQBdCQ6Qw8g2LZQWzVbOR77IIa4H2X%2BQFN9PZeJkPmkelMgVh4UaYQUsdpY1%2F5gKAGyS4QvCmXRr4Wfdrq271W8NSba0Qh%2Fs9n1L24wLSGIJ4ucqtvh5U5eSUIowTsEIOyBp2oQ7afrBOCImBANO4i9y8D94Eg6Yqim4qjPoTgaMLfBlL4%2FGSGu%2F6PcXRjFbUD%2BopPQ2seC00qSZV%2F8t5bY0WaONKEVsV6juzNtAiHry7MvScedC4hAIdSLS6SRiAyPL7okEJ%2BeXi2%2BF%2FD%2BOFSj1jgtfBlGHGTAW03%2FFu3k8LFLmqr79%2B1HPpHY4TRXrU4jtCb2WLahvTxUfSUVkDmeN9nVipPfmXUeUbV9%2BbRKLDI04N8I0mv0otXsmNAZl14p5bOtXknPuDfsk4GuPy0MIVaBhE2lKZXG6GOrS6U%2BblP3d7bQFO8ImGI3OdC9pJ%2BzwAUo1%2FLCBlVIAd%2BnJgc91C4mo7OMCLiqYfg90jvCtylQLF1NwQ5%2B3bKAF5OlH0Hy9L7FBtL3kh5bawwCzbfa0h9c1p0Dn0QYQmOfNr4r%2Fy%2BdYDkGRHUG2OANuDHx2y2KoG24YXE9TcfA2rr%2FPktkuAHnwA2PyeeJ45B%2Bpn0R63AVcjbvSoyWt2m8Alp30zAbZiIm9dP29CuRgOD%2BU4%2BHXauK04Xj%2FVr05ThLuakac1jqPEzc4%2ByQi2a3apFfT7WnMuHKrwGJkZHyAuCTnpz9rpIVg7iJ5OH2AuguVMUMgR1H1%2F6ibc3sduBkZZT3ce%2FVzB%2FvpecIPZXPRf75VW3766kxE5ZOxCO%2FMChpulg9KjXiA8a3WhmC1kzaUY9GqG6dMtT7t9SN1mRKb0BaiomsIHRxWn%2Fu0gSkA4W4P3cKFP%2FRlYcBu5En2haBXEsxzGQWfvbetN9Ihw1tmiJuoWA22uCCTq8KV%2FFaC3fnThaiU3AZworpG1jGC3ppvNWeXVsRArXwGsYsXHdKwFxOM9D1SmikFf6W%2BvpzBKE%2FiotR16dI0X%2FrKEBZFaNwgn598PVwXVdDFi1TijtmBmxbYYDitwJ5JLX9u1rmYHCxxZsS4XxgpbSIG2YfsFC1WQgzLplV5NMhUfE2jkh9egad9D4knPQf4cd7tXKXBE245gWtSOrEvkBd9vhkuMjA9MT1k4x0Mnl12pJ8845hx2Gm7Mut59XJ5PWR7VecpWImx%2Fcwnz%2F2cuYW457oUS1hPNZbrplEg7N9wp4%2FwzrRktLS9pQ2bqNUMiNPaKvt1Wcn1AgvIUmM9CtYzI98PQSajSi%2FaiKQe5IbGNZus7qG99eD%2Bc9zoitQZaEDblRxmA9D88gs9D9rSI3Ik4EoKyHok9OegxTY2BEPRU186Y7dvMrzE4jbx3Bfn%2Fl7QXZdmcFaYdPofUXSZXFpSiHCMxTq99bdlqncW8WvlbovpBtkuxzG8C86N5NXnJnE89Fxbw6v7dWKE8c8gBU3CJ4mvMDIyEiuVxJI2Z3FnLMkoeYcoLkH2k%2FspsyBpIfGvFBk%2Fify3mL14cy7VMZvz2VUIsXkdUk3vWJ1BJGmQ6DT3XwxeCeegnbhH5VuZOkkqMPf7NnDNXN4JLEw749ysGX1z96hJudtOiaAndJ53J9EtZWz5Iu50iB%2BxEoVmJGC%2FHPGnsH%2BoOcZ8jlK2uWRAegyCc8Boz11IbM3ffYDYxDhp0CRurNfM2BG0%2BGooK4BXniJl9Iu37BmhjgA%2B5XB2JHIplKbKbAutjp6tbtQBRosvnjgqYoGtinkSpGs2%2BVEQJg7Yjya0Hw7hbDq%2FJgURmYO5QDIvE9mJ6yHB2THRsjinFRb37c4AkNUEw3zZVVTk2FOm26O6qmxBGlYnlDItamNUfnZJxxRUTbo78pfumovf9WAU9ca7oO4HKIsTGnmyHZeRUFN5hSZXOly1aXxsVyETqSoLOQ7s258BmyOvF3eV3te%2FLm2uEp6vtVbPkOZYvU3jP%2F037JUWjE95eill79y55BC7C3mQqZyOQTwYbxHghaudW%2BBlZCpdp%2BovdnzcfffskeVN97X3JKdNNGXVi3aJKFA%2BLC6AiwPvRExBXDTDjikdgXMaAhvFLcT8Wv7DZTbaNs%2FzhMSYFaThjifK6MA4s9OKTbxJjHFPiXlNx3rO3CUkEGY0M4zPZoIPV792pAi6SKlzmUlcrVWRGfaQblzDhzfARnHBWGXmAu%2Ffx4X8F1QjC8aLSzVE%2Fyy6LJqBBdIW%2FF2L2LNeMspe3fIOdi2yJOaWUJ2D8LCnRXUQtpCVj1nLUiV29Y3qv26081dclEyLmQ2uN9EG6O8DeSme8Lq%2FubnxLPq87DrOUQQfmziMBTAyQIfQV33DKW8Htvz4EL9tUFs8s0ig2QCtU78SyfmVBms5n59Bet9ioeynk9diXylmTfd2YkRUiJmXhGu2pnD9YL3x2ylp3Z9FQ6hsvHH40rRdIpQqmAkst84YIVU6fA%2FT65hrLcOaDGxdhcoNNWMqZAn7GpWhhwcut9dhoxjMOuSzb%2B2XrYX%2FauQRix2mgr2vYAgJfEYRdgiEKwwpZHHCtGVJWSZSM3rkB17B9fBBsFegJe29ykpW4SPf%2BcZZsxpazivYFm8Q%2BrBOxfjthKXSWJ9okvD9hApba0rNu4fabed3VADPM%2FIJbLsuJu7wv%2BHJZtTBwT4P%2BTqeF8uEzlRh3rvHZHHmZlt4zQhE2ZmgsR%2BwShYIx%2FZ8qMTFBJ%2FzK7wTRHFkwLpbCkEX5w9um%2B2uMBP5mhm5MgsUQjdSEYnMlWzPUFnFTkQA52oo47xGwkrBkDmMw2xHOGSu%2FJat6lXQjnqDYfV8MbfR7WV3%2BEUNthWVtU8sUEqTypB2WOQKIymsHhBxd8JDX6x5gdlN8g3NfoWOhciW0LjfthTZr10A6LHEz39jwYACJPGD8Dv5NaHeAwJACRBgQmyWtUuiQbaXkNWHfI7Ir4x2gH7D0%2BIU4Bv3C3ZKTvJEdT9dEzZwjwQgON1KfppmQyWHygqsPf4kSNRHX05URlbyMAvKr8rCk2owA8oiwk%2BOTXuvGw1K%2BwSBPI1KthftwMvMK2vipjTqqyrT3JifbblWCf05cnrdYD8X0LBC%2B8NQ6dYwDMBaPt86Iahp2v%2FR%2FCsJJroZ2sngq%2Fx6YMmNJtX5984jf51z9CtFxfEiUKg%2FXhmcwmAYLEOCrTuJlXYWsaQBGaLE47pAoBzaRe%2FC93rEHkKS93wa6LTvYW7OEOuLtggQC0X%2BC58WWuUo6uxD6vNrtX0rWMS2XNE1ndN0s%2FQtd0WBJtosnAq%2B2NmT%2F2UHdY%2FZ66IMjg8SzkoetFkQfi9y531W8flpxIJTR7IXNhw%2BCfSQwtrFvnBYDclUrQDb2%2Bpz739kaTZ%2BNXMbGr5RXMNMnrN84C55DCJ5qtlnZCdilFjt77Y%2B1Yi0jH3FJEE9lgkA7B%2BfFvCCGAAC4gYzSMmA4IcX7EnIlrbSObdpXvFSpqMccr1liwm1AAiTuuvxvtrELq7TiAX3by%2Fi4vw57ZmUFqGh9svfTi3wrgNNjPbWDeu6lLooGKZmJl%2FDh9Mh52nwgpDg3kje0Uj1P8zCNG3hn8TvBKoP3ge6iSt3yNXMD5q1SRXJVFz2QcGdJIW02A65VjiwA52rEZpc1dm7RDC3Ew9JNjw5Ex1RBPM769r5J4NSk8z6LJUrEf2znCvSOVe7p%2BVTdXo8edm5c8RvPUk99KhBVjU4SGaWDsfsQHXuLSw7kIPELj0%2BxDNWheJPtscc79Wpfxtilr8Ya4K6Dy3fFNXuiOC6i%2F3qvH2L3SoNw2fFr0A%2BBq5Oi04knEBA1eqnq%2BPx12WKN0fPnyYMMkbrg2tiqZWmHJIlSz0vnXbNXB6cL6m55%2Byb74OyoxiTGwaVDtpW%2BkLjc9YQXe6p2vgICfPYUehOzHntJSnhosDjb6fQbCWHBTtMgrkP2RIyuLKcMhPLLQ64Uys4KB33xUoEtG0HEC%2FTIj6KvAN52pF3cL15yNY5bcTwxMLpJ2MNHGD2DKuNvCJvF%2BDR8QkwuASkBl8kl9CSCP7A6Ae8s%2FZ8Z3uBxBet5YuGXtWskkaPSo3kF4otUY8bnTUH6XJRSyKR4KuZwpIl974OYsadr9sgilLpUe2zrQPOum4noxc1Xf4UX6hozFxD%2BkDSC3VXyD0Zi2fUidm%2BU37a27tymBI2dXf%2FxVkFx2Z4ixyrrv8xGkh2LbTaMjQcXkK5iChVl5nYbbFCUC3S6ZXvmbkifYOGUOMJ011DOeTjXmExYMXnWHwJq%2FzZqiUC%2FDpvdNOc1LGJfmLf2DOliTkFMmJDZIFqfEFH5nyVSC8XYkrUa80k6LayMPkV5VplFvDQoDNa75lplkLbpneEZMtzS%2F3P8QqWHMK7YGRIm6I2Ypm08Kw3TMFZL3oUpWPAgii%2FF9qv2PLFZJI49WliSwTlqm9cCY1k1oJkyAaUp4K5rVgor%2BSDHALB3CjAvLSwVkXnTG7DLHsFBssBsCM6t2S8EZrMik1EgOKUOZ%2B%2BwmGwWFot4MqV02gwdPnsu6msvqGDTzLZtKuwF7q1BlgkwS%2Bt8sYjZ%2FtpziiYLF6rLU7sg2vWMWGxXlw5JvMB64zG9P3Xb%2Bq9ic9imdieh6%2F8QdbCI1ZcgFGQIslEhIySMJEQPaK3iawiPCm0Y%2Fae2%2B%2Ff3ChSho3lHIYsyFCpl9d3rwtSAlUJn5HgNG1JSFKpSPiLty1B%2BNF%2FHDr4J0rPv97kYOtNh1DJKWjnXuP0KlTHvr2rVB5eaMXQDPbWke5cDhBq%2FzpCAOtoqSIGDHw5BUEhBAdUqlNqOuEKS8uq3rGBPuRhPfxcfmNMbR88xv0iC2xAP9PGIPiL9%2BLrXJej%2FJdoFmpjtWRwUT7jkRyWVWzzjQ7jAZ8XiXB6JgsfClfjq%2F1QwPT4Y7UABJfETkJZf57yvm1Pq9Mh06BKnWnrZ6wbVl0Nzc%2Fd8tgBEP%2BN3wv2rYsiRVtt8OZ0Pd402QqOrOEmi77oPRMZnWOqeDNXwdjsBAodoqf5fQXUvJBkrXkfjcjIAF%2BaWcgcts0nIdEH95N7D6EgUZcbF75NY%2FKkGkWL7QFvAXnfyVaLaGQH3%2Fbt28Z1yo1XkGptJgAdalAdhaDaNCLl0OgJN6IiQKHqULU8JCck5%2BW9DflHyC02uakycHToWPNceQ1m2%2BTRBIsLuqjjzM%2FNW6jcjWSmj6egeLaW10w4m94oKpurGtfZKHw%2F96kXP8Dfn9jwBN8JK8brepiYtcJMsqh4Y9WKckTNiL%2B%2BOHKVwIifWmZuVIPOz2r9hGNJMdJGry%2FwjcK4qQ6ThZ40xDgS8BGXrHmo4NwqjrooxD5blfUxgvRsUZAqBU%2BwGAvmraFQ4p%2BpK60JY7XV2s4m8I0LxLkhOcSU7TB%2B%2F20Mpgqp%2FGqjZWaEp2B4%2BZX3zDWr4GNWGKaKNss%2Bc0T9ZbkNFIhOvfxdLpwgQHD47PZXN4JZfCAUSY0hqGwe%2BP49u0SqiSJagbLpujLz0wps0DilHrw0F%2BzH7IuM8PO%2Ba5ki0yo0pQHZTS6Ag%2FilryTt%2FZ28PUc6y3QvtAQPJOvDCwMkQJoKZgx2jQN1yGjUCXa3LbbDFG0OycY%2BzyXlQ6PhOKg57KTA%2F1QMdlLo%2B1%2B2Igqcy4esrgHoDFrTKbKcXaQDwy7QTeZViqp1A%2BxZR4VQyuvo9hhr4mOI%2FS8lD74sYooX%2BaD6hp%2FYhbNLK55K90TwyojBIw2eyeWd3tFTbyhx5CcrnOgMZDf9Omhnh3F5rqZ0aRWiU0O3KxFIdkUzFaZLod9r2v%2FWOvKuDb9GZnRx81YzkgmpBqJhbKZAkg7AkOIfkk%2FDnSk8s4hKrY%2Bdc4OVGTNWxMoXatmq7yE%2BvCZbWV8FWGcll3H2KJdRm9GD6RTkZ0EAwt13ECZPyQBqEieeAMldBgXhJIuoebLwZvIu9Y%2FNI2UB%2BU6StB8ObO2SoC%2BVkiLkcC2vIEdQEV2W23UV%2FN6Z%2FuKHJ3%2B37nGlAMLXWrKWfRWSSC7Qi9ITSXLI39v90lNO3I%2FtdYJ%2BZekQhK7HnlKWWDNLudwpyjQOveM53BYicGlVrgzNcRhbuEbqx8Uta8P3JtB%2BS8SjmM6qKU19YhPupkgrU4E47oTX1%2BstmsMR7k1h1kL7nXorWTm0pV58SI9%2FvhLzOaSjTKuXuhcTeR0SlcFsfB6uRiPfW%2FaiM10HE62YhDPQ1DoX%2BOAEZ9wYDc47QwNRu7gvFGQFqlniED35OeUB7%2FP4PVuhqIppTDntFHhAhZxTLl3ybXSvX50uq6nwnqHQpHSc4e605bwcGxjCPAwIreuyg6CVQsjjeEDP0uy3G%2FqHIye4OhnL9OuQcWKVIez7omG%2FTUcmtx6Nig3tfhn59Xs93YWHAsdGDqQjacn4FpGlI5vIKqIJU3doioEMuk%2BUSI4ogtv3X3vUuaGH0liO1Kn3iOs50rSVTGf0uI2uQG76lkUUDaa%2BI6w5qLyTmpYsezDpAwPk9ktEzCCmQQR1B3EhrbXBvRpB7YkAWg9e2dtQAV1H%2FBwPuJX5QI%2FoILpfH%2Bl3DX%2BS3KCCkoXwPCYI1cYpdJBMy4Y5k3tFqVOseU6HAdqS%2Bg%2BPk7%2BOZEpBB9D7YiU%2F7zi1tRbo7M3pWYN8kgcejTVb%2BdXTWTi3EaD1VlSYj3zp1SfWgksGZI0FW9Ht4EsHeZKMFVmKFhcE%2BazaTHkV7t%2B1T28jE4gKrgcxIZ287VVcJgYz%2Fb%2BV1gRFiuu3YThhU3RyX8snskywLPWmT4WSYBHAyYLxzS7noAIFR1WIS%2BwSkHk5o7fRdYhAZCFRaj23I0kz0X4khhugP4OMumVvv3tRNyQ%2BMy%2Ffz9O1YywSXtNh059BHQlzEQ4%2Fs7jFcrk2udU%2BmlVKGrGFOVi5D1WasKRQgJh7%2BMH5KqhLqwvkKw39oiBjLWwo5MeLw7wb8LwuNlgejtuFyNqLFTpcV3lwVOOEzn2HJTvbrSmbrLWPN20IyO2Ic4rHty1%2F5SICLejtyMm%2FATDrl3j7JmeU3If90r3ZJDkAPlQBijrImEaNfR481ZRaRfu1%2F34V%2FcWAQCj8Rv9UcuBEhPG3oR%2BvHWR608MC9fWXNxh4sxArNZBV0tu02bgfK2dnqjT7hvhcFC7mzXR3X8hfkYglKa2zBqtHTchZ8TM3%2FwPU%2BEshtZws7rk1TPx3y%2FVPCDP6SSXtg9ICELcRR6PGX%2BGQvkVr%2Fhd%2FvorgjDM4lrqjPd3lPx5fikZSXplhMWpwHr%2BgnTBqXPM0wOg6tUQ%2FJvQ39ESt5ySEex14rXDOwtTOGCb6IybS5YBORn0rA5YInvK6%2BvmattexquNmPM9y1buPG67gC5JjgB%2FXN6CmCFi1djZExHHWm7%2Bn9pWFoz9rMi8ZiheQNAA8Vd4xdLwc0Yoxzg6M%2BMVHny39jD27vfBiBWyaCAMPStdYzQ9oGCzNnXngs%2FnDVZkRKiH3ofo6HQf2sLfECiYT0mKSKcQ9opyqb96jhcxD8cos3wDVYusMJJwmgok%2B%2FTAsv%2BJ%2BI2KXZwRVf7M0zNH19XVkuBcAms9EppnjnJyYRJjTPY4FmZZsB69kP27zPLdnMiH4%2FmszBxr8o3jShE%2B1C9Qx60RqmrBc2AXCKiZ%2BJSl6gp%2FfmteG2LgGMpsS%2BysPuMaydAQMwJHX6ObOPuPq3IfGybBYXKgfTrDIPRHvwMEEz5TMkM1Z%2F5ugOlco1lHQt0fvU26kj3fb08HhD0EX29PWrxuvWbmSfCHjOqgCb0B8zq6rhL3q3yRf0xp3xbdmrQqBC0tjqSMH86TKnNfl9%2FNZ2i1g7LpfX6Z4eFiZ5quVvE9Q85CVNhRgq8lJQ7Dp15Z3ckIdnHT9qqmswB%2B%2FzRD4B8X93oBZfoJrO5PQoUQg%2FjGsdklviFSrYfFneMAGUfhBndWVWPnVKmBBnyXXdRedMAyAZXWWqP32RkZD9%2F3gWF46d6tn5ZN4xiyDo62OVZu9r4OjIV9V1Z4ze0kH4rh%2BlZluL9q3vsh9RGvm95TYPUCrYcyJm85jhVynPjwW8wN6POZ701JsHx2yxPcvaV1H9%2Be1703lQa56GvaBMwFJsIvSe5U%2FbyBCU6ztSIeXdqBVYiavw%2Bp4fLzoPCeMaXHkisxueTpKy4SKAeAO1NlaW00508AJcIkS3JGqXIuXdROfvmeWpeOa6v2p4Jg8axJ07r6oLSPx4uLTSUtEO6e0mvDNIpDl2xIMRSwUHnEMTheapD%2FpjxNFhr0%2F8AGFkGWzmf2UKhAklp4m1aLAj2Sfjr19f0KlNbUCNV3ldJHfC2O%2FUQZ3YZxTf2vcNhme0%2FpTZbwD72g9rlbrnDNQXpb0YWBCPB76wKbKGUt%2FfdSFjp7%2FNhlYbA%2FuNl9G3MALJ68114Ss6dHayJTK9PX1UYM%2FPKUcqs9isC89C8GKOz3f2RVYpxc2OR9drvheEOU0EjoXz3tQlXAjHhp%2BKUmK57S4qV1os6%2F8P%2F2a6Rt5fjcULDPVtOMvoSIp3GfTarLc4G6PdJzUuJjHMLm3UPxCj8cdzEIGmXkS%2F66B39dS7vyvsIyydVQ9HQnwV9QEzAv%2F0bk99yrRB9N%2FDd2VX4OUDxuRXJAVeCLJDbt1BXZoO9jwTSJFM5VQu6MasBZjiCqihJz5GX3xpFwk9W3vQoJqfGz85SqU0b6X0G2cR9zj%2BlQd23a7YOMsLgaWHPEi36MAtgT0A33HAdib3YbssbuNa0OQXm9syfXW%2BdlKColWAvez%2Fyuc14O0e0NiV25bwLkRGX6LcFeUbG8W9vQRYWdFYIdeU64BPm4BGB2dVaGqSes2xlEwr%2BRnC05wGJ4S6nWpDhnXmnCd2Ti9rDzi%2FzfWlIPw7dbFfEkMSuK3q48MfgFRv3ehDeVh986Xmgsth317lN4LigqlKhoRr5qeUfiY5Wo3zJThF3i8hy%2FJyBJXVCj1vxV7QMsPAQZjhOwOtgg0RihBInobFWTnuWMVDxvQxZTLqt%2FH%2Fe3wIsQ7sFl9d8B9mmXPaZOZSNQPP5sp8rnk%2BAKf2ELT7rWnUk0y1OseaJ6MW9W8TmWnLpTiEuIbA93XQUDoMm6Ba2NKZ6uuA7X3XlUmOGIAp55Q3YS8AZl%2FshwD6Pym%2BV9rBsEVz1pTcgvYdCQI09XS3jmaR1fHzw%2F3oU7n0DFVaTAoep3imLWwbquyPipT12DPlvyaBeW4lFPVNw7%2BaIcyPj7RsTPli7JI2imjRkt3XZC1ZcmwQwWwgZ2q7pHzJOze%2ByU5nA190hiMfBTOh1Odxpjs%2BW2tuWHWH8kMBZ9NrR3GHRE5G0EZvfdOm88tsWXZRVgcpaOI47yhVzE1iap48I7ZXUhoenkqR19B0dZNDoUaRsFdCwdn6a%2BEG8Me%2BLll1bovSdQYps%2B2xY4BB43QB0siREyXLUMlDEBQLBAxOIOmSOH7tyU5LVJ121Q7L1wuhIEHAC0QqAjdzIKaEOmd1T%2FhXqEQoHhoKTTj4ZSNdSmoXYEi2%2B6xlKpTA92OSIVTwzC1LMtl7dttOfE8oNfVoKJkpIjX%2F7ks9gikEQTqTIY4Zmlw%2BH3WGqSpC%2BpY4Zn8w8EBVhddFk3Jn0%2BZgVvJ94OS3L5gvKX69nHS%2BazXSuive%2BBqSANMqtZSOVhbSY8crzTQG3W8f7ifeX6ivxEZpqOdP%2F680AUGinT0Keog%2BZ%2BeU55oukliKPi0xp8UWJ6Om%2Fi7izZvL0vrOlVx36noM83SL3hGK%2FG86p2EQlMb%2FHoz07zVYjAf3batBTTUaqZYoDtjkzDsVt8gAwY%2BI%2BaBE1DUXiCNxXSmvthoOz5ZTP5BuZJXeZ%2FOkDDRTH1zLAOWkFSqIqwJd8Dr0l4k1pGUZL25WtRfVvt8SDz5wv2JLxKkq%2BB0Ut6B2jOwaMl1VLqL4nZV1zZRC4UeJHDqcN2kjhD9O%2BCJOB6lEYRVFwrDb5Z%2Bu7E4xdYUQGwzezn6XRH%2BDgAqvZYcVX3yAGGJ9%2Fl6dgKZXez%2FMjvUqfb0FsWKaUOXb%2FHVndIyeehkcC%2FQPcW074rNdAOhyyj6NO7c6VPKXK%2Fm%2BfWFOyKzzg4PtykcqcvBkpaOFb0FjxlcSY4xBo7MXDUaS5011bNRDwUQuaoyjEuTCnnx%2F7aVK2F6uIKYKuT%2FRkfWKG2wNkt%2BOCT3FAnX%2Fq9QmOmPeuoR7EEB0%2FKiUyNY5dVAYGZ3hMNG76vtJrWIvlIDyXw%2B5ZdtzMEN5xyLMuvUmlXPuS%2F7ZeKEpcSsP1k%2BXBvRdR9kXvBQXR6oPpxFCJLV90A05ZFRfP2kOjxc0ojOlm9o6m7qwLSQtxtwKWL4tTqCOMiirOcfGSGYS0do3yrVGd1bpsJV8sVuQziH%2BQ%2FHvkEDzhdEvIxssw1RfzcrPCxkGKAkBpSnCSlHzDD9mY5fHBVciZekqTkiUa8PPTMd4wq4%2BYJxl6zO4pmvvKeah0bTmY2UvzGGdsjav%2F2a3BGSau86gayZPVI05WTbt8TPX3LfPr8vDnXemaaUpoBDZkUh8oiMi9JGHPsQ%2BqtoV8oJRqOQgdvgrWlOE%2BOjXLkUsOQuLqWJ9OK%2FqTDDMZUrbeOErVVL2LE5S6JACwPeLUoMEXp1JkyKppyzFzFQZUqRDx0P37vWMOFNAKTUV92mpxr2XObtAAYrYGvdIDraNkulMPKP1a03THmwupUxC%2B5eUyA5QbonkLcbWWvyDhXUO3nafRYrfHwxN4wKfuIsdt4tL0pF4oPx%2FvGuRgu9hK0kVRbN9RzzznxV1CUwa9qX0v21rhvbE4AnyNPVh3HdcfCH51TAnpMRJ%2FwOWyE1giy2X5Bb04QHRWQHsg2MMWX%2FsScWJ3fYjrfFzBkE2%2BESThnNgFFbyFFsh9sVCpTQATJfi4i8Jqb6TCFU0C37ItFHjva7kv072k5Os7BU3cZI0IKz7fz2jbL2AEv3qGeaHqj1zTkdLp8%2BrrOpVrhowY4%2B%2F5O0xq5M5TUjKLQckuWsToTUVKidfs7afK%2FmNxVUzusvO1MVJE6nmMBYoTRb6Gva5C1HeGwJ2%2B90AjgohkrvI4%2B%2BF3iP8KlbIoO%2BTwVux63j20m0gKbVblJU3naG5adoT47JdWyjNzAvBjnYsWKpzSkv6lpDIhsGqxNYQ2jFhYMxXud3eIv0ako95i0uyABccZ9Aw5sIallf85%2BGWhPbMg3%2FGumpQfVS6uRMRlzYWIN4gzWjEf6JIxiYXdZRq3kH9OK%2F8tvlDohToTTFDhOGE1W2As5WlCMQug8sE6Theo5BhURN5b1rK0d9BbbkgwWb19deZq2r0VrtzSzDbgpooltT0VuuVIPDeNe30Qcs06Z7feCoLNUB1NODMi4IftIa6W6MlPug42WSft037ip0aXxtxVM6ZfOLnAbGfTMfVl3lFJSVj4F4VEmJuCRjfKDb8IeeTf0Wugalh%2BGNM82dgHRDRbUvktVMfpZmvYJv%2FdspXPrI3TZqT84yvOvQwj0lK%2B58Mf62ebxAC2fya4EVeyTOCOce1Hj01L%2Fh1GfeIxlyHz90FbFdC4q7sgTLrNIQORG4Fvwv5RQb28enKeOjAdkTEFCHSGDGZy6KKjuV3%2Bv%2BcvvFs90k%2ByjJTg%2F8HY0ZCT9oyYcnGdyZkJAkPqK2I2JIi2d4DQajxO%2BYVn0i1HtW5jlSbuctTQ9suv1DsJLkCSORJNSqN%2BHuUeteGTxcnWEeslRJCF6VrFtE5mN4OWttlVD2EblwY7bop%2BY%2Be8b8a4LXbwYGKExssTGnNsx%2BbwDbzsPS5cWCR8R4ySIfuPT%2FMeKPZb7lcsIjFab0NKOhOyjvaXVw1Zns8U07thHVhQ6ENX%2BNOd22vQ6lussew6i%2BdIQmaolnOmy7R0U4FG0B1yXj6tKX8H806nBRrfUNbS94XiQZS3QwJyagPOGXAWcpPvAAxDkAkwd6%2BM4gmosO6Ea9mudGxJ4G%2Bjkx%2FATvBb3UoePLWOyZNnuZ%2FDJbFgElaAYkRBNYhDWKl7VXMTK1A17XKG%2FWXvTJF%2F8l1UQTimlQakIqxRI6ooUwmBYCeMNnstUKvN3ureNDAVo3j74ApH%2BwqyUmqwuUCn5CYPvcyzPCpzZx0Wu2SCHu1NF4%2FshiTCJanXyQZt3gHWI5sjxSp7ez1COjNLSFKENT2BvcNK54jfNZexNvQVCklYCmOwqW9I9clPvpO9Ev4EUutmPPT3QjRY%2BjXwXL3Jfk%2BT47Ys9OiB1nqn5C4C3%2B2zxkFx6qE7G0Bkup9p748hCR0%2FAoZFiMZDuXxsaM7Y8yh8tyEzmkJ07ggz17WpFW7%2Fyqiq%2BGJoHWJ2qpCOwlXlOBNDxttOzQHDspk2mXacNibt1sLpc2LlAZMu9ytPGJkmcuSNVBQqbyCXf9BS%2FlDQ9COMhKdjmEy7%2BXafM8oM1Yikv9cJb1Nqqigoq%2FpNQsxhoVygIZVGBRIJW7B35boEGXmhI4aIBdylP6w5m4gA4ciREJGfT6xXeMW1v2dlekmkFtxRGYxSAbFsO2xS%2FjkuiJJr72mGD6W%2FsXb2MZJvRkTGIOVlZ2x8PHhZf%2FpCj6RExIZqCjIbP0oac02Z1bPAdMRAfwFufm2cRNiF%2FJtW4%2BMTqwGslhkXpgGSqk60qrBiQhIYTX4uiWqsPiFAjk8zf4NcMW%2FRoBRYP0OTbAPXSSqIHTbmdMVghR4z258ZjQEThWSzH%2BpLrDlbfcdf9FK2ghS3teNJM2Na5l9YYKQUHOixESpzbKlagfe6Z%2BckgYGNl7Ra08Fr3jDi9DjQQUhBPaxNPNj3UqA8lQ09JERnrJvKiOSb0YUjPBxbKvUNjeNud755LpA82j1ZL%2B0KnXzRUuWPLlTYgX8LskRzM3yWHxUIf8YIHw%2FcxcpJL48SyVvftBqMEf5QNk57z3iN%2Btd0uF7HNS2YXlTIzVKXilbcV4b3YtWlOFfIZpW9%2FVBfzh9dPfB%2BxMOr%2Ba8diFAZZm9NWCY5gHfo0m%2F9qoliIcaSDSC9%2BUxxzDXCis4WvC4pTxdjhyWAh0Tkoy0wx7faA0HqHG4KxtuxKkovxLt9j2TxwmpHv1uc2Cws9G8kftv0fnBvfxPrqD1udWsY9oA56ayyPGlB9VBo9KASlzRo%2FePw0FKoDGta%2FJ1foQabDYOf6tfPycw%2FAv2e0tk4FI9hizKrX3ojsDsN68rAjHdJ1BUM%2Bpvk9rZ6fnYE2hORBdu%2BXUNFDUK9gTL6hMTSIyD1ROvCuZDqHv6EQkfZa78TKC36K4Ta9cC%2BhXADVG%2FWSxShReKNh7psw0iiIxYjxoslh6qr2oYsheXgyZkl3XdXr4oVEfX2M1rLvdmh00A6mRo01gATHmdq8OsxhhWczuQMuAeRTsAydbraeTuMM2OQ%2BRb5M382zvKlzlvvqNs1%2BPKf5HKKFWBKSAKEoWqlTPVt2yYaV%2FLGYR0SFqRKN31FTY6N7TfLRmx0SshI3%2BN4XwLL0wJLMdSKTc%2FchcEBMYBMAT06ydSR77%2FPKkYtWg%2BVMWt967Bf9xfQjhdOgM30mDV313Lu5ISY41exmKBvnGvrwWMgNv989kWJ6PNpX23%2FLlJaVPigxEdO0YKQc2EvpTeMAUC52uc9exx5KV80yoKzNGVOBfLndYYpUXN2pvgImQy%2Bba%2FNzU7218pvzDC61D2l0quc3XhIOW9H7D6E6t8v43dTf7SJwn%2Fajhq8L%2FvSUdxlKBTXZDkkDdhI6KUavcM4pllG4hnsi5VL3Ru7mnAKBbSH0bZKsALW8TnNkmye%2BCCs2REnLmwuSs1i2XCP0LhlZOcu%2ByF1qjCUYRGU%2FuRxJrW7OK6SVWu15N0bmMGR6SFyeZfLBwrMz2thHgp%2F7ThFHQWuhJeECzMlA4A1SGmy0oD63kJgqQ6j%2BM%2FCvOwozOs%2B2IbDH67nnnD9QP1AS92EbNeSPOz8vEr8ckLeZM6NSJBmRTiYkVKDSt5SobhMF4l6BnoW6%2F05V%2FIJhQs8mctqnAcx7otgOSfSqhZ%2BBftzHVknLLhKViEfkvSkgUk4ZDuncXdiEalUcYUS4glS2xF8kklF6qF3ZFENlzpA791%2BNsy7%2BMhE7CixhVhJPciXgQcpFnoylL19NPXfQs3vNC545TNpqFEUG2opxKRVpSVFaW7ISN1bNqfxE8KY4duPNSlTN0WCOuAjmgVU93z3IkUr43xbJPeX7uc%2FcJ4FvZII7pyRVFJrRT8mUtejAbotVNbOb%2FMz6Pij8YSUZ%2Bifx0MAe05Ur3SsqwpDhx61uxjyxo4Pyc6EIrCU19XZQ%2F4ZbVNt6oBVeyfcDqhcdnZjUCqta%2BaxFVGTglsSZ5WYQCiVbvm5kX6Ca2BBijRGZS9JKkidp%2BncdfA6F6GK6gixT8V0%2Fyb4LyjBP9Yr6anNQ9EBRVvr%2FMBJUwxB8XlQbHlR5XI%2FZvofbJ2O195OrtoD78ulEDbK%2FNfKXYRnVwSY7JecgQbPPmlMa2B2Ert%2F4%2ByvdApskEkR64Xgy%2Fi0GO4O3nUtlylQr3qoYlFob0LZZ%2FPDtcPUPoahzfhcAbnHHrz0RckKmzsp2IfmZJRfGH8ypq9duGMbWKBcQJiQDOlW9YVOxcxc%2Fc0oEtsooLFsc8%2BLGmlt%2BoWqvSXP54fzIogsmPyH%2Fujrl0n%2FTe9VJyLmGGWXwLLvCqZqIwAMliZYnnPO7My4rhPjCRzc4FZ2uF7Ya%2FH%2FxaIoJSOIv3dbFD2n3NrC7UVaUf%2BvetAKa4c1Q4RM48YtQ9N3%2F9%2BSjuZCTjSAC8F1Va1W3o8t33%2FvbE4bS8ufOD1gQfP8b1qy34OWs6pPUsnDDcznuSDwTPmmSSvJn5ykOQk2PqfUozl3RLX5XcasVF3gtXYv79w1qHMPHZfS7dlMVUMTJSEHdZOHQjWe9eeKU%2F388mEqK%2Ff3cWwLeYIP2FSZlLIcCkQ3cQ%2B7SDqV6U%2BqMPbuJ5x0ZaHZwXPEI9zA8l8JFfce%2BXpNUo30fm2To5YrpgXoI9I95U3dIrx9f5SAh%2F%2FGSweK%2Byose0Nk4Uy0rzp8Nn9XMVBwpCdH2dKy7fHxxlItt9SyGB%2FE3EWaWFNUHYbYZPanjuZbvhlLzCsl31zMGwZBwwKip%2FqfDEHQkC1PRNmpbTz2n1hS9c1AHEp%2B91BmtDcv%2FlXW0Lav46JMW%2Br23cZ7PL5USG2gsGCbYfrS%2F004up4n%2BveCj7WfCL2R4c2iOM%2B5fccnBw40MF2hmUxx%2BGtJT%2Fle6eZl9tYfRYdg5lwydR04x1n%2FSdMKQt3pHOOX%2BINAcMZPAjgLNvDW3jNmh%2BPb1%2Fpizqa08PBfSnuSRlkmtUQ%2FtFZbdcGTbyRIqUpe183UmbgyCR1DcwQM%2BFiOOmjIOs%2BMYA4VcI1jWFWNYVWq3tg1OXeheqElo6iCgxDuiPHjrHE%2F7msUY5%2BemoNk7oQV%2BsMVQji4p9gsXfdZLeRtnyNagI2dhm%2BLsGLC55D7sQhTmiGfsMd7s1kr4UoP0%2BPHmZR%2FAIdXYDFkxbjd%2FcVsSPqF1GS5YjFCVXkiuWaXCjSMd6SOo2sTK5YyxGCsKh8fuDDZ6KFKq%2BrKb1CR5vr3oIDhXXEjllWM8e43Txbx1WKMfFpEDYLckZY0lz6l9tBbLMUEZ8csS7Sb6OLqR4F6%2FC3eVHpJk%2F3OSMC4SXqjfmOrPs6nwkLpUD4GwuA%2FxIxLfee6b2kxoo5fjzXHxxu2cp11C%2FeFkxlmrfQHB11M8nVj4FpZXrJU9aT3JVYBqNZMYVELKovC%2Br2JjoMu3Y8%2FMU7RHDvXqkPYwT0fsNy9nK9YUftzcAl22eiXlOLFMIW19QpvcKNBcfiB2HgGT3YjSn5YtpHTrpJors3O2ORhMwhwDEnO3gFvRmKUi1t7dD2QU1Y8oyZKcIrqhlC9apW%2BMZd92T2L0AUgZqBksJZnr4ONF%2ByS%2BMRTeflxJpRWGvT5I1ew1v5ynnttm6yOH14UIP%2FhJngprXEIXgQJ6upv5xyBlaQDU41j%2Bcq9s7vIZJEwN7KOccRDBlUoBoeIyaWlNkcjR%2FBooIV6RfSO%2Fn2beNZhvMHkY2wQCpWnLzFWi55zXDNQO%2B2bJ%2FKhsa73Afn3wHdz5FQ3CBiHcHCf5eRLP9exEocL2tPfMB4%2BZLDQ9IXJoh42R83qwu9QdVX6twEaWNAenjgDf76gRvuHCV4bjipvFDQmj4NrDtEWWqfJ5CNGx%2FpI9Ytu5A14bBluhr%2BTWLwGHvKbkZ5sJtK%2BmEnTxqiSMIjMQ2uJkG%2BNSt5yWQ7OQJydYPDD4iYmC30XV9SECqAazGih9VFEKfrg%2BzL8ysPfVC%2F%2Fa9OzrzAvYCtuot7c2brMdG%2FAurBGMEywiWRx2pUTVRBCd42ez8zg5D4yCwNBNeeSNBswQ7JNvL6xtwARA%2FsIvUO6h2ek7xRZTMyn9bApOuAp8bC7h0erLBeLEZWVor810PmjoSDKNLTbBX9irGBeFqDrfTCQL5U4x2OJRbDShcl2a0NEhBTfYk6VWJmvWKI69fkot2WxvUH25cf6Y68rNOWi1wQY3l6SIN080BqrXL8k%2F1pPf9Mi%2BwuBhBhidhv%2B0CffxJiHvnsaRmWZvN0NKTqPeDCT3i2IRQPpb0x5d6yCOjrzJ13O8sase3YXK4I822nRiao2Ng5BQJ6NCoDXnwzb0yQd3xMUts%2FF4F6JLCHCtKbRHWli5uzyapplfs0I456WDvb1aSMx2M%2FnS3k4WAlMeh%2BUZPVtjc31I8y0vOWRmour9v5EO%2BPvtW9FSZIB6TQGwpDgSuc%2FV6iyAnh1U%2FNNDVvB5qARcFn79YD67gxjvBua3Ya08ieHW1idTkABu4Sif8o%2FUwzpPI8MWdm%2BC4soHLYUqJHeuTbdDJnOJx%2BNb1sw8l60OKF89OdVeR70227%2BvKIENBf6hrVvDsz1T7verQTgYpN98MOmNTx2TL%2BqcnDRVP1ZORRN%2FCqAQqqcEYYG47Qi39kiD1Q%2BVnjV%2B7tmzPFLU1%2FAwvs%2FPLl9wRwRt%2Bp3Wkdx1YjHjP2EkFczUEbGOnEvnkgXKKcVrJY3lppWoLzPZq2oSBXQCoOCKngBVQwc7kADrcHw7myjIAfNv10RNFtQWngJ25RMLWoWjlV9%2BicotvOS0Ve%2F70YRexN5LCRHv4F40O1jvrA76OLqy%2BMqyDpI7CGdAsU4iQ7jPbOddTnHnSCiEBZSAB22r1Zvs0qoefq4TyP6eD8TKLL9Uac2AAOvjEK4T6Mf%2FVQFGtazPdN1GBXL0mydEHqJyyNXHFZltaFOUExSAXakIZeNaOOuXsuUK6KBOSd1Q59yQUeTsuQybfbEI2E80HsuUQTidiLk0Q5CfUzvFHjYy%2BcgOjpu8SL8VtHQCtZOcOdOb%2BxH29hN6ffruzFEqeQuqXHHuz0Uh9tvA14wAowZNU1CnK1g307aDt3QGk6dVHWFity3cH%2BMCJqlmkpsmnS3wwbPMUgXLfLldONTLmIkd%2FHk25SLqSiQh8Y2GSyz06tTm%2BH3cd%2BG64m8003rFG3%2BVZJbc3QHg95Qai007ma4cS4XONnqCf7qrFVK%2FUxVZpuZySzC2uZKGAh5vftWG7msVRy8kwIdbA4fL5jmw9uCI9KleJPqoCslkY3m%2FOOZ4aB8FqcskRjQ5wKt199%2FUQcxRLpyb6ok2GwATWqhK8WOgAB49ALlUQGsM1aYt77YwTqcCEAClfiJ%2B%2BjlFQYNaBZNYxS6b9UPTyVDxBYeXv1BqYSLrsKs7E%2F0wpHtW9LHNnn%2B%2BSEnNMfwI%2FZY1GBBIx3o2eq2ya5Gyhbg1%2F0FSUpcwnOVl8wxRzWd3d6QI6mO3Nd%2Bht1%2FxAqbx070UwjORfbxTjAID3R%2BhPBhea%2FB%2BqCRAowPtvyYtKgWx8sCVafC5J%2BrI102g5wz2mVTPe6AGrirX124PUb5YMhYVRbJgYPlyE4UCMIgubwu7bM680Q9WKIJQiK%2BPx1C%2F3ZU7Q9nccUfuUReAtJUw61Wr4MtpHWSFoLW5wzM9TEmVR83mpcnXGwLtf7Y74ReSZrc%2FkG16z3%2BEhQRPo4ucvidP4fm7YcUs%2F7lOwzw%2BZYea5z76jr50BbZ%2F5IgSU6k9BmeKCt8tCQDswnAGMQCpF%2FDM2KYbHVzAtSVT2zLNCLQvUq1kBVK1rOXA7MKWJr%2BC7xCV0FJ351Onzloe5B9VR4oe%2Fmmx0hSLYUQ8vP%2FGaG78B%2F5r%2BsfUHbFaCKZGOuJMG%2FOwqC7EFlR9hKwJKt%2BN4Qsd3K1Ux%2BRzT9Q2damunXZQU2YnObMJf0dPEPd17wf3c8PMBp8o7YYNeWlQ8y6xqG5%2B7iJm%2FJf2BnYzRA9Hc5QY82pDffnbjewUAYT2e6ykR3E48Rk9KC1VZNnoSXcSFyKX2RLSGSgYrpjzkw5AGp8kGRW2vtnUQHy287fxHU2wtW9Dhiza2GZMtd1X%2B2kfaVG2YdcHjbUqHuAsVGomtsWUsShP3BYV4h4Jpa4NQlJ1Yq9IzE5Z7QFFXcPvo45RlWgjiP22Uw%2FK3AbKnazGr4UD9swDzfb2SP1Z8Kw15GDVirZzOSKPEF9bXBXt3PzfyPHuuYCxFLK8lyTs8QG56BMojNpOFhSi0h4sq9pDRvbTntfN7bwlyZEzLd1r2gLHlag4ayGcYLFNZawn8kpoEohB%2BZG7bqdqgCXRwwbRkJQ1fIiumrxacQEWrqwwaGDDXcIxNrYUDdFg0mKQklwY2Ad5uBSVwKtqK7zWm%2BiVS3WnA1o4uOaJKYIQBrEx0kMLCJ2aIgR5axuE6JKS9WIpajZP02I9jQEQppvuRTAbs7q4DEl%2FQ8TN2Lk12oC5LUJ0eIJ%2BxJeqKO7K3Sqymf1M7OyuliEnJUlgTmjE%2BymjsYyFhu37pL2AhmkIj2bWJC1sfnE0uJOSolXE2TV%2FF1x29Mh27NziT6HXrjNW4IbZjMfeyG3jgo8UxQpGlZASNA%2FH5W5U8%2BOAIL6GfCpXfWRTBKJv0UAlDTKd%2FSHyn2QgKxvEkRuvL9eYDgwTx0gYCiW%2BzR7jH596hGd3n%2BtZ4KAjhAOJizXGK1443h22JFEe24F5oRBuEvASFIKzdgaBmPuuYTCKvX4TgFofiHh5kBdBTPelgG7TEE6sC9EFQXhSsk1umM3Ekee3jWJ6cGxtb7pKrWVcqEDcI%2F8%2B9KjaIDPu1Iw%2FawWi4ckyFCoVNnsK3mJsWJ5VzELf0Dejg8fEuB6ic%2BTS5Kj8oQOPxErBEavHsAj9i8f569pkRn%2BC2IkDevlhYM0kZhJbN9NDz9aLaRXSLa6EjVfVx0tknqeVBHv7T21u9sT8yi7uboMRGNLhL%2F3Lx8Jkg0p0h8IJJ03gwVyuGYCyfTCE6b04G5I0q7gPtmg5wXZJMjJpMLveMOkNPYPtJpD3MpcAmhhaM7N16nI6jTlQjhAgietfw8PZ%2FklD9ZIkip7xvhgtQOJeSS89tsxqu4HyaZW0OViKcGufHAsahYk%2BmwdBgl4rLtFXSZpLPP89urJjGcE2tegy5mTt1Ies64zau1Y0Ev1XAw98QKyWWVFpV8j6c0FgBQ3PAIKcMqXbLGyVbwvS%2B8RcbI36IH51k0HtbKDN9fOwlnwBn2BbcGo6icMBLi%2BLfLv5db3iyajreGEoXH30YDjou5HDdWd4ZMFWbvE%2BYjmegdt%2Fmwaxrx5e6mli%2BJEOQZQC4D0J46U6vfeVCR%2BFrilJDB0hZ0ectImAqToTsMg5RGBRH9452O8M8Nq1lbTJofFKBnPYbnzzUxw9UfPPdQ2mGYh8yRrCaJ7W%2BKpWbsq3%2BSYpTYnn%2BR7SH4Qa0cQBStOiTsUWgWRW0aGjSHQo2SRtZ%2B1yLWnqwTYpRAuaWgq4%2BtJmYh%2BJhAxBFvrPTOtYNSt%2FnQXZzI3h4eG3xzo3avfqbsGcdSfL5%2BBfAnFxcEuzWFqv9%2Ba9yCRr7cuFzCvYpjJEw502hjcrnl6UNAv%2FjXvuHhSX2s1AayzhkElTsDqjN9DMwNTSAoKpmxQI8AeJfu%2FYY6EkkBoUn1mS76QnGo%2FcFAKhT0GfUPM9F9fVjlcr3uzVUr7Ia7t1safAOcsRtHUbm4vyn%2Fpru2fGQ%2BTDVm%2F40GgwERK3Wz%2BWl63YmWkLwMr41FnTPzTiJWU1DMpUsuVqjjb3T2eGLyeEqbwulHWbhAOMVjenv77fQJEe0XYZHxH2HdqLlhgx1LRqk%2FF%2FTzqCqy9LpARKhFHizJFTvY3xg%2FBfUxT%2F%2FXBEhFi7EoVOUw8xXfuZooTFY7p4XP7WJmG06rwYQsjjt%2BFcuTHMq03YkZURWTB0JGVB2lXQz9ECB8d4Ly6MoWunwLifJfyLn1R5aPboKBPKVWBsLGLQGxeHhhTALFYpify2bf8A%2BF3RG3mxHo8rn2HNNXwMb0cKpkoCe8ALjF%2FzZHZg3e3e5CsGZRQn861YLalRk%2F5hA2kfe%2BiV8t8Bkurh5UnieL%2BW3IQFiHlgcH9Qemj1%2B5gfyn06qljDR1Q5lpiL7GeN6wVJFK9hc0ePmYIwY8UegI1KGYTlRu9DU9tWRzG%2BT%2FrViU2raefOfL5ubM36%2BIqk4LnZRFW1pEiytKgnBn5jM0%2BsmWD8W2AGD1pQqMmqVKTUDDYpPD34Whta18F3rEP8S5lHglj7aiJlur2z5XIibqa5y0%2BaeNWDYTLJjMRw9fFVBuS2pEYm%2BhC0NG%2FF0CU4kw497sk5nHIYuUr4z7yNbGa4RdzT%2BFKRlGOHI7kuINy0M3EKgR%2F0Pj3%2BZcOLzwXqOn9raU%2BN1UwyamE3%2FeIeUrhjcNClxoeVIrcBqIAP%2F1CYJutTsMiUVMeHMzu55vEslBQ49W%2B6rKH1ZMx7bVG3m4WszYOYuyhnosfnqDJ1DH1mH6pOLHyRpfwRbtThhXlVJ4%2BHN6zFPFByvooPM2VKpxadRQBrJ9u2DhP2MGjOXPq0ORhgHhUq%2FwEz9%2F3JMsrFBUyR%2B%2BDQbSkggda6sKF7SqROfWAI96216v2UChp9uZTlXcabgzSe6qQICvciQ11oS%2BPnxuxIgNwM8tW9LgqgkGNnkJ5Dw1%2Bv8mUhZvzS1qOSMvmO8X%2FT8qAaW%2FebYRzdmlBqC5gAu0%2FCnu%2BhoG4Zgz2XBDkdz%2FQAgYJiCKsUU96b0Kiw1FamTK203hqs%2BqT0rWBLkUzUz8v7uEudUF6df5c2CpF7y8VO8AcSRs7hfyGky5U7L8RQhiZGHGKkyqPwCSujT3sQP4RvE%2FZuP5otievmTCxx6OAxM5bg7muECnIbyn8Uh%2FBnXvRm9czA%2FNS4BW94cTO3fLcDwD0s%2Bkh8orVDwu15W18l9mbuSrfJgl95B1IqKxcchA3qvKlQb5GTUlYbsoZpnPkffX%2B33o5DV9WDGGhmSsgLZV3JHEH4%2FnLypWv73oB1%2Blxy3qOHrS46vsXc84wolYhi2JwEdy1ZD1SxMSmbyQ0d1%2BMrqClkeEZtlKRN3YmctwNKIYhZRccUyc9DKczVf0uJDrxyTc5%2BFsv2LcBTLke84thGW1l6u9yQyO4RXzCUWcBb0ox8X52cc%2Fpz73zEvkF6tQO%2BfewcYzFUU2m9wdR3WQQDfqtwhT5QFeFla0sLYTvmZ9Uzc23lzrBqiuVU5nQNXCdR4tuMJRT%2FQ%2FZ2zr6j0rPWFwIHT8UTv2L0AJz97LZgds2udCElV5pfdopD5CZQX%2FvXKbddnqACbRdNBZfR85oUjdDjNy3O0xx6QqtjjQkXocIsC3iJTkMlEleMeWGsBTX0s7DCg62vUZygiwlzikyOnUZO9j2V%2BiEHXZuBMMJm15ExgNxYQfGWNZJc%2FulokwxyDz5PzqjWK97ZE11iK2rWPSEDnjotMGXIENKpVdhfzrSBj8kthNtvnXlz3W0quwMbyeuOVqlcJeDWZs50QDJDUD9sr2TX5O8riw07xxyMB0RErJfN2qmscY4ASQDAvDsRDdOwJiepaX%2FXs2j%2FOM42Pyrabpoeh5U3fgnaf0arC4vDdXq6paLk4NAoGR3D0b%2BNbM%2BzWIMnT29Jwv4yWY%2BCEWzjfOwDmfNCgCW1Ofv%2FoejayCSpz9spjYMSW30P%2FEt0BygrtwBBF7CNCfEQC1odCqYTYQM1YI5MPQ2JyuLxRRqpwI4M6uNt8nXpz9iBU8L4CglJvh959Ttn%2BoS1Q9VbbM1aDXW8wXYwuxWoDwM%2FVvuwK7bk3uW2kftD%2FEW%2BsmcLKGL%2FNinbYhEGRNpvNZ9B7mWjSQccEp34ZuBwlXG1yynnNsT6VpDKvDTtm8DQXXRowjLpVCG7OS0B3bedv4jQscLsbEoCZ3VJ5S63AH0itw%2Bu9sl%2Fn3iM%2FVZQzbt8gatflVgnlBc8hzjrEM5uONVLC1ObFvIdmT4EXDLu0%2FdDhGoQlsqNzVTg6TGjmGCiXpv5oEc%2Fk%2F%2BgQQCDR9Eqaen1vK39Hffzrmy0TsARY4E9tpSI%2FCAlZTRpAL93yBXB%2FrOn%2F2rVYdfDoN%2BxbB%2FNh1dCZHQm2X6jVxS6tkbGg0%2ByuPwxp10gZ5ssNRagkKeNpUwWXUf%2FH4kd76Uk6J4sd4aMD%2F8g7x0E4v21J3RmsdnB1t%2FnLtCeSD9bM0bih5Ik4%2BxQBLuJUO3voNtx9d%2FYtyok31euJb3KN0Z3lQ6rELDH8SsnUaQ3l%2B2Nb0jnTsoCvyBD6Xg9sClmnG0H5jv%2Fhqf3%2BBEQ4W37ymUH9SsM2cbQAw5gy8fvVXblyOS9doitYrEYGAAJt0Lzj4j1bxvWUxm2cTzU1RitfINLLZg6nXn9r0m8P1dJv9ZTUDvjbT9sFYfBhA06vJtZybjslFlRXoVAn%2Bfc3DDVSA2wW%2B6TQyKTMp%2BMIk1rhJtcsFo6rcQRjtztZstpSs%2BSVdVwBuvsylXUcZ06rjLP5mIuL%2BpjvEmlu4wci3kaZyI1BxpAGcOeaR7IZ6m5II22ZSEiUMU%2FNA0CgV5Mrq%2F5kgx82kq7OQ6osZof28Dm2RnqW54t27K532aqpzKRfC5vigd8AWV2dNSh5BVxcGoaCCDHPfCOhL2iDCdECCVkKsIwKGcsrYWBct5mEcsr1BvcB0e4aYVxomaNBlIIxx5K%2BSP57xFa4uquYn8s%2BfsXzbh3meyKb5QmdFgoOs7M10PIcgS9LvZBu2g6hAoC7WPDghQ33m13o%2F90s88%2BpB8uK7NCDh8TtiQJG%2FByhp8h5VuhwNp7yySZoMlwdFtI%2BUyCigE47mHZuexE79HzmDWSDIqyDCWGXVRablvc6PLp2OfdOEq0QXfvdDNnUe%2BRfLYlaouIVbq0OIgpO0gxtZ7vWPTQp0Oj0OG1KTeVbE1%2BtzTdZwRjNdVxAMDs7vlkfySgNLAWA4%2F4TnmUUhdTzhumL0WHO4RkZjdnfDNGR2a05n7tiB29V1tsuFGm%2B8px1T232tH4ZoXTZroX6nTp8E2Y%2BJlvfyITXX6McR%2BindR9uBsSZbpO%2F%2FesfJLndz2O6oShVRgYVN0Hzrx5LxYD2fcVQi%2FSfmYepWbVIU%2FGFcQ2%2Bu5fmZ0G2QrLiIe0U7GpXyMHteFy1MxjgXpoKVpRaa6z6wFPXr7nz8gS6nMs%2F1%2Ff9rc1tEXZmGHNLG4XpmoZ4G58CTFrBUtygZnWV24U2xTVgxEtFKGD9%2B%2B3NtFFCvpxDBTFGNVDpsbEM3koZxnzBkFIOO7QenFWnh2bbmsFSEBX9w5xrwpncg2bY26SsFRlggb1XwT9FNq39LtI0iybZBrtgNbJUzqbJy6j%2BJigZ3piBfYpLosKN7p6wLPexk9WlR0Bzx4lbgDErVbHTVBQmrWA42E8L59tpsOg9miDbb%2B8Za6H%2BYHnKJ2oNSj5DdZmBYSi7AqnoeIvvac9bwO4drTgr9BCWpY1P9VeMlfRf%2B7KjXjB0vS0eqbwHQQF9PZWtiB2uyT%2FH5vFqWMg6yfO4qjgD7tgi7ZZ9Dpizv3Sxn8GaHT5XIDOpz9scZFy5e2qGcP1lBb2OYCkZ6ahRR1%2F8JUip3cXXsY4tMuKgEXAv6MvNM%2BaSYjtJBdasF%2FCBkA%2B84Qptu0H3QmVZLqsWj0HGotbmSF2zBXpVGC2DJuLR570aDjfDeKsUBKggo2WWKq49u8KxvqO9rA6ApF7ceAZc0rGfFx8tw8eo%2BZz%2BuvY9uHxqxCI9hQCPrtEOcvCvY2cghmVldKgPBjyHOU7Yu0G97M%2BMJZbWKlEkhXosLGxPNTtQgiyDBk%2BVrq7rJ41Rp%2BcHSw%2FXu3ZaOUSnB3seV7%2FoxQe3f8tslZonL%2FgJMR4L%2F6nIoxqDUe1b5TP6I4Pt0tj0W4X2sOc2%2BNIauoKuRqfOdmh2wyXdQWjb15zEAN9ImsW2Vxc9MZXrtIZoo8MxiHQUkyIql8IAFKQsDCKZCn9bPmi1caA%2FTUdEk%2BNY%2BAQsCRrmccoEGgYQC71VvAvRp6fHsjHroVBDSi5rx6c21%2Fl%2FgO1mRU7%2Fb%2F0LCkytOJN%2Byf11fkHVH1ylg3KqVglTEpDUx%2BYa4VKUOw21OIZNhUr9i7m2jlUagUv4OzkKyOhdnBSy1TOkgepWtq1zi9hhha31jC1RExpI2Q%2B8M%2BaCrgxOptUqo%2F47ZD3lF%2BD5LO4U9KrUgdn9yTZJ9bv%2B13uh2LgIqN2WLo2JJygcM3qGN42gZt%2BC9XJ%2FXe%2BiWTArGjSkmCP%2BvPpkLAyDVKc9oFyYjY0vAdTTKwTci0a1Z4NNd3mT9MQB9GOn9lvJrxk1Jwl0kwd%2Fc5CPBqUoW4N3pGTB25w0tvTDlSEUcAR9Y%2FvGpUBoJcQXQNQOzb7SkUpG%2BUT3fil2XM26cf1A1unwtH%2FT15v%2FzbMT%2Fv%2FSnIQIeLPO%2FN1w92jmrgQNtpZKJYtZsZytDyEik7g2la284N2%2FMm9O%2Ffp5PJDnsBSteZ4RALua%2FtgdolWAAaZVDJvUdcuYu8Zs%2Ba9uFxMMRiv0sv70c6LYx%2BIVLX7xWtmYQ9V2mgBpIyOPBRDFTe8S9vSIvZ9xXtPsGJLw%2Fy8nDJA008qjtd7cJlakIqJRTXoGRtBuFEWVaWBxQvMvQynF5AFBxXOmf1QeDZSe5AdLobKKTkUuTTOGX%2FlIpJGgNZtQh5w95MEbz3oaXLVIXPjIB2ALtxIQQPZA6MFuFN0pjfBWDhulijwpAwo2qoLMeWHhn3sLifMUvNkNcTcVkNow3sksSpJ4JXGirDUjk4oo8I0pJgRvhePR9zoN36%2BD1nr%2FHZ7wxaITnASO7btoH6qgHlDfceIjZXdZpeySZdcUQLu3vnzJX8oVqy8cU9482du%2FcCiBqn7Kfoo1luWy0eTpRUTEaOTBOJCPqRcZlMXBBbvIoqUynH8wjxkj8jAZNzlCQL8QlwjP3txrYqabhZgiPGzqkqdWE03BC0%2BDDR%2BhjFQ1Wt%2FLBMlkk9mOCUoBSBLQXznSmQBzhbFOifRB0Jl%2FC529IbpH35ouSGIuqvD52bMgWQ0RFd%2BNi5iUYkhOUhlxGNZ9IZqOxKs%2FJ6h2OosUQy6aPosuJSt5YcicdEGFvLkLRLz4fmSU9tpbsEMGKdb3ecdW9sM7nFUVxi1zQhpWE2FawAXe%2F0H1tBonhg4mgrtkP6ax3x7%2FSJdIhMHKM84FlewEtzYdpQxKwqme7haXiq7BT1ecfcYKrx0Yq5i788PNcQ%2BUInE%2BqXjf3VUW3HA%2BF11s%2BcT%2FF98%2FvirJdTTSO4dvRaSML%2BzMpaU6mX7F%2F8oezwF9GMOqqeAzmGjNX50m7RpQfMa4V6Ug%2BiXOH4O6e1biB8ywmiOx5ZfV%2BRLiRLAsH8uCEoxB2PWJWZOlHXURq69YBGIe9cC6qTF6UgK%2F42SN3ad6kn%2B%2Bp4odYEgpN0tXHXOvW8W4af2BGXriEif7yRfsv5hBJtzFy7Y2x3ds9aXXC2CivlYv8eQ%2BndoIQb6NXWGLE%2BobRvnXDmXSXtFTG3YuWLf1W%2BmebfHeHPMlVtBy%2BgvttZh8K9Qt%2B74Q1EnZwiw4ouEIVJqMD8jwsWAEu1RU51KGfq2vG%2BoTYIrhIPCZZ3luEtRivImfqmTfRuxtxr4Jk1lnysZXBaMiaUk2OzXp92MlCFAeoi44UMpBDyCXi%2Fx32HXZknKRFliigClUVIDjfXNijHmzEfFOFnlw4Jz7BCpdzIAyUolQp4RFo%2BLyNcQxLE0HE%2BVlzFJRYTrqUt%2FuxkddKSpKqHMfI7UHaEsr0H5e2sxReJxOdWKDcyr5B305gya4icOeeQOq9PDYt3Abz9YUnxSN%2Bv%2BuwrWltsUJu33gefBfMeKi8%2FSUKClGaf4nilYRzCSELPZ%2Fb6v8%2Fp2yMGkDzpefM3IQJeODk8a4dklQNgJRH7QI9frEcNjFnYlYFP1jUiwTESHBC0LHoKOULSiIe02%2F3RPL%2FHmqAKioKTL0w5hSMtWn06SwUQ50I59ovnteCuln3rskL9ZYpdvBtDPiCSXX%2FMhaDxZJRK3K1pMTGOyBq8L3cDtxrf0CZ27aniethrpSrS0TVUYhpb6Fgis2Yd8uKhz2GaJOdOk%2FryRCyw9xK1Gtxrlro45COl1noas6VSAEsKJB%2BZhHFWWjzsWHV5e0WgjklRpX8g0YmHIxWKdBEPbNqAMB03WgBEywhhvEhGahouZ%2FioTKM2geA9zh5wKvz8BdJ1SL4rdeF1rCcZSymqKyVdCUun6JO5EitjNH%2FpoyVwqev6arFtCJnd575uG9svMniS8PFRAIE9kTlws3ybBSjHohKwz%2BYzmVszLYSUNcUTH08jllhsiQPSaJPdhS23etPwNerMVlTi%2BgHCLFJFjJ8geKg20wHta0Y5zp8rpHzP53X7F73Dc6lIOAlJV5u80cWUlIGu3uhArSl7Mv4f%2BQJFD9n4Q7qcqeoTJ5fAYcAnrOyYg69Ha62njttGf%2Bgl8aLKD8y%2B%2F7zY9jxsYEj7yG2EVg6AFBXS%2FYn%2FLB0Yaw0RTVr8BiqWRd8DO2z%2BJbtcvXHhiwLsqHkGByunislkc%2FcI5tslLsfbC08ihElF9qG2OJ798cKxEvaf%2FXv2H%2BTz5Nc8Fi%2Bb9CX2n2Yy3tq1mBD23lxWxohbt2avMFsNR4jHd17RI0p8G0Fvn8Vlx2KldBxk%2BxHeI1nCkCeVHWB1VMqcvXewp%2FlqdZ6cY2JpKgPcngUPEWiEk7Ud3tS68Ak5jLXmI0UOJnoo68Rqr7D4pL%2FIig%2BocTyWDJlpJEnHbt7JdVVIYMC%2FKvzAObeMwixLl7cZ0ePWFPLvx0vJra8ukNt0MrWv8aevpAK0VUoYVD2wCJKCu%2FxOJ1g%2F6gObRvzDFouGCfZRL7U1WWcvIevg9ZH6JHL25Y8gBL7%2BB%2BcGM0dviHAb%2BSDzKPT1LPm2padDBWuSXmaRl0PJYXNFCEpIiPvab%2FrifiHmYBQfIUsQf2aAfXdeETxJFoMoSxCHUFQyuOawF1JQVp1catGsC8H1HjrQ%2FaYIroDtqeDxetXqxl2HEAMSJEJi74G92xZK1Fq0yYwL%2FMi5YnrCdH2ar%2FB9C2ilpB%2BIPZSIDZKJ0FoFveMPy0oJBWADWOwzSGoIJ9C9PSb9aKCYiHB0k6NpxduJgIoR6qfs%2Brp0tzQVtcWuW47D1w6PrX38EngGeMLlGoadPEoO6BvdDDDstGgpQiwqfnrnQmKB6dHkKms826LHoogG2%2Fd9JlArLfATh9m%2Bi%2FfFYGsDUmiNqxP1NlC9TigFIdVrR%2FSypda8CQ5xEwQ6jy0avz%2F3rJ%2BSP6K9EByNPX5vjA2PsTt9EGBa5xpws7fjgbf0e3KU%2FAAPp96HDri5OIvwW1K5wb2iYJ2tDbRj8ykHJbpvNgO7y6n6VxeLdUl8kl6NA1ArJWz9FRRBVoiUh521VUkAj36oGtnohg9O%2BRoaNt4WyLy2oWkUKHVOgTlhIJ%2BY%2BHquR5x5LmyaZ8YTRm2KEiFlBe260srNGZPW%2FVWzbW%2Bf%2Bv2pj4h2BC%2F52EbLxfDhpJQANTtMrdSiZq7IiqDTUnVA2hcm6iBvoWt3VQgSc%2F%2B1yuwP536ahsNgYr8FUaB4Xv3Us2OQj%2B4OT8o1jytMeS89UJQej6HclJpMS6jQKgiOKsPm3%2B1nG2dTHr142xco575gOHZkEboJEWBq2NRqsdl6C1HSUS5IVmbbDMEWqjMZ%2BEhlrDZpPs%2FJMisSMOaQZWvJefUeW9r%2F3%2FMdsZYfx22loa50396h4Rn9im%2BtwPY%2FYxjDLTZCr%2BS2%2BXbFNhLgVzmfJA2V4PdwYcxvqRCPggHKTwV%2B8cE4qjy8Bkjn18GTpVARJA4b4HOHhwcdxuG93%2Fql29912aHbeb8FDDxBZgSLYr2rHJ7uArMJCW9mwmq3cBufdNKKF7n0Ur7wsbkNCkLjXp7%2B91ElmKtOUmkmTD2O83lhSNZdhWqQOB%2FTsBax7UAj0n%2BsKTau6aeIao2j89Y02jBOmu%2F%2BddGY916C1PHkZkWnv8rar8dxfSSIRKjb0RQazJ41wt7hx9FcKGDVVBcER4rutV97KIFGOvowviUqdFhadJ1DKngKNAx8iMu4lyGrnEXV%2FysAnufmQa3G64wb3tLIJcZLtLPfnpwVepfv21ZtJwHwp95Y41vIsHZl9wXzkxp7DPnAfYzM%2FELTQnJhi%2Fr%2FcR3WYJoApDf786xGkYE%2Bfa4%2FwOdzcbQVCrdEJ%2FvcQt8KTNmgWPNAK9aNHFOIYdpylThwNd4GxTCQBfGHuWEu6ueZjhPzmu%2BpCNbcEmjnYYLEexdVPcfMKMpWRPW5%2B9B4X39dywlKYEkcp3RmcLkFpREdgctCyVIkcpFw%2Bb31fE1zaWjhTspcUUqQL3K2bD0rwziaT7hyhYG0QYmQ9CbGzcJVmYkIej%2F9UKtSwbCdJKme4f0Ko4Z5cnSui2x1rx2XLb4rnYi%2FLzVS01Q2IfckAkKPTBuCRQM93FlXSvOuZaEnGa5I6i5cwzGcYLGxz3UGNuptsZXj7PqqpMmyyeHruwNj%2B9ULB%2BK4cTOPSAVDp33AFAweeqv8VR6IXl1BCTS75Gcq9zv9XCt2%2BKeaOP%2FGonVgak%2BkLidOfYatvOzolu0IBLGYgk1vKBR0cwOMG%2BO0CH5BfQfKZzkUGSH2bE3IJenAL4Dqoay1alvbHdJjtj6N7GAAaNUGjL%2B74Sq8vmVep0HX2%2BWTJs0MFHKSD9CD2gscRub6%2FhuIOGiWMvA50mNRBaGuTbfcI%3D&__VIEWSTATEGENERATOR=90059987&v=$_1&t=0&p=0&prov=h"
            },
            "headers": null
        },
        {
            "target_id": "6515862e6bdbcfa9997f1270",
            "request_id": "6515862e6bdbcfa9997f1273",
            "host": "concorsi.gdf.gov.it",
            "ip": "2.45.152.69",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/Account/CancellaRichiedente",
            "body": {
                "type": "str",
                "value": "CodFiscale=$_14&Email=$_1%40GMAIL.COM"
            },
            "headers": null
        },
        {
            "target_id": "6515862e6bdbcfa9997f1270",
            "request_id": "6515862e6bdbcfa9997f1274",
            "host": "concorsi.gdf.gov.it",
            "ip": "2.45.152.69",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/Portale-Concorsi/ElencoCompletoNews.aspx",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6515862e6bdbcfa9997f1270",
            "request_id": "6515862e6bdbcfa9997f1275",
            "host": "concorsi.gdf.gov.it",
            "ip": "2.45.152.69",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/Portale-Concorsi/Concorsi-Pubblici.aspx",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6515862e6bdbcfa9997f1270",
            "request_id": "6515862e6bdbcfa9997f1276",
            "host": "concorsi.gdf.gov.it",
            "ip": "2.45.152.69",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/Portale-Concorsi/Concorsi-Interni.aspx",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6515862e6bdbcfa9997f1270",
            "request_id": "6515862e6bdbcfa9997f1277",
            "host": "concorsi.gdf.gov.it",
            "ip": "2.45.152.69",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/Portale-Concorsi/ArchivioStorico.aspx",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6515862e6bdbcfa9997f1270",
            "request_id": "6515862e6bdbcfa9997f1278",
            "host": "concorsi.gdf.gov.it",
            "ip": "2.45.152.69",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/Portale-Concorsi/Contatti.aspx",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6515862e6bdbcfa9997f1270",
            "request_id": "6515862e6bdbcfa9997f1279",
            "host": "concorsi.gdf.gov.it",
            "ip": "2.45.152.69",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/Portale-Concorsi/Ispettori.aspx",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6515862e6bdbcfa9997f1270",
            "request_id": "6515862e6bdbcfa9997f127a",
            "host": "concorsi.gdf.gov.it",
            "ip": "2.45.152.69",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/Portale-Concorsi/SediConcorsuali.aspx",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6515862e6bdbcfa9997f1270",
            "request_id": "6515862e6bdbcfa9997f127b",
            "host": "concorsi.gdf.gov.it",
            "ip": "2.45.152.69",
            "type": "tcp",
            "method": "syn",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6515862e6bdbcfa9997f1270",
            "request_id": "6515862e6bdbcfa9997f127c",
            "host": "concorsi.gdf.gov.it",
            "ip": "2.45.152.69",
            "type": "tcp",
            "method": "ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6515862e6bdbcfa9997f1270",
            "request_id": "6515862e6bdbcfa9997f127d",
            "host": "concorsi.gdf.gov.it",
            "ip": "2.45.152.69",
            "type": "tcp",
            "method": "syn_ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6515862e6bdbcfa9997f1270",
            "request_id": "6515862e6bdbcfa9997f127e",
            "host": "concorsi.gdf.gov.it",
            "ip": "2.45.152.69",
            "type": "tcp",
            "method": "PING",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6515862e6bdbcfa9997f1270",
            "request_id": "6515862e6bdbcfa9997f127f",
            "host": "concorsi.gdf.gov.it",
            "ip": "2.45.152.69",
            "type": "nginx_loris",
            "method": "udp_flood",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6515879f6bdbcfa9997f16e8",
            "request_id": "6515879f6bdbcfa9997f16eb",
            "host": "www.gdf.gov.it",
            "ip": "2.45.152.31",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/it/search?SearchableText=$_1&path.query=%2Fit",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6515879f6bdbcfa9997f16e8",
            "request_id": "6515879f6bdbcfa9997f16ed",
            "host": "www.gdf.gov.it",
            "ip": "2.45.152.31",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/api/@search?SearchableText=$_1**&b_size=10000&metadata_fields:list=geolocation&metadata_fields:list=street&metadata_fields:list=zip_code&metadata_fields:list=city&metadata_fields:list=province&portal_type:list=Reparto",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6515879f6bdbcfa9997f16e8",
            "request_id": "6515879f6bdbcfa9997f16ee",
            "host": "www.gdf.gov.it",
            "ip": "2.45.152.31",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/it/gdf-comunica/notizie-ed-eventi/comunicati-stampa/anno-202$_3",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6515879f6bdbcfa9997f16e8",
            "request_id": "6515879f6bdbcfa9997f16f3",
            "host": "www.gdf.gov.it",
            "ip": "2.45.152.31",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/it/cosa-facciamo",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6515879f6bdbcfa9997f16e8",
            "request_id": "6515879f6bdbcfa9997f16f4",
            "host": "www.gdf.gov.it",
            "ip": "2.45.152.31",
            "type": "tcp",
            "method": "syn",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6515879f6bdbcfa9997f16e8",
            "request_id": "6515879f6bdbcfa9997f16f5",
            "host": "www.gdf.gov.it",
            "ip": "2.45.152.31",
            "type": "tcp",
            "method": "ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6515879f6bdbcfa9997f16e8",
            "request_id": "6515879f6bdbcfa9997f16f6",
            "host": "www.gdf.gov.it",
            "ip": "2.45.152.31",
            "type": "tcp",
            "method": "syn_ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6515879f6bdbcfa9997f16e8",
            "request_id": "6515879f6bdbcfa9997f16f7",
            "host": "www.gdf.gov.it",
            "ip": "2.45.152.31",
            "type": "tcp",
            "method": "PING",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6515879f6bdbcfa9997f16e8",
            "request_id": "6515879f6bdbcfa9997f16f8",
            "host": "www.gdf.gov.it",
            "ip": "2.45.152.31",
            "type": "nginx_loris",
            "method": "udp_flood",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6515879f6bdbcfa9997f16e8",
            "request_id": "6638e83a0560c5f1ff54cdb3",
            "host": "www.gdf.gov.it",
            "ip": "2.45.152.31",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/api/it/gdf-comunica/urp-e-stampa/ufficio-stampa/@customer-satisfaction-add",
            "body": {
                "type": "str",
                "value": "{\"conferma_email\":\"\",\"vote\":\"nok\",\"comment\":\"$_1\"}"
            },
            "headers": null
        },
        {
            "target_id": "6515879f6bdbcfa9997f16e8",
            "request_id": "6638e8500560c5f1ff54cdbd",
            "host": "www.gdf.gov.it",
            "ip": "2.45.152.31",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/api/it/gdf-comunica/documenti-e-pubblicazioni/modulistica/esposto-denuncia-e-querela/@customer-satisfaction-add",
            "body": {
                "type": "str",
                "value": "{\"conferma_email\":\"\",\"vote\":\"ok\",\"comment\":\"$_1\"}"
            },
            "headers": null
        },
        {
            "target_id": "6515879f6bdbcfa9997f16e8",
            "request_id": "6638e85f0560c5f1ff54cdc8",
            "host": "www.gdf.gov.it",
            "ip": "2.45.152.31",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/api/it/area-riservata/personale-in-servizio/@customer-satisfaction-add",
            "body": {
                "type": "str",
                "value": "{\"conferma_email\":\"\",\"vote\":\"nok\",\"comment\":\"$_1\"}"
            },
            "headers": null
        },
        {
            "target_id": "6515879f6bdbcfa9997f16e8",
            "request_id": "6638e87d0560c5f1ff54cdd6",
            "host": "www.gdf.gov.it",
            "ip": "2.45.152.31",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/it/gdf-comunica/notizie-ed-eventi/comunicati-stampa/anno-202$_3",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "65ccd66c72add7dc6ffe97c7",
            "request_id": "65ccd66c72add7dc6ffe97c8",
            "host": "www.milanomalpensa-airport.com",
            "ip": "31.131.240.67",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/ols-auth/v1/auth/login",
            "body": {
                "type": "str",
                "value": "{\"username\":\"$_1@gmail.com\",\"password\":\"$_9\"}"
            },
            "headers": null
        },
        {
            "target_id": "65ccd66c72add7dc6ffe97c7",
            "request_id": "65ccd66c72add7dc6ffe97c9",
            "host": "www.milanomalpensa-airport.com",
            "ip": "31.131.240.67",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/ols-flights/v1/en/operative/flights/lists?movementType=A&dateFrom=2024-0$_3-$_5+00%3A00&dateTo=2024-0$_3-$_5+23%3A59&loadingType=P&airportReferenceIata=mxp&mfFlightType=P",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "65ccd66c72add7dc6ffe97c7",
            "request_id": "65ccd66c72add7dc6ffe97ca",
            "host": "www.milanomalpensa-airport.com",
            "ip": "31.131.240.67",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/ols-flights/v1/en/operative/flights/lists?movementType=D&dateFrom=2024-0$_3-$_5+00%3A00&dateTo=2024-0$_3-$_5+23%3A59&loadingType=P&airportReferenceIata=mxp&mfFlightType=P",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "65ccd66c72add7dc6ffe97c7",
            "request_id": "65ccd82172add7dc6ffe9839",
            "host": "www.milanomalpensa-airport.com",
            "ip": "31.131.240.67",
            "type": "tcp",
            "method": "syn",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "65ccd66c72add7dc6ffe97c7",
            "request_id": "65ccd83972add7dc6ffe983f",
            "host": "www.milanomalpensa-airport.com",
            "ip": "31.131.240.67",
            "type": "tcp",
            "method": "ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "65ccd66c72add7dc6ffe97c7",
            "request_id": "65ccd85572add7dc6ffe9849",
            "host": "www.milanomalpensa-airport.com",
            "ip": "31.131.240.67",
            "type": "tcp",
            "method": "syn_ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "65ccd66c72add7dc6ffe97c7",
            "request_id": "65ccd87872add7dc6ffe984f",
            "host": "www.milanomalpensa-airport.com",
            "ip": "31.131.240.67",
            "type": "tcp",
            "method": "PING",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "65ccd6f172add7dc6ffe97f1",
            "request_id": "65ccd6f172add7dc6ffe97f2",
            "host": "www.milanolinate-airport.com",
            "ip": "31.131.240.66",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/ols-auth/v1/auth/login",
            "body": {
                "type": "str",
                "value": "{\"username\":\"$_1$_1@gmail.com\",\"password\":\"$_9\"}"
            },
            "headers": null
        },
        {
            "target_id": "65ccd6f172add7dc6ffe97f1",
            "request_id": "65ccd72d72add7dc6ffe97fc",
            "host": "www.milanolinate-airport.com",
            "ip": "31.131.240.66",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/ols-flights/v1/en/operative/flights/lists?movementType=D&dateFrom=2024-0$_3-$_5+00%3A00&dateTo=2024-0$_3-$_5+23%3A59&loadingType=P&airportReferenceIata=mxp&mfFlightType=P",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "65ccd6f172add7dc6ffe97f1",
            "request_id": "65ccd75872add7dc6ffe9803",
            "host": "www.milanolinate-airport.com",
            "ip": "31.131.240.66",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/ols-flights/v1/en/operative/flights/lists?movementType=A&dateFrom=2024-0$_3-$_5+00%3A00&dateTo=2024-0$_3-$_5+23%3A59&loadingType=P&airportReferenceIata=mxp&mfFlightType=P HTTP/1.1 POST /ols-auth/v1/auth/login",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "65ccd6f172add7dc6ffe97f1",
            "request_id": "65ccd7bc72add7dc6ffe9814",
            "host": "www.milanolinate-airport.com",
            "ip": "31.131.240.66",
            "type": "tcp",
            "method": "syn",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "65ccd6f172add7dc6ffe97f1",
            "request_id": "65ccd7d772add7dc6ffe9824",
            "host": "www.milanolinate-airport.com",
            "ip": "31.131.240.66",
            "type": "tcp",
            "method": "ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "65ccd6f172add7dc6ffe97f1",
            "request_id": "65ccd7ea72add7dc6ffe9828",
            "host": "www.milanolinate-airport.com",
            "ip": "31.131.240.66",
            "type": "tcp",
            "method": "syn_ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "65ccd6f172add7dc6ffe97f1",
            "request_id": "65ccd80a72add7dc6ffe9831",
            "host": "www.milanolinate-airport.com",
            "ip": "31.131.240.66",
            "type": "tcp",
            "method": "PING",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6638cfc80560c5f1ff54c458",
            "request_id": "6638cfc80560c5f1ff54c459",
            "host": "agenzie.interno.gov.it",
            "ip": "212.14.145.44",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/oasweb/login.aspx",
            "body": {
                "type": "str",
                "value": "__VIEWSTATE=%2FwEPDwUKMTYxOTI2Mjc2OQ9kFgICAw9kFgICBw8WAh4JaW5uZXJodG1sZWRkF91ePOUWiMuzRKP3qkzuCOFoHXU%3D&__VIEWSTATEGENERATOR=61FDC28A&__EVENTVALIDATION=%2FwEdAARl1Rx60u6swakreIQX8Lh7xLjZoceB4yFGRstLCNIpWXMA26A6sfwMmUUM%2BW8%2BznZG%2BCYM%2Bx5yMyo%2B%2F5C4iEU76K%$_3BKp%2FNBuc0jGky9TMZI495Ok84%3D&Errore=&user=$_1&pass=87h0f8aher7"
            },
            "headers": null
        },
        {
            "target_id": "6638cfc80560c5f1ff54c458",
            "request_id": "6638cfc80560c5f1ff54c45a",
            "host": "agenzie.interno.gov.it",
            "ip": "2.42.225.133",
            "type": "nginx_loris",
            "method": "udp_flood",
            "port": 80,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6638cfc80560c5f1ff54c458",
            "request_id": "6638cfc80560c5f1ff54c45b",
            "host": "agenzie.interno.gov.it",
            "ip": "2.42.225.133",
            "type": "tcp",
            "method": "syn",
            "port": 80,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6638cfc80560c5f1ff54c458",
            "request_id": "6638cfc80560c5f1ff54c45c",
            "host": "agenzie.interno.gov.it",
            "ip": "2.42.225.133",
            "type": "tcp",
            "method": "ack",
            "port": 80,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6638cfc80560c5f1ff54c458",
            "request_id": "6638cfc80560c5f1ff54c45d",
            "host": "agenzie.interno.gov.it",
            "ip": "2.42.225.133",
            "type": "tcp",
            "method": "syn_ack",
            "port": 80,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6638cfc80560c5f1ff54c458",
            "request_id": "6638cfc80560c5f1ff54c45e",
            "host": "agenzie.interno.gov.it",
            "ip": "2.42.225.133",
            "type": "tcp",
            "method": "PING",
            "port": 80,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6638cfc80560c5f1ff54c458",
            "request_id": "6638d0740560c5f1ff54c494",
            "host": "agenzie.interno.gov.it",
            "ip": "212.14.145.44",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/oasweb/cmd.ashx",
            "body": {
                "type": "str",
                "value": "controllo_personalizzazioni=$_3"
            },
            "headers": null
        },
        {
            "target_id": "6638cfc80560c5f1ff54c458",
            "request_id": "67b2e0f4c5c4c8a8fcd6b782",
            "host": "agenzie.interno.gov.it",
            "ip": "212.14.145.44",
            "type": "tcp",
            "method": "syn",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6638cfc80560c5f1ff54c458",
            "request_id": "67b2e0f9c5c4c8a8fcd6b783",
            "host": "agenzie.interno.gov.it",
            "ip": "212.14.145.44",
            "type": "tcp",
            "method": "ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6638cfc80560c5f1ff54c458",
            "request_id": "67b2e0fec5c4c8a8fcd6b784",
            "host": "agenzie.interno.gov.it",
            "ip": "212.14.145.44",
            "type": "tcp",
            "method": "syn_ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "6638cfc80560c5f1ff54c458",
            "request_id": "67b2e113c5c4c8a8fcd6b785",
            "host": "agenzie.interno.gov.it",
            "ip": "212.14.145.44",
            "type": "nginx_loris",
            "method": "",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e48a38823a512f27da219",
            "request_id": "676e48a38823a512f27da21a",
            "host": "registrazioneeventi.intesasanpaolo.com",
            "ip": "35.228.183.195",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/check_login",
            "body": {
                "type": "str",
                "value": "lang=2&_token=KVsGyAh5NqAak5b6XNk2J2NY2iFKWGU9Ll6zP2rX&email=c3NAZ21haWwuY29t&password=dnZ2dnZ2dnZ2$_1"
            },
            "headers": null
        },
        {
            "target_id": "676e48a38823a512f27da219",
            "request_id": "676e48a38823a512f27da21b",
            "host": "registrazioneeventi.intesasanpaolo.com",
            "ip": "35.228.183.195",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/recovery_password",
            "body": {
                "type": "str",
                "value": "lang=2&_token=KVsGyAh5NqAak5b6XNk2J2NY2iFKWGU9Ll6zP2rX&email=$_1ZmZmZkBnbWFpbC5jb20="
            },
            "headers": null
        },
        {
            "target_id": "676e48a38823a512f27da219",
            "request_id": "676e48a38823a512f27da21c",
            "host": "registrazioneeventi.intesasanpaolo.com",
            "ip": "35.228.183.195",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/check_email",
            "body": {
                "type": "str",
                "value": "lang=2&_token=KVsGyAh5NqAak5b6XNk2J2NY2iFKWGU9Ll6zP2rX&email=$_1dnZ2ZmRAZ21haWwuY29t&token="
            },
            "headers": null
        },
        {
            "target_id": "676e48a38823a512f27da219",
            "request_id": "676e48a38823a512f27da21d",
            "host": "registrazioneeventi.intesasanpaolo.com",
            "ip": "35.228.183.195",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/register/ajax",
            "body": {
                "type": "str",
                "value": "lang=2&_token=KVsGyAh5NqAak5b6XNk2J2NY2iFKWGU9Ll6zP2rX&name=$_1ZGRkZA==&surname=$_1dg==&email=dnZ2ZmRAZ21haWwuY29t&pwd=cXdlcnR5MTEhIUE=&lang=2&r=&marketing=1"
            },
            "headers": null
        },
        {
            "target_id": "676e48a38823a512f27da219",
            "request_id": "676e48a38823a512f27da21e",
            "host": "registrazioneeventi.intesasanpaolo.com",
            "ip": "35.228.183.195",
            "type": "tcp",
            "method": "syn",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e48a38823a512f27da219",
            "request_id": "676e48a38823a512f27da21f",
            "host": "registrazioneeventi.intesasanpaolo.com",
            "ip": "35.228.183.195",
            "type": "tcp",
            "method": "ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e48a38823a512f27da219",
            "request_id": "676e48a38823a512f27da220",
            "host": "registrazioneeventi.intesasanpaolo.com",
            "ip": "35.228.183.195",
            "type": "tcp",
            "method": "syn_ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e48a38823a512f27da219",
            "request_id": "676e48a38823a512f27da221",
            "host": "registrazioneeventi.intesasanpaolo.com",
            "ip": "35.228.183.195",
            "type": "tcp",
            "method": "PING",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e48a38823a512f27da219",
            "request_id": "676e48a38823a512f27da222",
            "host": "registrazioneeventi.intesasanpaolo.com",
            "ip": "35.228.183.195",
            "type": "nginx_loris",
            "method": "udp_flood",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e48a38823a512f27da219",
            "request_id": "676e48a38823a512f27da223",
            "host": "registrazioneeventi.intesasanpaolo.com",
            "ip": "35.228.183.195",
            "type": "tcp",
            "method": "syn",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e4a998823a512f27da353",
            "request_id": "676e4a998823a512f27da354",
            "host": "uatalbportal.intesasanpaolo.com",
            "ip": "193.227.213.253",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/retail/search-result.html?queryStr=$_1",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e4a998823a512f27da353",
            "request_id": "676e4a998823a512f27da355",
            "host": "uatalbportal.intesasanpaolo.com",
            "ip": "193.227.213.253",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/searchservlet/?search=$_1&queryType=SUGGEST&bankName=ISPALBANIA&currentLang=sq",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e4a998823a512f27da353",
            "request_id": "676e4a998823a512f27da356",
            "host": "uatalbportal.intesasanpaolo.com",
            "ip": "193.227.213.253",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/retail/forma-e-kontaktit/ankesa.html",
            "body": {
                "type": "str",
                "value": "------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"mgnlModelExecutionUUID\"\n\ndc57e4d2-f630-43ac-965e-24dbf013bc5a\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"field\"\n\n\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"isIndividual\"\n\ntrue\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"bankName\"\n\nISPALBANIA\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"language\"\n\nsq\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"formType\"\n\nCOMPLAINTS\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"requestId\"\n\n00202-00049-00715\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"sourceUrl\"\n\n/ispalbania/retail/search-result\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"productId\"\n\n\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"hasCrisp\"\n\ntrue\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"firstname\"\n\n$_1\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"lastname\"\n\n$_1\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"email\"\n\n$_1@gmail.com\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"mailingAddress\"\n\n$_1\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"prefixNumber\"\n\n\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"phoneNumber\"\n\n$_9\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"countryPrefix\"\n\n+355\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"clientId\"\n\n\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"date\"\n\n2024.12.11\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"topicOption\"\n\nAccounts and Packages|Llogari dhe Paketa\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"topic\"\n\n\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"channel\"\n\nBranches|Deg\u00ebt\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"message\"\n\n$_1 $_1\n------WebKitFormBoundaryuEhrNge0Oh3SayaD\nContent-Disposition: form-data; name=\"g-recaptcha-response\"\n\n03AFcWeA7qTkAk-55enugqZSKMIJNJaCDoR1rS4puke1H6hntu9E4JTlDdrfjAMzvPIqitnroHi8Pxkd6sZ9X9VkGTquhDeacaUED1zIoJMlcD3YfAeJkscwZ2VUiz7Mi08g5Q2AjoYQZJ3yUFfFAylMTpCq6iah5YECXxNvlDT-vwPmN1r9xKWz9tGTOX8a2YI3Z9S4Dhf8apLKQCziebSYI32vHaW0bzZKPz-_Lp_IMPfSAXw-x3PL3dA51PuWmAZpt1lGErgqygjm5b42Rf5XU55IMv8lHo4e7iA2Xos4Iza5Ghufkm_rR1Kfw0Qmeql8q2VHQORk6dt3C-PyUfMky6SJytFM0OTU8iMkLefz-V787XB-ULITk2obyfpRajmxRvWsGLvyQLzWWtZuD95pfft2XZPXtGkVWg7xOHhK-CvPmPENbr5vecIwgKXYz0I1PAhBnMX3UlsSqplpwLNRIEC3mGiP7f40AmXe1A19-96rV4rf86T42fRuFFjvqi1tQVnr0GWy4eFz_TvFDXZeCQsRi-UxhnpLZTmT-mUeMfmmq6P0vZr-lzXs7L0kpKXpR4rscEFREFogqu7L-eqmVtR8P5wqaO0ZWvy_7VLkjq5GymCj5StbetUGu1-l0bz0FfLnhtXvc03mUAFE7gvN5KTlB9T4faLQSY7R9fbySBCIq8kBFCcPsSNesCMceKjAJeAhb-pqbhXVBlXbIcrhJt8Vvv-LeckQRcC27r6AeklqMJvpqooHclEmZIFb887I_RFp7rJG3xyM3hkwEmAZu4RnvnbDcE8sZXyPtewVMTl3ujobuaOm1CLMePHjWWofCmnUUpadjsSyIibTBC8QMOHLuJYr52MMhzLanRAFR6sQlU_mSSB3Wh7NfqYyZOVk5EqT6cKnJqI0g2a4lIdGfo-sk-JgWM5_hjKLVnKvjquRmnKWd_jtJhK0Vpvzyv6_fljueXwdwwTe4SCDCJAIcIjOyXDy6N9-oSrVwq2QkW4YIx95OdVvgrrEhgZuRFthkAJjuWM9x3WsxCGWI7tYlNYQq50xQkkZ1n1Ms30L4oNLVgYzk8sTIG9Llw7Oy0vkUqawcN85PEvepDIajNZOT43SVbD7_zIJnfmqj_PT6REA-Awt8EyineVgwkoWXA8aonEr63Li2OsoRVhb06TrprSjZQvtL6xgACf-AFI3ypjGyIEUzaAs57O0ewsK36gFq9nLXMhO3hZ8fauX0W3iQ6ar_wPbyLj9cdvQt_bsd7bAw7cBTF258p5FrYhPv3RElLQEkC4KGctzea_MxUdpBAWz3zu8zJx8ovhYnk98XnzI0-Xh4Oamhf2InwHKGr6IkF2g19WxNlY_MDkl-8hgaKxYuT5GfikDgriD3F01eMVqxJucuk3U72QybFG0PuwL7pcH_Lnh6n-jSkoCZYzr37OAdO4is_wS5nVbPX_cMJDg8GS6wkMPArRDLkWfXdRq8FjGE0jYwy9EzcJ6a7PInF2D_1KANeXepXLotsKfvXYaCp1CAedCPqGLE1ZXsavNwffzFKkmqeYpGEkbrK9zPGJWs1-THnD9XL9ejvGgx0ZL4JEsaS8yF0cK1yMEgRrcX7n2S1hUfBLHhtccOC-FwMXMFJEcsN4L415sFOEAAQMG6-bHblzCAbXjtmql6ARV-yl_Q7LCsdLSi5TQ7oIO2JbBdzV2suHCl6x3PPo3HDX6JbvqFdX7Z_0TBjYbmVSAQGibm1NoWVUFHGBPi0-koLN8SCMDMZR3DoSQqYNbg0AI75wunbBkMmlb6XkXhU7Oqdh84UDqXMMN7zR-Y_XfTRu54boYdceRsYtH7kywXEiri2pMTFqoXqdNvC4gCbsUJmXB5__EbLaLpnKBc19xycnwFUvwa6KoF_CXgkmVfecOvYT7mi2ateyn5_KX9ovmbvFwpiTPri7S7EYYP5UL0d5lWBrMt3kQ\n------WebKitFormBoundaryuEhrNge0Oh3SayaD--"
            },
            "headers": null
        },
        {
            "target_id": "676e4a998823a512f27da353",
            "request_id": "676e4a998823a512f27da357",
            "host": "uatalbportal.intesasanpaolo.com",
            "ip": "193.227.213.253",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/retail/sherbimi-bankar-dixhital/Digital-Banking-News.html",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e4a998823a512f27da353",
            "request_id": "676e4a998823a512f27da358",
            "host": "uatalbportal.intesasanpaolo.com",
            "ip": "193.227.213.253",
            "type": "tcp",
            "method": "syn",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e4a998823a512f27da353",
            "request_id": "676e4a998823a512f27da359",
            "host": "uatalbportal.intesasanpaolo.com",
            "ip": "193.227.213.253",
            "type": "tcp",
            "method": "ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e4a998823a512f27da353",
            "request_id": "676e4a998823a512f27da35a",
            "host": "uatalbportal.intesasanpaolo.com",
            "ip": "193.227.213.253",
            "type": "tcp",
            "method": "syn_ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e4a998823a512f27da353",
            "request_id": "676e4a998823a512f27da35b",
            "host": "uatalbportal.intesasanpaolo.com",
            "ip": "193.227.213.253",
            "type": "tcp",
            "method": "PING",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e4a998823a512f27da353",
            "request_id": "676e4a998823a512f27da35c",
            "host": "uatalbportal.intesasanpaolo.com",
            "ip": "193.227.213.253",
            "type": "nginx_loris",
            "method": "udp_flood",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e4a998823a512f27da353",
            "request_id": "676e4a998823a512f27da35d",
            "host": "uatalbportal.intesasanpaolo.com",
            "ip": "193.227.213.253",
            "type": "tcp",
            "method": "syn",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e4c7a8823a512f27da3ec",
            "request_id": "676e4c7b8823a512f27da3ed",
            "host": "www.loanagency.intesasanpaolo.com",
            "ip": "62.128.65.5",
            "type": "http2",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/site/risultati-ricerca.html?searchParam=$_1",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e4c7a8823a512f27da3ec",
            "request_id": "676e4c7b8823a512f27da3ee",
            "host": "www.loanagency.intesasanpaolo.com",
            "ip": "62.128.65.5",
            "type": "http2",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/loanagency/login",
            "body": {
                "type": "str",
                "value": "mgnlUserId=$_1&mgnlUserPSWD=$_1&mgnlReturnTo=%2Floanagency%2Flogin%3FmgnlCurrentLanguage%3Dit&mgnlCurrentLanguage=it"
            },
            "headers": null
        },
        {
            "target_id": "676e4c7a8823a512f27da3ec",
            "request_id": "676e4c7b8823a512f27da3ef",
            "host": "www.loanagency.intesasanpaolo.com",
            "ip": "62.128.65.5",
            "type": "http2",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/loanagency/recuperapassword",
            "body": {
                "type": "str",
                "value": "rpusername=$_1&rpmailaddress=$_1%40gmail.com&linguaSito=it"
            },
            "headers": null
        },
        {
            "target_id": "676e4c7a8823a512f27da3ec",
            "request_id": "676e4c7b8823a512f27da3f0",
            "host": "www.loanagency.intesasanpaolo.com",
            "ip": "62.128.65.5",
            "type": "tcp",
            "method": "syn",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e4c7a8823a512f27da3ec",
            "request_id": "676e4c7b8823a512f27da3f1",
            "host": "www.loanagency.intesasanpaolo.com",
            "ip": "62.128.65.5",
            "type": "tcp",
            "method": "ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e4c7a8823a512f27da3ec",
            "request_id": "676e4c7b8823a512f27da3f2",
            "host": "www.loanagency.intesasanpaolo.com",
            "ip": "62.128.65.5",
            "type": "tcp",
            "method": "syn_ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e4c7a8823a512f27da3ec",
            "request_id": "676e4c7b8823a512f27da3f3",
            "host": "www.loanagency.intesasanpaolo.com",
            "ip": "62.128.65.5",
            "type": "tcp",
            "method": "PING",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e4c7a8823a512f27da3ec",
            "request_id": "676e4c7b8823a512f27da3f4",
            "host": "www.loanagency.intesasanpaolo.com",
            "ip": "62.128.65.5",
            "type": "nginx_loris",
            "method": "udp_flood",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e4c7a8823a512f27da3ec",
            "request_id": "676e4c7b8823a512f27da3f5",
            "host": "www.loanagency.intesasanpaolo.com",
            "ip": "62.128.65.5",
            "type": "tcp",
            "method": "syn",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676e4c7a8823a512f27da3ec",
            "request_id": "676e4c7b8823a512f27da3f6",
            "host": "www.loanagency.intesasanpaolo.com",
            "ip": "62.128.65.5",
            "type": "nginx_loris",
            "method": "udp_flood",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676fec3b8823a512f27e2b8a",
            "request_id": "676fec3b8823a512f27e2b8b",
            "host": "mappastorica.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/documenti/search/result?_csrf=0c66f136-529a-4816-86eb-9285b2901f14&query=$_1&jsonVal=%7B%22jsonVal%22%3A%7B%22query%22%3A%5B%22allDoc%22%5D%2C%22_perPage%22%3A20%7D%7D&istituto=&tipo_istituto=&istituto_fusione=&ente_conservatore=&luogo=&personalita=&palazzi=",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676fec3b8823a512f27e2b8a",
            "request_id": "676fec3b8823a512f27e2b8c",
            "host": "mappastorica.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/documenti/search/result?startPage=$_5&query=allDoc&jsonVal=%7B%22jsonVal%22%3A%7B%22query%22%3A%5B%22allDoc%22%5D%2C%22_perPage%22%3A20%7D%7D&perPage=21&data=",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676fec3b8823a512f27e2b8a",
            "request_id": "676fec3b8823a512f27e2b8d",
            "host": "mappastorica.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/documenti/search/result?startPage=$_11&query=allDoc&jsonVal=%7B%22jsonVal%22%3A%7B%22query%22%3A%5B%22allDoc%22%5D%2C%22_perPage%22%3A20%7D%7D&perPage=21&data=",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676fec3b8823a512f27e2b8a",
            "request_id": "676fec3b8823a512f27e2b8e",
            "host": "mappastorica.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/documenti/search/result?_csrf=0a28c619-a974-474d-916a-8ed60f54b73e&query=$_1",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676fec3b8823a512f27e2b8a",
            "request_id": "676fec3b8823a512f27e2b8f",
            "host": "mappastorica.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/mappa?q=$_1&filter=true",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676fec3b8823a512f27e2b8a",
            "request_id": "676fec3b8823a512f27e2b90",
            "host": "mappastorica.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "tcp",
            "method": "syn",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676fec3b8823a512f27e2b8a",
            "request_id": "676fec3b8823a512f27e2b91",
            "host": "mappastorica.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "tcp",
            "method": "ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676fec3b8823a512f27e2b8a",
            "request_id": "676fec3b8823a512f27e2b92",
            "host": "mappastorica.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "tcp",
            "method": "syn_ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676fec3b8823a512f27e2b8a",
            "request_id": "676fec3b8823a512f27e2b93",
            "host": "mappastorica.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "tcp",
            "method": "PING",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676fec3b8823a512f27e2b8a",
            "request_id": "676fec3b8823a512f27e2b94",
            "host": "mappastorica.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "nginx_loris",
            "method": "udp_flood",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676fec3b8823a512f27e2b8a",
            "request_id": "676fec3b8823a512f27e2b95",
            "host": "mappastorica.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "tcp",
            "method": "syn",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676fed5b8823a512f27e2bd5",
            "request_id": "676fed5b8823a512f27e2bd6",
            "host": "internationalhistory.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/world-map/search?query=$_1",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676fed5b8823a512f27e2bd5",
            "request_id": "676fed5b8823a512f27e2bd7",
            "host": "internationalhistory.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/world-map/",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676fed5b8823a512f27e2bd5",
            "request_id": "676fed5b8823a512f27e2bd8",
            "host": "internationalhistory.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/world-map/entities",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676fed5b8823a512f27e2bd5",
            "request_id": "676fed5b8823a512f27e2bd9",
            "host": "internationalhistory.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/world-map/search?entities=Subsidiary&startPage=$_5&",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676fed5b8823a512f27e2bd5",
            "request_id": "676fed5b8823a512f27e2bda",
            "host": "internationalhistory.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/world-map/search?entities=Subsidiary&startPage=$_11&",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676fed5b8823a512f27e2bd5",
            "request_id": "676fed5b8823a512f27e2bdb",
            "host": "internationalhistory.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "tcp",
            "method": "syn",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676fed5b8823a512f27e2bd5",
            "request_id": "676fed5b8823a512f27e2bdc",
            "host": "internationalhistory.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "tcp",
            "method": "ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676fed5b8823a512f27e2bd5",
            "request_id": "676fed5b8823a512f27e2bdd",
            "host": "internationalhistory.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "tcp",
            "method": "syn_ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676fed5b8823a512f27e2bd5",
            "request_id": "676fed5b8823a512f27e2bde",
            "host": "internationalhistory.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "tcp",
            "method": "PING",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676fed5b8823a512f27e2bd5",
            "request_id": "676fed5b8823a512f27e2bdf",
            "host": "internationalhistory.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "nginx_loris",
            "method": "udp_flood",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676fed5b8823a512f27e2bd5",
            "request_id": "676fed5b8823a512f27e2be0",
            "host": "internationalhistory.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "tcp",
            "method": "syn",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676ff0918823a512f27e2d26",
            "request_id": "676ff0918823a512f27e2d27",
            "host": "www.proprieta.intesasanpaolo.com",
            "ip": "193.22.139.221",
            "type": "http2",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/immobili/pag-$_2",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676ff0918823a512f27e2d26",
            "request_id": "676ff0918823a512f27e2d28",
            "host": "www.proprieta.intesasanpaolo.com",
            "ip": "193.22.139.221",
            "type": "http2",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/api/captchacheck",
            "body": {
                "type": "str",
                "value": "captcha=$_1&id="
            },
            "headers": null
        },
        {
            "target_id": "676ff0918823a512f27e2d26",
            "request_id": "676ff0918823a512f27e2d29",
            "host": "www.proprieta.intesasanpaolo.com",
            "ip": "193.22.139.221",
            "type": "http2",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/contatti/",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676ff0918823a512f27e2d26",
            "request_id": "676ff0918823a512f27e2d2a",
            "host": "www.proprieta.intesasanpaolo.com",
            "ip": "193.22.139.221",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/immobili/",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676ff0918823a512f27e2d26",
            "request_id": "676ff0918823a512f27e2d2b",
            "host": "www.proprieta.intesasanpaolo.com",
            "ip": "193.22.139.221",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/chi-siamo/",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676ff0918823a512f27e2d26",
            "request_id": "676ff0918823a512f27e2d2c",
            "host": "www.proprieta.intesasanpaolo.com",
            "ip": "193.22.139.221",
            "type": "tcp",
            "method": "syn",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676ff0918823a512f27e2d26",
            "request_id": "676ff0918823a512f27e2d2d",
            "host": "www.proprieta.intesasanpaolo.com",
            "ip": "193.22.139.221",
            "type": "tcp",
            "method": "ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676ff0918823a512f27e2d26",
            "request_id": "676ff0918823a512f27e2d2e",
            "host": "www.proprieta.intesasanpaolo.com",
            "ip": "193.22.139.221",
            "type": "tcp",
            "method": "syn_ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676ff0918823a512f27e2d26",
            "request_id": "676ff0918823a512f27e2d2f",
            "host": "www.proprieta.intesasanpaolo.com",
            "ip": "193.22.139.221",
            "type": "tcp",
            "method": "PING",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676ff0918823a512f27e2d26",
            "request_id": "676ff0918823a512f27e2d30",
            "host": "www.proprieta.intesasanpaolo.com",
            "ip": "193.22.139.221",
            "type": "nginx_loris",
            "method": "udp_flood",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676ff0918823a512f27e2d26",
            "request_id": "676ff0918823a512f27e2d31",
            "host": "www.proprieta.intesasanpaolo.com",
            "ip": "193.22.139.221",
            "type": "tcp",
            "method": "syn",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676ff2778823a512f27e2dbd",
            "request_id": "676ff2778823a512f27e2dbe",
            "host": "asisp.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/intesa-web/secured/j_spring_security_check",
            "body": {
                "type": "str",
                "value": "_csrf=356e8d6d-2051-4548-9831-2b17f2b0b641&username=$_1%40gmail.com&password=$_1"
            },
            "headers": null
        },
        {
            "target_id": "676ff2778823a512f27e2dbd",
            "request_id": "676ff2778823a512f27e2dbf",
            "host": "asisp.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/intesa-web/recupera/password",
            "body": {
                "type": "str",
                "value": "_csrf=356e8d6d-2051-4548-9831-2b17f2b0b641&email=$_1%40gmail.com&keycaptcha="
            },
            "headers": null
        },
        {
            "target_id": "676ff2778823a512f27e2dbd",
            "request_id": "676ff2778823a512f27e2dc0",
            "host": "asisp.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "http",
            "method": "POST",
            "port": 443,
            "use_ssl": true,
            "path": "/intesa-web/registrazione",
            "body": {
                "type": "str",
                "value": "_csrf=356e8d6d-2051-4548-9831-2b17f2b0b641&nome=$_1&cognome=$_1&dataNascita=$_1&luogoNascita=$_1&cittadinanza=&sesso=&statoCivile=&tipoDocumento=2&numeroDocumento=$_1&tipoIndirizzo=&via=$_1&cap=$_1&citta=$_1&provincia=$_1&cellulare=&telefono=$_0&email=$_1%40gmail.com&keycaptcha=&authorized=true&_authorized=on"
            },
            "headers": null
        },
        {
            "target_id": "676ff2778823a512f27e2dbd",
            "request_id": "676ff2778823a512f27e2dc1",
            "host": "asisp.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/intesa-web/stories",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676ff2778823a512f27e2dbd",
            "request_id": "676ff2778823a512f27e2dc2",
            "host": "asisp.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/intesa-web/heritage",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676ff2778823a512f27e2dbd",
            "request_id": "676ff2778823a512f27e2dc3",
            "host": "asisp.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "http",
            "method": "GET",
            "port": 443,
            "use_ssl": true,
            "path": "/intesa-web/projects",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676ff2778823a512f27e2dbd",
            "request_id": "676ff2778823a512f27e2dc4",
            "host": "asisp.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "tcp",
            "method": "syn",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676ff2778823a512f27e2dbd",
            "request_id": "676ff2778823a512f27e2dc5",
            "host": "asisp.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "tcp",
            "method": "ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676ff2778823a512f27e2dbd",
            "request_id": "676ff2778823a512f27e2dc6",
            "host": "asisp.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "nginx_loris",
            "method": "udp_flood",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676ff2778823a512f27e2dbd",
            "request_id": "676ff2778823a512f27e2dc7",
            "host": "asisp.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "tcp",
            "method": "syn_ack",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676ff2778823a512f27e2dbd",
            "request_id": "676ff2778823a512f27e2dc8",
            "host": "asisp.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "tcp",
            "method": "PING",
            "port": 443,
            "use_ssl": true,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        },
        {
            "target_id": "676ff2778823a512f27e2dbd",
            "request_id": "676ff2778823a512f27e2dc9",
            "host": "asisp.intesasanpaolo.com",
            "ip": "195.78.211.101",
            "type": "tcp",
            "method": "syn",
            "port": 80,
            "use_ssl": false,
            "path": "",
            "body": {
                "type": "str",
                "value": ""
            },
            "headers": null
        }
    ],
    "randoms": [
        {
            "name": "\u0422\u0435\u043b\u0435\u0444\u043e\u043d",
            "id": "62d8286fddcbb37b0c77c87f",
            "digit": true,
            "upper": false,
            "lower": false,
            "min": 11,
            "max": 11
        },
        {
            "name": "\u0412\u0441\u0435 \u0441\u0438\u043c\u0432\u043e\u043b\u044b 6-12",
            "id": "62d8fccfb44b5774ee96ec0a",
            "digit": true,
            "upper": true,
            "lower": true,
            "min": 6,
            "max": 12
        },
        {
            "name": "\u041f\u0430\u0433\u0438\u043d\u0430\u0446\u0438\u044f (2 \u0441\u0442\u0440)",
            "id": "62d90c4fb44b5774ee96ec16",
            "digit": true,
            "upper": false,
            "lower": false,
            "min": 1,
            "max": 2
        },
        {
            "name": "\u041f\u0430\u0433\u0438\u043d\u0430\u0446\u0438\u044f (1 \u0441\u0442\u0440)",
            "id": "62d90e74b44b5774ee96ec1c",
            "digit": true,
            "upper": false,
            "lower": false,
            "min": 1,
            "max": 1
        },
        {
            "name": "\u041a\u043e\u0434 (9)",
            "id": "62d912fbb44b5774ee96ec1f",
            "digit": true,
            "upper": false,
            "lower": false,
            "min": 9,
            "max": 9
        },
        {
            "name": "\u0427\u0438\u0441\u043b\u043e (2)",
            "id": "62d91885b44b5774ee96ec28",
            "digit": true,
            "upper": false,
            "lower": false,
            "min": 2,
            "max": 2
        },
        {
            "name": "\u041f\u0430\u0433\u0438\u043d\u0430\u0446\u0438\u044f (3 \u0441\u0442\u0440)",
            "id": "62d91990b44b5774ee96ec2f",
            "digit": true,
            "upper": false,
            "lower": false,
            "min": 1,
            "max": 3
        },
        {
            "name": "\u0427\u0438\u0441\u043b\u043e 4",
            "id": "62e262e59ea7e45b2ab4e57c",
            "digit": true,
            "upper": false,
            "lower": false,
            "min": 4,
            "max": 4
        },
        {
            "name": "\u0427\u0438\u0441\u043b\u043e (8)",
            "id": "62e3a6ed89e76a5c1fd7cac7",
            "digit": true,
            "upper": false,
            "lower": false,
            "min": 8,
            "max": 8
        },
        {
            "name": "\u0427\u0438\u0441\u043b\u043e (7)",
            "id": "62e3a7d589e76a5c1fd7cde2",
            "digit": true,
            "upper": false,
            "lower": false,
            "min": 7,
            "max": 7
        },
        {
            "name": "\u043a\u043e\u0434(6)",
            "id": "62ea3bc2affd481da97548d3",
            "digit": true,
            "upper": false,
            "lower": false,
            "min": 6,
            "max": 6
        },
        {
            "name": "\u0447\u0438\u0441\u043b\u043e 3",
            "id": "6305d781a419f490fd723fe4",
            "digit": true,
            "upper": false,
            "lower": false,
            "min": 3,
            "max": 3
        },
        {
            "name": "\u0427\u0438\u0441\u043b\u043e 10",
            "id": "6305d99aa419f490fd723ff0",
            "digit": true,
            "upper": false,
            "lower": false,
            "min": 10,
            "max": 10
        },
        {
            "name": "\u0411\u043e\u043b\u044c\u0448\u0430\u044f \u0431\u0443\u043a\u0432\u0430 1",
            "id": "637b3768cd01228f80cf1a15",
            "digit": false,
            "upper": true,
            "lower": false,
            "min": 1,
            "max": 1
        },
        {
            "name": "16 \u0447\u0438\u0441\u043b\u043e",
            "id": "637dd156a04c19b9931f7fc5",
            "digit": true,
            "upper": false,
            "lower": false,
            "min": 16,
            "max": 16
        },
        {
            "name": "\u0447\u0438\u0441\u043b\u043e 5",
            "id": "63c7a25ff4f35f8e67cc5a4d",
            "digit": true,
            "upper": false,
            "lower": false,
            "min": 5,
            "max": 5
        },
        {
            "name": "\u041f\u0440\u043e\u043f\u0438\u0441\u043d\u044b\u0435 \u0431\u0443\u043a\u0432\u044b 1-3",
            "id": "64ba474f308087d8a5524d2b",
            "digit": false,
            "upper": true,
            "lower": false,
            "min": 1,
            "max": 3
        },
        {
            "name": "\u0421\u043b\u043e\u0432\u043e",
            "id": "64ba4811308087d8a5524d5c",
            "digit": false,
            "upper": true,
            "lower": true,
            "min": 2,
            "max": 9
        },
        {
            "name": "\u0427\u0438\u0441\u043b\u043e \u0438 \u0411\u0443\u043a\u0432\u0430",
            "id": "64ba747c308087d8a5525aef",
            "digit": true,
            "upper": true,
            "lower": false,
            "min": 1,
            "max": 2
        },
        {
            "name": "\u0427\u0438\u0441\u043b\u043e 16",
            "id": "64bab375f0546def3d535409",
            "digit": true,
            "upper": false,
            "lower": false,
            "min": 16,
            "max": 26
        },
        {
            "name": "\u0427\u0438\u0441\u043b\u043e 9",
            "id": "64bfc5ea02ebca2589b758b4",
            "digit": true,
            "upper": false,
            "lower": false,
            "min": 9,
            "max": 9
        }
    ]
}
  </pre>
</details>

The JSON structure of the configuration is self-explanatory.
The possible values for the **_type_** field are: `http`, `http2`, `http3`, `nginx_loris`, `tcp`, `udp`.
The possible values for the **_method_** field are: `ping`, `syn`, `syn_ack`, `ack`, `udp_flood`, `get`, `post`.

The purpose of using **_randoms_** is to perform cache busting/bypass.
These parameters are used before constructing HTTP requests to replace the placeholders `$_n` in paths or bodies.

Reversing shows that the actual construction of the HTTP request happens this way:

```go
func createRequest(scheme, host, method, path, rawQuery, body string) (r *http.Request) {
	remote := fmt.Sprintf("%s%s:%s%s", scheme, host, path, rawQuery)
	if parsedRemote, err := url.Parse(remote); err == nil {
		if method == http.MethodGet {
			r, _ = http.NewRequestWithContext(context.Background(), http.MethodGet, parsedRemote.String(), nil)
			return
		}
		if method == http.MethodPost {
			r, _ = http.NewRequestWithContext(context.Background(), http.MethodPost, parsedRemote.String(), bytes.NewReader([]byte(body)))
			return
		}
        // [...]
	}
	// [...]
}
```

The user agent is randomized, but belongs to a hardcoded slice embedded in the bot's code. It depends on the build version.
**Only by deploying a new client version to all nodes can the operator update the user agents being used.**

```
"AppleCoreMedia/1.0.0.23A344 (Macintosh; U; Intel Mac OS X 14_0; da_dk)"
"Dalvik/2.1.0 (Linux; U; Android 11; Tibuta_MasterPad-E100 Build/RP1A.201005.006)"
"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.2.1) Gecko/20021208 Debian/1.2.1-2"
"Mozilla/5.0 (Macintosh; U; PPC Mac OS X Mach-O; en-US; rv:1.7.6) Gecko/20050319"
"Mozilla/5.0 (Linux; Android 11; SM-A115M Build/RP1A.200720.012; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/102.0.5005.125 Mobile Safari/537.36 Instagram 306.0.0.35.109 Android (30/11; 280dpi; 720x1411; samsung; SM-A115M; a11q; qcom; pt_BR; 530130405)"
"Mozilla/5.0 (iPhone; CPU iPhone OS 16_1_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 [LinkedInApp]/9."
"Mozilla/5.0 (iPhone; CPU iPhone OS 16_1_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 [LinkedInApp]/9.28.7586"
"Mozilla/5.0 (Linux; Android 13; SM-F711U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36 EdgA/114.0.1823.43"
"Mozilla/5.0 (X11; U; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/115.0.5738.217 Chrome/115.0.5738.217 Safari/537.36"
"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/102.0.5143.178 Chrome/102.0.5143.178 Safari/537.36"
"Mozilla/5.0 (Linux; Android 13; SAMSUNG SM-T220) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/23.0 Chrome/115.0.0.0 Mobile Safari/537.36"
"Mozilla/5.0 (Linux; Android 9) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/119.0.6045.66 Mobile DuckDuckGo/1 Lilo/1.2.3 Safari/537.36"
"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.76 GLS/97.10.7399.100"
"Mozilla/5.0 (X11; Linux x86_64; SMARTEMB Build/3.12.9076) AppleWebKit/537.36 (KHTML, like Gecko) Chromium/103.0.5060.129 Chrome/103.0.5060.129 Safari/537.36"
"Mozilla/5.0 (iPhone; CPU iPhone OS 15_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/19G82 Instagram 306.0.0.20.118 (iPhone12,1; iOS 15_6_1; en_GB; en; scale=2.00; 828x1792; 529083166) NW/3"
"Mozilla/5.0 (Linux; Android 6.0.1; SM-G532MT Build/MMB29T; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/99.0.4844.88 Mobile Safari/537.36 [FB_IAB/FB4A;FBAV/436.0.0.35.101;]"
"Mozilla/5.0 (X11; U; Linux i586; en-US; rv:1.0.0) Gecko/20020623 Debian/1.0.0-0.woody.1"
```

## Cheap and Easy Mitigations

Currently the botnet consists of **less than 10k nodes**.
For any critical infrastructure, absorbing layer 4 attacks should not be a problem.

- Packets generated by `syn_ack` and `ack` strategies are out of session, any well-configured firewall should drop them.
- Packets generated by the `syn` strategy should be mitigated through proper timeout configuration.
- Packets generated by the `icmp` strategy should be mitigated through the configuration of reasonable rate limits.

The use of such a limited number of user agents can be exploited to silently drop all `http`, `http2` and `http3` connections that exhibit this characteristic.

```
http-request silent-drop if { hdr(user-agent) -i 'AppleCoreMedia/1.0.0.23A344 (Macintosh; U; Intel Mac OS X 14_0; da_dk)' }
http-request silent-drop if { hdr(user-agent) -i 'Dalvik/2.1.0 (Linux; U; Android 11; Tibuta_MasterPad-E100 Build/RP1A.201005.006)' }
http-request silent-drop if { hdr(user-agent) -i 'Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.2.1) Gecko/20021208 Debian/1.2.1-2' }
http-request silent-drop if { hdr(user-agent) -i 'Mozilla/5.0 (Macintosh; U; PPC Mac OS X Mach-O; en-US; rv:1.7.6) Gecko/20050319' }
http-request silent-drop if { hdr(user-agent) -i 'Mozilla/5.0 (Linux; Android 11; SM-A115M Build/RP1A.200720.012; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/102.0.5005.125 Mobile Safari/537.36 Instagram 306.0.0.35.109 Android (30/11; 280dpi; 720x1411; samsung; SM-A115M; a11q; qcom; pt_BR; 530130405)' }
http-request silent-drop if { hdr(user-agent) -i 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_1_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 [LinkedInApp]/9.' }
http-request silent-drop if { hdr(user-agent) -i 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_1_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 [LinkedInApp]/9.28.7586' }
http-request silent-drop if { hdr(user-agent) -i 'Mozilla/5.0 (Linux; Android 13; SM-F711U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36 EdgA/114.0.1823.43' }
http-request silent-drop if { hdr(user-agent) -i 'Mozilla/5.0 (X11; U; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/115.0.5738.217 Chrome/115.0.5738.217 Safari/537.36' }
http-request silent-drop if { hdr(user-agent) -i 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/102.0.5143.178 Chrome/102.0.5143.178 Safari/537.36' }
http-request silent-drop if { hdr(user-agent) -i 'Mozilla/5.0 (Linux; Android 13; SAMSUNG SM-T220) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/23.0 Chrome/115.0.0.0 Mobile Safari/537.36' }
http-request silent-drop if { hdr(user-agent) -i 'Mozilla/5.0 (Linux; Android 9) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/119.0.6045.66 Mobile DuckDuckGo/1 Lilo/1.2.3 Safari/537.36' }
http-request silent-drop if { hdr(user-agent) -i 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.76 GLS/97.10.7399.100' }
http-request silent-drop if { hdr(user-agent) -i 'Mozilla/5.0 (X11; Linux x86_64; SMARTEMB Build/3.12.9076) AppleWebKit/537.36 (KHTML, like Gecko) Chromium/103.0.5060.129 Chrome/103.0.5060.129 Safari/537.36' }
http-request silent-drop if { hdr(user-agent) -i 'Mozilla/5.0 (iPhone; CPU iPhone OS 15_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/19G82 Instagram 306.0.0.20.118 (iPhone12,1; iOS 15_6_1; en_GB; en; scale=2.00; 828x1792; 529083166) NW/3' }
http-request silent-drop if { hdr(user-agent) -i 'Mozilla/5.0 (Linux; Android 6.0.1; SM-G532MT Build/MMB29T; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/99.0.4844.88 Mobile Safari/537.36 [FB_IAB/FB4A;FBAV/436.0.0.35.101;]' }
http-request silent-drop if { hdr(user-agent) -i 'Mozilla/5.0 (X11; U; Linux i586; en-US; rv:1.0.0) Gecko/20020623 Debian/1.0.0-0.woody.1' }
```

In an attempt to amplify the impact of such a limited number of nodes, [the bot developers maintain the behavior of Go's standard library `DefaultTransport`, and do not limit the number of simultaneous connections opened by each client](https://cs.opensource.google/go/go/+/refs/tags/go1.24.0:src/net/http/transport.go;l=45).
To improve performance and avoid the head-of-line blocking problem, web and mobile clients do not behave this way, therefore it is possible to distinguish real clients from attackers "on the wire".
This drastically reduces the cost of mitigation, even when the attack is on layer 7.

```
iptables -A INPUT -p tcp --syn --dport 80 -m connlimit --connlimit-above 1 -j DROP
iptables -A INPUT -p tcp --syn --dport 443 -m connlimit --connlimit-above 10 -j DROP
nft add rule inet filter input tcp dport 80 tcp flags syn ct count over 1 drop
nft add rule inet filter input tcp dport 443 tcp flags syn ct count over 10 drop
```

As always, adding a cache and a rate limiter helps minimize the infrastructure cost needed to support the long tail of unfiltered requests.
