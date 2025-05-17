import { getUsernames } from './infrastructure/getUsernames.js';

const goTo = url => location.href = url;
const formatName = inputEl => inputEl.value.trim().toLowerCase();

export async function initIntroForm() {
  const form  = document.getElementById('introForm');
  const input = document.getElementById('username');

  form.onsubmit = async event => {
    event.preventDefault();

    const name = formatName(input);
    const existingNames = await getUsernames();
    if (existingNames.includes(name)) {
      alert('This name has been taken');
      return;
    }

    sessionStorage.setItem('surveyUser', name);
    goTo('compare.html');
  };
}

window.addEventListener('DOMContentLoaded', initIntroForm);
