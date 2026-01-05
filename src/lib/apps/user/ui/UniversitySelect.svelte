<script lang="ts">
	import { Search, Plus, GraduationCap } from 'lucide-svelte';
	import { UNIVERSITIES } from '../models';

	interface Props {
		value: string;
		onchange: (value: string) => void;
		label?: string;
		placeholder?: string;
		required?: boolean;
	}

	let { value = $bindable(''), onchange, label, placeholder, required }: Props = $props();

	let searchQuery = $state(value);
	let isFocused = $state(false);

	const suggestions = $derived(
		UNIVERSITIES.filter((uni) => uni.toLowerCase().includes(searchQuery.toLowerCase()))
	);

	function selectUniversity(uni: string) {
		value = uni;
		searchQuery = uni;
		onchange(value);
		isFocused = false;
	}

	function handleInput(e: Event) {
		const target = e.target as HTMLInputElement;
		searchQuery = target.value;
		value = searchQuery; // Allow custom input
		onchange(value);
	}

	function handleBlur() {
		// Small delay to allow clicking on suggestion
		setTimeout(() => {
			isFocused = false;
		}, 200);
	}
</script>

<div class="form-control w-full">
	{#if label}
		<div class="label">
			<span class="label-text font-medium">{label}</span>
		</div>
	{/if}

	<div class="relative w-full">
		<div
			class={[
				'input input-md flex w-full items-center gap-2',
				isFocused && 'border-primary input-primary'
			]}
		>
			<GraduationCap size={18} class="opacity-40" />
			<input
				type="text"
				value={searchQuery}
				oninput={handleInput}
				onfocus={() => (isFocused = true)}
				onblur={handleBlur}
				placeholder={placeholder || 'Search or enter university...'}
				class="grow border-none bg-transparent outline-none focus:ring-0"
				{required}
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
							<li>
								<button
									type="button"
									onclick={() => selectUniversity(suggestion)}
									class="rounded-md"
								>
									<GraduationCap size={14} class="opacity-50" />
									{suggestion}
								</button>
							</li>
						{/each}
					{/if}

					{#if searchQuery && !UNIVERSITIES.some((u) => u.toLowerCase() === searchQuery.toLowerCase())}
						<li class={suggestions.length > 0 ? 'mt-1 border-t border-base-200 pt-1' : ''}>
							<button
								type="button"
								onclick={() => selectUniversity(searchQuery)}
								class="rounded-md font-medium text-primary"
							>
								<Plus size={14} />
								Use "{searchQuery}"
							</button>
						</li>
					{/if}
				</ul>
			</div>
		{/if}
	</div>
</div>
