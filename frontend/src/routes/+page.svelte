<script>
	import { DiaTextReveal } from "$lib/components/magic/dia-text-reveal";
	import Navbar from "$lib/components/navbar.svelte";
</script>

<Navbar />

<div class="flex min-h-64 items-center justify-center p-8">
  <h1 class="text-[clamp(1rem,7vw,5.5rem)] text-center font-semibold tracking-tight leading-none text-white">
    <span class="block text-foreground">Make learning</span>{" "}
    <span class="block align-bottom leading-tight text-inherit font-extrabold pb-3">
      <DiaTextReveal
        repeat
        repeatDelay={1.2}
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
</div>

<h2 class="text-3xl font-semibold tracking-tight text-white text-center">With Quizzr</h2>