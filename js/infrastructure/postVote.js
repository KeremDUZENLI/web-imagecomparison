export async function postVote(data) {
  const res = await fetch('/api/votes', {
    method: 'POST',
    headers: {'Content-Type':'application/json'},
    body: JSON.stringify(data)
  });
  if (!res.ok) throw new Error('Could not post vote');
  return await res.json();
}
