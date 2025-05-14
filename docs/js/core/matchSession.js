export class MatchSession {
  constructor(images, minVotes, defaultRating, initialRatings = {}) {
    this.images = images;
    this.minVotes = minVotes;
    this.matchesDone = 0;
    this.ratings = Object.fromEntries(
      images.map(img => [img, initialRatings[img] ?? defaultRating])
    );
  }

  nextPair() {
    return pickPair(this.images, this.ratings);
  }

  applyVote() {
    this.matchesDone++;
  }

  isDone() {
    return this.matchesDone >= this.minVotes;
  }

  getRatings() {
    return this.ratings;
  }
}

export function pickPair(images, ratings) {
  const sorted = [...images].sort((a, b) => ratings[a] - ratings[b]);
  const i = Math.floor(Math.random() * (sorted.length - 1));
  return [sorted[i], sorted[i + 1]];
}
