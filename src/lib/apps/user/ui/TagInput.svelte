<script lang="ts">
	import { Plus, X, Search } from 'lucide-svelte';
	import { TAGS } from '../models';

	interface Props {
		selectedTags: string[];
		onchange: (tags: string[]) => void;
		label?: string;
	}

	let { selectedTags = $bindable([]), onchange, label }: Props = $props();

	let searchQuery = $state('');
	let isFocused = $state(false);

	const suggestions = $derived(
		TAGS.filter(
			(tag) =>
				tag.toLowerCase().includes(searchQuery.toLowerCase()) &&
				!selectedTags.includes(tag.toLowerCase())
		)
	);

	function addTag(tag: string) {
		const normalizedTag = tag.trim().toLowerCase();
		if (normalizedTag && !selectedTags.includes(normalizedTag)) {
			selectedTags = [...selectedTags, normalizedTag];
			onchange(selectedTags);
		}
		searchQuery = '';
	}

	function removeTag(tag: string) {
		selectedTags = selectedTags.filter((t) => t !== tag);
		onchange(selectedTags);
	}

	function handleKeydown(e: KeyboardEvent) {
		if (e.key === 'Enter' && searchQuery) {
			e.preventDefault();
			addTag(searchQuery);
		} else if (e.key === 'Backspace' && !searchQuery && selectedTags.length > 0) {
			removeTag(selectedTags[selectedTags.length - 1]);
		}
	}
</script>

<div class="form-control w-full">
	{#if label}
		<div class="label">
			<span class="label-text font-medium">{label}</span>
		</div>
	{/if}

	<div class="space-y-3">
		<div class="relative w-full">
			<div
				class={[
					'input input-md flex w-full items-center gap-2',
					isFocused && 'border-primary input-primary'
				]}
			>
				<Search size={18} class="opacity-40" />
				<input
					type="text"
					bind:value={searchQuery}
					onfocus={() => (isFocused = true)}
					onblur={() => setTimeout(() => (isFocused = false), 200)}
					onkeydown={handleKeydown}
					placeholder={label || 'Search or add tags...'}
					class="grow border-none bg-transparent outline-none focus:ring-0"
				/>
			</div>

			{#if isFocused && (suggestions.length > 0 || searchQuery)}
				<div
					class="absolute top-full left-0 z-[100] mt-1 max-h-60 w-full overflow-y-auto rounded-xl border border-base-200 bg-base-100 shadow-2xl"
				>
					<ul class="menu w-full menu-sm p-2">
						{#if suggestions.length > 0}
							<li
								class="menu-title px-2 pb-1 text-[10px] font-bold tracking-wider uppercase opacity-50"
							>
								Suggestions
							</li>
							{#each suggestions as suggestion}
								<li class="">
									<button type="button" onclick={() => addTag(suggestion)} class="rounded-md">
										<Search size={14} class="opacity-50" />
										{suggestion}
									</button>
								</li>
							{/each}
						{/if}

						{#if searchQuery && !suggestions.includes(searchQuery.toLowerCase()) && !selectedTags.includes(searchQuery.toLowerCase())}
							<li class={suggestions.length > 0 ? 'mt-1 border-t border-base-200 pt-1' : ''}>
								<button
									type="button"
									onclick={() => addTag(searchQuery)}
									class="rounded-md font-medium text-primary"
								>
									<Plus size={14} />
									Add "{searchQuery.toLowerCase()}"
								</button>
							</li>
						{/if}
					</ul>
				</div>
			{/if}
		</div>

		{#if selectedTags.length > 0}
			<div class="max-h-32 overflow-y-auto rounded-lg border border-base-200/50 bg-base-200/20 p-2">
				<div class="flex flex-wrap gap-2">
					{#each selectedTags as tag}
						<span
							class="badge flex items-center gap-1.5 badge-outline px-3 py-4 font-medium transition-all badge-secondary hover:scale-105"
						>
							<span class="text-sm">{tag}</span>
							<button
								type="button"
								class="rounded-full p-0.5 transition-colors hover:bg-black/10"
								onclick={() => removeTag(tag)}
							>
								<X size={14} />
							</button>
						</span>
					{/each}
				</div>
			</div>
		{/if}
	</div>
</div>
