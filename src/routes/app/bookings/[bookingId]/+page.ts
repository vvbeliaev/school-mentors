import { pb, Collections } from '$lib';
import { error } from '@sveltejs/kit';

export async function load({ params }) {
	try {
		const booking = await pb.collection(Collections.Bookings).getOne(params.bookingId, {
			expand: 'slot,slot.mentor,mentee'
		});

		return {
			booking
		};
	} catch (err) {
		throw error(404, 'Booking not found');
	}
}
