import {API} from '../globals';
import {stationsInteraction} from './../interactions';
import {
  renderBusNumber, renderDiv, renderHeader, coreNode,
  addNavbarIfNotExists,
} from '../dom';
import cmpNavbar from './components/navbar';

export default (keys) => {
  fetch(`${API}/buses/${keys.busNumber}/stations`)
  .then(response => response.json())
  .then(stations => {
    const core = coreNode();
    core.classList.add('bus-list'); // REMEMBER TO DELETE
    core.className = 'bus-list';
    // Navbar
    addNavbarIfNotExists(`Linia nr ${keys.busNumber}`, '/', 'MzdikPWA');

    // Header - Kierunek
    const h2 = renderHeader(2, 'Kierunek');
    h2.classList.add('direction__header');
    core.appendChild(h2);

    // Each direction
    Object.keys(stations).forEach(direction => {
      const wrapper = renderDiv();
      wrapper.classList.add('directions');

      const headerWrapper = renderDiv();
      headerWrapper.classList.add('direction-header__wrapper');

      const h4 = renderHeader(4, stations[direction][stations[direction].length - 1].Name);
      h4.classList.add('direction__name');
      headerWrapper.appendChild(h4);

      const icon = document.createElement('i');
      icon.classList.add('toggleDirection');
      icon.classList.add('fas');
      icon.classList.add('fa-angle-down');
      headerWrapper.appendChild(icon);

      wrapper.appendChild(headerWrapper);

      const ul = document.createElement('ul');
      ul.classList.add('direction');
      ul.dataset.direction = direction;

      stations[direction].forEach((station, index) => {
        const li = document.createElement('li');
        li.dataset.num = index;
        li.dataset.delay = station.AverageDelay;

        const busIcon = document.createElement('i');
        busIcon.classList.add('fas');
        busIcon.classList.add('fa-bus');

        const textCnt = document.createElement('span');
        textCnt.textContent = station.Name;

        li.appendChild(busIcon);
        li.appendChild(textCnt);

        ul.appendChild(li);
      });

      wrapper.appendChild(ul);
      core.appendChild(wrapper);
    });

    stationsInteraction();
  });
};