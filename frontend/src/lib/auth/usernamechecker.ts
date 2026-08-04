import { PUBLIC_BACKEND } from "$env/static/public";

export default async function checkUsernameAvailability(username: string): Promise<number> {
  try {
    const response = await fetch(
      `${PUBLIC_BACKEND}/checkusername?username=${encodeURIComponent(username)}`,
      {
        method: "GET",
        headers: { "Content-Type": "application/json" }
      }
    );
    console.log("code", response.status)
      
    //0 is ok, 1 is taken, 2 is error
    if (response.status === 409) return 1; // Username choice has a conflict (already taken)
    if (response.status !== 200) return 2; // Error in the request, treat as unavailable
    const data = await response.json();
    return data.available === true ? 0 : 2;
  } catch {
    return 3;
  }
}
