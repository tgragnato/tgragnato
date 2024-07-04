---
title: D-Link botnet
description: Extracting binaries from a botnet of compromised NAS devices
layout: default
images:
  - loc: /images/2018-07-01-automatic-decompilation.webp
    caption: automatic decompilation
  - loc: /images/2018-07-01-implant.webp
    caption: implant.c
  - loc: /images/2018-07-01-C2-IP.webp
    caption: IP of the C&C
  - loc: /images/2018-07-01-VirusTotal.webp
    caption: VirusTotal results
---

## SEARCH-LAB security assessment

The 53 unique vulnerabilities [identified in 2014](/documents/2018-07-01-Advisory.pdf) are now actively being exploited in the wild.
Remote attackers can execute arbitrary code and gain full control over the device.

## Implants

![automatic decompilation](/images/2018-07-01-automatic-decompilation.webp){:loading="lazy"}

The automatic decompilation of the implants is trivial.

![implant.c](/images/2018-07-01-implant.webp){:loading="lazy"}

The names of the functions are clearly visible and give an exact indication of what the implant does,
the IP of the C&C is directly exposed, the download of the secondary stage was not protected.

![IP of the C&C](/images/2018-07-01-C2-IP.webp){:loading="lazy"}

Despite the reverse engineering being trivial on them all detection amongst av-engines varies greatly,
as is shown by the VirusTotal results.

![VirusTotal results](/images/2018-07-01-VirusTotal.webp){:loading="lazy"}

etp.db is an encrypted sqlite database, and is the only thing that the attackers protected.

## Samples

[2pn4yRMh8T1G.so - FXx3CyvEoz9E.so - IWBtRNuqVS.so : e29afce3afe992b57c7b03660d4ec5fcdcf2b694f580d221121d9cd9e04e15d5](/samples/e29afce3afe992b57c7b03660d4ec5fcdcf2b694f580d221121d9cd9e04e15d5)

[DsGn8v1r.so - ZI42LA3R.so : 8f4967b653b6e7e00943e7a96d9126a0b734b9c0613029487179fab76f4aa4c0](/samples/8f4967b653b6e7e00943e7a96d9126a0b734b9c0613029487179fab76f4aa4c0)

[8sy2b96E9K.so : dfa64db65f854441bd094acf3a48ecf858848c79f8f312e6fec0ade8e8fe43dd](/samples/dfa64db65f854441bd094acf3a48ecf858848c79f8f312e6fec0ade8e8fe43dd)

[wLl7EhHw.so : a725aad4f05628d2c49d46ee6070aec0a01f70cbac720ccee54cbab5524bee53](/samples/a725aad4f05628d2c49d46ee6070aec0a01f70cbac720ccee54cbab5524bee53)

[ilQqCI+J.dms : a14dbeea55221daf5344ab5f5ca98dd209f415cae6cb168c62a39db32b2d77fa](/samples/a14dbeea55221daf5344ab5f5ca98dd209f415cae6cb168c62a39db32b2d77fa)

[etp.db : b2b88659556adf9510507ee6187e3af4b52bd49e04cb4247527dd60cf4ce72bc](/samples/b2b88659556adf9510507ee6187e3af4b52bd49e04cb4247527dd60cf4ce72bc)

[etp2.db : 3f7e34d6c13d10730a24f3834db01f290961e4a594b4b13f927c387d49f5368f](/samples/3f7e34d6c13d10730a24f3834db01f290961e4a594b4b13f927c387d49f5368f)

[etp3.db : c4e3e2443dbc6578cc8f7e4614dc9591e315da65aaa68a2fee81ef8a623504a3](/samples/c4e3e2443dbc6578cc8f7e4614dc9591e315da65aaa68a2fee81ef8a623504a3)