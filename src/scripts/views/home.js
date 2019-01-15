import {API} from '../globals';
import {renderBusNumber, renderDiv, renderHeader, coreNode} from '../dom';

export default () => {
  fetch(`${API}/buses`)
  .then(response => response.json())
  .then(buses => {
    const parent = renderDiv();
    parent.id = "numbers";

    const header = renderHeader(1, 'Wybierz numer linii');
    const core = coreNode();

    core.appendChild(header);
    core.appendChild(parent);

    buses.forEach(bus => {
      parent.appendChild(renderBusNumber(bus));
    });
  });
};