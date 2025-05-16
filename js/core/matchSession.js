export class MatchSession {
  constructor(images, minVotes) {
    this.images      = images;
    this.minVotes    = minVotes;
    this.matchesDone = 0;
    this.currentPair = null;
    this.lastPair    = [];
  }

  nextPair() {
    const count = this.images.length;
    let first, second;

    do {
      first = Math.floor(Math.random() * count);
      do {
        second = Math.floor(Math.random() * count);
      } while (second === first);
    } while (
      this.lastPair.includes(this.images[first]) ||
      this.lastPair.includes(this.images[second])
    );

    this.currentPair = [this.images[first], this.images[second]];
    this.lastPair = this.currentPair;
    return this.currentPair;
  }

  applyVote() {
    this.matchesDone++;
  }

  canFinish() {
    return this.matchesDone >= this.minVotes;
  }
}
