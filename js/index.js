import { getUsernames } from './infrastructure/getUsernames.js';
import { postSurvey }   from './infrastructure/postSurvey.js';

const goTo = url => location.href = url;
const setSession = (key, value) => {sessionStorage.setItem(key, value)};
const formatName = input => input.value.trim().toLowerCase().replace(/\s+/g, '').replace(/[^a-z0-9.]/g, '');

async function initIntroForm() {
  const form             = document.getElementById('intro_form');
  const inputUsername    = document.getElementById('username');
  const selectAge        = document.getElementById('age');
  const selectGender     = document.getElementById('gender');
  const selectVRExp      = document.getElementById('experience');
  const selectDomainExpt = document.getElementById('profession');

  form.addEventListener('submit', async event => {
    event.preventDefault();

    const survey = {
      username: formatName(inputUsername),
      age: selectAge.value,
      gender: selectGender.value,
      experience: selectVRExp.value,
      profession: selectDomainExpt.value
    };

    let existingNames;
    try {
      existingNames = await getUsernames();
    } catch (error) {
      alert('Could not connect to the server');
      return;
    }

    if (existingNames.includes(survey.username)) {
      alert('This name has been taken');
      return;
    }

    try {
      await postSurvey(survey);
    } catch (error) {
      alert('Could not save survey');
      return;
    }

    setSession('surveyUser', survey.username);
    setSession('surveyAge', survey.age);
    setSession('surveyGender', survey.gender);
    setSession('surveyVRExperience', survey.experience);
    setSession('surveyDomainExpertise', survey.profession);

    goTo('compare.html');
  });
}

window.addEventListener('DOMContentLoaded', initIntroForm);
