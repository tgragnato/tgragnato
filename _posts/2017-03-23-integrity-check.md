---
title: macOS firmware integrity check
description: Looking how Apple's security features are implemented across architectures
layout: default
---

Background context:

- [Negative Rings in Intel Architecture: The Security Threats That You’ve Probably Never Heard Of - Not Actual Protection Rings, But Conceptual Privilege Levels Susceptible To Exploitation](https://medium.com/swlh/negative-rings-in-intel-architecture-the-security-threats-youve-probably-never-heard-of-d725a4b6f831)

- [Vault 7 - a series of documents that WikiLeaks began to publish on 7 March 2017, detailing the activities and capabilities of the United States Central Intelligence Agency to perform electronic surveillance and cyber warfare](https://en.wikipedia.org/wiki/Vault_7)

---

## From "[Ask HN: How to clean Mac firmware?](https://news.ycombinator.com/item?id=13940284)"

> Given the news release today about Mac firmware being infected, how would you go about detecting this and cleaning it up?
> Is Apple going to do anything about it?

The released documents are very old, so we don't really know what's the actual state of the art.
This is what I did (any express or implied warranty is disclaimed):

`diskutil list`

```
/dev/disk0 (internal):
   #:                       TYPE NAME                    SIZE       IDENTIFIER
   0:      GUID_partition_scheme                         xxx.x GB   disk0
   1:                        EFI EFI                     314.6 MB   disk0s1
   2:          Apple_CoreStorage <Root_Name>             xxx.x GB   disk0s2
   3:                 Apple_Boot Recovery HD             650.0 MB   disk0s3
```

`sudo diskutil mount /dev/disk0s1`
```
Volume EFI on /dev/disk0s1 mounted
```

`ls -R /Volumes/EFI/EFI/APPLE`
```
EXTENSIONS FIRMWARE UPDATERS
/Volumes/EFI/EFI/APPLE/EXTENSIONS: Firmware.scap
/Volumes/EFI/EFI/APPLE/FIRMWARE: MBxx_xxxx_Bxx_LOCKED.fd
/Volumes/EFI/EFI/APPLE/UPDATERS: MULTIUPDATER USBCVA
...
```

Be sure there are no unwanted extensions. `shasum -a 256 <MBxx_xxxx_Bxx_LOCKED.fd>`
A sha hash is returned, check it against [https://github.com/gdbinit/firmware_vault](https://github.com/gdbinit/firmware_vault). If there's a mismatch, you may be affected. 

---

## From "[Xeno Kovah: macOS 10.13 EFI firmware integrity check (twitter.com/xenokovah)](https://news.ycombinator.com/item?id=15325829)"

So I hear macOS 10.13 comes out soon. Let's talk about what's up if you ever see this prompt.

![Your computer has detected a potential problem. Send a report to Apple.](/images/2017-03-23-EFIcheck-Prompt.jpg)

This comes from `/usr/libexec/firmwarecheckers/eficheck/eficheck`, a tool that's included in 10.13. eficheck runs once a week, and checks if measurements of EFI match some known-good measurements. (Yes, I've gone native, so I'm going to use the term "EFI" like I used to use "BIOS" to just refer to our x86 firmware.)

The tl;dr is that if you see it click send, unless you're running a hackintosh, in which case don't, because yours is garbage data to us.

But now a brief history/story about eficheck (which was a collaboration between @NikolajSchlej, @coreykal, and myself).

A design requirement we were given early on by Privacy was that we can't just scoop up everyone's firmware and send it back for analysis. Even if EFI is supposed to be signed so there's no legit way for customers to modify it, and even if we're stripping out the nvram vars.

So that led to what I'm sure folks in the security community will recognize as a suboptimal security design, but we had to start somewhere. (I look forward to some \*ahem\*, "researcher", doing a presentation about bypassing eficheck as if we don't know that's possible.)

In my perfect world we would have rolled this out way back in 10.12 as just a silent data collector. It would have been designed to \<sarcasm\>"surprise and delight"\</sarcasm\> any potential attackers. But we just kept getting requirements added and people just kept expecting it to follow typical project guidelines like seeding (the nerve!).

It was the first requirement, to not just collect data, but compare data against a known good baseline, which pushed the schedule the most.

It turns out, collecting all possible binaries for all EFIs is not as trivial as one might hope.
We found early on that there wasn't a good single archive of everything we had ever shipped.
Much later on (as in, this past June), I found a much better archive in build system records as I was being forced to learn the build system.
But even that had apparently had some data loss back around 2014, so some binaries occasionally had to be dredged up from the depths.
There were also the occasional cases where the version that went on in the factory only existed in build records, not the EFI team records.

(That's what @stroughtonsmith was unlucky enough to run into here: [stroughtonsmith/896020333098151937](https://twitter.com/stroughtonsmith/status/896020333098151937))
But even once one has all the data, there are foibles to it, which need to be compensated for. Like Apple-custom sections.

...

> [https://twitter.com/xenokovah/status/912011064304267264](https://twitter.com/xenokovah/status/912011064304267264)

---

## Update (nov 2022)

In macOS 13 the binary is amusingly still there, even on Apple Silicon. (the build system may not be as sophisticated as one could naively think)

![eficheck-standalone decompilation](/images/2017-03-23-Decompilation.png)

The firmware check is now done in a much more robust way, during the bootstrap process, during the `Low Level Bootstrap`.

[![Apple Silicon Bootstrap](/images/2017-03-23-AppleSilicon-Boot.png)](https://support.apple.com/guide/security/boot-process-secac71d5623/web)
