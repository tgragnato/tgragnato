---
sitemap: false
---
const urlsToCache = [
{% for file in site.static_files %}{% unless file.path contains '/samples/' %}  '{{ file.path | relative_url | jsonify }}',
{% endunless %}{% endfor %}
{% for page in site.pages %}  '{{ page.url | relative_url | jsonify }}',
{% endfor %}
{% for post in site.posts %}  '{{ post.url | relative_url | jsonify }}'{% unless forloop.last %},
{% endunless %}{% endfor %}
];
const CACHE_EPOCH = '{{ "now" | date: "%s" }}';

self.addEventListener('install', event => {
  event.waitUntil(
    caches.open(CACHE_EPOCH)
      .then(cache => {
        return Promise.all(
          urlsToCache.map(url => 
            cache.add(url).catch(error => {
              console.warn('[service-worker] Failed to cache:', url, error);
            })
          )
        );
      })
      .then(() => self.skipWaiting())
  );
});

self.addEventListener('fetch', event => {
  if (event.request.method !== 'GET') return;
  if (!event.request.url.startsWith('http')) return;
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