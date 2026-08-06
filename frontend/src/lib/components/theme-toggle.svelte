<script lang="ts">
	import { MoonStar, SunMedium } from '@lucide/svelte';
	import { onMount } from 'svelte';

	const storageKey = 'quizzr-theme';

	let isDark = $state(false);
	let hasManualTheme = $state(false);

	function applyTheme(theme: 'dark' | 'light') {
		document.documentElement.classList.toggle('dark', theme === 'dark');
	}

	function setTheme(theme: 'dark' | 'light') {
		isDark = theme === 'dark';
		hasManualTheme = true;
		window.localStorage.setItem(storageKey, theme);
		applyTheme(theme);
	}

	function toggleTheme() {
		setTheme(isDark ? 'light' : 'dark');
	}

	onMount(() => {
		const storedTheme = window.localStorage.getItem(storageKey);
		const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');

		if (storedTheme === 'dark' || storedTheme === 'light') {
			hasManualTheme = true;
			isDark = storedTheme === 'dark';
			applyTheme(storedTheme);
			return;
		}

		const syncSystemTheme = () => {
			if (hasManualTheme) {
				return;
			}

			const theme = mediaQuery.matches ? 'dark' : 'light';
			isDark = theme === 'dark';
			applyTheme(theme);
		};

		syncSystemTheme();
		mediaQuery.addEventListener('change', syncSystemTheme);

		return () => mediaQuery.removeEventListener('change', syncSystemTheme);
	});
</script>

<button
	type="button"
	onclick={toggleTheme}
	class="inline-flex h-11 w-11 items-center justify-center rounded-full border border-border bg-card text-foreground shadow-[0_10px_24px_-18px_rgba(31,32,32,0.7)] transition hover:-translate-y-0.5 hover:bg-accent/70"
	aria-label={isDark ? 'Switch to light mode' : 'Switch to dark mode'}
	title={isDark ? 'Switch to light mode' : 'Switch to dark mode'}
>
	{#if isDark}
		<SunMedium class="size-5" />
	{:else}
		<MoonStar class="size-5" />
	{/if}
</button>