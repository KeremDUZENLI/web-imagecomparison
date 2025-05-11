const IMAGES_FOLDER = 'images/';


export function initDOM() {
  return {
    imgA:       document.getElementById('imgA'),
    imgB:       document.getElementById('imgB'),
    btnA:       document.getElementById('buttonChooseA'),
    btnB:       document.getElementById('buttonChooseB'),
    progress:   document.getElementById('progress'),
    results:    document.getElementById('results'),
    container:  document.getElementById('imageContainer'),
  };
}

export function showPair(dom, pair, matchesDone, maxMoves) {
  dom.imgA.src = IMAGES_FOLDER + pair[0];
  dom.imgB.src = IMAGES_FOLDER + pair[1];
  dom.progress.textContent = `Match ${matchesDone + 1} of ${maxMoves}`;
}

export function showFinal(dom, ratings) {
  dom.container.style.display = 'none';
  dom.progress.style.display = 'none';

  const sorted = Object.entries(ratings)
    .sort(([, r1], [, r2]) => r2 - r1);

  dom.results.innerHTML = '<h2>Final Ranking</h2>' +
    sorted.map(([file, score]) =>
      `<div class="result-row">
         <img src="${IMAGES_FOLDER + file}" alt="${file}" />
         <div>${file} — ${score.toFixed(1)} ELO</div>
       </div>`
    ).join('');
}
