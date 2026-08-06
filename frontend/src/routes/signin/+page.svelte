<script>
	import Navbar from "$lib/components/navbar.svelte";
</script>

<Navbar />

<main class="px-4 pb-16 pt-8 sm:px-6 lg:px-8">
	<section class="mx-auto grid max-w-6xl gap-8 lg:grid-cols-[1.02fr_.98fr]">
		<div class="flex items-end rounded-[2rem] border border-white/70 bg-white/55 p-8 shadow-[0_28px_80px_-48px_rgba(31,32,32,0.7)] backdrop-blur-xl lg:p-10">
			<div class="max-w-xl space-y-5">
				<p class="text-sm font-semibold uppercase tracking-[0.35em] text-primary">Welcome back</p>
				<h1 class="text-[clamp(2.75rem,7vw,5rem)] font-black leading-none tracking-tight text-foreground">
					Jump back into the game room.
				</h1>
				<p class="max-w-lg text-lg leading-8 text-foreground/75">
					Pick up where you left off, keep the room moving, and get back to the part people
					remember.
				</p>
				<div class="grid gap-3 sm:grid-cols-2">
					<div class="rounded-3xl bg-foreground/5 p-4">
						<p class="text-sm font-semibold text-primary">Fast access</p>
						<p class="mt-1 text-sm leading-6 text-foreground/70">One quick sign in and you are back in control.</p>
					</div>
					<div class="rounded-3xl bg-foreground/5 p-4">
						<p class="text-sm font-semibold text-primary">Focused layout</p>
						<p class="mt-1 text-sm leading-6 text-foreground/70">Calm surfaces, clear hierarchy, no clutter.</p>
					</div>
				</div>
			</div>
		</div>

		<div class="rounded-[2rem] border border-white/70 bg-white/70 p-6 shadow-[0_28px_80px_-48px_rgba(31,32,32,0.7)] backdrop-blur-xl sm:p-8">
			<div class="space-y-2">
				<p class="text-sm font-semibold uppercase tracking-[0.35em] text-primary">Sign in</p>
				<h2 class="text-3xl font-black tracking-tight text-foreground">Welcome back, teacher.</h2>
				<p class="text-sm leading-6 text-foreground/70">Use your credentials to continue.</p>
			</div>

			<form method="POST" class="mt-8 flex flex-col gap-4">
				<div class="space-y-2">
					<label class="text-sm font-semibold text-foreground" for="email">Email</label>
					<input
						type="text"
						name="email"
						placeholder="Email"
						class="w-full rounded-2xl border-0 bg-background/70 px-4 py-3 text-foreground shadow-inner shadow-white/40 ring-1 ring-inset ring-foreground/10 placeholder:text-foreground/40 focus:ring-2 focus:ring-primary"
					/>
				</div>
				<div class="space-y-2">
					<label class="text-sm font-semibold text-foreground" for="password">Password</label>
					<input
						type="password"
						name="password"
						placeholder="Password"
						class="w-full rounded-2xl border-0 bg-background/70 px-4 py-3 text-foreground shadow-inner shadow-white/40 ring-1 ring-inset ring-foreground/10 placeholder:text-foreground/40 focus:ring-2 focus:ring-primary"
					/>
				</div>
				<button
					type="submit"
					class="mt-2 inline-flex items-center justify-center rounded-full bg-foreground px-5 py-3 text-sm font-semibold text-background shadow-[0_18px_36px_-24px_rgba(31,32,32,0.8)] transition hover:-translate-y-0.5"
				>
					Sign in
				</button>
			</form>

			<p class="mt-6 text-sm text-foreground/75">
				Don't have an account?
				<a href="/signup" class="font-semibold text-primary hover:underline">Sign up</a>
			</p>
		</div>
	</section>
</main>
