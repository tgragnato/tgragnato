---
layout: default
title: goflow
description: >
  The high-scalability sFlow/NetFlow/IPFIX collector used internally at Cloudflare [syslog transport]
go-import: tgragnato.it/goflow git https://github.com/tgragnato/goflow
go-source: tgragnato.it/goflow https://github.com/tgragnato/goflow https://github.com/tgragnato/goflow/tree/main{/dir} https://github.com/tgragnato/goflow/blob/main{/dir}/{file}#L{line}
---

Go package hosted at [tgragnato/goflow](https://github.com/tgragnato/goflow).

[![Go](https://github.com/tgragnato/goflow/actions/workflows/go.yml/badge.svg)](https://github.com/tgragnato/goflow/actions/workflows/go.yml)
[![CodeQL](https://github.com/tgragnato/goflow/actions/workflows/codeql.yml/badge.svg)](https://github.com/tgragnato/goflow/actions/workflows/codeql.yml)
[![Codecov](https://codecov.io/gh/tgragnato/goflow/graph/badge.svg)](https://codecov.io/gh/tgragnato/goflow)

- golang 1.23+ & bumped dependencies
- The Kafka transport has been removed; the new default transport is syslog
- GeoIP enrichment is now performed directly within the process
- The `sampler_hostname` field is enriched with the reverse DNS of sampler IP sources
