export class MatchSession {
  constructor(images, minVotes) {
    this.images = images;
    this.minVotes = minVotes;
    this.matchesDone = 0;
  }

  nextPair() {
    const count = this.images.length;
    const first = Math.floor(Math.random() * count);
    let second = Math.floor(Math.random() * count);
    while (second === first) {
      second = Math.floor(Math.random() * count);
    }
    return [this.images[first], this.images[second]];
  }

  applyVote() {
    this.matchesDone++;
  }

  isDone() {
    return this.matchesDone >= this.minVotes;
  }
}
