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

<main class="px-4 pb-16 pt-8 sm:px-6 lg:px-8">
	<section class="mx-auto grid max-w-6xl gap-8 lg:grid-cols-[.98fr_1.02fr]">
		
		<div class="rounded-xl border border-border bg-card/90 p-6 shadow-[0_24px_60px_-36px_rgba(30,35,43,0.55)] backdrop-blur-xl sm:p-8">
			<div class="space-y-2">
				<p class="text-sm font-semibold uppercase tracking-[0.35em] text-primary">Sign up</p>
				<h2 class="text-3xl font-black tracking-tight text-foreground">Start your Quizzr account.</h2>
				<p class="text-sm leading-6 text-foreground/70">Keep it simple and jump right in.</p>
			</div>

			<form onsubmit={handleSubmit} class="mt-8 flex flex-col gap-4">
				<div class="space-y-2">
					<label class="text-sm font-semibold text-foreground" for="email">Email</label>
					<input
						type="text"
						name="email"
						placeholder="Email"
						required
						class="w-full rounded-md border-0 bg-background/80 px-4 py-3 text-foreground shadow-inner shadow-white/40 ring-1 ring-inset ring-foreground/10 placeholder:text-foreground/40 focus:ring-2 focus:ring-primary"
					/>
				</div>

				<div class="space-y-2">
					<label class="text-sm font-semibold text-foreground" for="username">Username</label>
					<input
						type="text"
						name="username"
						placeholder="Username"
						required
						class="w-full rounded-md border-0 bg-background/80 px-4 py-3 text-foreground shadow-inner shadow-white/40 ring-1 ring-inset ring-foreground/10 placeholder:text-foreground/40 focus:ring-2 focus:ring-primary"
					/>
					<p class="text-sm leading-6 text-foreground/65">
						Only letters, numbers, and underscores are allowed. Must be between 3 and 20 characters.
					</p>
				</div>

				{#if isChecking }
					<p class="text-sm font-medium text-primary">Checking username availability...</p>
				{/if}
				{#if usernameStatus}
					<p class="rounded-md bg-secondary/20 px-4 py-3 text-sm font-medium text-foreground">{usernameStatus}</p>
				{/if}

				<div class="space-y-2">
					<label class="text-sm font-semibold text-foreground" for="password">Password</label>
					<input
						type="password"
						name="password"
						placeholder="Password"
						required
						class="w-full rounded-md border-0 bg-background/80 px-4 py-3 text-foreground shadow-inner shadow-white/40 ring-1 ring-inset ring-foreground/10 placeholder:text-foreground/40 focus:ring-2 focus:ring-primary"
					/>
					<p class="text-sm leading-6 text-foreground/65">
						Aim for at least 8 characters with a mix of letters, numbers, and symbols.
					</p>
				</div>

				<button
					type="submit"
					class="mt-2 inline-flex items-center justify-center rounded-md bg-foreground px-5 py-3 text-sm font-semibold text-background shadow-[0_18px_36px_-24px_rgba(30,35,43,0.8)] transition hover:-translate-y-0.5"
				>
					Create account
				</button>

				{#if signupStatus}
					<p class="rounded-md bg-secondary/25 px-4 py-3 text-sm font-medium text-foreground">{signupStatus}</p>
				{/if}
			</form>
		</div>
	</section>
</main>
