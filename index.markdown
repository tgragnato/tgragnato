---
layout: default
---

Email: tgragnato「単価記号」icloud

PGP: A283 C77F BE2B 15D8 7EF4 16E9 B7E4 732B 4DE7 4631

[https://keys.openpgp.org/vks/v1/by-fingerprint/A283C77FBE2B15D87EF416E9B7E4732B4DE74631](https://keys.openpgp.org/vks/v1/by-fingerprint/A283C77FBE2B15D87EF416E9B7E4732B4DE74631)

[https://gitlab.torproject.org/tgragnato.gpg](https://gitlab.torproject.org/tgragnato.gpg) - [http://eweiibe6tdjsdprb4px6rqrzzcsi22m4koia44kc5pcjr7nec2rlxyad.onion/tgragnato.gpg](http://eweiibe6tdjsdprb4px6rqrzzcsi22m4koia44kc5pcjr7nec2rlxyad.onion/tgragnato.gpg)

---

{% for post in site.posts %}
[{% if post.date %}{{ post.date | date: "%d-%m-%Y" }} - {% endif %} {{ post.title }}{% if post.description %} - {{ post.description }}{% endif %}]({{ post.url }})
{% endfor %}