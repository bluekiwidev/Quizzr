<script lang="ts">
	import Navbar from '$lib/components/authednavbar.svelte';
	import { PUBLIC_BACKEND } from '$env/static/public';
    import { AnimatedCircularProgressBar } from "$lib/components/magic/animated-circular-progress-bar";
	import { onMount, onDestroy } from "svelte";

	let wins: number | string = $state('loading...');
	let losses: number | string = $state('loading...');
	let totalGames: number | string = $state('loading...');

	let value = $state(0);
    let targetLevel;
	let interval: ReturnType<typeof setInterval>;

	onMount(() => {
		const handleIncrement = () => {
			value = value === 50 ? 50 : value + 10;
		};
		handleIncrement();
		interval = setInterval(handleIncrement, 100);
	});
	onDestroy(() => {
		if (interval) {
			clearInterval(interval);
		}
	});

	async function getStats(): Promise<[number, number, number, number] | [-1, -1, -1, -1]> { //Returns [wins, losses, totalGames, level] or [-1, -1, -1] on error
			try {
				console.log('Requesting stats from backend...');
				const response = await fetch(`${PUBLIC_BACKEND}/getstats`, {
					method: 'get',
					credentials: 'include',
				});

				// Return codes:
				if (response.status === 201) {
					const data = await response.json();
					return data as [number, number, number, number]; //All good
				}
				if (response.status === 500) return [-1, -1, -1, -1]; //Server error
				return [-1, -1, -1, -1];
			} catch {
				return [-1, -1, -1, -1];
			}
		}

	onMount(async () => {
		const stats = await getStats();
		if (Array.isArray(stats) && stats.length === 4 && stats[0] === -1 && stats[1] === -1 && stats[2] === -1 && stats[3] === -1) {
			wins = 'Error';
			losses = 'Error';
			totalGames = 'Error';
			targetLevel = 'Error';
		} else if (Array.isArray(stats) && stats.length === 4) {
			[wins, losses, totalGames, targetLevel] = stats;
		} else {
			wins = 'Error';
			losses = 'Error';
			totalGames = 'Error';
			targetLevel = 'Error';
		}
	});
</script>

<Navbar />

<main class="px-4 pt-8 pb-16 sm:px-6 lg:px-8">
	<section class="mx-auto grid max-w-6xl gap-6 lg:grid-cols-[1.15fr_.85fr]">
		<div
			class="rounded-xl border border-border bg-card/90 p-8 shadow-[0_24px_60px_-36px_rgba(30,35,43,0.55)] backdrop-blur-xl lg:p-10"
		>
			<p class="text-sm font-semibold tracking-[0.35em] text-primary uppercase">Stats</p>
			<p class="mt-5 max-w-2xl text-lg leading-8 text-foreground/75">
				How many points you have untill your next level and how many games you have won, lost and played in total.
			</p>
            <AnimatedCircularProgressBar
	            {value}
	            gaugePrimaryColor="rgb(79 70 229)"
            	gaugeSecondaryColor="rgba(0, 0, 0, 0.1)"
            />

			<div class="mt-8 grid gap-3 sm:grid-cols-3">
				<div class="rounded-lg border border-border bg-background/70 p-4">
					<p class="text-xs font-semibold tracking-[0.3em] text-primary uppercase">Games won</p>
					<p class="mt-2 text-lg font-bold text-foreground">{wins}</p>
				</div>
				<div class="rounded-lg border border-border bg-background/70 p-4">
					<p class="text-xs font-semibold tracking-[0.3em] text-primary uppercase">Games loss</p>
					<p class="mt-2 text-lg font-bold text-foreground">{losses}</p>
				</div>
				<div class="rounded-lg border border-border bg-background/70 p-4">
					<p class="text-xs font-semibold tracking-[0.3em] text-primary uppercase">Total games</p>
					<p class="mt-2 text-lg font-bold text-foreground">{totalGames}</p>
				</div>
			</div>
		</div>

		<div class="grid gap-6">
			<div
				class="rounded-xl border border-border bg-[linear-gradient(135deg,rgba(53,90,138,0.95),rgba(30,35,43,0.92))] p-6 text-background shadow-[0_24px_60px_-36px_rgba(30,35,43,0.7)] sm:p-8 dark:bg-[linear-gradient(135deg,rgba(146,173,209,0.2),rgba(47,57,71,0.7))]"
			>
				<p class="text-xs font-semibold tracking-[0.35em] text-background/80 uppercase">
					Spotlight
				</p>
				<p class="mt-4 text-3xl leading-tight font-black">
					Lead with clear prompts, bright contrast, and a rhythm that keeps people playing.
				</p>
			</div>

			<div class="grid gap-4 sm:grid-cols-2">
				<div
					class="rounded-lg border border-border bg-card/90 p-5 shadow-[0_18px_40px_-30px_rgba(30,35,43,0.55)] backdrop-blur"
				>
					<p class="text-sm font-semibold text-primary">Sessions</p>
					<p class="mt-2 text-2xl font-black text-foreground">Polished</p>
					<p class="mt-2 text-sm leading-6 text-foreground/70">
						Everything is framed to feel easy to start and easy to follow.
					</p>
				</div>
				<div
					class="rounded-lg border border-border bg-card/90 p-5 shadow-[0_18px_40px_-30px_rgba(30,35,43,0.55)] backdrop-blur"
				>
					<p class="text-sm font-semibold text-primary">Visual language</p>
					<p class="mt-2 text-2xl font-black text-foreground">Warm</p>
					<p class="mt-2 text-sm leading-6 text-foreground/70">
						Soft cards, bold headings, and cleaner spacing bring the new look together.
					</p>
				</div>
			</div>
		</div>
	</section>
</main>
