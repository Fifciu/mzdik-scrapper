import routeHome from './views/home';
import { clearCoreNode } from './dom';

export const Router = function(routes){
  this.currentArgs = {};

  const findByPath = (path) => {
    const splitted = path.split('/');
    const params = Object.create(null);

    for (let j = 0, routesLength = routes.length; j < routesLength; j++) {
      let routeSplitted = routes[j].path.split('/');

      if (splitted.length === routeSplitted.length) {
        for (let i = 0; i < routeSplitted.length; i++) {
          if (routeSplitted[i].substr(0, 1) === ":") {
            // dynamic
            // saving url params to object
            params[routeSplitted[i].substr(1)] = splitted[i];
          }
          else {
            // static
            // break loop if it doesn't match
            if (routeSplitted[i] != splitted[i]) {
              break;
            }
          }

          if (i == (routeSplitted.length - 1)) {
            this.currentArgs = params;
            return j;
          }
        }
      }
    }
    this.currentArgs = {};
    return -1;
  };

  const launchView = (viewIndex) => {
    clearCoreNode();
    routes[viewIndex].renderFn(this.currentArgs);
  };

  const path = window.location.pathname;
  const args = path.split('/');

  let routeIndex = findByPath(path);

  if(routeIndex === -1 && path != '/'){
    window.location.href = '/';
  }

  // Render
  if(routeIndex >= 0) {
    launchView(routeIndex);
  }

  return {
    move: (path, state, title) => {
      let routeIndex = findByPath(`${path}`);

      if(routeIndex === -1){
        window.location.href = '/';
      } else {
        launchView(routeIndex);
        window.history.pushState(state, title, path);
        document.title = title;
      }
    },

    params: () => this.currentArgs
  };
};