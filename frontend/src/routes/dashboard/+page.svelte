<script>
    import Navbar from "$lib/components/authednavbar.svelte";
</script>

<Navbar />

<main class="px-4 pb-16 pt-8 sm:px-6 lg:px-8">
    <section class="mx-auto grid max-w-6xl gap-6 lg:grid-cols-[1.15fr_.85fr]">
		<div class="rounded-xl border border-border bg-card/90 p-8 shadow-[0_24px_60px_-36px_rgba(30,35,43,0.55)] backdrop-blur-xl lg:p-10">
            <p class="text-sm font-semibold uppercase tracking-[0.35em] text-primary">Dashboard</p>
            <h1 class="mt-4 text-[clamp(2.5rem,6vw,4.75rem)] font-black leading-none tracking-tight text-foreground">
                Your next session is ready to go.
            </h1>
            <p class="mt-5 max-w-2xl text-lg leading-8 text-foreground/75">
                The board is calm, the palette is warm, and the game room is set up to feel fast and focused.
            </p>

            <div class="mt-8 grid gap-3 sm:grid-cols-3">
                <div class="rounded-lg border border-border bg-background/70 p-4">
                    <p class="text-xs font-semibold uppercase tracking-[0.3em] text-primary">Energy</p>
                    <p class="mt-2 text-lg font-bold text-foreground">High</p>
                </div>
                <div class="rounded-lg border border-border bg-background/70 p-4">
                    <p class="text-xs font-semibold uppercase tracking-[0.3em] text-primary">Focus</p>
                    <p class="mt-2 text-lg font-bold text-foreground">Sharp</p>
                </div>
                <div class="rounded-lg border border-border bg-background/70 p-4">
                    <p class="text-xs font-semibold uppercase tracking-[0.3em] text-primary">Flow</p>
                    <p class="mt-2 text-lg font-bold text-foreground">Ready</p>
                </div>
            </div>
        </div>

        <div class="grid gap-6">
			<div class="rounded-xl border border-border bg-[linear-gradient(135deg,rgba(53,90,138,0.95),rgba(30,35,43,0.92))] p-6 text-background shadow-[0_24px_60px_-36px_rgba(30,35,43,0.7)] dark:bg-[linear-gradient(135deg,rgba(146,173,209,0.2),rgba(47,57,71,0.7))] sm:p-8">
                <p class="text-xs font-semibold uppercase tracking-[0.35em] text-background/80">Spotlight</p>
                <p class="mt-4 text-3xl font-black leading-tight">Lead with clear prompts, bright contrast, and a rhythm that keeps people playing.</p>
            </div>

            <div class="grid gap-4 sm:grid-cols-2">
				<div class="rounded-lg border border-border bg-card/90 p-5 shadow-[0_18px_40px_-30px_rgba(30,35,43,0.55)] backdrop-blur">
                    <p class="text-sm font-semibold text-primary">Sessions</p>
                    <p class="mt-2 text-2xl font-black text-foreground">Polished</p>
                    <p class="mt-2 text-sm leading-6 text-foreground/70">Everything is framed to feel easy to start and easy to follow.</p>
                </div>
				<div class="rounded-lg border border-border bg-card/90 p-5 shadow-[0_18px_40px_-30px_rgba(30,35,43,0.55)] backdrop-blur">
                    <p class="text-sm font-semibold text-primary">Visual language</p>
                    <p class="mt-2 text-2xl font-black text-foreground">Warm</p>
                    <p class="mt-2 text-sm leading-6 text-foreground/70">Soft cards, bold headings, and cleaner spacing bring the new look together.</p>
                </div>
            </div>
        </div>
    </section>
</main>


