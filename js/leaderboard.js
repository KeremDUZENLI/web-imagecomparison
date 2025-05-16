async function loadLeaderboard() {
  try {
    const res = await fetch('/api/ratings');
    if (!res.ok) throw new Error('Failed to load leaderboard');

    const ratings = await res.json();
    const top10 = ratings.sort((a, b) => b.elo - a.elo).slice(0, 10);

    const container = document.getElementById('leaderboard');
    container.innerHTML = '';

    top10.forEach((item, idx) => {
      const div = document.createElement('div');
      div.className = 'leaderboard-item';
      div.innerHTML = `
        <div class="rank">#${idx + 1}</div>
        <img src="images/${item.image}" alt="Image ${item.image}">
        <div class="elo">${item.elo} Elo</div>
      `;
      container.appendChild(div);
    });
  } catch (err) {
    console.error(err);
    document.getElementById('leaderboard').textContent =
      'Could not load leaderboard.';
  }
}

window.addEventListener('load', loadLeaderboard);
