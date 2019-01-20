import {API} from '../globals';
import {renderDiv, renderHeader, coreNode, addNavbarIfNotExists} from '../dom';
import cmpNavbar from './components/navbar';
import cmpDaytype from './components/daytype';

export default (keys) => {

  fetch(`${API}/buses/${keys.busNumber}/timetable/${keys.way.toLowerCase()}`)
  .then(response => response.json())
  .then(dayTypes => {
    let target = 'CasualDay';
    const core = coreNode();
    core.className = 'cont-schedule';

    const day = new Date().getDay();
    if(day < 1){
      target = 'Saints';
    } else if(day > 5) {
      target = 'Saturday';
    }

    // Navbar
    const path = window.location.pathname.split('/');
    const oldPath = `/${path[1]}`;

    addNavbarIfNotExists(`Linia nr ${keys.busNumber}`, oldPath, 'MzdikPWA');

    //console.log(dayTypes[target]);

    const schedule = renderDiv();
    schedule.classList.add('schedule');

    const getDayname = (t) => {
      let dayname = 'Dzień powszedni';
      switch(t){
        case 'Saints':
          dayname = 'Niedziela';
          break;
        case 'Saturday':
          dayname = 'Sobota';
          break;
      }

      return dayname;
    };



    const properDaytype = cmpDaytype(getDayname(target), dayTypes[target], keys.delay);
    schedule.appendChild(properDaytype);

    const dayTypesKeys = Object.keys(dayTypes);
    for(let i = 0; i < dayTypesKeys.length; i++){
      if(dayTypesKeys[i] === target)
        continue;

      const anotherDaytype = cmpDaytype(getDayname(dayTypesKeys[i]), dayTypes[dayTypesKeys[i]], keys.delay);
      schedule.appendChild(anotherDaytype);
    }

    core.appendChild(schedule);
  });
};