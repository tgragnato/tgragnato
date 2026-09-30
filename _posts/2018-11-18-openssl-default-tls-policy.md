---
title: An OpenSSL Default TLS Policy Fragment
title_marker: OpenSSL
description: A historical system-default configuration for protocol versions, cipher suites, and TLS options
layout: default
lang: en
---

This is a fragment for OpenSSL's system-default TLS configuration. It sets a minimum protocol version, raises the cipher security level, lists TLS 1.3 cipher suites, and requests additional options such as record padding.

It is not a complete OpenSSL configuration: the top-level directive that activates the configuration section may need to be supplied by the surrounding file. Support and meaning of individual commands also depend on the OpenSSL version and how the application loads its configuration. Treat this as a historical policy snapshot, not a drop-in hardening recipe.

```
[default_conf]
ssl_conf = ssl_sect

[ssl_sect]
system_default = system_default_sect

[system_default_sect]
MinProtocol = TLSv1.2
CipherString = DEFAULT@SECLEVEL=2

# Hardening
Ciphersuites = TLS_CHACHA20_POLY1305_SHA256:TLS_AES_256_GCM_SHA384:TLS_AES_128_GCM_SHA256
Options = ServerPreference,PrioritizeChaCha
RecordPadding = 16384
```