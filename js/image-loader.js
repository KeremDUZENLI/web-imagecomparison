const IMAGES_JSON_PATH = 'images.json';


export async function loadImages() {
  const res = await fetch(IMAGES_JSON_PATH);
  if (!res.ok) throw new Error('Failed to load images.json');
  return res.json();
}
