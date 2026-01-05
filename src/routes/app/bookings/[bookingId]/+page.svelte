<script lang="ts">
	import { pb } from '$lib';
	import {
		CheckCircle,
		Calendar,
		Clock,
		User,
		ArrowRight,
		Video,
		Timer,
		AlertCircle,
		XCircle,
		CreditCard
	} from 'lucide-svelte';
	import { Button } from '$lib/shared/ui';
	import { onMount } from 'svelte';

	let { data } = $props();
	const booking = $derived(data.booking);
	const mentor = $derived(booking.expand?.slot?.expand?.mentor);

	let timeLeft = $state('');
	let isExpired = $state(false);

	function updateTimer() {
		const created = new Date(booking.created).getTime();
		const expires = created + 15 * 60 * 1000;
		const now = new Date().getTime();
		const diff = expires - now;

		if (diff <= 0) {
			timeLeft = '00:00';
			isExpired = true;
			return;
		}

		const minutes = Math.floor(diff / 1000 / 60);
		const seconds = Math.floor((diff / 1000) % 60);
		timeLeft = `${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`;
	}

	onMount(() => {
		if (booking.status === 'pending') {
			updateTimer();
			const interval = setInterval(() => {
				updateTimer();
				if (isExpired) clearInterval(interval);
			}, 1000);
			return () => clearInterval(interval);
		}
	});

	function formatDateTime(isoString: string) {
		const date = new Date(isoString);
		return {
			date: date.toLocaleDateString('en-US', { weekday: 'long', month: 'long', day: 'numeric' }),
			time: date.toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit', hour12: false })
		};
	}

	const dt = $derived(formatDateTime(booking.expand?.slot?.start));

	const statusConfig: Record<
		string,
		{
			title: string;
			description: string;
			icon: typeof Timer;
			colorClass: string;
			bgClass: string;
			ringClass: string;
		}
	> = $derived(
		{
			pending: {
				title: 'Finish Your Booking',
				description:
					'This slot is reserved for you for the next 15 minutes. Please complete the payment to confirm your session.',
				icon: Timer,
				colorClass: 'text-warning',
				bgClass: 'bg-warning/10',
				ringClass: 'ring-warning/5'
			},
			confirmed: {
				title: 'Booking Confirmed!',
				description: `Your session with ${mentor?.name} is successfully scheduled.`,
				icon: CheckCircle,
				colorClass: 'text-success',
				bgClass: 'bg-success/10',
				ringClass: 'ring-success/5'
			},
			expired: {
				title: 'Booking Expired',
				description: 'The 15-minute reservation period has ended. The slot has been released.',
				icon: AlertCircle,
				colorClass: 'text-base-content/40',
				bgClass: 'bg-base-200',
				ringClass: 'ring-base-100'
			},
			canceled: {
				title: 'Booking Canceled',
				description: 'This booking has been canceled.',
				icon: XCircle,
				colorClass: 'text-error',
				bgClass: 'bg-error/10',
				ringClass: 'ring-error/5'
			}
		}[booking.status as keyof typeof statusConfig] || (statusConfig.pending as any)
	);
</script>

<div class="container mx-auto max-w-2xl px-4 py-20 text-center">
	<div class="mb-8 flex justify-center">
		<div
			class="rounded-full {statusConfig.bgClass} p-6 {statusConfig.colorClass} ring-8 {statusConfig.ringClass}"
		>
			<statusConfig.icon size={64} />
		</div>
	</div>

	<h1 class="font-display text-4xl font-black lg:text-5xl">{statusConfig.title}</h1>
	<p class="mt-4 text-lg opacity-60">
		{statusConfig.description}
	</p>

	{#if booking.status === 'pending' && !isExpired}
		<div class="mt-6 flex justify-center">
			<div
				class="flex items-center gap-3 rounded-2xl bg-warning/10 px-6 py-3 font-mono text-2xl font-black text-warning"
			>
				<Timer size={24} />
				{timeLeft}
			</div>
		</div>
	{/if}

	<div class="mt-12 rounded-3xl border border-base-200 bg-base-100 p-8 shadow-sm">
		<div class="grid grid-cols-1 gap-8 md:grid-cols-2">
			<div class="flex flex-col items-center gap-2 md:items-start">
				<div class="flex items-center gap-2 text-xs font-bold tracking-widest uppercase opacity-40">
					<Calendar size={14} />
					Date
				</div>
				<div class="text-xl font-bold">{dt.date}</div>
			</div>
			<div class="flex flex-col items-center gap-2 md:items-start">
				<div class="flex items-center gap-2 text-xs font-bold tracking-widest uppercase opacity-40">
					<Clock size={14} />
					Time
				</div>
				<div class="text-xl font-bold">{dt.time} ({booking.expand?.slot?.durationMinutes} min)</div>
			</div>
		</div>

		<div class="mt-8 flex flex-col items-center gap-4 border-t border-base-200 pt-8 md:flex-row">
			<div class="h-12 w-12 overflow-hidden rounded-xl bg-base-200">
				{#if mentor?.avatar}
					<img
						src={pb.files.getURL(mentor, mentor.avatar)}
						alt={mentor.name}
						class="h-full w-full object-cover"
					/>
				{:else}
					<div
						class="flex h-full w-full items-center justify-center bg-primary/10 font-bold text-primary"
					>
						{mentor?.name?.charAt(0)}
					</div>
				{/if}
			</div>
			<div class="text-center md:text-left">
				<div class="text-left text-sm font-medium opacity-40">Your Mentor</div>
				<div class="font-bold">{mentor?.name}</div>
			</div>
			<div class="mt-4 flex-1 text-center md:mt-0 md:text-right">
				<div class="text-sm font-medium opacity-40">Amount</div>
				<div class="font-mono text-xl font-bold">
					${(booking.agreedPriceCents / 100).toFixed(2)}
				</div>
			</div>
		</div>
	</div>

	<div class="mt-12 flex flex-col gap-4 sm:flex-row sm:justify-center">
		{#if booking.status === 'pending' && !isExpired}
			<Button color="primary" size="lg" class="gap-2 px-12">
				<CreditCard size={20} />
				Checkout
			</Button>
			<Button variant="ghost" size="lg" href="/app/bookings">Cancel</Button>
		{:else if booking.status === 'confirmed'}
			<Button href="/app/mentors" variant="outline" size="lg">Browse more mentors</Button>
			<Button href="/app/bookings" color="primary" size="lg" class="gap-2">
				Go to My Bookings
				<ArrowRight size={18} />
			</Button>
		{:else}
			<Button href="/app/mentors" color="primary" size="lg">Browse Mentors</Button>
			<Button href="/app/bookings" variant="ghost" size="lg">Back to My Bookings</Button>
		{/if}
	</div>

	{#if booking.status === 'confirmed'}
		<div class="mt-12 rounded-2xl bg-primary/5 p-6 text-sm opacity-70">
			<div class="mb-2 flex items-center justify-center gap-2 font-bold text-primary">
				<Video size={16} />
				Video Link
			</div>
			Meeting link will be available 5 minutes before the session starts in your dashboard.
		</div>
	{/if}
</div>
