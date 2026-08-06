<script>
	import Navbar from "$lib/components/navbar.svelte";
</script>

<Navbar />

<main class="px-4 pb-16 pt-8 sm:px-6 lg:px-8">
	<section class="mx-auto grid max-w-6xl gap-8 lg:grid-cols-[1.02fr_.98fr]">
		<div class="rounded-[2rem] border border-border bg-card/85 p-6 shadow-[0_28px_80px_-48px_rgba(31,32,32,0.7)] backdrop-blur-xl sm:p-8">
			<div class="space-y-2">
				<p class="text-sm font-semibold uppercase tracking-[0.35em] text-primary">Sign in</p>
				<h2 class="text-3xl font-black tracking-tight text-foreground">Welcome back peeps</h2>
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
