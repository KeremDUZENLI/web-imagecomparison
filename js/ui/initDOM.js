export function initDOM() {
  return {
    imgA:      document.getElementById('imgA'),
    imgB:      document.getElementById('imgB'),
    btnA:      document.getElementById('buttonChooseA'),
    btnB:      document.getElementById('buttonChooseB'),
    btnFinish: document.getElementById('buttonFinish'),
    progress:  document.getElementById('progress'),
    results:   document.getElementById('results'),
    container: document.getElementById('imageContainer'),
  };
}
