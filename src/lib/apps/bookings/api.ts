import { pb, Collections } from '$lib';

class BookingsAPI {
	async createBooking(slotId: string) {
		const userId = pb.authStore.record?.id;
		if (!userId) throw new Error('User not authenticated');

		return await pb.collection(Collections.Bookings).create({
			mentee: userId,
			slot: slotId
		});
	}

	async getBooking(bookingId: string) {
		return await pb.collection(Collections.Bookings).getOne(bookingId, {
			expand: 'slot,slot.mentor'
		});
	}
}

export const bookingsApi = new BookingsAPI();
