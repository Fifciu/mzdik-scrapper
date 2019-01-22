const CACHE_STATIC_NAME = 'static-v10';
const CACHE_DYNAMIC_NAME = 'dynamic-v2';
const STATIC_FILES = [
  '/'
];

self.addEventListener('install', function(event) {
  // Perform install steps
  event.waitUntil(
      caches.open(CACHE_STATIC_NAME)
      .then(function(cache) {
        console.log('Opened cache');
        return cache.addAll(STATIC_FILES);
      })
  );
});

self.addEventListener('activate', function(event) {
  console.log('[Service Worker] Activating Service Worker ....', event);
  event.waitUntil(
      caches.keys()
      .then(keyList => {
        return Promise.all(keyList.map(val => {
          if(val !== CACHE_STATIC_NAME && val !== CACHE_DYNAMIC_NAME){
            console.log("[SW] Removing old cache.", val);
            return caches.delete(val);
          }
        }));
      })
  );
  return self.clients.claim();
});