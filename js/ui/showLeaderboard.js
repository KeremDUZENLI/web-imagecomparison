export function showLeaderboard(container, items) {
  container.innerHTML = '';
  items.forEach((item, i) => {
    const div = document.createElement('div');
    div.className = 'leaderboard-item-inline';
    div.innerHTML = `
      <div>#${i + 1}</div>
      <img src="images/${item.image}" alt="${item.image}">
      <div>${item.image}</div>
      <div>${item.elo} ELO</div>
    `;
    container.appendChild(div);
  });
}
