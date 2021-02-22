# FASTGate targeted root-RCE and CGN bypass [PoC]

This is the combination of already available exploit and attack techniques.

- [taviso/rbndr](README_rbndr.md)
- [Depau/fastgate-python](README_python.md)
- [Nimayer/fastgate-toolkit](README_toolkit.md)

## Local demonstration

Serve [exploit.html](exploit.html) as the index of a local webserver.
Craft a tailored `payload.shell` to verify the execution.
Open a web browser and visit `7f000001.c0a801fe.rbndr.us`
(switch between localhost and 192.168.1.254).

## Notes

Nothing prevents remote exploitation, root is gained.
Shell shoveling and verification is voluntarily omitted.

See [exploit.coffee](exploit.coffee) to inspect the code.

## Links

- http://www.fastweb.it/forum/servizi-rete-fissa-tematiche-tecniche/urgente-vulnerabilita-modem-fastgate-0-00-47-t22720.html
- https://www.exploit-db.com/exploits/44606/
- https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-6023
- https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-20122
- https://github.com/taviso/rbndr
- https://github.com/Depau/fastgate-python
- https://github.com/Nimayer/fastgate-toolkit
