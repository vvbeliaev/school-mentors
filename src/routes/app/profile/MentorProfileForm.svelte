<script lang="ts">
	import { GraduationCap, DollarSign, Tag, Info, Pencil } from 'lucide-svelte';

	import { userStore, userApi, TagInput } from '$lib/apps/user';
	import { Modal, Button, Input, Textarea } from '$lib';

	const user = $derived(userStore.user);

	let isOpen = $state(false);
	let isSubmitting = $state(false);

	let bio = $state('');
	let university = $state('');
	let degree = $state('');
	let graduationYear = $state('');
	let hourRate = $state('');
	let tags = $state<string[]>([]);

	$effect(() => {
		if (isOpen && user) {
			bio = user.bio || '';
			university = user.university || '';
			degree = user.degree || '';
			graduationYear = user.graduationYear?.toString() || '';
			hourRate = (user.hourRateCents ? user.hourRateCents / 100 : 0).toString();
			tags = user.tags ? [...user.tags] : [];
		}
	});

	async function handleSubmit() {
		isSubmitting = true;
		try {
			await userApi.submitMentorApplication({
				bio,
				university,
				degree,
				graduationYear: parseInt(graduationYear) || undefined,
				hourRateCents: Math.round(parseFloat(hourRate) * 100) || 0,
				tags
			});
			isOpen = false;
		} catch (err) {
			console.error('Failed to submit mentor profile:', err);
		} finally {
			isSubmitting = false;
		}
	}
</script>

{#if !user?.isMentor}
	<div class="card bg-linear-to-br from-primary/10 to-primary/5 shadow-sm">
		<div class="card-body">
			<div class="flex flex-col items-center gap-4 text-center">
				<div class="rounded-full bg-primary/20 p-4 text-primary">
					<GraduationCap size={40} />
				</div>
				<div class="space-y-2">
					<h3 class="text-xl font-bold">Become a Mentor</h3>
					<p class="mx-auto max-w-sm text-sm text-base-content/70">
						Share your knowledge, help other students, and earn money. Fill out your mentor profile
						and start helping others.
					</p>
				</div>
				<Button class="mt-2" onclick={() => (isOpen = true)}>Set up Mentor Profile</Button>
			</div>
		</div>
	</div>
{:else}
	<div class="card bg-base-100 shadow-sm">
		<div class="card-body">
			<div class="flex items-start justify-between">
				<div class="space-y-1">
					<h3 class="text-lg font-bold">Mentor Profile</h3>
					<div class="flex items-center gap-2">
						{#if user?.isVerified}
							<div class="badge badge-sm badge-success">Verified</div>
						{:else}
							<div class="badge badge-sm badge-warning">Pending Verification</div>
						{/if}
					</div>
				</div>
				<Button variant="ghost" size="sm" onclick={() => (isOpen = true)}>Edit Profile</Button>
			</div>

			<div class="mt-4 grid gap-4 sm:grid-cols-2">
				<div class="flex items-center gap-3 text-sm">
					<GraduationCap class="text-base-content/50" size={18} />
					<div class="min-w-0 flex-1">
						<p class="text-[10px] font-medium tracking-wider text-base-content/60 uppercase">
							Education
						</p>
						<p class="truncate">{user?.university || 'Not set'}</p>
						{#if user?.degree}
							<p class="truncate text-xs text-base-content/50">{user.degree}</p>
						{/if}
					</div>
				</div>

				<div class="flex items-center gap-3 text-sm">
					<DollarSign class="text-base-content/50" size={18} />
					<div>
						<p class="text-[10px] font-medium tracking-wider text-base-content/60 uppercase">
							Hourly Rate
						</p>
						<p>${(user?.hourRateCents || 0) / 100}/hr</p>
					</div>
				</div>
			</div>

			{#if user?.tags && user.tags.length > 0}
				<div class="mt-4 flex flex-wrap gap-2">
					{#each user.tags as tag}
						<div class="badge badge-outline badge-sm">{tag}</div>
					{/each}
				</div>
			{/if}
		</div>
	</div>
{/if}

<Modal bind:open={isOpen} class="max-w-2xl">
	<div class="space-y-6 p-2">
		<div class="space-y-1">
			<h2 class="text-2xl font-bold">Mentor Profile Setup</h2>
			<p class="text-sm text-base-content/70">
				Fill out your mentor card. This information will be visible to students.
			</p>
		</div>

		<div class="grid gap-6">
			<!-- Personal Section -->
			<div class="space-y-4">
				<h3
					class="flex items-center gap-2 text-sm font-semibold tracking-wider text-base-content/50 uppercase"
				>
					<Info size={14} /> Academic Details
				</h3>
				<div class="grid gap-4 sm:grid-cols-2">
					<Input
						label="University"
						bind:value={university}
						placeholder="e.g. Stanford University"
						required
					/>
					<Input
						label="Degree / Major"
						bind:value={degree}
						placeholder="e.g. Computer Science"
						required
					/>
					<Input
						label="Graduation Year"
						type="number"
						bind:value={graduationYear}
						placeholder="e.g. 2025"
						required
					/>
					<Input
						label="Hourly Rate ($)"
						type="number"
						step="0.01"
						bind:value={hourRate}
						placeholder="e.g. 25.00"
						required
					/>
				</div>
			</div>

			<!-- Bio Section -->
			<div class="space-y-4">
				<h3
					class="flex items-center gap-2 text-sm font-semibold tracking-wider text-base-content/50 uppercase"
				>
					<Pencil size={14} /> About You
				</h3>
				<Textarea
					label="Bio"
					bind:value={bio}
					placeholder="Tell students about your experience, expertise, and how you can help them..."
					required
				/>
			</div>

			<!-- Tags Section -->
			<div class="space-y-4">
				<h3
					class="flex items-center gap-2 text-sm font-semibold tracking-wider text-base-content/50 uppercase"
				>
					<Tag size={14} /> Skills & Topics
				</h3>
				<TagInput label="" bind:selectedTags={tags} onchange={(newTags) => (tags = newTags)} />
			</div>
		</div>

		<div class="flex justify-end gap-3 border-t border-base-200 pt-4">
			<Button variant="ghost" onclick={() => (isOpen = false)} disabled={isSubmitting}>
				Cancel
			</Button>
			<Button onclick={handleSubmit} disabled={isSubmitting}>
				{#if isSubmitting}
					<span class="loading loading-sm loading-spinner"></span>
					Saving...
				{:else}
					{user?.isMentor ? 'Save Changes' : 'Submit for Verification'}
				{/if}
			</Button>
		</div>
	</div>
</Modal>
