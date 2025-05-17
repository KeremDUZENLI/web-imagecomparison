import { MIN_VOTES }      from './env/constants.js';

import { MatchSession }   from './core/matchSession.js';
import { postVote }       from './infrastructure/postVote.js';
import { initCompareDOM } from './ui/initCompareDOM.js';
import { loadImages }     from './ui/loadImages.js';
import { showPair }       from './ui/showPair.js';

let dom, currentUser, images, session;

const goTo = url => location.href = url;
const getSession = key => sessionStorage.getItem(key);
const setSession = (key, value) => sessionStorage.setItem(key, value);

const getJSON = key => {
  try { return JSON.parse(sessionStorage.getItem(key)); }
  catch { return null; }
};
const setJSON = (key, value) => sessionStorage.setItem(key, JSON.stringify(value));

const persistPair = pair => {
  setJSON('currentPair', pair);
  setJSON('lastPair',    pair);
};

const restorePair = () => {
  const pair = getJSON('currentPair');
  if (!pair) return null;

  session.currentPair = pair;
  session.lastPair    = getJSON('lastPair') || pair;
  showPair(dom, pair, session.matchesDone, MIN_VOTES);
  return true;
};

async function initCompare() {
  dom = initCompareDOM();

  dom.btnA.onclick      = () => handleChoice(0);
  dom.btnB.onclick      = () => handleChoice(1);
  dom.btnFinish.onclick = finishSession;

  attachWindowEvents();
}

function attachWindowEvents() {
  window.addEventListener('beforeunload', unloadHandler);
  window.addEventListener('pageshow',     pageshowHandler);
}

async function renderCompare() {
  currentUser = getSession('surveyUser');
  if (!currentUser) return goTo('index.html');

  if (!images) {
    images  = await loadImages();
    session = new MatchSession(images, MIN_VOTES);
  }

  session.matchesDone = initVoteState();

  if (restorePair()) return;

  loadNextPair();
}

function initVoteState() {
  const lastUser = getSession('votesUser');
  if (lastUser !== currentUser) {
    setSession('votesCount', '0');
    setSession('votesUser',  currentUser);
  }
  const count = Number(getSession('votesCount'));
  return Number.isNaN(count) ? 0 : count;
}

async function handleChoice(idx) {
  const [winner, loser] = [
    session.currentPair[idx],
    session.currentPair[1 - idx],
  ];

  await postVote({ username: currentUser, imageWinner: winner, imageLoser: loser });

  session.applyVote();
  setSession('votesCount', session.matchesDone);

  sessionStorage.removeItem('currentPair');
  loadNextPair();
}

function finishSession() {
  sessionStorage.removeItem('currentPair');
  setSession('votesCount', session.matchesDone);
  goTo('finish.html');
}

function unloadHandler(e) {
  if (session && !session.canFinish()) {
    const warning = 'You have unfinished votes — are you sure you want to leave?';
    e.preventDefault();
    e.returnValue = warning;
    return warning;
  }
}

function pageshowHandler(event) {
  const navType = performance.getEntriesByType('navigation')[0]?.type;
  if (event.persisted || navType === 'back_forward') {
    renderCompare();
  }
}

function loadNextPair() {
  const done      = session.matchesDone;
  const canFinish = session.canFinish();

  dom.btnFinish.style.display = canFinish ? 'inline-block' : 'none';
  dom.progress.textContent    = canFinish
    ? `You've reached ${MIN_VOTES} votes — you may Finish or keep voting.`
    : `Match ${done + 1} of ${MIN_VOTES}`;

  const pair = session.nextPair();
  session.lastPair = pair;
  persistPair(pair);
  showPair(dom, pair, done, MIN_VOTES);
}

window.addEventListener('DOMContentLoaded', async () => {
  await initCompare();
  await renderCompare();
});
