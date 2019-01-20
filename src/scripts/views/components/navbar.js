export default (description) => {
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
  span.textContent = `${description}`;
  nav.appendChild(span);

  return nav;
};