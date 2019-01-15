import {API} from '../globals';
import {stationsInteraction} from './../interactions';
import {renderBusNumber, renderDiv, renderHeader, coreNode} from '../dom';

export default (keys) => {
  fetch(`${API}/buses/${keys.busNumber}/stations`)
  .then(response => response.json())
  .then(stations => {
    const core = coreNode();
    core.classList.toggle('bus-list'); // REMEMBER TO DELETE

    // Navbar
    const nav = document.createElement('nav');
    nav.classList.add('navbar');

    const icon = document.createElement('i');
    icon.id = 'backIcon';
    icon.classList.add('fas');
    icon.classList.add('fa-chevron-circle-left');
    nav.appendChild(icon);

    const header = document.createElement('strong');
    header.textContent = 'MzdikPWA';
    nav.appendChild(header);

    const span = document.createElement('span');
    span.textContent = `Linia nr ${keys.busNumber}`;
    nav.appendChild(span);

    core.appendChild(nav);

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