export async function getRatings(n = 10, apiEndpoint = '/api/ratings') {
  const res = await fetch(apiEndpoint);
  if (!res.ok) {
    throw new Error('Could not load rankings');
  }
  const ratings = await res.json();
  return ratings.sort((a, b) => b.elo - a.elo).slice(0, n);
}
