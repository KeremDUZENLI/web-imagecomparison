export async function getRatings(topN = 10, apiEndpoint = '/api/ratings') {
  const res = await fetch(apiEndpoint);
  if (!res.ok) throw new Error('Could not get rankings');
  const ratings = await res.json();
  return ratings.sort((a, b) => b.elo - a.elo).slice(0, topN);
}
