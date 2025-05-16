export async function fetchUsernames(apiEndpoint = '/api/users') {
  const res = await fetch(apiEndpoint);
  if (!res.ok) throw new Error('Could not load usernames');
  return await res.json();
}
