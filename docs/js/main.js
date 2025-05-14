import { MatchSession } from './core/matchSession.js';
import { loadImages }   from './infrastructure/imageLoader.js';
import { loadRatings }  from './infrastructure/ratingLoader.js';
import { logVote }      from './infrastructure/voteLogger.js';
import { showPair }     from './ui/displayPair.js';
import { showFinal }    from './ui/displayResults.js';
import { initDOM }      from './ui/domElements.js';
import constants         from '../env/constants.js';

let dom, session, currentPair, userName;

const MIN_VOTES      = constants.MIN_VOTES;
const DEFAULT_RATING = constants.DEFAULT_RATING;

window.addEventListener('load', async () => {
  dom = initDOM();

  try {
    const images = await loadImages();
    const ratings = await loadRatings();

    session = new MatchSession(images, MIN_VOTES, DEFAULT_RATING, ratings);

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
    showPair(dom, currentPair, session.matchesDone, MIN_VOTES);
  }
}

async function handleChoice(idx) {
  const [winner, loser] = [currentPair[idx], currentPair[1 - idx]];
  const oldRound = {
    winner: session.getRatings()[winner],
    loser:  session.getRatings()[loser],
  };

  try {
    const { vote, ratings } = await logVote({
      userName,
      imageA: currentPair[0],
      imageB: currentPair[1],
      imageWinner: winner,
      imageLoser:  loser,
      eloWinnerPrevious: oldRound.winner,
      eloWinnerNew:      session.getRatings()[winner],
      eloLoserPrevious:  oldRound.loser,
      eloLoserNew:       session.getRatings()[loser],
    });

    session.applyVote();
    session.ratings = ratings;
    renderNextPair();
  } catch (err) {
    alert(`Vote failed: ${err.message}`);
  }
}
