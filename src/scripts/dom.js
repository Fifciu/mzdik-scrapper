import { loadStations } from './interactions';
import cmpNavbar from './views/components/navbar';
import {routerObj} from './index';

let cNode;

export const coreNode = () => {
  cNode = cNode || document.querySelector('div#app');
  return cNode;
};

export const clearCoreNode = () => {
  return new Promise((resolve) => {
    transitionOut();

    setTimeout(() => {
      const kids = cNode.childNodes;

      for(let i = 0; i < kids.length; i++){
        if(kids[i].tagName === "NAV"){
          continue;
        }

        cNode.removeChild(kids[i]);
        i--;
      }

      resolve();
    }, 1000);
  });
};

export const transitionOut = () => {
  coreNode();
  cNode.classList.add('mzd-transition-out');
};

export const transitionIn = () => {
  coreNode();
  cNode.classList.remove('mzd-transition-out');
};

export const addNavbarIfNotExists = (content, iconUrl, moveTitle) => {
  coreNode();
  const nav = cNode.querySelector('nav.navbar');

  const setIconUrl = (el, url, title) => {
   /* el.onclick = (ev) => {
      routerObj.move(url, {}, title);
    };*/
  };

  if(nav === null){
    const navbar = cmpNavbar(content);
    const icon = navbar.querySelector('i#backIcon');
    if(iconUrl === undefined){
      icon.classList.add('hidden');
    } else if(icon.classList.contains('hidden')){
      icon.classList.remove('hidden');
      setIconUrl(icon, iconUrl, moveTitle);
    }
    cNode.appendChild(navbar);
  } else {
    const icon = nav.querySelector('i#backIcon');
    if(iconUrl === undefined){
      icon.classList.add('hidden');
    } else if(icon.classList.contains('hidden')){
      icon.classList.remove('hidden');
      setIconUrl(icon, iconUrl, moveTitle);
    }
    const span = nav.querySelector('span');
    if(span === null && content.trim() !== ''){
      const elSpan = document.createElement('span');
      span.textContent = content;
      nav.appendChild(elSpan);
    } else if(span !== null){
      span.textContent = content;
    }
  }
};

export const renderBusNumber = (number) => {

  const el = renderDiv();
  el.classList.add('number');
  el.textContent = number;

  el.addEventListener('click', () => {
    loadStations(number);
  });

  return el;
};

export const renderDiv = () => {
  const el = document.createElement('div');
  return el;
};

export const renderHeader = (n, content) => {
  const el = document.createElement(`h${n}`);
  el.textContent = content;

  return el;
};