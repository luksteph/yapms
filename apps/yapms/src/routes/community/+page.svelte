<script lang="ts">
	import image from '$lib/assets/images/countries/usa.webp';
	import { PocketBaseStore } from '$lib/stores/PocketBase';
	const myarr = [1, 1, 1, 1, 1, 1, 1, 1];

	const communityUserMaps = $PocketBaseStore.collection('community_user_maps').getFullList({
		expand: 'community_user.user'
	});
	async function test() {
		console.log(await communityUserMaps);
	}
	test();
</script>

<div class="navbar bg-base-200">
	<div class="navbar-start">
		<a href="/" class="btn btn-primary">Home</a>
	</div>
	<h1 class="text-2xl font-bold navbar-center">YAPms Community</h1>
	<div class="navbar-end">
		{#if $PocketBaseStore.authStore.isValid}
			YES
		{:else}
			now
		{/if}
	</div>
</div>

{#await communityUserMaps}
	loading..
{:then maps}
	<div class="flex flex-row flex-wrap gap-5 p-5">
		{#each maps as map}
			<div class="card card-compact w-96 bg-base-100 shadow-lg">
				<figure><img src={image} alt={map.title} /></figure>
				<div class="card-body">
					<div class="flex flex-row justify-between">
						<h2 class="card-title">{map.title}</h2>
						<span class="font-example text-lg"
							>Crafted by {map.expand?.community_user.expand.user.username}</span
						>
					</div>
					<p>{map.description}</p>
					<div class="card-actions justify-end">
						<button class="btn btn-primary btn-sm">Load</button>
					</div>
				</div>
			</div>
		{/each}
	</div>
{/await}
