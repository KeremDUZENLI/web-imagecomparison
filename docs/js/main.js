import { MatchSession } from './core/matchSession.js';
import { loadImages }   from './infrastructure/loadImages.js';
import { loadUserName } from './infrastructure/loadUserName.js';
import { postVote }     from './infrastructure/postVote.js';
import { showPair }     from './ui/showPair.js';
import { initDOM }      from './ui/initDOM.js'

let dom, userName, images, session;
const MIN_VOTES = 10;

async function bootstrap() {
  dom = initDOM();

  userName = await loadUserName();
  if (!userName) {
    return location.reload();
  }

  try {
    images = await loadImages();
    session = new MatchSession(images, MIN_VOTES);
  } catch (e) {
    console.error(e);
    alert('Initialization failed.');
    return;
  }

  dom.btnA.onclick = () => handleChoice(0);
  dom.btnB.onclick = () => handleChoice(1);
  dom.btnFinish.onclick = finishSession;

  loadNext();
}

async function handleChoice(idx) {
  const pair = session.currentPair;
  const winner = pair[idx];
  const loser  = pair[1 - idx];

  await postVote({ userName, imageWinner: winner, imageLoser: loser });
  session.applyVote();
  loadNext();
}

function loadNext() {
  const canFinish = session.canFinish();

  dom.btnFinish.style.display = canFinish ? 'inline-block' : 'none';
  dom.progress.textContent = canFinish
    ? `You've reached ${MIN_VOTES} votes — you may Finish or keep voting.`
    : `Match ${session.matchesDone + 1} of ${MIN_VOTES}`;

  const pair = session.nextPair();
  showPair(dom, pair, session.matchesDone, MIN_VOTES);
}

function finishSession() {
  dom.container.innerHTML = '<h2>Thanks! You’ve finished all votes.</h2>';
  dom.progress.textContent = '';
  dom.btnA.disabled = dom.btnB.disabled = true;
  dom.btnFinish.disabled = true;
}

window.addEventListener('load', bootstrap);
