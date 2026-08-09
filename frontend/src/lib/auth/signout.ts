import { PUBLIC_BACKEND } from "$env/static/public";

export default async function sendSignoutRequest(): Promise<number> {
    try {
        console.log("Sending signout request to backend...");
        const response = await fetch(
          `${PUBLIC_BACKEND}/submitsignout`,
          {
            method: "POST",
            credentials: 'include'
          }
        );

        // Return codes:
        if (response.status === 204) return 0; //All good
        if (response.status === 500) return 1; //Server error
        return response.status;
      } catch {
        return 2;
      }
}