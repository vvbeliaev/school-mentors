<script lang="ts">
	import { page } from '$app/state';
	import { goto } from '$app/navigation';
	import { mentorsStore } from '$lib/apps/user';
	import { MentorCard, MentorsFilter } from '$lib/apps/user/ui';
	import type { MentorFilters } from '$lib/apps/user/models';
	import { Loader2, SearchX, Info } from 'lucide-svelte';
	import { Button } from '$lib/shared/ui';
	import { debounce } from '$lib/shared/utils';

	let { data } = $props();

	// Local state for filters to keep UI responsive
	let filters: MentorFilters = $derived({ ...data.filters });

	// Sync local filters with URL changes (e.g. back button)
	$effect(() => {
		filters = { ...data.filters };
	});

	// Load data when data.filters changes (stable URL state)
	$effect(() => {
		mentorsStore.load(data.filters);
	});

	// Sync local filters with URL after a short delay
	const debouncedUpdateUrl = debounce((newFilters: MentorFilters) => {
		const params = new URLSearchParams();
		if (newFilters.q) params.set('q', newFilters.q);
		if (newFilters.university) params.set('university', newFilters.university);
		if (newFilters.tags && newFilters.tags.length > 0)
			params.set('tags', newFilters.tags.join(','));
		if (newFilters.minPrice) params.set('minPrice', newFilters.minPrice.toString());
		if (newFilters.maxPrice) params.set('maxPrice', newFilters.maxPrice.toString());
		if (newFilters.sort && newFilters.sort !== '-created') params.set('sort', newFilters.sort);

		const search = params.toString();
		goto(`?${search}`, { replaceState: true, keepFocus: true, noScroll: true });
	}, 300);

	function handleFilterChange(newFilters: MentorFilters) {
		filters = { ...newFilters };
		debouncedUpdateUrl(filters);
	}
</script>

<div class="container mx-auto space-y-8 px-4 py-8">
	<div class="flex flex-col gap-2">
		<h1 class="font-display text-4xl font-black">Find your Mentor</h1>
		<p class="text-lg opacity-60">
			Connect with students from top universities who've already walked your path.
		</p>
	</div>

	<div
		class="sticky top-0 z-30 -mx-4 bg-base-100/80 px-4 py-4 backdrop-blur-md lg:relative lg:bg-transparent lg:p-0"
	>
		<MentorsFilter {filters} onchange={handleFilterChange} />
	</div>

	{#if mentorsStore.loading && mentorsStore.mentors.length === 0}
		<div class="flex h-64 items-center justify-center">
			<Loader2 class="size-8 animate-spin text-primary" />
		</div>
	{:else}
		{#if mentorsStore.isFallback}
			<div class="alert rounded-2xl border-none bg-info/10 alert-info">
				<Info class="size-5 text-info" />
				<div class="flex flex-col">
					<span class="font-bold">No exact matches found</span>
					<span class="text-sm"
						>We couldn't find mentors matching all your filters. Here are some of our verified
						mentors instead.</span
					>
				</div>
			</div>
		{/if}

		{#if mentorsStore.mentors.length > 0}
			<div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
				{#each mentorsStore.mentors as mentor (mentor.id)}
					<MentorCard {mentor} />
				{/each}
			</div>

			{#if mentorsStore.page < mentorsStore.totalPages}
				<div class="flex justify-center pt-8">
					<Button
						variant="soft"
						onclick={() => mentorsStore.loadNextPage()}
						disabled={mentorsStore.loading}
					>
						{#if mentorsStore.loading}
							<Loader2 class="mr-2 size-4 animate-spin" />
						{/if}
						Load More Mentors
					</Button>
				</div>
			{/if}
		{:else}
			<div class="flex flex-col items-center justify-center space-y-4 py-20 text-center">
				<div class="rounded-full bg-base-200 p-6">
					<SearchX class="size-12 opacity-20" />
				</div>
				<div class="space-y-1">
					<h3 class="text-xl font-bold">No mentors found</h3>
					<p class="mx-auto max-w-xs opacity-60">
						We couldn't find any mentors at the moment. Try adjusting your filters or check back
						later.
					</p>
				</div>
				<Button variant="outline" onclick={() => handleFilterChange({})}>Clear all filters</Button>
			</div>
		{/if}
	{/if}
</div>
