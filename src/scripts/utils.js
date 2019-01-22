const compareHours = (a, b) => {
  const timeA = a.split('.').map(Number);
  const timeB = b.split('.').map(Number);

  if(timeA[0] > timeB[0]){
    return 0;
  } else if(timeA[0] < timeB[0]){
    return 1;
  } else {
    if(timeA[1] >= timeB[1]){
      return 0;
    } else {
      return 1;
    }
  }
};

export const offsetToHour = (hour, offset) => {
  const time = hour.split('.').map(Number);
  offset = Number(offset);

  time[1] += offset;
  if(time[1] >= 60){
    time[1] %= 60;
    time[0]++;
    if(time[0] > 23){
      time[0] = 0;
    }
  }

  for(let i = 0; i < 2; i++){
    if(time[i] < 10){
      time[i] = `0${time[i]}`;
    }
  }

  return time.join('.');
};

export const compareWithOffset = (a, b, aOffset) => compareHours(offsetToHour(a, aOffset), b);