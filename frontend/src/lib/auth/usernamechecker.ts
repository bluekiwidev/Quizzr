export default async function checkUsernameAvailability(username: string) {
		const response = await fetch(`http://localhost:3000/checkusername?username=${encodeURIComponent(username)}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        })
        const data = await response.json();
        if (data.available) {
            return "Username is available";
        } else {
            return "Username is not available";
        }
}