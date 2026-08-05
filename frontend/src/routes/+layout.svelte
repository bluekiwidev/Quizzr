<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';

	let { children } = $props();
</script>

<title>Quizzr - Make learning fun again</title>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>
{@render children()}
