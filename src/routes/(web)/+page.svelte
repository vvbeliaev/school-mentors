<script lang="ts">
	import {
		GraduationCap,
		ArrowRight,
		Menu,
		X,
		CheckCircle,
		Video,
		Calendar,
		Search,
		Star,
		Verified,
		DollarSign,
		Sparkles,
		Clock
	} from 'lucide-svelte';

	import { Logo, Sidebar, ThemeController } from '$lib';

	let mobileMenuOpen = $state(false);

	const navigation = [
		{ name: 'Mentors', href: '#mentors' },
		{ name: 'How it Works', href: '#how-it-works' },
		{ name: 'Pricing', href: '#pricing' },
		{ name: 'Become a Mentor', href: '/app/auth/sign-up' }
	];

	const mentors = [
		{
			name: 'Sarah Jenkins',
			university: 'Stanford University',
			degree: 'Computer Science',
			rate: 45,
			rating: 4.9,
			reviews: 124,
			tags: ['Python', 'Algorithms', 'AI'],
			avatar: 'https://i.pravatar.cc/150?u=sarah',
			accent: 'bg-primary/10'
		},
		{
			name: 'David Chen',
			university: 'MIT',
			degree: 'Mathematics',
			rate: 55,
			rating: 5.0,
			reviews: 89,
			tags: ['Calculus', 'Linear Algebra', 'Physics'],
			avatar: 'https://i.pravatar.cc/150?u=david',
			accent: 'bg-secondary/10'
		},
		{
			name: 'Elena Rodriguez',
			university: 'Oxford University',
			degree: 'Economics & History',
			rate: 40,
			rating: 4.8,
			reviews: 215,
			tags: ['Microeconomics', 'Academic Writing'],
			avatar: 'https://i.pravatar.cc/150?u=elena',
			accent: 'bg-accent/10'
		}
	];

	const steps = [
		{
			title: 'Search',
			desc: 'Find mentors from top universities who succeeded where you want to go.',
			icon: Search
		},
		{
			title: 'Book',
			desc: 'Pick a 15, 30, or 60-minute slot that fits your schedule.',
			icon: Calendar
		},
		{
			title: 'Succeed',
			desc: 'Get direct advice via our integrated HD video platform.',
			icon: Video
		}
	];
</script>

<svelte:head>
	<title>Univerlink</title>
	<meta
		name="description"
		content="Connect with top college students for personalized 1-on-1 coaching. Get the inside track from those who just did it."
	/>
</svelte:head>

