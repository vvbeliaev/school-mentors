import { pb, Collections, type SlotsRecord } from '$lib';

class SlotsAPI {
	async createSlot(data: Partial<SlotsRecord>) {
		const userId = pb.authStore.record?.id;
		if (!userId) throw new Error('User not authenticated');

		return await pb.collection(Collections.Slots).create({
			...data,
			mentor: userId,
			booked: false
		});
	}

	async deleteSlot(slotId: string) {
		const userId = pb.authStore.record?.id;
		if (!userId) throw new Error('User not authenticated');

		// PocketBase will handle permission check via API Rules
		return await pb.collection(Collections.Slots).delete(slotId);
	}
}

export const slotsApi = new SlotsAPI();
