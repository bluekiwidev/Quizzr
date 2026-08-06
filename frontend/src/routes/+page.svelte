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
          <span class="block text-[#556890]">Make learning</span>
          <span class="block overflow-visible pb-2 leading-[1.05] text-foreground">
            <DiaTextReveal
              repeat
              repeatDelay={1.2}
              colors={["#556890", "#f4b9b8", "#1f2020", "#d7bf87"]}
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
          class="inline-flex items-center justify-center rounded-full bg-foreground px-6 py-3 text-sm font-semibold text-background shadow-[0_20px_40px_-22px_rgba(31,32,32,0.8)] transition hover:-translate-y-0.5"
        >
          Get started
        </a>
        <a
          href="/signin"
          class="inline-flex items-center justify-center rounded-full border border-foreground/15 bg-white/60 px-6 py-3 text-sm font-semibold text-foreground transition hover:border-foreground/25 hover:bg-white/80"
        >
          I already have an account
        </a>
      </div>
    </div>

    <div class="relative">
      <div class="absolute -left-8 top-8 hidden h-24 w-24 rounded-full bg-secondary/70 blur-3xl lg:block"></div>
      <div class="absolute -right-6 bottom-10 hidden h-28 w-28 rounded-full bg-primary/25 blur-3xl lg:block"></div>
      <div class="relative overflow-hidden rounded-[2rem] border border-border bg-card/80 p-6 shadow-[0_30px_80px_-45px_rgba(31,32,32,0.7)] backdrop-blur-xl sm:p-8">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-xs font-semibold uppercase tracking-[0.35em] text-primary">Game room</p>
            <h2 class="mt-3 text-2xl font-bold tracking-tight text-foreground">A brighter way to learn</h2>
          </div>
          <div class="rounded-2xl bg-foreground px-3 py-2 text-xs font-semibold uppercase tracking-[0.3em] text-background">
            Live
          </div>
        </div>

        <div class="mt-8 rounded-[1.75rem] border border-border bg-[linear-gradient(135deg,rgba(254,236,192,0.95),rgba(244,185,184,0.32))] p-6 shadow-inner shadow-white/40 dark:bg-[linear-gradient(135deg,rgba(39,40,41,0.96),rgba(85,104,144,0.25))]">
          <p class="text-sm font-semibold uppercase tracking-[0.3em] text-primary">This is Quizzr</p>
          <p class="mt-4 text-3xl font-black leading-tight text-foreground sm:text-4xl">
            Quick quizzes, bright reactions, and a room that feels worth joining.
          </p>
          <div class="mt-6 grid gap-3 sm:grid-cols-2">
            <div class="rounded-3xl bg-background/70 p-4">
              <p class="text-xs font-semibold uppercase tracking-[0.3em] text-primary">Momentum</p>
              <p class="mt-2 text-lg font-bold text-foreground">Keeps players focused</p>
            </div>
            <div class="rounded-3xl bg-background/70 p-4">
              <p class="text-xs font-semibold uppercase tracking-[0.3em] text-primary">Style</p>
              <p class="mt-2 text-lg font-bold text-foreground">Feels polished and modern</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</main>