---
layout: default
title: Snippets / Gists
index_description: Welcome to the Chronicles of Nonchalance!
description: >

    A snippet or a gist generally refers to a small piece or fragment of code, text, or information. <br><br>

    They are often used to demonstrate or illustrate a particular programming concept, syntax, or technique. <br>
    Their characteristic is of being small, self-contained pieces of code or information that can be easily shared or reused. <br><br>

    Why did I publish this page? Well, think of it as my digital attic – a curated collection of things I've decided I don't care about enough to keep to myself. <br><br>

    It's not a mess; it's avant-garde minimalism with a touch of 'selective enthusiasm'. Consider it a page-turner for the truly discerning audience who appreciates the art of not caring too much. <br>
    Welcome to the Chronicles of Nonchalance!

---

{% for snippet in site.snippets reversed %}
## {{ snippet.title }}
{{ snippet.content }}
{% endfor %}