<script>
	let planetPromise = getPlanet();
	async function getPlanet() {
		let planet;
		setInterval(() => {
			planet = "Mars";
		}, 1000);
		return planet;
	}
</script>

<p>
	{#await planetPromise}
		Loading planet...
	{:then planet}
		Your next planet is {planet}.
	{:catch someError}
		System error: {someError.message}.
	{/await}
</p>

<button on:click={() => (planetPromise = getPlanet())}>
	Explore the next planet
</button>
