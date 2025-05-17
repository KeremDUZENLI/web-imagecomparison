export async function postSurvey(data) {
  const res = await fetch('/api/surveys', {
    method: 'POST',
    headers: {'Content-Type':'application/json'},
    body: JSON.stringify(data)
  });
  if (!res.ok) throw new Error('Could not post survey');
  return await res.json();
}
