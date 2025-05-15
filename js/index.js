import { fetchExistingNames, validateName } from './infrastructure/fetchExistingNames.js';

export async function initIntroForm() {
  const form = document.getElementById('introForm');
  const input = document.getElementById('username');

  form.onsubmit = async e => {
    e.preventDefault();
    const name = input.value.trim();

    const existing = await fetchExistingNames();
    if (!validateName(name, existing)) {
      return alert('This name has been taken');
    }

    sessionStorage.setItem('surveyUser', name);
    location.href = 'compare.html';
  };
}

window.addEventListener('load', initIntroForm);
