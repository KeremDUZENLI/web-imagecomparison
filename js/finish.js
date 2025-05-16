import { MIN_VOTES, TOPN } from './env/contants.js';
import { getRatings }      from './infrastructure/getRatings.js';
import { setText }         from './ui/setText.js';
import { showLeaderboard } from './ui/showLeaderboard.js';

function initFinishPage() {
  const username       = sessionStorage.getItem('surveyUser');
  const votes          = Number(sessionStorage.getItem('votesCount')) || 0;
  if (!username || votes < MIN_VOTES) return window.location.href = 'index.html';

  const btnRestart     = document.getElementById('btnRestart');
  const btnContinue    = document.getElementById('btnContinue');
  const btnToggleBoard = document.getElementById('btnViewBoard');
  const boardContainer = document.getElementById('containerBoard');

  setText('thankyou',  `Thank you, ${username}!`);
  setText('voteCount', `You made ${votes} choice${votes === 1 ? '' : 's'}.`);

  btnRestart.addEventListener('click', () => {
    sessionStorage.removeItem('surveyUser');
    sessionStorage.removeItem('votesUser');
    sessionStorage.removeItem('votesCount');
    window.location.href = 'index.html';
  });

  btnContinue.addEventListener('click', () => {
    window.location.href = 'compare.html';
  });

  boardContainer.style.display = 'none';
  btnToggleBoard.textContent   = 'View Leaderboard';

  btnToggleBoard.addEventListener('click', async () => {
    const isHidden = boardContainer.style.display === 'none';

    if (isHidden) {
      boardContainer.style.display = 'block';
      btnToggleBoard.textContent   = 'Hide Leaderboard';

      try {
        const topItems = await getRatings(TOPN);
        showLeaderboard(boardContainer, topItems);
      } catch (err) {
        boardContainer.textContent = 'Could not load leaderboard.';
        console.error(err);
      }
    } else {
      boardContainer.style.display = 'none';
      btnToggleBoard.textContent   = 'View Leaderboard';
    }
  });
}

window.addEventListener('DOMContentLoaded', initFinishPage);
