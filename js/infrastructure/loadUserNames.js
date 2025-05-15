export async function loadUserNames(apiEndpoint = '/api/users') {
  let existing = [];
  try {
    const res = await fetch(apiEndpoint);
    if (!res.ok) throw new Error();
    existing = (await res.json()).map(u => u.toLowerCase());
  } catch {
    alert('Could not load existing usernames');
    return null;
  }

  while (true) {
    const raw = prompt('Enter a unique name:');
    if (!raw?.trim()) {
      alert('Name is required.');
      continue;
    }
    const name = raw.trim();
    if (existing.includes(name.toLowerCase())) {
      alert('That name is already taken');
      continue;
    }
    existing.push(name.toLowerCase());
    return name;
  }
}
