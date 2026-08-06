<script>
    import Navbar from "$lib/components/authednavbar.svelte";
</script>

<Navbar />

<main class="px-4 pb-16 pt-8 sm:px-6 lg:px-8">
    <section class="mx-auto grid max-w-6xl gap-6 lg:grid-cols-[1.15fr_.85fr]">
        <div class="rounded-[2rem] border border-white/70 bg-white/60 p-8 shadow-[0_28px_80px_-48px_rgba(31,32,32,0.7)] backdrop-blur-xl lg:p-10">
            <p class="text-sm font-semibold uppercase tracking-[0.35em] text-primary">Dashboard</p>
            <h1 class="mt-4 text-[clamp(2.5rem,6vw,4.75rem)] font-black leading-none tracking-tight text-foreground">
                Your next session is ready to go.
            </h1>
            <p class="mt-5 max-w-2xl text-lg leading-8 text-foreground/75">
                The board is calm, the palette is warm, and the game room is set up to feel fast and focused.
            </p>

            <div class="mt-8 grid gap-3 sm:grid-cols-3">
                <div class="rounded-3xl bg-foreground/5 p-4">
                    <p class="text-xs font-semibold uppercase tracking-[0.3em] text-primary">Energy</p>
                    <p class="mt-2 text-lg font-bold text-foreground">High</p>
                </div>
                <div class="rounded-3xl bg-foreground/5 p-4">
                    <p class="text-xs font-semibold uppercase tracking-[0.3em] text-primary">Focus</p>
                    <p class="mt-2 text-lg font-bold text-foreground">Sharp</p>
                </div>
                <div class="rounded-3xl bg-foreground/5 p-4">
                    <p class="text-xs font-semibold uppercase tracking-[0.3em] text-primary">Flow</p>
                    <p class="mt-2 text-lg font-bold text-foreground">Ready</p>
                </div>
            </div>
        </div>

        <div class="grid gap-6">
            <div class="rounded-[2rem] border border-white/70 bg-[linear-gradient(135deg,rgba(85,104,144,0.95),rgba(31,32,32,0.9))] p-6 text-background shadow-[0_28px_80px_-48px_rgba(31,32,32,0.8)] sm:p-8">
                <p class="text-xs font-semibold uppercase tracking-[0.35em] text-background/80">Spotlight</p>
                <p class="mt-4 text-3xl font-black leading-tight">Lead with clear prompts, bright contrast, and a rhythm that keeps people playing.</p>
            </div>

            <div class="grid gap-4 sm:grid-cols-2">
                <div class="rounded-[1.75rem] border border-white/70 bg-white/65 p-5 shadow-[0_18px_50px_-40px_rgba(31,32,32,0.6)] backdrop-blur">
                    <p class="text-sm font-semibold text-primary">Sessions</p>
                    <p class="mt-2 text-2xl font-black text-foreground">Polished</p>
                    <p class="mt-2 text-sm leading-6 text-foreground/70">Everything is framed to feel easy to start and easy to follow.</p>
                </div>
                <div class="rounded-[1.75rem] border border-white/70 bg-white/65 p-5 shadow-[0_18px_50px_-40px_rgba(31,32,32,0.6)] backdrop-blur">
                    <p class="text-sm font-semibold text-primary">Visual language</p>
                    <p class="mt-2 text-2xl font-black text-foreground">Warm</p>
                    <p class="mt-2 text-sm leading-6 text-foreground/70">Soft cards, bold headings, and cleaner spacing bring the new look together.</p>
                </div>
            </div>
        </div>
    </section>
</main>


