---
title: Parsing and Pretty-Printing JSON Log Payloads with a Go Syslog Server
title_marker: JSON Log Payloads Syslog
description: Learn how to set up a lightweight Go Syslog server using go-syslog to receive, extract, and format JSON-encoded log messages over UDP and TCP
layout: default
lang: en
---

Syslog remains one of the most widely adopted standards for transmitting system and application logs across networks. However, modern application observability heavily relies on structured formats like JSON. Combining the two gives you the best of both worlds: standard transport infrastructure with the flexibility of structured payloads.

In this post, we will walk through building a lightweight Syslog server in Go using the popular `gopkg.in/mcuadros/go-syslog.v2` library to ingest Syslog messages and format JSON payloads on the fly.

Go's concurrency primitives (goroutines and channels) make it an excellent choice for network servers. Handling concurrent UDP and TCP connections while processing log streams asynchronously requires minimal code and delivers high throughput with low memory overhead.

Ultimately, the primary purpose of this minimal setup is to serve as a practical debugging tool for developers. When configuring services to ship logs remotely, it can be tricky to ensure the payloads are formatted exactly as expected. By spinning up this lightweight server locally, developers can intercept incoming traffic and immediately validate that the Syslog messages being sent contain perfectly valid, well-formed JSON before deploying their application to a production environment.

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"gopkg.in/mcuadros/go-syslog.v2"
)

func main() {
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	server.SetFormat(syslog.Automatic)
	server.SetHandler(handler)
	server.ListenUDP("0.0.0.0:514")
	server.ListenTCP("0.0.0.0:514")

	server.Boot()

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			fmt.Println(logParts)
			message, ok := logParts["message"].(string)
			if !ok {
				message, ok = logParts["content"].(string)
				if !ok {
					fmt.Println("Error retrieving the message from logParts")
					continue
				}
			}

			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, []byte(message), "", "  "); err != nil {
				fmt.Printf("JSON formatting error: %s\n", err.Error())
			} else {
				fmt.Printf("Received JSON log: %s\n", prettyJSON.String())
			}
		}
	}(channel)

	server.Wait()
}
```
