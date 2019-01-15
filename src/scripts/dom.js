import { loadStations } from './interactions';

let cNode;

export const coreNode = () => {
  cNode = cNode || document.querySelector('div#app');
  return cNode;
};

export const clearCoreNode = () => {
  coreNode();
  while(cNode.hasChildNodes()){
    cNode.removeChild(cNode.firstElementChild||cNode.firstChild);
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