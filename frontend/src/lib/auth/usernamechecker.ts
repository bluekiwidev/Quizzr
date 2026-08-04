import { PUBLIC_BACKEND } from "$env/static/public";

export default async function checkUsernameAvailability(username: string): Promise<number> {
  try {
    const response = await fetch(
      `${PUBLIC_BACKEND}/checkusername?username=${encodeURIComponent(username)}`,
      {
        method: "GET",
      }
    );
    //0 is ok, 1 is taken, 2 is error
    if (response.status === 200) return 0; // Username choice is available
    if (response.status === 409) return 1; // Username choice has a conflict (already taken)
    return 2;
  } catch {
    return 2; // Error occurred while checking username availability
  }
}
