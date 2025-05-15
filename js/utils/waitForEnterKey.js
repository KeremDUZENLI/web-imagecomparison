export function waitForEnterKey() {
  return new Promise(resolve => {
    const handler = e => {
      if (e.key === 'Enter') {
        document.removeEventListener('keydown', handler);
        resolve();
      }
    };
    document.addEventListener('keydown', handler);
  });
}
