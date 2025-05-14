export async function loadRatings() {
  const res = await fetch('/api/ratings');
  if (!res.ok) throw new Error('Failed to load ratings');
  return res.json();
}
