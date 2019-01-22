import '../styles/index.scss';

import {Router} from './router';

import routeHome from './views/home.js';
import routeBus from './views/bus.js';
import routeSchedule from './views/schedule.js';

// Service worker
var deferredPrompt;

if('serviceWorker' in navigator) {
  navigator.serviceWorker.register('/service-worker.js')
      .then(() => {
        console.log('Worker registered!');
      })
      .catch(err => {
        console.log(err);
      });
}

window.addEventListener('beforeinstallprompt', event => {
  console.log('Beforeinstallprompt!');
  event.preventDefault();
  deferredPrompt = event;
  return false;
});


// Router
var router;

(() => {
  router = new Router([
    { path: '/', renderFn: routeHome },
    { path: '/:busNumber', renderFn: routeBus },
    { path: '/:busNumber/:way/:station/:delay', renderFn: routeSchedule },
  ]);
})();

export const routerObj = router;