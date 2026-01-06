import { pb, Collections, type SlotsResponse, type SlotExpand } from '$lib';

type EnhancedSlot = SlotsResponse<SlotExpand>;

class SlotsStore {
	private userId: string | null = null;

	private slots: EnhancedSlot[] = $state([]);

	mentorSlots = $derived(this.slots.filter((s) => s.mentor === this.userId));
	menteeSlots = $derived(this.slots.filter((s) => s.mentor !== this.userId));

	set(slots: EnhancedSlot[]) {
		this.slots = slots;
	}

	async loadMentorSlots(userId: string) {
		const slots = await pb.collection(Collections.Slots).getFullList<EnhancedSlot>({
			filter: `mentor = "${userId}" && booked = false && start >= "${new Date().toISOString()}"`,
			sort: 'start'
		});
		return slots;
	}

	async load(userId: string) {
		const mentorSlots = await pb
			.collection(Collections.Slots)
			.getFullList<EnhancedSlot>({ filter: `mentor = "${userId}"`, expand: 'bookings_via_slot' });

		const menteeSlots = await pb.collection(Collections.Slots).getFullList<EnhancedSlot>({
			filter: `bookings_via_slot.mentee = "${userId}"`,
			expand: 'bookings_via_slot'
		});

		this.userId = userId;
		console.log(mentorSlots, menteeSlots);
		return [...mentorSlots, ...menteeSlots];
	}

	subscribe() {
		return pb.collection(Collections.Slots).subscribe(
			'*',
			(e) => {
				const slot = e.record as EnhancedSlot;
				console.log(slot);
				switch (e.action) {
					case 'create':
						this.slots.unshift(slot);
						break;
					case 'update':
						this.slots = this.slots?.map((s) => (s.id === slot.id ? slot : s)) || [];
						break;
					case 'delete':
						this.slots = this.slots?.filter((s) => s.id !== slot.id) || [];
						break;
				}
			},
			{
				filter: `mentor = "${this.userId}"`,
				expand: 'bookings_via_slot'
			}
		);
	}

	unsubscribe() {
		pb.collection(Collections.Slots).unsubscribe();
	}

	clear() {
		this.userId = null;
		this.slots = [];
	}
}

export const slotsStore = new SlotsStore();
