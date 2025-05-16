import { MatchSession }   from './core/matchSession.js';
import { MIN_VOTES }      from './env/contants.js';
import { postVote }       from './infrastructure/postVote.js';
import { initCompareDOM } from './ui/initCompareDOM.js'
import { loadImages }     from './ui/loadImages.js';
import { showPair }       from './ui/showPair.js';

let dom, currentUser, images, session;

async function initCompare() {
  dom = initCompareDOM();

  dom.btnA.onclick      = () => handleChoice(0);
  dom.btnB.onclick      = () => handleChoice(1);
  dom.btnFinish.onclick = finishSession;

  window.addEventListener('beforeunload', unloadHandler);
  window.addEventListener('pageshow', pageshowHandler);
}

async function renderCompare() {
  const username = sessionStorage.getItem('surveyUser');
  if (!username) return location.href = 'index.html';
  currentUser = username;
  
  if (!images) {
    images  = await loadImages();
    session = new MatchSession(images, MIN_VOTES);
  }
  
  session.matchesDone = initVoteState();
  loadNext();
}

async function handleChoice(idx) {
  const pair   = session.currentPair;
  const winner = pair[idx];
  const loser  = pair[1 - idx];

  await postVote({ username: currentUser, imageWinner: winner, imageLoser: loser });

  session.applyVote();
  sessionStorage.setItem('votesCount', session.matchesDone);

  loadNext();
}

function finishSession() {
  sessionStorage.setItem('votesCount', session.matchesDone);
  location.href = 'finish.html';
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
  const nav = performance.getEntriesByType("navigation")[0]?.type;
  if (event.persisted || nav === "back_forward") {
    renderCompare();
  }
}

function initVoteState() {
  const user = sessionStorage.getItem('surveyUser');
  const last = sessionStorage.getItem('votesUser');

  if (last !== user) {
    sessionStorage.setItem('votesCount', '0');  
    sessionStorage.setItem('votesUser', user);  
  }
  
  const count = Number(sessionStorage.getItem('votesCount'));
  return isNaN(count) ? 0 : count; 
}

function loadNext() {
  const canFinish = session.canFinish();

  dom.btnFinish.style.display = canFinish ? 'inline-block' : 'none';
  dom.progress.textContent    = canFinish ? `You've reached ${MIN_VOTES} votes — you may Finish or keep voting.` : `Match ${session.matchesDone + 1} of ${MIN_VOTES}`;

  const pair = session.nextPair();
  showPair(dom, pair, session.matchesDone, MIN_VOTES);
}

window.addEventListener('DOMContentLoaded', async () => {
  await initCompare();
  await renderCompare();
});
