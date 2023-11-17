---
layout: default
title: tgragnato.it - Tommaso Gragnato
override_title: tgragnato - tgragnato.it
description: >

    Hi there 👋 I'm Tommaso <br><br>

    - 🔧 I am a backend developer currently engaged in platform engineering roles <br>
    - 💼 I strive to leverage my knowledge and expertise to create and manage robust and scalable infrastructures <br>
    - 🚀 My focus is on security, containerization and orchestration, monitoring and logging <br>
    - 🔭 I am presently working on projects that involve Go and/or PHP <br>
    - 🌱 I am actively engaged in expanding my skill set by delving into the realms of Ruby and Javascript <br>
    - 💬 I'm always eager to engage in discussions and share my knowledge, especially about networking and censorship circumvention

---

## Contacts

Email: tgragnato「単価記号」icloud

PGP: A283 C77F BE2B 15D8 7EF4 16E9 B7E4 732B 4DE7 4631

[https://keys.openpgp.org/vks/v1/by-fingerprint/A283C77FBE2B15D87EF416E9B7E4732B4DE74631](https://keys.openpgp.org/vks/v1/by-fingerprint/A283C77FBE2B15D87EF416E9B7E4732B4DE74631)

[https://gitlab.torproject.org/tgragnato.gpg](https://gitlab.torproject.org/tgragnato.gpg) - [http://eweiibe6tdjsdprb4px6rqrzzcsi22m4koia44kc5pcjr7nec2rlxyad.onion/tgragnato.gpg](http://eweiibe6tdjsdprb4px6rqrzzcsi22m4koia44kc5pcjr7nec2rlxyad.onion/tgragnato.gpg)

---

## Words

{% for post in site.posts %}
[{% if post.date %}{{ post.date | date: "%d-%m-%Y" }} - {% endif %} `{{ post.title }}`{% if post.description %} - {{ post.description }}{% endif %}]({{ post.url }})
{% endfor %}
{% for page in site.html_pages %}
{% if page.url != '/404.html' and page.url != '/' %}
[日付不明 - `{{ page.title }}`{% if page.index_description %} - {{ page.index_description }}{% elsif page.description %} - {{ page.description }}{% endif %}]({{ page.url }})
{% endif %}
{% endfor %}

---

## Git

[`tgragnato` - A ✨special✨ repository dedicated to my profile README and my Jekyll website](https://github.com/tgragnato/tgragnato)

[`ebmbtreebench` - A naive measurement tool for ebmb trees](https://github.com/tgragnato/ebmbtreebench)

[`magnetico` - An updated version of the famous "Autonomous (self-hosted) BitTorrent DHT search engine suite"](https://github.com/tgragnato/magnetico)

[`snowflake` - Pluggable Transport using WebRTC, inspired by Flashproxy. A custom fork with mine opinionated patches](https://github.com/tgragnato/snowflake)

[`amule` - An aMule fork. Personal companion of tgragnato/homebre-amule](https://github.com/tgragnato/amule)

[`homebrew-amule` - A custom homebrew tap for aMule](https://github.com/tgragnato/homebrew-amule)

[`pure` - A collection of software and tools that I use to manage my network and my storage ](https://github.com/tgragnato/pure)

[`pilsung` - 북쪽 필성。 붉은별 사용자용체계。 pilsung 암호화 알고리즘](https://github.com/tgragnato/pilsung)

[`dns323-toolchain` - D-Link DNS-323 cross compilation toolchain](https://github.com/tgragnato/dns323-toolchain)

[`migration` - Simple tool that mirrors Paradox database as SQLite at runtime. Tuned for Mosaico Sorgente Aperto.](https://github.com/tgragnato/migration)

[`socialhub` - GraphQLHub demo](https://github.com/tgragnato/socialhub)

[`owncloud-convert` - Converts videos stored in Owncloud at the click of a button](https://github.com/tgragnato/owncloud-convert)

---

## Quotes

> Humans are allergic to change. They love to say, "We've always done it this way." I try to fight that. `Grace Murray Hopper`

> We can lick gravity but sometimes the paperwork is overwhelming. `Wernher Magnus Maximilian Freiherr von Braun`

> It’s often easier to ask forgiveness than to ask for permission. `Grace Murray Hopper`
