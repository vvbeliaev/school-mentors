<script lang="ts">
	import { bookingsStore } from '$lib/apps/bookings';
	import { userStore } from '$lib/apps/user';
	import { pb } from '$lib';
	import { Calendar, Clock, User, ChevronLeft, ChevronRight, ExternalLink } from 'lucide-svelte';
	import { Button } from '$lib/shared/ui';

	const bookings = $derived(bookingsStore.menteeBookings);
	const totalItems = $derived(bookingsStore.totalItems);
	const perPage = $derived(bookingsStore.perPage);
	const page = $derived(bookingsStore.page);
	const totalPages = $derived(Math.ceil(totalItems / perPage));

	async function goToPage(newPage: number) {
		if (userStore.user?.id) {
			await bookingsStore.load(userStore.user.id, newPage);
		}
	}

	function formatDateTime(isoString: string) {
		const date = new Date(isoString);
		return {
			date: date.toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' }),
			time: date.toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit', hour12: false })
		};
	}

	function getStatusClass(status: string) {
		switch (status) {
			case 'confirmed':
				return 'badge-success';
			case 'pending':
				return 'badge-warning';
			case 'canceled':
				return 'badge-error';
			case 'expired':
				return 'badge-ghost opacity-50';
			default:
				return 'badge-neutral';
		}
	}
</script>

<div class="p-4 md:p-8">
	<div class="mb-8 flex items-center justify-between">
		<div>
			<h1 class="font-display text-3xl font-black">My Bookings</h1>
			<p class="text-base-content/60">View and manage your scheduled sessions.</p>
		</div>
	</div>

	{#if bookings.length === 0}
		<div
			class="flex min-h-[400px] flex-col items-center justify-center rounded-3xl border border-dashed border-base-300 bg-base-100/50 p-12 text-center"
		>
			<div class="mb-4 rounded-full bg-base-200 p-6 opacity-20">
				<Calendar size={48} />
			</div>
			<h3 class="text-xl font-bold">No bookings found</h3>
			<p class="mt-2 max-w-xs opacity-60">
				You haven't scheduled any sessions yet. Browse mentors to get started!
			</p>
			<Button href="/app/mentors" color="primary" class="mt-8">Find a Mentor</Button>
		</div>
	{:else}
		<div class="overflow-hidden rounded-2xl border border-base-200 bg-base-100 shadow-sm">
			<table class="table w-full table-zebra">
				<thead>
					<tr class="bg-base-200/50">
						<th>Mentor / Mentee</th>
						<th>Date & Time</th>
						<th>Price</th>
						<th>Status</th>
						<th class="text-right">Actions</th>
					</tr>
				</thead>
				<tbody>
					{#each bookings as booking}
						{@const isMentor = booking.expand?.slot?.mentor === userStore.user?.id}
						{@const otherUser = isMentor
							? booking.expand?.mentee
							: booking.expand?.slot?.expand?.mentor}
						{@const dt = formatDateTime(booking.expand?.slot?.start)}
						<tr class="hover">
							<td>
								<div class="flex items-center gap-3">
									<div class="placeholder avatar">
										<div class="w-10 rounded-xl bg-base-300">
											{#if otherUser?.avatar}
												<img
													src={pb.files.getURL(otherUser, otherUser.avatar)}
													alt={otherUser.name}
												/>
											{:else}
												<span class="text-xs">{otherUser?.name?.charAt(0) || 'U'}</span>
											{/if}
										</div>
									</div>
									<div>
										<div class="font-bold">{otherUser?.name || 'Deleted User'}</div>
										<div class="text-xs opacity-50">{isMentor ? 'Mentee' : 'Mentor'}</div>
									</div>
								</div>
							</td>
							<td>
								<div class="flex flex-col">
									<span class="flex items-center gap-1 font-medium">
										<Calendar size={12} class="opacity-50" />
										{dt.date}
									</span>
									<span class="flex items-center gap-1 text-xs opacity-60">
										<Clock size={12} class="opacity-50" />
										{dt.time} ({booking.expand?.slot?.durationMinutes} min)
									</span>
								</div>
							</td>
							<td>
								<div class="font-mono text-sm">${(booking.agreedPriceCents / 100).toFixed(2)}</div>
							</td>
							<td>
								<span class="badge badge-sm font-bold uppercase {getStatusClass(booking.status)}">
									{booking.status}
								</span>
							</td>
							<td class="text-right">
								<Button href="/app/bookings/{booking.id}" variant="ghost" size="sm" class="gap-1">
									Details
									<ExternalLink size={14} />
								</Button>
							</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>

		{#if totalPages > 1}
			<div class="mt-8 flex justify-center">
				<div class="join">
					<button
						class="btn join-item btn-sm"
						disabled={page === 1}
						onclick={() => goToPage(page - 1)}
					>
						<ChevronLeft size={16} />
					</button>
					<button class="no-animation btn pointer-events-none join-item btn-sm">
						Page {page} of {totalPages}
					</button>
					<button
						class="btn join-item btn-sm"
						disabled={page === totalPages}
						onclick={() => goToPage(page + 1)}
					>
						<ChevronRight size={16} />
					</button>
				</div>
			</div>
		{/if}
	{/if}
</div>
