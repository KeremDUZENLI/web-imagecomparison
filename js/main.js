import { pickPair, updateElo, shouldContinue } from './elo-calculator.js';
import { loadImages }                          from './image-loader.js';
import { initDOM, showPair, showFinal }        from './init-dom.js';


const MAX_MOVES             = 10;
const CONVERGENCE_THRESHOLD = 0.5;
const K_FACTOR              = 32;

let ratings      = {};
let images       = [];
let currentPair  = [];
let recentDeltas = [];
let matchesDone  = 0;

const dom = initDOM();
dom.btnA.addEventListener('click', () => handleChoice(0));
dom.btnB.addEventListener('click', () => handleChoice(1));

window.addEventListener('load', init);


async function init() {
  try {
    images = await loadImages();
    images.forEach(name => ratings[name] = 1500);

    const userName = prompt('Enter your name:');
    if (!userName) return location.reload();

    showNext();
  } catch (e) {
    console.error(e);
    alert('Failed to initialize.');
  }
}

function handleChoice(winnerIdx) {
  const loserIdx = 1 - winnerIdx;
  updateElo(currentPair[winnerIdx], currentPair[loserIdx], ratings, recentDeltas, K_FACTOR);
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