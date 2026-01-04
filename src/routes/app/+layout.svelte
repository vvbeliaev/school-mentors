<script lang="ts">
	import posthog from 'posthog-js';
	import { Heart, Plus, Settings, Menu, PanelRight, Search, Home } from 'lucide-svelte';
	import { afterNavigate } from '$app/navigation';
	import { page } from '$app/state';

	import { userStore } from '$lib/apps/user';
	import { Button, ThemeController, uiStore, Sidebar, swipeable } from '$lib';

	import Splash from './Splash.svelte';

	let { children, data } = $props();
	const globalPromise = $derived(data.globalPromise);

	const user = $derived(userStore.user);

	// Posthog identify and set person
	$effect(() => {
		console.log(user);

		if (!user) return;
		posthog.identify(user.id, {
			email: user.email,
			name: user.name
		});
		posthog.capture('user_authenticated', {
			email: user.email,
			name: user.name
		});
	});

	// Global user load
	$effect(() => {
		globalPromise.then(({ user }) => {
			if (user) userStore.user = user;
		});
	});

	// Real-time subscription
	$effect(() => {
		const userId = userStore.user?.id;
		if (!userId) return;
		userStore.subscribe(userId);
		return () => {
			userStore.unsubscribe();
		};
	});

	afterNavigate(() => {
		uiStore.setSidebarOpen(false);
	});
</script>

{#snippet sidebarHeader({ expanded }: { expanded: boolean })}
	{#if expanded}
		<a href="/app" class="flex items-center gap-2">
			<Heart class="size-6 text-primary" />
			<span class="font-semibold text-nowrap">Mentorship</span>
		</a>
	{/if}
{/snippet}

{#snippet sidebarContent({ expanded }: { expanded: boolean })}
	<div class="shrink-0 px-2 pt-4">
		<!-- <Button block class="rounded-xl" disabled={loading} square={!expanded} onclick={handleNewChat}>
			<Plus class="size-5" />
			{#if expanded}
				<span class="text-nowrap">New Chat</span>
			{/if}
		</Button> -->
	</div>
{/snippet}

{#snippet sidebarFooter({ expanded }: { expanded: boolean })}
	<div class="divider my-1"></div>

	{#if user}
		<div class="mb-2 flex justify-center px-2">
			<!-- <button
				class={['btn justify-start btn-ghost', expanded ? 'btn-block' : 'btn-square']}
				onclick={() => uiStore.toggleFeedbackModal()}
			>
				<MessageSquare class={expanded ? 'size-5' : 'size-6'} />
				{#if expanded}
					Feedback
				{:else}
					<span class="sr-only">Feedback</span>
				{/if}
			</button> -->
		</div>
	{/if}

	<div class={['mb-2 border-base-300', expanded ? 'px-2' : 'flex justify-center']}>
		<ThemeController {expanded} navStyle />
	</div>

	<div class="border-t border-base-300">
		{#if user && user.email}
			<a
				href="/app/profile"
				class={[
					'flex items-center gap-3 p-2 transition-colors hover:bg-base-200',
					!expanded && 'justify-center'
				]}
				title={!expanded ? 'Settings' : ''}
			>
				{#if userStore.avatarUrl}
					<img src={userStore.avatarUrl} alt={user.name} class="size-10 rounded-full" />
				{:else}
					<div class="flex size-10 items-center justify-center rounded-full bg-base-300">
						{user.name?.at(0)?.toUpperCase() ?? 'U'}
					</div>
				{/if}
				{#if expanded}
					<div class="flex-1 overflow-hidden">
						<div class="truncate text-sm font-semibold">{user.name || '<No Name>'}</div>
						<div class="truncate text-xs opacity-60">{user.email}</div>
					</div>
					<Settings class="size-5 opacity-60" />
				{/if}
			</a>
		{:else}
			<a
				href="/app/auth/sign-in"
				class={[
					'flex items-center gap-3 rounded-lg p-2 transition-colors hover:bg-base-300',
					!expanded && 'justify-center'
				]}
				title={!expanded ? 'Log in' : ''}
			>
				<div class="size-10 rounded-full bg-base-300"></div>
				{#if expanded}
					<div class="flex-1 overflow-hidden">
						<div class="truncate text-sm font-semibold">Log in</div>
					</div>
				{/if}
			</a>
		{/if}
	</div>
{/snippet}

{#await globalPromise}
	<Splash />
{:then}
	<div
		class="flex h-screen flex-col overflow-hidden bg-base-100 md:flex-row"
		use:swipeable={{
			isOpen: uiStore.sidebarOpen ?? false,
			direction: 'right',
			onOpen: () => uiStore.setSidebarOpen(true),
			onClose: () => uiStore.setSidebarOpen(false)
		}}
	>
		<!-- Sidebar -->
		<Sidebar
			open={uiStore.sidebarOpen ?? false}
			expanded={uiStore.sidebarExpanded ?? true}
			position="left"
			header={sidebarHeader}
			footer={sidebarFooter}
			onclose={() => uiStore.setSidebarOpen(false)}
			ontoggle={() => uiStore.toggleSidebarExpanded()}
		>
			{#snippet children({ expanded })}
				{@render sidebarContent({ expanded })}
			{/snippet}
		</Sidebar>

		<!-- Main Content -->
		<main class="flex-1 overflow-hidden">
			<div class="h-full max-w-[1440px]">
				{@render children()}
			</div>
		</main>

		<!-- Mobile Dock -->
		<div class="dock dock-sm border-t border-base-300 md:hidden">
			<a href="/app" class:dock-active={page.url.pathname === '/app'}>
				<Home class="size-5" />
				<span class="dock-label">Home</span>
			</a>
			<a href="/app" class:dock-active={false}>
				<Search class="size-5" />
				<span class="dock-label">Explore</span>
			</a>
			<a href="/app/profile" class:dock-active={page.url.pathname === '/app/profile'}>
				<Settings class="size-5" />
				<span class="dock-label">Profile</span>
			</a>
			<button class="hidden" onclick={() => uiStore.toggleRightSidebar()}>
				<PanelRight class="size-5" />
				<span class="dock-label">Panel</span>
			</button>
			<button class="hidden" onclick={() => uiStore.setSidebarOpen(true)}>
				<Menu class="size-5" />
				<span class="dock-label">Menu</span>
			</button>
		</div>
	</div>
{/await}
