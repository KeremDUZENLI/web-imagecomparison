const IMAGE_JSON_PATH = '../images.json';

export async function loadImages() {
  const res = await fetch(IMAGE_JSON_PATH);
  if (!res.ok) throw new Error('Failed to load image list.');
  return res.json();
}
