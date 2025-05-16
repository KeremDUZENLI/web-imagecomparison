import { MatchSession }  from './core/matchSession.js';
import { postVote }      from './infrastructure/postVote.js';
import { getCompareDOM } from './ui/getCompareDOM.js'
import { loadImages }    from './ui/loadImages.js';
import { showPair }      from './ui/showPair.js';

let dom, username, images, session;
const MIN_VOTES = 10;

async function setup() {
  dom = getCompareDOM();

  dom.btnA.onclick      = () => handleChoice(0);
  dom.btnB.onclick      = () => handleChoice(1);
  dom.btnFinish.onclick = finishSession;

  window.addEventListener('beforeunload', unloadHandler);
  window.addEventListener('pageshow', pageshowHandler);
}

async function render() {
  const currentUser = sessionStorage.getItem('surveyUser');
  if (!currentUser) return location.href = 'index.html';
  username = currentUser;
  
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

  await postVote({ username, imageWinner: winner, imageLoser: loser });

  session.applyVote();
  sessionStorage.setItem('votesCount', session.matchesDone);

  loadNext();
}

function initVoteState() {
  const user = sessionStorage.getItem('surveyUser');  
  const last = sessionStorage.getItem('votesUser'); 

  if (last !== user) {
    sessionStorage.setItem('votesCount', '0');  
    sessionStorage.setItem('votesUser', user);  
  }
  
  return parseInt(sessionStorage.getItem('votesCount') || '0', 10);  
}

function finishSession() {
  sessionStorage.setItem('votesCount', session.matchesDone);
  location.href = 'finish.html';
}

function loadNext() {
  const canFinish = session.canFinish();

  dom.btnFinish.style.display = canFinish ? 'inline-block' : 'none';
  dom.progress.textContent    = canFinish ? `You've reached ${MIN_VOTES} votes — you may Finish or keep voting.` : `Match ${session.matchesDone + 1} of ${MIN_VOTES}`;

  const pair = session.nextPair();
  showPair(dom, pair, session.matchesDone, MIN_VOTES);
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
    render();
  }
}

window.addEventListener('load', async () => {
  await setup();
  await render();
});
