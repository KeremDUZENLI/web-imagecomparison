export async function loadImages() {
  const res = await fetch('/images.json');
  if (!res.ok) throw new Error('Failed to load image list');
  return res.json();
}
