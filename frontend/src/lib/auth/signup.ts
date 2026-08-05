import { PUBLIC_BACKEND } from "$env/static/public";

export default async function sendSignupRequest(email: string, username: string, password: string): Promise<number> {
    try {
        console.log("Sending signup request to backend...");
        const response = await fetch(
          `${PUBLIC_BACKEND}/submitsignup`,
          {
            method: "POST",
            body: JSON.stringify(
                { Email: email, Username: username, Password: password }),
          }
        );

        // Return codes:
        if (response.status === 201) return 0; //All good
        if (response.status === 409) return 1; //Username or email already exists
        if (response.status === 400) return 2; //Invalid input
        if (response.status === 500) return 3; //Server error
        return 4;
      } catch {
        return 4;
      }
}