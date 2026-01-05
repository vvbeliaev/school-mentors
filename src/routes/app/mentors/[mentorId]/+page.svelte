<script lang="ts">
	import { pb } from '$lib';
	import type { MentorResponse } from '$lib/apps/user/mentors.svelte';
	import {
		GraduationCap,
		Star,
		Verified,
		DollarSign,
		Clock,
		BookOpen,
		Award,
		MessageSquare,
		Share2,
		ChevronLeft
	} from 'lucide-svelte';
	import { Button } from '$lib/shared/ui';
	import BookingWidget from '$lib/apps/slots/ui/BookingWidget.svelte';

	let { data } = $props();
	const mentor = $derived(data.mentor as MentorResponse);
	const slots = $derived(data.slots);

	// Add mock rating for visual beauty if not present
	const displayRating = $derived(mentor.rating || 4.9);

	const avatarUrl = $derived(
		mentor.avatar
			? pb.files.getURL(mentor, mentor.avatar, { thumb: '300x300' })
			: `https://ui-avatars.com/api/?name=${encodeURIComponent(mentor.name || 'Mentor')}&size=300&background=random`
	);

	const rate = $derived(mentor.hourRateCents ? (mentor.hourRateCents / 100).toFixed(0) : '0');
</script>

<div class="min-h-screen bg-base-100 pb-20">
	<!-- Top Navigation -->
	<!-- <div class="container mx-auto px-4 py-6">
		<Button href="/app/mentors" variant="ghost" size="sm" class="gap-2">
			<ChevronLeft size={18} />
			Back to mentors
		</Button>
	</div> -->

	<!-- Profile Hero Header -->
	<div class="relative overflow-hidden border-b border-base-200 bg-base-200/30 py-12 lg:py-20">
		<!-- Background Decorative Element -->
		<div class="absolute -top-24 -right-24 h-96 w-96 rounded-full bg-primary/5 blur-3xl"></div>
		<div class="absolute -bottom-24 -left-24 h-96 w-96 rounded-full bg-secondary/5 blur-3xl"></div>

		<div class="relative z-10 container mx-auto px-4">
			<div class="flex flex-col items-center gap-8 lg:flex-row lg:items-start lg:gap-16">
				<!-- Avatar Section -->
				<div class="relative">
					<div
						class="h-40 w-40 overflow-hidden rounded-[2.5rem] border-4 border-base-100 shadow-2xl lg:h-56 lg:w-56"
					>
						<img src={avatarUrl} alt={mentor.name} class="h-full w-full object-cover" />
					</div>
					{#if mentor.isVerified}
						<div
							class="absolute -right-2 -bottom-2 flex h-10 w-10 items-center justify-center rounded-2xl bg-success text-white shadow-lg ring-4 ring-base-100 lg:h-12 lg:w-12"
							title="Verified Mentor"
						>
							<Verified size={24} />
						</div>
					{/if}
				</div>

				<!-- Info Section -->
				<div class="flex-1 text-center lg:text-left">
					<div class="flex flex-col items-center gap-3 lg:flex-row lg:items-center">
						<h1 class="font-display text-4xl font-black lg:text-6xl">{mentor.name}</h1>
						<div
							class="flex items-center gap-1.5 rounded-full bg-warning/10 px-4 py-1.5 text-sm font-bold text-warning ring-1 ring-warning/20"
						>
							<Star size={16} fill="currentColor" />
							{displayRating}
						</div>
					</div>

					<div class="mt-4 flex flex-wrap justify-center gap-6 lg:justify-start">
						<div class="flex items-center gap-2 text-lg font-medium opacity-60">
							<GraduationCap size={20} class="text-primary" />
							{mentor.university}
						</div>
						{#if mentor.degree}
							<div class="flex items-center gap-2 text-lg font-medium opacity-60">
								<Award size={20} class="text-secondary" />
								{mentor.degree}
							</div>
						{/if}
					</div>

					{#if mentor.tags && Array.isArray(mentor.tags) && mentor.tags.length > 0}
						<div class="mt-8 flex flex-wrap justify-center gap-2 lg:justify-start">
							{#each mentor.tags as tag}
								<span
									class="badge badge-outline border-base-300 px-4 py-4 badge-md font-semibold opacity-70"
								>
									{tag}
								</span>
							{/each}
						</div>
					{/if}

					<div
						class="mt-10 flex items-center justify-center gap-8 border-t border-base-200 pt-8 lg:justify-start"
					>
						<div class="flex flex-col">
							<span class="text-xs font-bold tracking-widest uppercase opacity-40">Hourly Rate</span
							>
							<div class="flex items-baseline gap-1">
								<span class="text-xl font-bold text-primary">$</span>
								<span class="font-display text-4xl font-black text-primary">{rate}</span>
							</div>
						</div>
						<div class="h-10 w-px bg-base-200"></div>
						<!-- <div class="flex flex-col">
							<span class="text-xs font-bold tracking-widest uppercase opacity-40"
								>Response time</span
							>
							<div class="flex items-baseline gap-1">
								<span class="font-display text-3xl font-black italic">~2h</span>
							</div>
						</div> -->
					</div>
				</div>
			</div>
		</div>
	</div>

	<!-- Content Section -->
	<div class="container mx-auto mt-4 px-4 lg:mt-8">
		<div class="grid grid-cols-1 gap-12 lg:grid-cols-12">
			<!-- Left: Bio & Experience -->
			<div class="space-y-12 lg:col-span-7">
				<section>
					<h2 class="mb-6 flex items-center gap-3 font-display text-2xl font-bold lg:text-3xl">
						<BookOpen size={28} class="text-primary" />
						About Me
					</h2>
					<div class="prose prose-lg max-w-none leading-relaxed text-neutral">
						{mentor.bio || 'No bio provided.'}
					</div>
				</section>

				<!-- <section class="rounded-3xl bg-base-200/50 p-8">
					<h3 class="mb-4 text-xl font-bold">Why book with me?</h3>
					<ul class="space-y-4">
						<li class="flex items-start gap-3 font-medium">
							<div class="mt-1 rounded-full bg-success/20 p-1 text-success">
								<Verified size={14} />
							</div>
							Direct experience with {mentor.university} admissions process.
						</li>
						<li class="flex items-start gap-3 font-medium">
							<div class="mt-1 rounded-full bg-success/20 p-1 text-success">
								<Verified size={14} />
							</div>
							Tailored guidance based on your current level and goals.
						</li>
						<li class="flex items-start gap-3 font-medium">
							<div class="mt-1 rounded-full bg-success/20 p-1 text-success">
								<Verified size={14} />
							</div>
							Access to shared resources, notes, and study plans.
						</li>
					</ul>
				</section> -->
			</div>

			<!-- Right: Booking Widget -->
			<div class="lg:col-span-5">
				<div class="sticky top-24">
					<BookingWidget {slots} mentorName={mentor.name} />

					<div class="mt-8 flex items-center justify-center gap-4">
						<Button variant="ghost" size="sm" class="gap-2">
							<MessageSquare size={18} />
							Ask a question
						</Button>
						<div class="h-4 w-px bg-base-200"></div>
						<Button variant="ghost" size="sm" class="gap-2">
							<Share2 size={18} />
							Share profile
						</Button>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
