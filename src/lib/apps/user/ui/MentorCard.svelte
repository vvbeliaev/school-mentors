<script lang="ts">
	import { pb } from '$lib';
	import type { MentorResponse } from '../mentors.svelte';
	import { Star, Verified, GraduationCap, Clock, DollarSign } from 'lucide-svelte';
	import { Button } from '$lib/shared/ui';

	interface Props {
		mentor: MentorResponse;
	}

	let { mentor }: Props = $props();

	const avatarUrl = $derived(
		mentor.avatar
			? pb.files.getURL(mentor, mentor.avatar, { thumb: '100x100' })
			: `https://ui-avatars.com/api/?name=${encodeURIComponent(mentor.name || 'Mentor')}&background=random`
	);

	const rate = $derived(mentor.hourRateCents ? (mentor.hourRateCents / 100).toFixed(0) : '0');
</script>

<div
	class="group relative flex flex-col rounded-3xl border border-base-200 bg-base-100 p-5 shadow-sm transition-all hover:-translate-y-1 hover:shadow-md"
>
	<div class="flex items-center gap-4">
		<div class="relative">
			<img
				src={avatarUrl}
				alt={mentor.name}
				class="h-16 w-16 rounded-full border-2 border-base-100 object-cover shadow-sm ring-1 ring-base-200"
			/>
			{#if mentor.isVerified}
				<div
					class="absolute -right-1 -bottom-1 flex h-5 w-5 items-center justify-center rounded-full bg-success text-white ring-2 ring-base-100"
					title="Verified Mentor"
				>
					<Verified size={12} />
				</div>
			{/if}
		</div>
		<div class="flex-1 overflow-hidden">
			<div class="flex items-center justify-between gap-2">
				<h3 class="truncate font-display text-lg font-bold">{mentor.name}</h3>
				{#if mentor.rating}
					<div class="flex items-center gap-1 text-sm font-bold text-warning">
						<Star size={14} fill="currentColor" />
						{mentor.rating}
					</div>
				{/if}
			</div>
			<div class="flex items-center gap-1.5 text-sm font-medium opacity-60">
				<GraduationCap size={14} />
				<span class="truncate">{mentor.university || 'University not specified'}</span>
			</div>
		</div>
	</div>

	{#if mentor.tags && Array.isArray(mentor.tags) && mentor.tags.length > 0}
		<div class="mt-4 flex flex-wrap gap-1.5">
			{#each mentor.tags.slice(0, 3) as tag}
				<span class="badge badge-ghost badge-sm font-medium">{tag}</span>
			{/each}
			{#if mentor.tags.length > 3}
				<span class="badge badge-ghost badge-sm font-medium opacity-50"
					>+{mentor.tags.length - 3}</span
				>
			{/if}
		</div>
	{/if}

	{#if mentor.bio}
		<p class="mt-4 line-clamp-2 text-sm leading-relaxed opacity-70">
			{mentor.bio}
		</p>
	{/if}

	<div class="mt-auto pt-6">
		<div class="flex items-center justify-between border-t border-base-200 pt-4">
			<div class="flex flex-col">
				<span class="text-[10px] font-bold tracking-wider uppercase opacity-40">Hourly Rate</span>
				<div class="flex items-baseline gap-0.5">
					<span class="text-sm font-bold text-primary">$</span>
					<span class="font-display text-2xl font-black text-primary">{rate}</span>
				</div>
			</div>
			<Button href="/app/mentors/{mentor.id}" size="sm" variant="soft" color="primary">
				View Profile
			</Button>
		</div>
	</div>
</div>
