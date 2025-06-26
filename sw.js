---
sitemap: false
---
const urlsToCache = [
{% for file in site.static_files %}{% unless file.path contains '/samples/' %}  '{{ file.path }}',
{% endunless %}{% endfor %}
{% for page in site.pages %}  '{{ page.url }}',
{% endfor %}
{% for post in site.posts %}  '{{ post.url }}'{% unless forloop.last %},
{% endunless %}{% endfor %}
];
const CACHE_EPOCH = '{{ "now" | date: "%s" }}';

self.addEventListener('install', event => {
  event.waitUntil(
    caches.open(CACHE_EPOCH)
      .then(cache => {
        return cache.addAll(urlsToCache).catch(error => {
          console.error('[service-worker] Failed to cache: ', error);
        });
      })
      .then(() => self.skipWaiting())
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
    ).then(() => self.clients.claim())
  );
});