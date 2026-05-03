export async function loadImages() {
  const url = `./images/images.json?t=${new Date().getTime()}`;
  const res = await fetch(url);

  if (!res.ok) throw new Error("Could not load image list");
  return res.json();
}
