# D-Link botnet

## SEARCH-LAB security assessment

The 53 unique vulnerabilities identified in 2014 are now actively being exploited in the wild.
Remote attackers can execute arbitrary code and gain full control over the device.

## Implants

The [automatic](screen_1.png) decompilation of the [implants](screen_2.png) is trivial.

The names of the functions are clearly visible and give an exact indication of what the implant does,
the [IP of the C&C](screen_3.png) is directly exposed, the download of the secondary stage was not protected.

Despite the reverse engineering being trivial on them all detection amongst av-engines varies greatly,
as is shown by the [VirusTotal results](screen_4.png).

etp.db is an encrypted sqlite database, and is the only thing that the attackers protected.
