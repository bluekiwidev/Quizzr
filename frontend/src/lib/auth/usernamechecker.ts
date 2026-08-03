import { PUBLIC_BACKEND } from "$env/static/public";

export default async function checkUsernameAvailability(username: string): Promise<boolean> {
  try {
    const response = await fetch(
      `${PUBLIC_BACKEND}/checkusername?username=${encodeURIComponent(username)}`,
      {
        method: "GET",
        headers: { "Content-Type": "application/json" }
      }
    );

    if (response.status === 409) return false; // Username choice has a conflict (already taken)
    if (response.status !== 200) return false; // Error in the request, treat as unavailable
    const data = await response.json();
    return data.available === true;
  } catch {
    return false;
  }
}