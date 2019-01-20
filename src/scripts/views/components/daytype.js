import { renderDiv, renderHeader } from './../../dom';
import { compareWithOffset } from './../../utils';

export default (dayName, hours, offset) => {
  const div = renderDiv();
  div.classList.add('daytype');

  div.appendChild(
      renderHeader(3, dayName)
  );

  const ul = document.createElement('ul');
  ul.classList.add('days');

  hours.forEach(hour => {
    const li = document.createElement('li');
    li.classList.add('day');
    li.textContent = hour.Time;

    // Add offset!
    const date = new Date();
    const currentHour = `${date.getHours()}.${date.getMinutes()}`;
    if(compareWithOffset(hour.Time, currentHour, offset) === 1){
      li.classList.add('past');
    };

    ul.appendChild(li);
  });

  div.appendChild(ul);
  return div;
};