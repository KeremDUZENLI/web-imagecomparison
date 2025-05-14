import { MatchSession } from './core/matchSession.js';
import { loadImages }   from './infrastructure/loadImages.js';
import { postVote }     from './infrastructure/postVote.js';
import { showPair }     from './ui/showPair.js';
import { initDOM }      from './ui/initDOM.js'

let session, userName, dom;
const MIN_VOTES = 10;

async function bootstrap() {
  dom = initDOM();

  userName = prompt('Enter your name:');
  if (!userName) {
    return location.reload();
  }

  try {
    const images = await loadImages();
    session = new MatchSession(images, MIN_VOTES);
  } catch (e) {
    console.error(e);
    alert('Initialization failed.');
    return;
  }

  dom.btnA.onclick = () => handleChoice(0);
  dom.btnB.onclick = () => handleChoice(1);

  session.currentPair = loadNext();
}

async function handleChoice(idx) {
  const pair = session.currentPair;
  const winner = pair[idx];
  const loser  = pair[1 - idx];

  try {
    await postVote({
      userName,
      imageWinner: winner,
      imageLoser: loser,
    });
    session.applyVote();
    session.currentPair = loadNext();
  } catch (err) {
    alert(`Vote failed: ${err.message}`);
  }
}

function loadNext() {
  if (session.isDone()) {
    dom.container.innerHTML = '<h2>Thanks! You have completed all votes.</h2>';
    return;
  }
  const pair = session.nextPair();
  showPair(dom, pair, session.matchesDone, MIN_VOTES);
  return pair;
}

window.addEventListener('load', bootstrap);
