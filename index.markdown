---
layout: default
title: tgragnato.it - Tommaso Gragnato
override_title: tgragnato - tgragnato.it
index_description: Homepage
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
{% unless post.tag %}
[{{ post.date | date: "%d-%m-%Y" }} - `{{ post.title }}`{% if post.description %} - {{ post.description }}{% endif %}]({{ post.url }})
{% endunless %}
{% endfor %}
{% for page in site.html_pages %}
{% if page.url != '/404.html' and page.url != '/' %}
[日付不明 - `{{ page.title }}`{% if page.index_description %} - {{ page.index_description }}{% elsif page.description %} - {{ page.description }}{% endif %}]({{ page.url }})
{% endif %}
{% endfor %}

---

## Git

{% for git in site.data.git %}
[`{{ git.title }}` - {{ git.text }}]({{ git.href }})
{% endfor %}

---

## Quotes

{% for quote in site.data.quotes %}
> {{ quote.text }} `{{ quote.author }}`
{% endfor %}


---

## Advent Of Code 2024

{% for post in site.posts %}
{% if post.tag == "aoc24" %}
[{{ post.date | date: "%d-%m-%Y" }} - `{{ post.title }}`{% if post.description %} - {{ post.description }}{% endif %}]({{ post.url }})
{% endif %}
{% endfor %}

---

## Advent Of Code 2023

{% for post in site.posts %}
{% if post.tag == "aoc23" %}
[{{ post.date | date: "%d-%m-%Y" }} - `{{ post.title }}`{% if post.description %} - {{ post.description }}{% endif %}]({{ post.url }})
{% endif %}
{% endfor %}
