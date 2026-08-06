<script lang="ts">
	import ThemeToggle from './theme-toggle.svelte';
	import Logo from '$lib/assets/logo.svg';
</script>

<header class="sticky top-0 z-50 px-4 pt-4 sm:px-6">
	<div class="mx-auto flex max-w-6xl items-center justify-between gap-3 rounded-full border border-border bg-card/80 px-4 py-3 shadow-[0_20px_60px_-40px_rgba(31,32,32,0.7)] backdrop-blur-xl dark:shadow-[0_20px_60px_-40px_rgba(0,0,0,0.8)]">
		<a href="/dashboard" class="flex items-center gap-3">
			<span class="grid h-10 min-w-10 place-items-center rounded-2xl bg-foreground px-3">
				<img src={Logo} alt="Quizzr Logo" class="h-10 w-auto max-w-32 object-contain" />
			</span>
		</a>
		<a
			href="/dashboard"
			class="rounded-full border border-transparent px-4 py-2 text-sm font-semibold text-foreground transition hover:border-foreground/10 hover:bg-foreground/5"
		>
			Dashboard
		</a>
		<a
			href="/locker"
			class="rounded-full border border-transparent px-4 py-2 text-sm font-semibold text-foreground transition hover:border-foreground/10 hover:bg-foreground/5"
		>
			My locker
		</a>
		<a
			href="/quizzes"
			class="rounded-full border border-transparent px-4 py-2 text-sm font-semibold text-foreground transition hover:border-foreground/10 hover:bg-foreground/5"
		>
			My quizzes
		</a>
		<a
			href="/library"
			class="rounded-full border border-transparent px-4 py-2 text-sm font-semibold text-foreground transition hover:border-foreground/10 hover:bg-foreground/5"
		>
			Quiz library
		</a>
		<a
			href="/settings"
			class="rounded-full border border-transparent px-4 py-2 text-sm font-semibold text-foreground transition hover:border-foreground/10 hover:bg-foreground/5"
		>
			Settings
		</a>
		<div class="flex items-center gap-2 sm:gap-3">
			<ThemeToggle />
			<a
				href="/signout"
				class="rounded-full bg-foreground px-4 py-2 text-sm font-semibold text-background shadow-[0_10px_24px_-16px_rgba(31,32,32,0.7)] transition hover:-translate-y-0.5"
			>
				Sign out
			</a>
		</div>
	</div>
</header>