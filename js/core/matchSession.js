export class MatchSession {
  constructor(images, minVotes) {
    this.images      = images;
    this.minVotes    = minVotes;
    this.matchesDone = 0;
    this.currentPair = null;
  }

  nextPair() {
    const count = this.images.length;
    let first   = Math.floor(Math.random() * count);
    let second  = Math.floor(Math.random() * count);

    while (second === first) {
      second = Math.floor(Math.random() * count);
    }

    this.currentPair = [this.images[first], this.images[second]];
    return this.currentPair;
  }

  applyVote() {
    this.matchesDone++;
  }

  canFinish() {
    return this.matchesDone >= this.minVotes;
  }
}
