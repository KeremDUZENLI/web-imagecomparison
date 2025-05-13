const IMAGES_FOLDER = '../images/';

export function showFinal(dom, ratings) {
  dom.container.style.display = 'none';
  dom.progress.style.display = 'none';

  const sorted = Object.entries(ratings).sort(([, a], [, b]) => b - a);
  dom.results.innerHTML = '<h2>Final Ranking</h2>' + sorted.map(([img, score]) =>
    `<div class="result-row">
      <img src="${IMAGES_FOLDER + img}" alt="${img}" />
      <div>${img} — ${score} ELO</div>
    </div>`
  ).join('');
}
