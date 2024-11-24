---
sitemap: false
---
{% assign posts = site.posts | where_exp:'post','post.sitemap != false' %}
const urlsToCache = [
  '/manifest.json',
  '/style.css',
  '/fonts/GrechenFuemen-Regular.woff2',
  '/fonts/square-github.svg',
  '/fonts/square-hn.svg',
  '/fonts/square-linkedin.svg',
  '/fonts/square-signal.svg',
  '/fonts/tor.svg',
{% for post in posts %}  '{{ post.url }}'{% unless forloop.last %},{% endunless %}
{% endfor %}
];
const CACHE_EPOCH = '{{ "now" | date: "%s" }}';

self.addEventListener('install', event => {
  event.waitUntil(
    caches.open(CACHE_EPOCH)
      .then(cache => {
        return cache.addAll(urlsToCache);
      })
  );
});

self.addEventListener('fetch', event => {
  event.respondWith(
    caches.match(event.request)
      .then(response => {
        if (response) {
          return response;
        }
        return fetch(event.request).then(fetchResponse => {
          return caches.open(CACHE_EPOCH).then(cache => {
            if (!fetchResponse || !fetchResponse.ok) {
              return fetchResponse;
            }
            cache.put(event.request, fetchResponse.clone());
            return fetchResponse;
          });
        });
      })
  );
});

self.addEventListener('activate', event => {
  const cacheWhitelist = [CACHE_EPOCH];
  event.waitUntil(
    caches.keys().then(keyList =>
      Promise.all(
        keyList.map(key => {
          if (!cacheWhitelist.includes(key)) {
            return caches.delete(key);
          }
        })
      )
    )
  );
});