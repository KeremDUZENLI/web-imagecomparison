import { pickPair, updateElo, shouldContinue } from './elo-calculator.js';
import { loadImages }                          from './image-loader.js';
import { initDOM, showPair, showFinal }        from './init-dom.js';


const MAX_MOVES             = 10;
const CONVERGENCE_THRESHOLD = 0.5;
const K_FACTOR              = 32;

let matchesDone  = 0;
let userName     = '';
let ratings      = {};
let images       = [];
let currentPair  = [];
let recentDeltas = [];

const dom = initDOM();
dom.btnA.addEventListener('click', () => handleChoice(0));
dom.btnB.addEventListener('click', () => handleChoice(1));

window.addEventListener('load', init);


async function init() {
  try {
    images = await loadImages();
    images.forEach(name => ratings[name] = 1500);

    userName = prompt('Enter your name:');
    if (!userName) return location.reload();

    showNext();
  } catch (e) {
    console.error(e);
    alert('Failed to initialize.');
  }
}

async function logVote(voteData) {
  try {
    await fetch('http://localhost:3000/api/votes', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(voteData)
    });
  } catch (err) {
    console.error('Error sending vote:', err);
  }
}


async function handleChoice(winnerIdx) {
  const winnerImg = currentPair[winnerIdx];
  const loserImg  = currentPair[1 - winnerIdx];
  const prevWinnerElo = ratings[winnerImg];
  const prevLoserElo  = ratings[loserImg];

  updateElo(winnerImg, loserImg, ratings, recentDeltas, K_FACTOR);
  const newWinnerElo = ratings[winnerImg];
  const newLoserElo  = ratings[loserImg];

  await logVote({
    userName,
    imageA:            currentPair[0],
    imageB:            currentPair[1],
    imageWinner:       winnerImg,
    imageLoser:        loserImg,
    eloWinnerPrevious: prevWinnerElo,
    eloWinnerNew:      newWinnerElo,
    eloLoserPrevious:  prevLoserElo,
    eloLoserNew:       newLoserElo
  });
  showNext();
}

function showNext() {
  if (!shouldContinue(recentDeltas, matchesDone, MAX_MOVES, CONVERGENCE_THRESHOLD)) {
    showFinal(dom, ratings);
  } else {
    currentPair = pickPair(images, ratings);
    showPair(dom, currentPair, matchesDone, MAX_MOVES);
    matchesDone++;
  }
}
