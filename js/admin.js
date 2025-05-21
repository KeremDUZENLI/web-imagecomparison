import { getRatings } from './infrastructure/getRatings.js';
import { showLeaderboard } from './ui/showLeaderboard.js';

async function initLeaderboard() {
  const container = document.getElementById('leaderboard-container');

  try {
    const ratings = await getRatings();
    showLeaderboard(container, ratings);
  } catch (err) {
    container.textContent = 'Could not load leaderboard.';
    return
  }
}

window.addEventListener('DOMContentLoaded', initLeaderboard);
