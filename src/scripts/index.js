import '../styles/index.scss';

import {Router} from './router';

import routeHome from './views/home.js';
import routeBus from './views/bus.js';

var router;

(() => {
  router = new Router([
    { path: '/', renderFn: routeHome },
    { path: '/:busNumber', renderFn: routeBus },
    { path: '/:busNumber/:way', renderFn: ()=>{} },
  ]);
})();

export const routerObj = router;