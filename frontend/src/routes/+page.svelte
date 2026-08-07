<script>
	import { DiaTextReveal } from "$lib/components/magic/dia-text-reveal";
	import Navbar from "$lib/components/navbar.svelte";
</script>

<Navbar />

<main class="px-4 pb-16 pt-8 sm:px-6 lg:px-8">
  <section class="mx-auto grid max-w-6xl items-center gap-10 lg:grid-cols-1">
    <div class="space-y-8 text-center mx-auto">
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

      <div class="flex flex-wrap items-center justify-center gap-3">
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
  </section>
</main>