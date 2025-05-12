---
layout: default
title: snowflake
description: >
  Pluggable Transport using WebRTC, inspired by Flashproxy. <br>
  A custom fork with mine opinionated patches
go-import: tgragnato.it/snowflake git https://github.com/tgragnato/snowflake
go-source: tgragnato.it/snowflake https://github.com/tgragnato/snowflake https://github.com/tgragnato/snowflake/tree/main{/dir} https://github.com/tgragnato/snowflake/blob/main{/dir}/{file}#L{line}
prefetch:
  - goreportcard.com
  - codecov.io
  - raw.githubusercontent.com
---

Go package hosted at [tgragnato/snowflake](https://github.com/tgragnato/snowflake).

Documentation available at [pkg.go.dev](https://pkg.go.dev/tgragnato.it/snowflake).

[![Go](https://github.com/tgragnato/snowflake/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/tgragnato/snowflake/actions/workflows/go.yml)
[![CodeQL](https://github.com/tgragnato/snowflake/actions/workflows/codeql.yml/badge.svg?branch=main)](https://github.com/tgragnato/snowflake/actions/workflows/codeql.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/tgragnato/snowflake)](https://goreportcard.com/report/github.com/tgragnato/snowflake)
[![codecov](https://codecov.io/gh/tgragnato/snowflake/branch/main/graph/badge.svg)](https://codecov.io/gh/tgragnato/snowflake)

- golang 1.24+ & bumped dependencies
- custom transport for broker negotiation (TLS 1.3 with selected ciphersuites & groups, MultiPath TCP)
- custom DTLS fingerprint, different from any popular WebRTC implementation
- use the Setting Engine to reduce MulticastDNS noise
- use a context aware io.Reader that closes on errors in copyLoop
- extremely simple token handling
- client padding to evade TLS in DTLS detection
- introduction of a proxy option to force the NAT type as unrestricted
- coder/websocket in place of gorilla/websocket

![Schematic](https://raw.githubusercontent.com/tgragnato/snowflake/main/schematic.png)

We have more documentation in the [Snowflake wiki](https://gitlab.torproject.org/tpo/anti-censorship/pluggable-transports/snowflake/-/wikis/home) and at `https://snowflake.torproject.org/`.

## Structure of this Repository

- `broker/` contains code for the Snowflake broker
- `doc/` contains Snowflake documentation and manpages
- `client/` contains the Tor pluggable transport client and client library code
- `common/` contains generic libraries used by multiple pieces of Snowflake
- `proxy/` contains code for the Go standalone Snowflake proxy
- `probetest/` contains code for a NAT probetesting service
- `server/` contains the Tor pluggable transport server and server library code
