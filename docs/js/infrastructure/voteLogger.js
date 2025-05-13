export async function logVote(data) {
  try {
    console.log('Sending vote:', data);
    await fetch('/api/votes', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data)
    });
  } catch (err) {
    console.error('Vote logging failed:', err);
  }
}
