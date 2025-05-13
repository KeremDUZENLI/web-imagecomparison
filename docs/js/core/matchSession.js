import { pickPair, updateElo, shouldContinue } from './eloCalculator.js';

export class MatchSession {
  constructor(images, kFactor, maxMoves, convergenceThreshold, initialRatings = {}) {
    this.images = images;
    this.kFactor = kFactor;
    this.maxMoves = maxMoves;
    this.convergenceThreshold = convergenceThreshold;
    this.ratings = { ...initialRatings };
    this.recentDeltas = [];
    this.matchesDone = 0;

    images.forEach(img => {
      if (this.ratings[img] === undefined) {
        this.ratings[img] = 1500;
      }
    });
  }

  nextPair() {
    return pickPair(this.images, this.ratings);
  }

  applyVote(winner, loser) {
    updateElo(winner, loser, this.ratings, this.recentDeltas, this.kFactor);
    this.matchesDone++;
  }

  isDone() {
    return !shouldContinue(
      this.recentDeltas,
      this.matchesDone,
      this.maxMoves,
      this.convergenceThreshold
    );
  }

  getRatings() {
    return this.ratings;
  }
}
