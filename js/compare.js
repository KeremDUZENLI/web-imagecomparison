import { MIN_VOTES } from "./env/constants.js";

import { MatchSession } from "./core/matchSession.js";
import { postVote } from "./infrastructure/postVote.js";
import { loadImages } from "./ui/loadImages.js";
import { showPair } from "./ui/showPair.js";

let compareDOM = {};
let currentUser, images, session;

const goTo = (url) => (location.href = url);
const getSession = (key) => sessionStorage.getItem(key);
const setSession = (key, value) => sessionStorage.setItem(key, value);
const getJSON = (key) => JSON.parse(sessionStorage.getItem(key));
const setJSON = (key, value) =>
  sessionStorage.setItem(key, JSON.stringify(value));

const persistPair = (pair) => {
  setJSON("currentPair", pair);
  setJSON("lastPair", pair);
};
const restorePair = () => {
  const pair = getJSON("currentPair");
  if (!pair) return null;

  session.currentPair = pair;
  session.lastPair = getJSON("lastPair") || pair;
  compareDOM.btnFinish.style.display = session.canFinish()
    ? "inline-block"
    : "none";

  showPair(compareDOM, pair, session.matchesDone, MIN_VOTES);
  return true;
};

function initCompare() {
  compareDOM.imgA = document.getElementById("img_a");
  compareDOM.imgB = document.getElementById("img_b");
  compareDOM.btnA = document.getElementById("btn_a");
  compareDOM.btnB = document.getElementById("btn_b");
  compareDOM.btnFinish = document.getElementById("btn_finish");
  compareDOM.progress = document.getElementById("progress");

  compareDOM.btnA.addEventListener("click", () => handleChoice(0));
  compareDOM.btnB.addEventListener("click", () => handleChoice(1));
  compareDOM.btnFinish.addEventListener("click", finishSession);

  window.addEventListener("beforeunload", unloadHandler);
  window.addEventListener("pageshow", pageshowHandler);
}

async function handleChoice(idx) {
  const [winner, loser] = [
    session.currentPair[idx],
    session.currentPair[1 - idx],
  ];

  try {
    await postVote({
      username: currentUser,
      image_winner: winner,
      image_loser: loser,
    });
  } catch {
    alert("Could not save vote");
    return;
  }

  session.applyVote();
  setSession("votesCount", session.matchesDone);
  sessionStorage.removeItem("currentPair");
  loadNextPair();
}

function finishSession() {
  sessionStorage.removeItem("currentPair");
  setSession("votesCount", session.matchesDone);
  goTo("finish.html");
}

function unloadHandler(e) {
  if (session && !session.canFinish()) {
    const warning =
      "You have unfinished votes — are you sure you want to leave?";
    e.preventDefault();
    e.returnValue = warning;
    return warning;
  }
}

function pageshowHandler(event) {
  const navType = performance.getEntriesByType("navigation")[0]?.type;
  if (event.persisted || navType === "back_forward") {
    renderCompare();
  }
}

async function renderCompare() {
  currentUser = getSession("surveyUser");
  if (!currentUser) return goTo("index.html");

  if (!images) {
    images = await loadImages();
    session = new MatchSession(images, MIN_VOTES);
  }

  session.matchesDone = initVoteState();

  if (!restorePair()) {
    loadNextPair();
  }
}

function initVoteState() {
  const lastUser = getSession("votesUser");
  if (lastUser !== currentUser) {
    setSession("votesCount", "0");
    setSession("votesUser", currentUser);
  }
  const count = Number(getSession("votesCount"));
  return Number.isNaN(count) ? 0 : count;
}

function loadNextPair() {
  const done = session.matchesDone;
  const canFinish = session.canFinish();

  compareDOM.btnFinish.style.display = canFinish ? "inline-block" : "none";
  compareDOM.progress.textContent = canFinish
    ? `You've reached ${MIN_VOTES} votes — you may Finish or keep voting.`
    : `Match ${done + 1} of ${MIN_VOTES}`;

  const pair = session.nextPair();
  session.lastPair = pair;
  persistPair(pair);
  showPair(compareDOM, pair, done, MIN_VOTES);
}

window.addEventListener("DOMContentLoaded", async () => {
  initCompare();
  await renderCompare();
});
