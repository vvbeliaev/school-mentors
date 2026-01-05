<script lang="ts">
	import { Calendar, ChevronLeft, ChevronRight, Clock, Trash2, CheckCircle2 } from 'lucide-svelte';
	import { slotsStore, slotsApi } from '$lib/apps/slots';
	import { userStore } from '$lib/apps/user';
	import { Button } from '$lib';

	// 0, 0.5, 1, 1.5, ..., 23.5
	const TIME_SLOTS = Array.from({ length: 48 }, (_, i) => i / 2);
	const DAYS = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'];

	let currentWeekStart = $state(getStartOfWeek(new Date()));

	function getStartOfWeek(date: Date) {
		const d = new Date(date);
		d.setHours(0, 0, 0, 0);
		const day = d.getDay();
		const diff = d.getDate() - day + (day === 0 ? -6 : 1);
		return new Date(d.setDate(diff));
	}

	const weekDays = $derived(
		Array.from({ length: 7 }, (_, i) => {
			const d = new Date(currentWeekStart);
			d.setDate(d.getDate() + i);
			return d;
		})
	);

	const mentorSlots = $derived(slotsStore.mentorSlots);

	function isSlotAt(date: Date, time: number) {
		const hour = Math.floor(time);
		const minute = (time % 1) * 60;
		return mentorSlots.find((s) => {
			const sDate = new Date(s.start);
			return (
				sDate.getFullYear() === date.getFullYear() &&
				sDate.getMonth() === date.getMonth() &&
				sDate.getDate() === date.getDate() &&
				sDate.getHours() === hour &&
				sDate.getMinutes() === minute
			);
		});
	}

	async function toggleSlot(date: Date, time: number) {
		const existing = isSlotAt(date, time);
		if (existing) {
			if (existing.booked) {
				alert('Cannot delete a booked slot');
				return;
			}
			try {
				await slotsApi.deleteSlot(existing.id);
			} catch (err) {
				console.error('Failed to delete slot:', err);
			}
		} else {
			const start = new Date(date);
			const hour = Math.floor(time);
			const minute = (time % 1) * 60;
			start.setHours(hour, minute, 0, 0);
			try {
				await slotsApi.createSlot({
					start: start.toISOString(),
					durationMinutes: 30
				});
			} catch (err) {
				console.error('Failed to create slot:', err);
			}
		}
	}

	function nextWeek() {
		const next = new Date(currentWeekStart);
		next.setDate(next.getDate() + 7);
		currentWeekStart = next;
	}

	function prevWeek() {
		const prev = new Date(currentWeekStart);
		prev.setDate(prev.getDate() - 7);
		currentWeekStart = prev;
	}

	const formatHour = (t: number) => {
		const h = Math.floor(t);
		const m = (t % 1) * 60;
		return `${h.toString().padStart(2, '0')}:${m.toString().padStart(2, '0')}`;
	};
</script>

<div class="flex flex-col gap-6">
	<div class="flex items-center justify-between">
		<div class="flex items-center gap-4">
			<h2 class="flex items-center gap-2 text-xl font-bold">
				<Calendar size={24} class="text-primary" />
				Weekly Schedule
			</h2>
			<div class="flex items-center gap-1">
				<Button variant="ghost" size="sm" onclick={prevWeek}>
					<ChevronLeft size={18} />
				</Button>
				<span class="min-w-[120px] text-center text-sm font-medium">
					{weekDays[0].toLocaleDateString('en-US', { month: 'short', day: 'numeric' })} -
					{weekDays[6].toLocaleDateString('en-US', { month: 'short', day: 'numeric' })}
				</span>
				<Button variant="ghost" size="sm" onclick={nextWeek}>
					<ChevronRight size={18} />
				</Button>
			</div>
		</div>
		<div class="flex gap-4 text-xs">
			<div class="flex items-center gap-1.5">
				<div class="h-3 w-3 rounded-sm border border-base-content/10 bg-base-200"></div>
				<span>Unavailable</span>
			</div>
			<div class="flex items-center gap-1.5">
				<div class="h-3 w-3 rounded-sm border border-primary/30 bg-primary/20"></div>
				<span>Available</span>
			</div>
			<div class="flex items-center gap-1.5">
				<div class="h-3 w-3 rounded-sm border border-success/30 bg-success/20"></div>
				<span>Booked</span>
			</div>
		</div>
	</div>

	<div class="overflow-x-auto rounded-xl border border-base-200 bg-base-100 shadow-sm">
		<div class="min-w-[800px]">
			<!-- Header -->
			<div class="grid grid-cols-[80px_repeat(7,1fr)] border-b border-base-200 bg-base-200/50">
				<div
					class="flex items-center justify-center p-3 text-xs font-bold text-base-content/40 uppercase"
				>
					<Clock size={14} />
				</div>
				{#each weekDays as day, i}
					<div class="border-l border-base-200 p-3 text-center">
						<div class="text-[10px] font-bold tracking-wider text-base-content/50 uppercase">
							{DAYS[i]}
						</div>
						<div class="text-lg font-bold">
							{day.getDate()}
						</div>
					</div>
				{/each}
			</div>

			<!-- Grid -->
			<div class="max-h-[500px] overflow-y-auto">
				{#each TIME_SLOTS as time}
					<div class="grid grid-cols-[80px_repeat(7,1fr)] border-b border-base-200 last:border-0">
						<div
							class="flex items-center justify-center bg-base-200/20 p-2 text-center text-[10px] font-medium text-base-content/40"
						>
							{formatHour(time)}
						</div>
						{#each weekDays as day}
							{@const slot = isSlotAt(day, time)}
							<div class="min-h-[40px] border-l border-base-200 p-0.5">
								<button
									type="button"
									onclick={() => toggleSlot(day, time)}
									class={[
										'cursor-pointer',
										'group flex h-full w-full flex-col items-center justify-center gap-1 rounded-md text-[9px] font-bold transition-all',
										!slot && 'text-transparent hover:bg-primary/10 hover:text-primary/60',
										slot &&
											!slot.booked &&
											'border border-primary/20 bg-primary/10 text-primary hover:border-error/20 hover:bg-error/10 hover:text-error',
										slot &&
											slot.booked &&
											'cursor-not-allowed border border-success/20 bg-success/10 text-success'
									]}
								>
									{#if !slot}
										<span>+ ADD</span>
									{:else if slot.booked}
										<CheckCircle2 size={12} />
										<span>BOOKED</span>
									{:else}
										<div class="flex flex-col items-center gap-0.5">
											<Trash2 size={12} class="group-hover:hidden" />
											<span class="uppercase group-hover:hidden">Available</span>
											<Trash2 size={12} class="hidden group-hover:block" />
											<span class="hidden uppercase group-hover:block">Remove</span>
										</div>
									{/if}
								</button>
							</div>
						{/each}
					</div>
				{/each}
			</div>
		</div>
	</div>
</div>
