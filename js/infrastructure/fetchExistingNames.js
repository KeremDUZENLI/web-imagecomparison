export async function fetchExistingNames(apiEndpoint = '/api/users') {
  try {
    const res = await fetch(apiEndpoint);
    if (!res.ok) throw new Error();
    return (await res.json()).map(u => u.toLowerCase());
  } catch {
    alert('Could not load existing usernames');
    return [];
  }
}

export function validateName(name, existing) {
  return name.trim() !== '' && !existing.includes(name.toLowerCase());
}
