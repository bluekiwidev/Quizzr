<script lang="ts">
	import { TextMarquee } from "$lib/components/spell/text-marquee";

	let phrases = ["fun", "engaging", "active", "competitive", "easy", "interactive", "rewarding", "social", "immersive", "personalized"];
</script>

<TextMarquee items={phrases} height={120} itemHeight={40} speed={1}>
	{#snippet prefix()}
		<span class="text-muted-foreground/75 text-3xl font-medium">Make learning </span>
	{/snippet}

	{#snippet children(item)}
		<span class="text-3xl font-medium">{item}</span>
	{/snippet}
</TextMarquee>
