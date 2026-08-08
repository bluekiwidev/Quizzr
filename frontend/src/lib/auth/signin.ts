import { PUBLIC_BACKEND } from "$env/static/public";

export default async function sendSigninRequest(email: string, password: string): Promise<number> {
    try {
        console.log("Sending signin request to backend...");
        const response = await fetch(
          `${PUBLIC_BACKEND}/submitsignin`,
          {
            method: "POST",
            body: JSON.stringify(
                { Email: email, Password: password }),
            credentials: 'include'
          }
        );

        // Return codes:
        if (response.status === 200) return 0; //All good
        if (response.status === 404) return 1; //Error
        if (response.status === 500) return 3; //Server error
        return response.status;
      } catch {
        return 4;
      }
}