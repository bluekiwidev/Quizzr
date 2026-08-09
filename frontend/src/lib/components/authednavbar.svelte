<script lang="ts">
	import ThemeToggle from './theme-toggle.svelte';
	import Logo from '$lib/assets/logo.svg';
	import sendSignoutRequest from '$lib/auth/signout';

	async function handleSignOut() {
		
		if (await sendSignoutRequest() == 0) {
			window.location.href = "/"
		} else {
			console.log("Failed to sign out")
		}
	}
</script>

<header class="px-4 pt-4 sm:px-6">
	<div class="mx-auto flex max-w-6xl items-center justify-between gap-3 rounded-xl border border-border bg-card/90 px-4 py-3 shadow-[0_18px_40px_-28px_rgba(30,35,43,0.55)] backdrop-blur-xl dark:shadow-[0_18px_40px_-28px_rgba(0,0,0,0.7)]">
		<a href="/dashboard" class="flex items-center gap-3">
			<span class="grid h-10 min-w-10 place-items-center rounded-md bg-foreground px-3">
				<img src={Logo} alt="Quizzr Logo" class="h-10 w-auto max-w-32 object-contain" />
			</span>
		</a>
		<a
			href="/dashboard"
			class="rounded-md border border-transparent px-4 py-2 text-sm font-semibold text-foreground transition hover:border-foreground/10 hover:bg-foreground/5"
		>
			Dashboard
		</a>
		<a
			href="/locker"
			class="rounded-md border border-transparent px-4 py-2 text-sm font-semibold text-foreground transition hover:border-foreground/10 hover:bg-foreground/5"
		>
			My locker
		</a>
		<a
			href="/quizzes"
			class="rounded-md border border-transparent px-4 py-2 text-sm font-semibold text-foreground transition hover:border-foreground/10 hover:bg-foreground/5"
		>
			My quizzes
		</a>
		<a
			href="/library"
			class="rounded-md border border-transparent px-4 py-2 text-sm font-semibold text-foreground transition hover:border-foreground/10 hover:bg-foreground/5"
		>
			Quiz library
		</a>
		<a
			href="/settings"
			class="rounded-md border border-transparent px-4 py-2 text-sm font-semibold text-foreground transition hover:border-foreground/10 hover:bg-foreground/5"
		>
			Settings
		</a>
		<div class="flex items-center gap-2 sm:gap-3">
			<ThemeToggle />
			<button
				onclick={handleSignOut}
				class="rounded-md bg-foreground px-4 py-2 text-sm font-semibold text-background shadow-[0_10px_24px_-16px_rgba(30,35,43,0.7)] transition hover:-translate-y-0.5"
			>
				Sign out
			</button>
		</div>
	</div>
</header>