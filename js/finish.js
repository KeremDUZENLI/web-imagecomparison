const username = sessionStorage.getItem('surveyUser') || 'Guest';
const votes    = sessionStorage.getItem('votesCount') || '0';

document.getElementById('thankyou').textContent = `Thank you, ${username}!`;
document.getElementById('voteCount').textContent = `You made ${votes} choice${votes === '1' ? '' : 's'}.`;

document.getElementById('btnRestart').onclick = () => {
  sessionStorage.removeItem('surveyUser');
  sessionStorage.removeItem('votesUser');
  sessionStorage.removeItem('votesCount');
  location.href = 'index.html';
};

document.getElementById('btnContinue').onclick = () => {
  location.href = 'compare.html';
};

document.getElementById('btnLeaderboard').onclick = () => {
  location.href = '/api/ratings';
};