<div class="min-h-screen bg-base-100">
	<!-- Sticky Header -->
	<header
		class="fixed top-0 z-40 w-full border-b border-base-200 bg-base-100/80 backdrop-blur-md transition-all"
	>
		<div class="container mx-auto flex h-16 items-center justify-between px-6 lg:h-20">
			<!-- Logo -->
			<a href="/" class="flex items-center gap-2.5">
				<Logo />
			</a>

			<!-- Desktop Nav -->
			<nav class="hidden items-center gap-8 md:flex">
				{#each navigation as item}
					<a
						href={item.href}
						class="text-sm font-semibold opacity-70 transition-opacity hover:opacity-100"
					>
						{item.name}
					</a>
				{/each}
			</nav>

			<!-- Desktop Auth -->
			<div class="hidden items-center gap-3 md:flex">
				<ThemeController />
				<a href="/app/auth/sign-in" class="btn btn-ghost">Log in</a>
				<a href="/app/auth/sign-up" class="btn px-6 btn-primary">Sign up free</a>
			</div>

			<!-- Mobile Menu Button -->
			<button
				class="btn btn-square btn-ghost md:hidden"
				onclick={() => (mobileMenuOpen = true)}
				aria-label="Open menu"
			>
				<Menu size={24} />
			</button>
		</div>
	</header>

	<!-- Mobile Sidebar -->
	<Sidebar
		open={mobileMenuOpen}
		position="right"
		mobileOnly={true}
		onclose={() => (mobileMenuOpen = false)}
	>
		{#snippet header()}
			<div class="flex items-center gap-2">
				<GraduationCap class="size-6 text-primary" />
				<span class="font-display font-bold">Menu</span>
			</div>
		{/snippet}

		<div class="flex flex-col gap-2 p-4">
			{#each navigation as item}
				<a
					href={item.href}
					class="rounded-lg p-3 text-lg font-medium transition-colors hover:bg-base-200"
					onclick={() => (mobileMenuOpen = false)}
				>
					{item.name}
				</a>
			{/each}
			<div class="mt-8 flex flex-col gap-3 border-t border-base-200 pt-8">
				<a href="/app/auth/sign-in" class="btn btn-block btn-outline">Log in</a>
				<a href="/app/auth/sign-up" class="btn btn-block btn-primary">Sign up free</a>
			</div>
		</div>
	</Sidebar>

	<main class="pt-16 lg:pt-20">
		<!-- Hero Section: Linktree-inspired Split Layout -->
		<section id="mentors" class="relative overflow-hidden py-12 lg:py-24">
			<div class="container mx-auto px-6">
				<div class="grid grid-cols-1 items-center gap-16 lg:grid-cols-12">
					<!-- Left: Content -->
					<div class="lg:col-span-7">
						<h1
							class="font-display text-5xl leading-[1.05] font-black tracking-tighter sm:text-7xl lg:text-8xl"
						>
							The <span class="text-primary">Near-Peer</span> link to your future.
						</h1>
						<p class="mt-8 max-w-xl text-lg text-neutral lg:text-xl">
							Connect with top college students for personalized 1-on-1 coaching. Get the inside
							track from those who just did it.
						</p>

						<div class="mt-12 flex flex-col gap-4 sm:flex-row sm:items-center">
							<a href="/app/auth/sign-up" class="btn gap-3 px-10 btn-lg btn-primary">
								Start your journey <ArrowRight size={20} />
							</a>
						</div>
					</div>

					<!-- Right: Floating Mentor Cards (Wow effect) -->
					<div class="relative lg:col-span-5">
						<div class="relative flex flex-col gap-6 lg:scale-110">
							{#each mentors as mentor, i (mentor.name)}
								<div
									class="group relative z-10 flex flex-col rounded-3xl border border-base-200 bg-base-100 p-5 shadow-lg transition-all hover:-translate-y-1 hover:shadow-xl"
									style="transform: translateX({i * 1.5}rem)"
								>
									<div class="flex items-center gap-4">
										<div class="relative">
											<img
												src={mentor.avatar}
												alt={mentor.name}
												class="h-14 w-14 rounded-full border-2 border-white shadow-sm"
											/>
											<div
												class="absolute -right-1 -bottom-1 flex h-5 w-5 items-center justify-center rounded-full bg-success text-white ring-2 ring-white"
											>
												<Verified size={12} />
											</div>
										</div>
										<div class="flex-1">
											<div class="flex items-center justify-between">
												<h3 class="font-display text-lg font-bold">{mentor.name}</h3>
												<div class="flex items-center gap-1 text-sm font-bold text-warning">
													<Star size={14} fill="currentColor" />
													{mentor.rating}
												</div>
											</div>
											<p class="text-sm font-medium opacity-60">{mentor.university}</p>
										</div>
									</div>

									<div class="mt-4 flex flex-wrap gap-2">
										{#each mentor.tags as tag}
											<span class="badge badge-ghost badge-sm font-medium">{tag}</span>
										{/each}
									</div>

									<div class="mt-6 flex items-center justify-between border-t border-base-200 pt-4">
										<div class="flex flex-col">
											<span class="text-xs font-bold uppercase opacity-40">Hourly Rate</span>
											<span class="font-display text-2xl font-black text-primary"
												>${mentor.rate}</span
											>
										</div>
										<a href="/app/auth/sign-up" class="btn gap-2 px-4 btn-outline btn-sm">
											View Profile
										</a>
									</div>
								</div>
							{/each}

							<!-- Decorative Background Glow -->
							<div
								class="absolute -inset-4 -z-10 rounded-[3rem] bg-primary/20 blur-3xl lg:block"
							></div>
						</div>
					</div>
				</div>
			</div>
		</section>

		<!-- Features / How it Works -->
		<section id="how-it-works" class="bg-base-200/50 py-24">
			<div class="container mx-auto px-6">
				<div class="mb-20 grid grid-cols-1 items-end gap-8 lg:grid-cols-2">
					<div>
						<h2 class="font-display text-4xl font-black lg:text-6xl">
							Simplicity in <br /> every step.
						</h2>
					</div>
					<p class="text-lg text-neutral lg:text-xl">
						We removed the gatekeepers. Find a mentor, choose a slot, and start learning. No long
						contracts, no hidden fees.
					</p>
				</div>

				<div class="grid grid-cols-1 gap-8 md:grid-cols-3">
					{#each steps as step, i (step.title)}
						<div
							class="group rounded-[2.5rem] border border-base-300 bg-base-100 p-10 transition-all hover:border-primary/50"
						>
							<div
								class="mb-8 flex h-20 w-20 items-center justify-center rounded-3xl bg-primary/10 text-primary transition-transform group-hover:scale-110"
							>
								<step.icon size={40} />
							</div>
							<h3 class="mb-4 font-display text-3xl font-bold">{step.title}</h3>
							<p class="text-lg leading-relaxed text-neutral">
								{step.desc}
							</p>
						</div>
					{/each}
				</div>
			</div>
		</section>

		<!-- Pricing / Value Proposition -->
		<section id="pricing" class="py-24">
			<div class="container mx-auto px-6">
				<div
					class="relative overflow-hidden rounded-[3rem] bg-base-300 p-8 text-base-content lg:p-20"
				>
					<div class="relative z-10 grid grid-cols-1 items-center gap-16 lg:grid-cols-2">
						<div>
							<h2 class="mb-8 font-display text-4xl font-black lg:text-6xl">
								Transparent <br /> pricing.
							</h2>
							<ul class="space-y-6">
								<li class="flex items-center gap-4 text-xl font-medium">
									<CheckCircle class="text-secondary" /> Mentors set their own rates
								</li>
								<li class="flex items-center gap-4 text-xl font-medium">
									<CheckCircle class="text-secondary" /> Pay only for the time you need
								</li>
								<li class="flex items-center gap-4 text-xl font-medium">
									<CheckCircle class="text-secondary" /> 0% platform fee for students
								</li>
							</ul>
						</div>
						<div class="rounded-4xl border border-base-300 bg-base-100 p-10 backdrop-blur-md">
							<div class="text-center">
								<div class="mb-2 text-sm font-bold tracking-widest uppercase opacity-60">
									Average Session
								</div>
								<div class="mb-4 font-display text-7xl font-black">
									$35<span class="text-2xl opacity-40">/hr</span>
								</div>
								<p class="mb-10 text-lg opacity-70">
									Includes full access to video, screen sharing, and messaging.
								</p>
								<a href="/app/auth/sign-up" class="btn btn-block btn-lg btn-primary"
									>Find your mentor</a
								>
							</div>
						</div>
					</div>
				</div>
			</div>
		</section>

		<!-- CTA Banner -->
		<section class="container mx-auto px-6 pb-24">
			<div
				class="flex flex-col items-center rounded-[3rem] bg-primary px-8 py-20 text-center text-primary-content lg:py-32"
			>
				<h2 class="max-w-4xl font-display text-5xl font-black tracking-tighter lg:text-7xl">
					Don't just dream it. <br /> Get mentored for it.
				</h2>
				<p class="mt-8 max-w-2xl text-xl opacity-90">
					Join the 2,000+ students who accelerated their career with the help of those who've been
					there.
				</p>
				<div class="mt-12 flex flex-col gap-4 sm:flex-row">
					<a href="/app/auth/sign-up" class="btn px-12 btn-lg">Sign up free</a>
					<a href="/app/auth/sign-in" class="btn btn-ghost btn-lg">Log in</a>
				</div>
			</div>
		</section>
	</main>

	<!-- Minimal Footer -->
	<footer class="border-t border-base-200 py-12">
		<div
			class="container mx-auto flex flex-col items-center justify-between gap-8 px-6 md:flex-row"
		>
			<div class="flex items-center gap-2">
				<Logo />
			</div>

			<div class="flex gap-8 text-sm font-medium opacity-50">
				<a href="/legal/privacy-policy" class="transition-opacity hover:opacity-100">Privacy</a>
				<a href="/legal/terms-and-conditions" class="transition-opacity hover:opacity-100">Terms</a>
				<a href="/" class="hidden transition-opacity hover:opacity-100">Twitter</a>
			</div>

			<div class="text-sm opacity-30">© 2026 Univerlink</div>
		</div>
	</footer>
</div>

<style>
	:global(html) {
		scroll-behavior: smooth;
	}

	/* Linktree-style big text utilities */
	.tracking-tighter {
		letter-spacing: -0.05em;
	}
</style>
