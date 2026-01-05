<script lang="ts">
	import { Search, SlidersHorizontal, X } from 'lucide-svelte';
	import { Button, Input } from '$lib/shared/ui';
	import UniversitySelect from './UniversitySelect.svelte';
	import TagInput from './TagInput.svelte';
	import type { MentorFilters, SortOption } from '../models';
	import { SORT_OPTIONS } from '../models';

	interface Props {
		filters: MentorFilters;
		onchange: (filters: MentorFilters) => void;
	}

	let { filters = $bindable(), onchange }: Props = $props();

	let showMobileFilters = $state(false);
	let sort;

	function updateFilter<K extends keyof MentorFilters>(key: K, value: MentorFilters[K]) {
		filters[key] = value;
		onchange(filters);
	}

	function clearFilters() {
		filters = {
			q: '',
			university: '',
			tags: [],
			minPrice: undefined,
			maxPrice: undefined,
			sort: '-created'
		};
		onchange(filters);
	}

	const activeFiltersCount = $derived(
		Object.entries(filters).filter(([key, value]) => {
			if (key === 'sort') return false;
			if (Array.isArray(value)) return value.length > 0;
			return value !== undefined && value !== '';
		}).length
	);
</script>

<div class="space-y-4 px-2">
	<!-- Top Bar: Search & Sort & Mobile Toggle -->
	<div class="flex flex-col gap-4 lg:flex-row lg:items-center">
		<div class="relative flex-1">
			<Search class="absolute top-1/2 left-3 size-4 -translate-y-1/2 opacity-40" />
			<input
				type="text"
				placeholder="Search by name, bio, or university..."
				class="input input-md w-full bg-base-100 pl-10"
				bind:value={filters.q}
				oninput={() => updateFilter('q', filters.q)}
			/>
		</div>

		<div class="flex items-center gap-2">
			<select
				class="select min-w-[160px] bg-base-100 select-md"
				bind:value={filters.sort}
				onchange={() => updateFilter('sort', filters.sort)}
			>
				{#each SORT_OPTIONS as option}
					<option value={option.value}>{option.label}</option>
				{/each}
			</select>

			<Button
				variant="soft"
				color="neutral"
				class="lg:hidden"
				onclick={() => (showMobileFilters = !showMobileFilters)}
			>
				<SlidersHorizontal size={18} />
				Filters
				{#if activeFiltersCount > 0}
					<span class="badge badge-xs badge-primary">{activeFiltersCount}</span>
				{/if}
			</Button>
		</div>
	</div>

	<!-- Desktop Filters / Expandable Mobile Filters -->
	<div
		class={[
			'grid grid-cols-1 gap-4 lg:grid lg:grid-cols-4',
			showMobileFilters ? 'block' : 'hidden lg:grid'
		]}
	>
		<UniversitySelect
			value={filters.university || ''}
			onchange={(v) => updateFilter('university', v)}
			label="University"
			placeholder="All Universities"
		/>

		<TagInput
			selectedTags={filters.tags || []}
			onchange={(v) => updateFilter('tags', v)}
			label="Skills & Topics"
		/>

		<div class="form-control">
			<label class="label">
				<span class="label-text font-medium">Price Range ($)</span>
			</label>
			<div class="flex items-center gap-2">
				<input
					type="number"
					placeholder="Min"
					class="input input-md w-full bg-base-100"
					value={filters.minPrice ? filters.minPrice / 100 : ''}
					oninput={(e) => {
						const val = (e.target as HTMLInputElement).valueAsNumber;
						updateFilter('minPrice', isNaN(val) ? undefined : val * 100);
					}}
				/>
				<span class="opacity-30">-</span>
				<input
					type="number"
					placeholder="Max"
					class="input input-md w-full bg-base-100"
					value={filters.maxPrice ? filters.maxPrice / 100 : ''}
					oninput={(e) => {
						const val = (e.target as HTMLInputElement).valueAsNumber;
						updateFilter('maxPrice', isNaN(val) ? undefined : val * 100);
					}}
				/>
			</div>
		</div>

		<div class="flex flex-col justify-end">
			{#if activeFiltersCount > 0}
				<Button variant="ghost" color="error" class="gap-2" onclick={clearFilters}>
					<X size={16} />
					Clear Filters
				</Button>
			{:else}
				<div class="hidden h-12 lg:block"></div>
			{/if}
		</div>
	</div>
</div>
