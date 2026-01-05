<script lang="ts">
	import { X } from 'lucide-svelte';
	import { goto } from '$app/navigation';

	import { pb, Button, Modal } from '$lib';
	import { userStore } from '$lib/apps/user';
	import { Schedule } from '$lib/apps/slots';

	import ProfileCard from './ProfileCard.svelte';
	import MentorProfileForm from './MentorProfileForm.svelte';
	import LegalCard from './LegalCard.svelte';

	const user = $derived(userStore.user);

	let isScheduleOpen = $state(false);

	function logout() {
		pb!.authStore.clear();
		goto('/app/auth/sign-in');
	}
</script>

<div class="mx-auto w-full max-w-7xl space-y-4 overflow-y-auto p-4 sm:space-y-0 sm:p-6 lg:p-4">
	<div class="mb-4 flex items-center justify-between sm:mb-6">
		<h1 class="mx-auto text-2xl font-bold sm:text-3xl">Profile & Settings</h1>
	</div>

	<div class={['mx-auto space-y-4', userStore.user?.isMentor ? 'max-w-4xl' : 'max-w-xl']}>
		<ProfileCard />

		<MentorProfileForm />

		{#if user?.isMentor}
			<div class="card bg-linear-to-br from-primary/10 to-primary/5 shadow-sm">
				<div class="card-body">
					<div class="flex items-center justify-between gap-4">
						<div class="space-y-1">
							<h3 class="text-lg font-bold">Availability Schedule</h3>
							<p class="text-xs text-base-content/60">
								Manage your working hours and available slots for students.
							</p>
						</div>
						<Button onclick={() => (isScheduleOpen = true)}>Manage Schedule</Button>
					</div>
				</div>
			</div>

			<Modal bind:open={isScheduleOpen} class="max-w-5xl">
				<div class="p-2">
					<Schedule />
					<div class="mt-6 flex justify-end">
						<Button onclick={() => (isScheduleOpen = false)}>Close</Button>
					</div>
				</div>
			</Modal>
		{/if}

		<LegalCard />

		<Button class="mt-2 w-full" onclick={logout} variant="soft" color="error">
			<X class="mr-2 h-4 w-4" />
			Logout
		</Button>
	</div>
</div>
