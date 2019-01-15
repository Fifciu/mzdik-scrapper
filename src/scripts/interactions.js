import { API } from './globals';
import {routerObj} from './index';

export const loadStations = async (busNumber) => {
  let response = await fetch(`${API}/buses/${busNumber}/stations`);
  let data = await response.json();

  routerObj.move(`/${busNumber}`, { bus: busNumber }, `MzdikPWA - Linia nr ${busNumber}`);
};

export const stationsInteraction = () => {
  // Second page interactions
  [...document.querySelectorAll('.direction-header__wrapper')].forEach(el => {
    el.addEventListener('click', ev => {
      const parent = parent || el.parentElement.querySelector('ul.direction');
      let busy = false;

      if(!busy){
        if(!parent.classList.contains('active')){
          parent.style.display = 'block';
          busy = true;
          setTimeout(()=>{
            parent.classList.toggle('active');
            el.classList.toggle('active');
          }, 1);

          setTimeout(() => {
            busy = false;
          }, 1000);
        } else {
          el.classList.toggle('active');
          parent.classList.toggle('active');
          busy = true;
          setTimeout(() => {
            parent.style.display = 'none';
            busy = false;
          }, 1000);
        }
      }

    });
  });

  [...document.querySelectorAll('.directions > .direction > li')].forEach(el => {
    el.addEventListener('click', ev => {
      alert(`${el.textContent} / ${el.dataset.num} / ${el.parentElement.dataset.direction}`);
    });
  });
};