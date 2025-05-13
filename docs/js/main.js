import { MatchSession } from './core/matchSession.js';
import { loadImages }   from './infrastructure/imageLoader.js';
import { logVote }      from './infrastructure/voteLogger.js';
import { showPair }     from './ui/displayPair.js';
import { showFinal }    from './ui/displayResults.js';
import { initDOM }      from './ui/domElements.js';

const MAX_MOVES = 10;
const K_FACTOR = 32;
const CONVERGENCE_THRESHOLD = 0.5;

let dom, session, currentPair, userName;

window.addEventListener('load', async () => {
  dom = initDOM();

  try {
    const images = await loadImages();
    const ratings = await loadRatings();

    session = new MatchSession(images, K_FACTOR, MAX_MOVES, CONVERGENCE_THRESHOLD, ratings);

    userName = prompt('Enter your name:');
    if (!userName) return location.reload();

    bindUIEvents();
    renderNextPair();
  } catch (e) {
    console.error(e);
    alert('Initialization failed.');
  }
});

function bindUIEvents() {
  dom.btnA.addEventListener('click', () => handleChoice(0));
  dom.btnB.addEventListener('click', () => handleChoice(1));
}

function renderNextPair() {
  if (session.isDone()) {
    showFinal(dom, session.getRatings());
  } else {
    currentPair = session.nextPair();
    showPair(dom, currentPair, session.matchesDone, MAX_MOVES);
  }
}

async function handleChoice(winnerIndex) {
  const winner = currentPair[winnerIndex];
  const loser = currentPair[1 - winnerIndex];
  const oldRatings = { winner: session.getRatings()[winner], loser: session.getRatings()[loser] };

  session.applyVote(winner, loser);
  const newRatings = { winner: session.getRatings()[winner], loser: session.getRatings()[loser] };

  await logVote({
    userName,
    imageA: currentPair[0],
    imageB: currentPair[1],
    imageWinner: winner,
    imageLoser: loser,
    eloWinnerPrevious: oldRatings.winner,
    eloWinnerNew: newRatings.winner,
    eloLoserPrevious: oldRatings.loser,
    eloLoserNew: newRatings.loser
  });

  renderNextPair();
}

async function loadRatings() {
  const res = await fetch('/api/ratings');
  if (!res.ok) throw new Error('Failed to load ratings');
  return res.json();
}
