<script lang="ts">
	import Navbar from "$lib/components/navbar.svelte";
	import checkUsernameAvailability from "$lib/auth/usernamechecker";
	import sendSignupRequest from "$lib/auth/signup";

	// Logic for username status
	let username = $state("");
 	let usernameStatus = $state();
	let signupStatus = $state();
	let isChecking = $state(false);

	async function handleSubmit(event: SubmitEvent) {
		event.preventDefault();
		const form = event.currentTarget as HTMLFormElement;
		const formData = new FormData(form);

		const email = String(formData.get("email"));
		const username = String(formData.get("username"));
		const password = String(formData.get("password"));

		isChecking = true;
		const availability = await checkUsernameAvailability(username);
		isChecking = false;

		const signupResult = await sendSignupRequest(email, username, password);

		if (availability === 0) {
			if (signupResult === 0) {
				window.location.href = "/signin";
			} else if (signupResult === 1) {
				signupStatus = "Username or email already exists.";
			} else if (signupResult === 2) {
				signupStatus = "Invalid input.";
			} else if (signupResult === 3) {
				signupStatus = "Error occurred while signing up.";
			}
		} else if (availability === 1) {
			usernameStatus = "Username is already taken.";
		} else {
			usernameStatus = "Error checking username availability.";
		}
	}

</script>

<Navbar />

<div class="flex w-[clamp(20rem,40vw,50rem)] items-center justify-center p-8 bg-primary place-self-center rounded-lg">
  <h1 class="text-3xl font-bold text-white display-block text-nowrap m-4">Sign Up</h1>
  <form onsubmit={handleSubmit} class="flex flex-col gap-4 w-full">
	<input
	  type="text"
	  name="email"
	  placeholder="Email"
	  required
	  class="bg-primary text-white border border-white rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-emerald-500 focus:border-transparent"
	/>
	<input
	  type="text"
	  name="username"
	  placeholder="Username"
	  required
	  class="bg-primary text-white border border-white rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-emerald-500 focus:border-transparent"
	/>
	<p class="text-sm text-white">Only letters, numbers, and underscores are allowed. Must be between 3 and 20 characters. </p>
	{#if isChecking }
	  <p class="text-sm text-white">Checking username availability...</p>
	{/if}
	<p>{usernameStatus}</p>
	<input
	  type="password"
	  name="password"
	  placeholder="Password"
	  required
	  class="bg-primary text-white border border-white rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-emerald-500 focus:border-transparent"
	/>
	<p class="text-sm text-white">Please make it at least 8 characters long, with capital letters, numbers, and special characters. But I'm only a sign and have not enforced it so do whatever.</p>
	<button
	  type="submit"
	  class="bg-emerald-700 text-white font-semibold py-2 px-4 rounded-md hover:bg-emerald-600 transition-colors duration-300"
	>
	  Sign In
	</button>
	<p class="text-red-500">{signupStatus}</p>
  </form>
</div>
