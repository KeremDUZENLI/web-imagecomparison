import { MIN_VOTES, TOPN } from './env/constants.js';

import { getRatings }      from './infrastructure/getRatings.js';
import { setText }         from './ui/setText.js';
import { showLeaderboard } from './ui/showLeaderboard.js';

let cachedLeaderboard = null;

const CACHE_LBOARD = 'leaderboardCache';
const CACHE_VOTES  = 'leaderboardCacheVotes';

const goTo = url => window.location.href = url;
const readSession   = key => sessionStorage.getItem(key);
const writeSession  = (key, value) => sessionStorage.setItem(key, value);
const clearSession  = (...keys) => keys.forEach(k => sessionStorage.removeItem(k));

function initFinishPage() {
  const username = readSession('surveyUser');
  const votes    = Number(readSession('votesCount')) || 0;
  if (!username || votes < MIN_VOTES) {
    return goTo('index.html');
  }

  rehydrateLeaderboard(votes);

  const btnRestart     = document.getElementById('btn_restart');
  const btnContinue    = document.getElementById('btn_continue');
  const btnToggleBoard = document.getElementById('btn_view_board');
  const boardContainer = document.getElementById('container_board');

  setText('thanks',  `Thank you, ${username}!`);
  setText('vote_count', `You made ${votes} choice${votes === 1 ? '' : 's'}`);

  btnRestart.addEventListener('click', () => {
    clearSession('surveyUser', 'votesUser', 'votesCount', CACHE_LBOARD, CACHE_VOTES);
    cachedLeaderboard = null;
    goTo('index.html');
  });

  btnContinue.addEventListener('click', () => goTo('compare.html'));
  boardContainer.style.display = 'none';
  btnToggleBoard.textContent   = 'View Leaderboard';
  btnToggleBoard.addEventListener('click', () => toggleLeaderboard(boardContainer, btnToggleBoard, votes));
}

function rehydrateLeaderboard(votes) {
  const storedVotes = Number(readSession(CACHE_VOTES)) || 0;
  const rawCache    = readSession(CACHE_LBOARD);

  if (rawCache && storedVotes === votes) {
    try {
      cachedLeaderboard = JSON.parse(rawCache);
    } catch {
      clearSession(CACHE_LBOARD, CACHE_VOTES);
    }
  } else {
    clearSession(CACHE_LBOARD, CACHE_VOTES);
  }
}

async function toggleLeaderboard(boardEl, toggleBtn, votes) {
  const isHidden = boardEl.style.display === 'none';
  boardEl.style.display = isHidden ? 'block' : 'none';
  toggleBtn.textContent = isHidden ? 'Hide Leaderboard' : 'View Leaderboard';

  if (isHidden) {
    if (!cachedLeaderboard) {
      rehydrateLeaderboard(votes);

      if (!cachedLeaderboard) {
        try {
          cachedLeaderboard = await getRatings(TOPN);
          writeSession(CACHE_LBOARD, JSON.stringify(cachedLeaderboard));
          writeSession(CACHE_VOTES, votes);
        } catch {
          boardEl.textContent = 'Could not load leaderboard.';
          return;
        }
      }
    }
    showLeaderboard(boardEl, cachedLeaderboard);
  }
}

window.addEventListener('DOMContentLoaded', initFinishPage);
