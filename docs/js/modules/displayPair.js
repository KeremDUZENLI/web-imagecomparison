const IMAGES_FOLDER = 'images/';

export function showPair(dom, pair, current, total) {
  dom.imgA.src = IMAGES_FOLDER + pair[0];
  dom.imgB.src = IMAGES_FOLDER + pair[1];
  dom.progress.textContent = `Match ${current + 1} of ${total}`;
}
