<script lang="ts">
	import { userStore } from '$lib/apps/user';
	import { bookingsStore } from '$lib/apps/bookings';
	import { pb } from '$lib';
	import {
		Calendar,
		Clock,
		ExternalLink,
		ArrowRight,
		Sparkles,
		ShieldCheck,
		Zap
	} from 'lucide-svelte';
	import { Button } from '$lib/shared/ui';

	const user = $derived(userStore.user);
	const bookings = $derived(bookingsStore.menteeBookings as any[]);

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
	{#if !user}
		<div class="flex min-h-[70vh] flex-col items-center justify-center text-center">
			<div class="relative mb-8">
				<div
					class="absolute -inset-4 animate-pulse rounded-full bg-linear-to-r from-primary/20 to-secondary/20 blur-2xl"
				></div>
				<div
					class="relative flex h-24 w-24 items-center justify-center rounded-3xl bg-base-100 shadow-xl"
				>
					<Sparkles size={48} class="text-primary" />
				</div>
			</div>

			<h1 class="max-w-2xl font-display text-4xl font-black text-pretty md:text-6xl">
				Unlock Your Potential with <span
					class="bg-linear-to-r from-primary to-secondary bg-clip-text text-transparent"
					>Near-Peer Mentorship</span
				>
			</h1>

			<p class="mt-6 max-w-xl text-lg text-base-content/60">
				Connect with college students who've already walked the path you're on. Get personalized
				guidance, SAT prep, and career advice.
			</p>

			<div class="mt-10 flex flex-col items-center gap-4 sm:flex-row">
				<Button
					href="/app/auth/sign-up"
					color="primary"
					size="lg"
					class="px-8 shadow-xl shadow-primary/20"
				>
					Create Account
					<ArrowRight size={20} />
				</Button>
				<Button href="/app/mentors" variant="ghost" size="lg">Browse Mentors</Button>
			</div>

			<div class="mt-16 grid grid-cols-1 gap-8 md:grid-cols-3">
				<div class="flex flex-col items-center gap-2">
					<div class="rounded-2xl bg-success/10 p-3 text-success">
						<ShieldCheck size={24} />
					</div>
					<div class="font-bold">Verified Mentors</div>
					<div class="text-sm opacity-60">Hand-picked students from top universities.</div>
				</div>
				<div class="flex flex-col items-center gap-2">
					<div class="rounded-2xl bg-warning/10 p-3 text-warning">
						<Zap size={24} />
					</div>
					<div class="font-bold">Instant Booking</div>
					<div class="text-sm opacity-60">Schedule sessions in seconds, not days.</div>
				</div>
				<div class="flex flex-col items-center gap-2">
					<div class="rounded-2xl bg-primary/10 p-3 text-primary">
						<Calendar size={24} />
					</div>
					<div class="font-bold">Flexible Learning</div>
					<div class="text-sm opacity-60">1-on-1 sessions on your own schedule.</div>
				</div>
			</div>
		</div>
	{:else}
		<div class="mb-8">
			<h1 class="font-display text-3xl font-black">Welcome back, {user.name}!</h1>
			<p class="text-base-content/60">Here is an overview of your activity.</p>
		</div>

		<div class="grid grid-cols-1 gap-6 lg:grid-cols-3">
			<div class="lg:col-span-2">
				<div class="mb-4 flex items-center justify-between">
					<h2 class="text-xl font-bold">Your Bookings</h2>
					<Button href="/app/bookings" variant="ghost" size="sm">View All</Button>
				</div>

				{#if bookings.length === 0}
					<div
						class="flex min-h-[300px] flex-col items-center justify-center rounded-3xl border border-dashed border-base-300 bg-base-100/50 p-8 text-center"
					>
						<div class="mb-4 rounded-full bg-base-200 p-4 opacity-20">
							<Calendar size={32} />
						</div>
						<h3 class="text-lg font-bold">No bookings yet</h3>
						<p class="mt-2 max-w-xs text-sm opacity-60">
							Your scheduled sessions will appear here.
						</p>
						<Button href="/app/mentors" color="primary" class="mt-6">Find a Mentor</Button>
					</div>
				{:else}
					<div class="overflow-hidden rounded-2xl border border-base-200 bg-base-100 shadow-sm">
						<table class="table w-full table-zebra">
							<thead>
								<tr class="bg-base-200/50">
									<th>Mentor / Mentee</th>
									<th>Date & Time</th>
									<th>Status</th>
									<th class="text-right"></th>
								</tr>
							</thead>
							<tbody>
								{#each bookings.slice(0, 5) as booking}
									{@const isMentor = booking.expand?.slot?.mentor === userStore.user?.id}
									{@const otherUser = isMentor
										? booking.expand?.mentee
										: booking.expand?.slot?.expand?.mentor}
									{@const dt = formatDateTime(booking.expand?.slot?.start)}
									<tr class="hover">
										<td>
											<div class="flex items-center gap-3">
												<div class="placeholder avatar">
													<div class="w-8 rounded-lg bg-base-300">
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
													<div class="text-sm font-bold">{otherUser?.name || 'Deleted User'}</div>
												</div>
											</div>
										</td>
										<td>
											<div class="flex flex-col text-xs">
												<span class="font-medium">{dt.date}</span>
												<span class="opacity-60">{dt.time}</span>
											</div>
										</td>
										<td>
											<span
												class="badge badge-xs font-bold uppercase {getStatusClass(booking.status)}"
											>
												{booking.status}
											</span>
										</td>
										<td class="text-right">
											<Button href="/app/bookings/{booking.id}" variant="ghost" size="xs">
												<ExternalLink size={14} />
											</Button>
										</td>
									</tr>
								{/each}
							</tbody>
						</table>
					</div>
				{/if}
			</div>

			<div class="space-y-6">
				<div class="rounded-3xl bg-primary p-6 text-primary-content shadow-xl shadow-primary/20">
					<h3 class="text-xl font-bold">Find a Mentor</h3>
					<p class="mt-2 text-sm opacity-80">
						Browse our catalog of verified student mentors to find the perfect match for your needs.
					</p>
					<Button href="/app/mentors" class="mt-6 w-full bg-white text-primary hover:bg-white/90">
						Browse All Mentors
						<ArrowRight size={18} />
					</Button>
				</div>

				<div class="rounded-3xl bg-base-200 p-6">
					<h3 class="text-lg font-bold">Quick Tips</h3>
					<ul class="mt-4 space-y-3 text-sm">
						<li class="flex gap-2">
							<div class="mt-1 h-1.5 w-1.5 shrink-0 rounded-full bg-primary"></div>
							<span>Complete your profile to get better recommendations.</span>
						</li>
						<li class="flex gap-2">
							<div class="mt-1 h-1.5 w-1.5 shrink-0 rounded-full bg-primary"></div>
							<span>Book sessions at least 24h in advance.</span>
						</li>
					</ul>
				</div>
			</div>
		</div>
	{/if}
</div>
