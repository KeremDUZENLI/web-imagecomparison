export async function getRatings(topN = null) {
  const res = await fetch('/api/ratings');
  if (!res.ok) {
    throw new Error('Could not get rankings');
  }

  const ratings = await res.json();
  const sorted = ratings.sort((a, b) => b.elo - a.elo);

  return (typeof topN === 'number')
    ? sorted.slice(0, topN)
    : sorted;
}
