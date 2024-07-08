---
title: BPF HTTP Replay
description: Midpoint network interception for HTTP traffic mirroring
layout: default
lang: en
---

<iframe src="https://www.youtube-nocookie.com/embed/PbipefyfkNY?vq=hd1080&rel=0&color=white" width="100%" height="560" title="Fiber Tapping - Monitoring Fiber Optic Connections" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" frameborder="0" allowfullscreen></iframe>

## Intercept - Resend

In the security industry an intercept-resend or replay attack can be implemented either by intercepting the signal using a network tap or by monitoring leakage from a bent fiber.

The optical signal can then be sampled and analyzed. The knowledge acquired about the signal can then be used to craft a replica pulse that is injected into the fiber.
The basis for this vulnerability is the exact cloning of signals in classical mechanical theory.

A network tap is thus a hardware device or software tool that allows you to monitor and capture network traffic passing through a specific point in a network.
It provides a way to intercept and analyze network packets without disrupting the normal flow of traffic. Network taps are commonly used for network troubleshooting, security monitoring, and performance analysis.

Network taps can be passive, simply copying the traffic, or active, allowing for additional functionality such as filtering or modifying the packets.

## Midpoint Network Interception

Midpoint Network Interception refers to a technique or mechanism used in networking tools and software to capture and analyze network packets.
It involves intercepting network traffic at a specific point in the network, typically at a midpoint between the source and destination of the packets.

cBPF, or classic Berkeley Packet Filter, is a low-level packet filtering mechanism used in various networking tools and software.
It allows users to define rules for capturing network packets based on specific criteria, such as source or destination IP address, port number, protocol, and more.

As such, cBPF is a dual use technology that can be used to implement midpoint intercept-resend techniques.

## The Idea

Being a dual-use tool, BPF is not only useful for executing network attacks but also for copying real-world traffic in real time and forwarding it to another endpoint.

The idea is to obtain a quick reaction capability by coding a simple traffic shadowing tool that is easy to deploy and not bound to any specific technological stack.

```go
package main

import (
	"bufio"
	"bytes"
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket/tcpassembly/tcpreader"
)

var (
	iface         = flag.String("i", "eth0", "Interface to get packets from")
	fname         = flag.String("r", "", "Filename to read from, overrides -i")
	snaplen       = flag.Int("s", 1600, "SnapLen for pcap packet capture")
	filter        = flag.String("f", "tcp dst port 80", "BPF filter for pcap")
	filterUrl     = flag.String("u", "/", "Filter with a substring the URL to replay")
	replayHost    = flag.String("h", "[::1]:8080", "Host to replay the requests to")
	connections   = flag.Int("c", 1000, "Max number of connections to keep in the pool")
	logAllPackets = flag.Bool("v", false, "Logs every packet in great detail")
	directHttp    *http.Client
)

type httpStreamFactory struct{}

type httpStream struct {
	net, transport gopacket.Flow
	r              tcpreader.ReaderStream
}

func (h *httpStreamFactory) New(net, transport gopacket.Flow) tcpassembly.Stream {
	hstream := &httpStream{
		net:       net,
		transport: transport,
		r:         tcpreader.NewReaderStream(),
	}
	go hstream.run()

	return &hstream.r
}

func (h *httpStream) run() {
	buf := bufio.NewReader(&h.r)
	for {
		req, err := http.ReadRequest(buf)
		if err == io.EOF {
			return
		} else if err != nil {
			log.Println("Error reading stream", h.net, h.transport, ":", err)
		} else {
			if !strings.Contains(req.URL.Path, *filterUrl) {
				req.Body.Close()
				continue
			}

			bodyBytes, err := io.ReadAll(req.Body)
			if err != nil {
				log.Println(err.Error())
				req.Body.Close()
				continue
			}

			request := http.Request{
				Method: req.Method,
				URL: &url.URL{
					Scheme:     "http",
					Host:       *replayHost,
					Path:       req.URL.Path,
					RawPath:    req.URL.RawPath,
					ForceQuery: true,
					RawQuery:   req.URL.RawQuery,
				},
				Proto:  req.Proto,
				Header: req.Header,
				Body:   io.NopCloser(bytes.NewReader(bodyBytes)),
			}
			go httpReplay(request)
		}
	}
}

func httpReplay(request http.Request) {
	resp, err := directHttp.Do(&request)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}
}

func main() {
	flag.Parse()

	directHttp = &http.Client{Transport: &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   time.Second,
			KeepAlive: time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          *connections,
		MaxIdleConnsPerHost:   *connections,
		MaxConnsPerHost:       *connections,
		ForceAttemptHTTP2:     true,
		IdleConnTimeout:       time.Second,
		ExpectContinueTimeout: time.Second,
		DisableKeepAlives:     true,
	}}

	var (
		handle *pcap.Handle
		err    error
	)
	if *fname != "" {
		log.Printf("Reading from pcap dump %q", *fname)
		handle, err = pcap.OpenOffline(*fname)
	} else {
		log.Printf("Starting capture on interface %q", *iface)
		handle, err = pcap.OpenLive(*iface, int32(*snaplen), true, pcap.BlockForever)
	}
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := handle.SetBPFFilter(*filter); err != nil {
		log.Fatal(err.Error())
	}

	streamFactory := &httpStreamFactory{}
	streamPool := tcpassembly.NewStreamPool(streamFactory)
	assembler := tcpassembly.NewAssembler(streamPool)

	log.Println("reading in packets")
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packets := packetSource.Packets()
	ticker := time.Tick(time.Minute)
	for {
		select {
		case packet := <-packets:
			if packet == nil {
				return
			}
			if *logAllPackets {
				log.Println(packet)
			}
			if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
				log.Println("Unusable packet")
				continue
			}
			tcp := packet.TransportLayer().(*layers.TCP)
			assembler.AssembleWithTimestamp(packet.NetworkLayer().NetworkFlow(), tcp, packet.Metadata().Timestamp)

		case <-ticker:
			assembler.FlushOlderThan(time.Now().Add(time.Second * -30))
		}
	}
}
```

## References

[https://www.kernel.org/doc/html/latest/bpf/index.html](https://www.kernel.org/doc/html/latest/bpf/index.html)

[https://www.haproxy.com/blog/haproxy-traffic-mirroring-for-real-world-testing](https://www.haproxy.com/blog/haproxy-traffic-mirroring-for-real-world-testing)

[https://pkg.go.dev/github.com/google/gopacket](https://pkg.go.dev/github.com/google/gopacket)
