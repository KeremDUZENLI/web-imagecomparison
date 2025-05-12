export function pickPair(images, ratings) {
  const sorted = [...images].sort((a, b) => ratings[a] - ratings[b]);
  const i = Math.floor(Math.random() * (sorted.length - 1));
  return [sorted[i], sorted[i + 1]];
}

export function updateElo(winner, loser, ratings, recentDeltas, kFactor) {
  const Ra = ratings[winner], Rb = ratings[loser];
  const Ea = 1 / (1 + 10 ** ((Rb - Ra) / 400));
  const delta = kFactor * (1 - Ea);

  ratings[winner] = Math.round(ratings[winner] + delta);
  ratings[loser]  = Math.round(ratings[loser] - delta);

  recentDeltas.push(Math.abs(delta) * 2);
  if (recentDeltas.length > 50) recentDeltas.shift();
}

export function shouldContinue(recentDeltas, matchesDone, maxMoves, convergenceThreshold) {
  if (matchesDone >= maxMoves) return false;
  if (recentDeltas.length === 0) return true;
  const avgDelta = recentDeltas.reduce((sum, d) => sum + d, 0) / recentDeltas.length;
  return avgDelta >= convergenceThreshold;
}
