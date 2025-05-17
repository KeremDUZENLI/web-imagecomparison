export async function getUsernames(apiEndpoint = '/api/users') {
  const res = await fetch(apiEndpoint);
  if (!res.ok) throw new Error('Could not get usernames');
  return await res.json();
}
