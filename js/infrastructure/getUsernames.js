export async function getUsernames() {
  const res = await fetch('/api/users');
  if (!res.ok) throw new Error('Could not get usernames');
  return await res.json();
}
