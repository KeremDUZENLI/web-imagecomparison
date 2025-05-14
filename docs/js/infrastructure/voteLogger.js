export async function logVote(data) {
  const res = await fetch('/api/votes', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data),
  });
  if (!res.ok) {
    const { error } = await res.json().catch(() => ({ error: 'Unknown' }));
    throw new Error(error);
  }
  return res.json();
}
