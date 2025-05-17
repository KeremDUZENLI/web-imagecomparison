export async function loadImages() {
  const res = await fetch('./images/images.json');
  if (!res.ok) throw new Error('Could not load image list');
  return res.json();
}
