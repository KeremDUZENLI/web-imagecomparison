export async function fetchExistingNames(apiEndpoint = '/api/users') {
  const res = await fetch(apiEndpoint);
  if (!res.ok) throw new Error('Could not load existing usernames');
  return await res.json();
}

export function validateName(name, existing) {
  const candidate = name.trim().toLowerCase();
  return candidate !== '' && !existing.includes(candidate);
}
