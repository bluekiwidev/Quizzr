<script lang="ts">
    import ThemeToggle from '$lib/components/theme-toggle.svelte';
    import Logo from '$lib/assets/logo.svg';
    import { page } from '$app/state';
</script>

<header class="px-4 pt-4 sm:px-6">
	<div class="mx-auto flex max-w-6xl items-center justify-between gap-3 rounded-xl border border-border bg-card/90 px-4 py-3 shadow-[0_18px_40px_-28px_rgba(30,35,43,0.55)] backdrop-blur-xl dark:shadow-[0_18px_40px_-28px_rgba(0,0,0,0.7)]">
		<a href="/" class="flex items-center gap-3">
			<span class="grid h-10 min-w-10 place-items-center rounded-md bg-foreground px-3">
				<img src={Logo} alt="Quizzr Logo" class="h-10 w-auto max-w-32 object-contain" />
			</span>
		</a>
        <p class="text-sm font-semibold text-foreground/75">What are you doing here?</p>
		<nav class="flex items-center gap-2 sm:gap-3">
			<a
				href="/join"
				class="rounded-md border border-transparent px-4 py-2 text-sm font-semibold text-foreground transition hover:border-foreground/10 hover:bg-foreground/5"
			>
				Join a game
			</a>
			<ThemeToggle />
			<a
				href="/signin"
				class="rounded-md border border-transparent px-4 py-2 text-sm font-semibold text-foreground transition hover:border-foreground/10 hover:bg-foreground/5"
			>
				Sign in
			</a>
            
			<a
				href="/signup"
				class="rounded-md bg-foreground px-4 py-2 text-sm font-semibold text-background shadow-[0_10px_24px_-16px_rgba(30,35,43,0.7)] transition hover:-translate-y-0.5"
			>
				Sign up
			</a>
		</nav>
	</div>
</header>

<h1 class="mx-auto mt-8 max-w-3xl rounded-lg bg-white/80 dark:bg-gray-800/80 border border-gray-200 dark:border-gray-700 p-6 text-center text-2xl sm:text-3xl font-extrabold text-gray-900 dark:text-white shadow">
	{page.status} {page.error.message}
</h1>