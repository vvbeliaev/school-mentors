<script lang="ts">
	import { pb } from '$lib';
	import type { SlotsResponse } from '$lib/shared/pb/pocketbase-types';
	import { Calendar, Clock, ChevronRight, CheckCircle2, Loader2 } from 'lucide-svelte';
	import { Button } from '$lib/shared/ui';
	import { bookingsApi } from '$lib/apps/bookings/api';
	import { goto } from '$app/navigation';

	interface Props {
		slots: SlotsResponse[];
		mentorName: string;
	}

	let { slots, mentorName }: Props = $props();

	let selectedSlotId = $state<string | null>(null);
	let isSubmitting = $state(false);

	const sortedSlots = $derived(
		[...slots].sort((a, b) => new Date(a.start).getTime() - new Date(b.start).getTime())
	);

	const groupedSlots = $derived(
		sortedSlots.reduce(
			(acc, slot) => {
				const date = new Date(slot.start).toLocaleDateString('en-US', {
					weekday: 'short',
					month: 'short',
					day: 'numeric'
				});
				if (!acc[date]) acc[date] = [];
				acc[date].push(slot);
				return acc;
			},
			{} as Record<string, SlotsResponse[]>
		)
	);

	async function handleBook() {
		if (!selectedSlotId) return;

		isSubmitting = true;
		try {
			const booking = await bookingsApi.createBooking(selectedSlotId);
			// Redirect to a booking confirmation or payment page
			// For now, let's just go to a success state or back to mentors
			goto(`/app/bookings/${booking.id}`);
		} catch (err) {
			console.error('Failed to book slot:', err);
			alert('Failed to book the slot. It might have been taken already.');
		} finally {
			isSubmitting = false;
		}
	}

	function formatTime(isoString: string) {
		return new Date(isoString).toLocaleTimeString('en-US', {
			hour: '2-digit',
			minute: '2-digit',
			hour12: false
		});
	}
</script>

<div class="flex flex-col rounded-3xl border border-base-200 bg-base-100 p-6 shadow-xl lg:p-8">
	<div class="mb-6 flex items-center justify-between">
		<h3 class="font-display text-xl font-bold">Available Slots</h3>
		<span class="badge badge-outline font-bold badge-primary">{slots.length} available</span>
	</div>

	{#if slots.length === 0}
		<div class="flex flex-col items-center justify-center py-12 text-center">
			<Calendar size={48} class="mb-4 opacity-10" />
			<p class="font-medium opacity-50">No upcoming slots found.</p>
			<p class="text-sm opacity-40">Check back later for new availability.</p>
		</div>
	{:else}
		<div class="max-h-[400px] space-y-6 overflow-y-auto pr-2">
			{#each Object.entries(groupedSlots) as [date, daySlots]}
				<div class="space-y-3">
					<div class="flex items-center gap-2">
						<span class="text-xs font-bold tracking-widest uppercase opacity-40">{date}</span>
						<div class="h-px flex-1 bg-base-200"></div>
					</div>
					<div class="grid grid-cols-2 gap-2 sm:grid-cols-3">
						{#each daySlots as slot}
							<button
								type="button"
								onclick={() => (selectedSlotId = slot.id)}
								class={[
									'group flex flex-col items-center justify-center rounded-xl border p-3 transition-all',
									selectedSlotId === slot.id
										? 'border-primary bg-primary/10 text-primary shadow-sm'
										: 'border-base-200 hover:border-primary/50 hover:bg-base-200/50'
								]}
							>
								<Clock size={16} class="mb-1 opacity-40 group-hover:opacity-100" />
								<span class="text-sm font-bold">{formatTime(slot.start)}</span>
								<span class="text-[10px] opacity-40">{slot.durationMinutes} min</span>
							</button>
						{/each}
					</div>
				</div>
			{/each}
		</div>

		<div class="mt-8 border-t border-base-200 pt-6">
			<Button
				variant="solid"
				color="primary"
				size="lg"
				block
				disabled={!selectedSlotId || isSubmitting}
				onclick={handleBook}
			>
				{#if isSubmitting}
					<Loader2 size={20} class="mr-2 animate-spin" />
					Processing...
				{:else}
					Book with {mentorName}
					<ChevronRight size={18} class="ml-2" />
				{/if}
			</Button>
			<p class="mt-4 text-center text-xs opacity-40">
				Instant confirmation. Cancel for free up to 24h before.
			</p>
		</div>
	{/if}
</div>
