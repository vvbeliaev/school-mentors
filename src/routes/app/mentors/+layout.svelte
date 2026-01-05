<script lang="ts">
	import { mentorsStore } from '$lib/apps/user';

	let { children, data } = $props();
	const mentorsPromise = $derived(data.mentorsPromise);

	$effect(() => {
		mentorsPromise.then((mentorsRes) => {
			mentorsStore.set(
				mentorsRes.items,
				mentorsRes.page,
				mentorsRes.totalPages,
				mentorsRes.totalItems
			);
		});
	});
</script>

{@render children()}
