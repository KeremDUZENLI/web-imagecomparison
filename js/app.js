const MAX_MOVES = 10;
const CONVERGENCE_THRESHOLD = 0.5;
const K = 32;

const imgA = document.getElementById('imgA');
const imgB = document.getElementById('imgB');
const progress = document.getElementById('progress');
const resultsDiv = document.getElementById('results');
const namePrompt = document.getElementById('namePrompt');

const ratings = {};
const recentDeltas = [];

let images = [];
let currentPair = [];
let userName = null;
let matchesDone = 0;
images.forEach(f => ratings[f] = 1500);

async function loadImages() {
  const response = await fetch('images.json');
  if (!response.ok) throw new Error('Failed to load image list');
  return await response.json();
}

window.addEventListener('load', async () => {
  try {
    images = await loadImages();
    images.forEach(f => ratings[f] = 1500);

    userName = prompt('Please enter your name:');
    if (!userName) {
      alert('Name is required to proceed');
      window.location.reload();
    } else {
      showPair();
    }
  } catch (error) {
    alert('Failed to load images.');
    console.error(error);
  }
});

function showPair() {
  currentPair = pickPair();
  imgA.src = `images/${currentPair[0]}`;
  imgB.src = `images/${currentPair[1]}`;
  progress.textContent = `Match ${matchesDone + 1} of ${MAX_MOVES}`;
}
function pickPair() {
  const sortedImgs = [...images].sort((a, b) => ratings[a] - ratings[b]);
  const i = Math.floor(Math.random() * (sortedImgs.length - 1));
  return [ sortedImgs[i], sortedImgs[i + 1] ];
}

function updateElo(winner, loser) {
  const Ra = ratings[winner];
  const Rb = ratings[loser];
  const Ea = 1 / (1 + 10 ** ((Rb - Ra) / 400));
  const Eb = 1 - Ea;

  const deltaA = K * (1 - Ea);
  const deltaB = K * (0 - Eb);

  ratings[winner] += deltaA;
  ratings[loser]  += deltaB;

  recentDeltas.push(Math.abs(deltaA) + Math.abs(deltaB));
  if (recentDeltas.length > 50) recentDeltas.shift();
}

function afterEachMatch() {
  matchesDone++;
  const avgDelta = recentDeltas.reduce((s, x) => s + x, 0) / recentDeltas.length;
  if (matchesDone >= MAX_MOVES || avgDelta < CONVERGENCE_THRESHOLD) {
    showFinalRanking();
  } else {
    showPair();
  }
}

function showFinalRanking() {
  document.getElementById('imageContainer').style.display = 'none';
  progress.style.display = 'none';
  const sorted = Object.entries(ratings).sort(([, a], [, b]) => b - a);
  const title = document.createElement('h2');
  title.textContent = 'Final Ranking (Highest ELO First)';
  resultsDiv.appendChild(title);

  sorted.forEach(([file, score]) => {
    const img = document.createElement('img');
    img.src = `images/${file}`;
    img.alt = file;
    const label = document.createElement('div');
    label.textContent = `${file} — ${score.toFixed(1)} ELO`;
    const container = document.createElement('div');
    container.appendChild(img);
    container.appendChild(label);
    resultsDiv.appendChild(container);
  });
}

const buttonChooseA = document.getElementById('buttonChooseA');
const buttonChooseB = document.getElementById('buttonChooseB');
buttonChooseA.addEventListener('click', () => {updateElo(currentPair[0], currentPair[1]); afterEachMatch()});
buttonChooseB.addEventListener('click', () => {updateElo(currentPair[1], currentPair[0]); afterEachMatch()});

// function sendResult(winner) {
//   const payload = {
//     user: userName,
//     imgA: images[currentPair.a],
//     imgB: images[currentPair.b],
//     choice: winner,
//     timestamp: new Date().toISOString()
//   };
//   fetch('https://formspree.io/f/YOUR_FORM_ID', {
//     method: 'POST',
//     headers: { 'Content-Type': 'application/json' },
//     body: JSON.stringify(payload)
//   }).catch(err => console.error('Submission error', err));
// }
