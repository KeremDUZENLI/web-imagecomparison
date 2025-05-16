import { fetchRatings }    from './infrastructure/fetchRatings.js';
import { setText }         from './ui/setText.js';
import { showLeaderboard } from './ui/showLeaderboard.js';


window.addEventListener('DOMContentLoaded', () => {
  const username       = sessionStorage.getItem('surveyUser') || 'Guest';
  const votes          = sessionStorage.getItem('votesCount')   || '0';
  const btnViewBoard   = document.getElementById('btnViewBoard');
  const containerBoard = document.getElementById('containerBoard');
  const TOPN           = 10;

  setText('thankyou',   `Thank you, ${username}!`);
  setText('voteCount',  `You made ${votes} choice${votes === '1' ? '' : 's'}.`);

  document.getElementById('btnRestart').onclick = () => {
    sessionStorage.clear();
    location.href = 'index.html';
  };

  document.getElementById('btnContinue').onclick = () => {
    location.href = 'compare.html';
  };

  containerBoard.style.display = 'none';
  btnViewBoard.onclick = async () => {
    const isHidden = containerBoard.style.display === 'none';

    if (isHidden) {
      containerBoard.style.display = 'block';
      btnViewBoard.textContent = 'Hide Leaderboard';

      try {
        const topN = await fetchRatings(TOPN);
        showLeaderboard(containerBoard, topN);
      } catch {
        containerBoard.textContent = 'Could not load Leaderboard';
      }
    } else {
      containerBoard.style.display = 'none';
      btnViewBoard.textContent = 'View Leaderboard';
    }
  };
});
