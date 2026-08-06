<script lang="ts">
	import ThemeToggle from './theme-toggle.svelte';
</script>

<header class="sticky top-0 z-50 px-4 pt-4 sm:px-6">
	<div class="mx-auto flex max-w-6xl items-center justify-between gap-3 rounded-full border border-border bg-card/80 px-4 py-3 shadow-[0_20px_60px_-40px_rgba(31,32,32,0.7)] backdrop-blur-xl dark:shadow-[0_20px_60px_-40px_rgba(0,0,0,0.8)]">
		<a href="/dashboard" class="flex items-center gap-3">
			<span class="grid size-10 place-items-center rounded-2xl bg-primary text-primary-foreground text-sm font-black tracking-[0.3em]">
				Q
			</span>
			<div class="leading-tight">
				<p class="text-xs font-semibold uppercase tracking-[0.35em] text-primary">Quizzr</p>
				<p class="text-sm text-muted-foreground">Your game room</p>
			</div>
		</a>

		<div class="flex items-center gap-2 sm:gap-3">
			<ThemeToggle />
			<a
				href="/dashboard"
				class="rounded-full bg-foreground px-4 py-2 text-sm font-semibold text-background shadow-[0_10px_24px_-16px_rgba(31,32,32,0.7)] transition hover:-translate-y-0.5"
			>
				Dashboard
			</a>
		</div>
	</div>
</header>