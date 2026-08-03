<script lang="ts">
	import Navbar from "$lib/components/navbar.svelte";
	import checkUsernameAvailability from "$lib/auth/usernamechecker";
	import sendSignupRequest from "$lib/auth/signup";


	// Logic for username status
	let usernameAvailable = $state("")

	function displayUsernameStatus(username: string) {
		const status = checkUsernameAvailability(username);
		usernameAvailable = status;
	}
	
</script>

<Navbar />

<div class="flex w-[clamp(20rem,30vw,45rem)] items-center justify-center p-8 bg-primary place-self-center rounded-lg">
  <h1 class="text-3xl font-bold text-white display-block">Sign In</h1>
  <form method="POST" class="flex flex-col gap-4 w-full">
	<input
	  type="text"
	  name="email"
	  placeholder="Email"
	  class="bg-primary text-white border border-white rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-emerald-500 focus:border-transparent"
	/>
	<input
	  type="text"
	  name="username"
	  placeholder="Username"
	  class="bg-primary text-white border border-white rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-emerald-500 focus:border-transparent"
	  oninput={(e) => displayUsernameStatus(e.target.value)}
	/>
	<input
	  type="password"
	  name="password"
	  placeholder="Password"
	  class="bg-primary text-white border border-white rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-emerald-500 focus:border-transparent"
	/>
	<button
	  type="submit"
	  class="bg-emerald-700 text-white font-semibold py-2 px-4 rounded-md hover:bg-emerald-600 transition-colors duration-300"
	>
	  Sign In
	</button>
  </form>
</div>
