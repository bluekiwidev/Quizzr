<script>
	import { DiaTextReveal } from "$lib/components/magic/dia-text-reveal";
	import Navbar from "$lib/components/navbar.svelte";
</script>

<Navbar />

<main class="px-4 pb-16 pt-8 sm:px-6 lg:px-8">
  <section class="mx-auto grid max-w-6xl items-center gap-10 lg:grid-cols-[1.12fr_.88fr]">
    <div class="space-y-8">
      <div class="max-w-3xl space-y-5">
        <h1 class="text-[clamp(3rem,9vw,6.75rem)] font-black leading-[0.9] tracking-tight text-foreground">
          <span class="block text-primary">Make learning</span>
          <span class="block overflow-visible pb-2 leading-[1.05] text-foreground">
            <DiaTextReveal
              repeat
              repeatDelay={1.2}
              colors={["#355a8a", "#92add1", "#1e232b", "#d6dde6"]}
              text={[
                "engaging",
                "active",
                "competitive",
                "easy",
                "interactive",
                "rewarding",
                "social",
                "immersive",
                "fun again"
              ]}
            />
          </span>
        </h1>

        <p class="max-w-2xl text-lg leading-8 text-foreground/75 sm:text-xl">
          Quizzr turns boring lessons (no offence) into engaging, interactive games that inspire students to learn, compete, and have fun. Teachers can create games in seconds, and students can join with a simple code.
        </p>
      </div>

      <div class="flex flex-wrap items-center gap-3">
        <a
          href="/signup"
          class="inline-flex items-center justify-center rounded-md bg-foreground px-6 py-3 text-sm font-semibold text-background shadow-[0_20px_40px_-22px_rgba(30,35,43,0.8)] transition hover:-translate-y-0.5"
        >
          Get started
        </a>
        <a
          href="/signin"
          class="inline-flex items-center justify-center rounded-md border border-foreground/15 bg-white/80 px-6 py-3 text-sm font-semibold text-foreground transition hover:border-foreground/25 hover:bg-white"
        >
          I already have an account
        </a>
      </div>
    </div>

    <div class="relative">
      <div class="absolute -left-6 top-10 hidden h-16 w-16 rounded-md bg-secondary/70 lg:block"></div>
      <div class="absolute -right-6 bottom-8 hidden h-20 w-20 rounded-md bg-primary/25 lg:block"></div>
      <div class="relative overflow-hidden rounded-xl border border-border bg-[linear-gradient(135deg,rgba(248,244,236,0.96),rgba(227,233,240,0.86))] p-6 shadow-[0_24px_60px_-36px_rgba(30,35,43,0.55)] backdrop-blur-xl sm:p-8 dark:bg-[linear-gradient(135deg,rgba(26,32,40,0.96),rgba(47,57,71,0.45))]">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-xs font-semibold uppercase tracking-[0.35em] text-primary">Game room</p>
            <h2 class="mt-3 text-2xl font-bold tracking-tight text-foreground">A brighter way to learn</h2>
          </div>
          <div class="rounded-md bg-foreground px-3 py-2 text-xs font-semibold uppercase tracking-[0.3em] text-background">
            Live
          </div>
        </div>

        <div class="mt-8 rounded-xl border border-border bg-[linear-gradient(135deg,rgba(248,244,236,0.92),rgba(230,234,240,0.72))] p-6 shadow-inner shadow-white/40 dark:bg-[linear-gradient(135deg,rgba(26,32,40,0.96),rgba(47,57,71,0.45))]">
          <p class="text-sm font-semibold uppercase tracking-[0.3em] text-primary">This is Quizzr</p>
          <p class="mt-4 text-3xl font-black leading-tight text-foreground sm:text-4xl">
            Quick quizzes, sharp blocks, and a room that feels worth joining.
          </p>
          <div class="mt-6 grid gap-3 sm:grid-cols-2">
            <div class="rounded-lg bg-background/80 p-4">
              <p class="text-xs font-semibold uppercase tracking-[0.3em] text-primary">Momentum</p>
              <p class="mt-2 text-lg font-bold text-foreground">Keeps players focused</p>
            </div>
            <div class="rounded-lg bg-background/80 p-4">
              <p class="text-xs font-semibold uppercase tracking-[0.3em] text-primary">Style</p>
              <p class="mt-2 text-lg font-bold text-foreground">Feels polished and modern</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</main>