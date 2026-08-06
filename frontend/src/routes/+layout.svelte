<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';

	let { children } = $props();
</script>

<title>Quizzr - Make learning fun again</title>

<svelte:head>
	<link rel="icon" href={favicon} />
	<script>
		(() => {
			const storageKey = 'quizzr-theme';
			const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
			const storedTheme = window.localStorage.getItem(storageKey);
			const isDark = storedTheme ? storedTheme === 'dark' : mediaQuery.matches;

			document.documentElement.classList.toggle('dark', isDark);
		})();
	</script>
</svelte:head>
{@render children()}
