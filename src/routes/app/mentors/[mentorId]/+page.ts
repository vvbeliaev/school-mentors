import { pb, Collections } from '$lib';
import type { UsersResponse, SlotsResponse } from '$lib/shared/pb/pocketbase-types';
import { error } from '@sveltejs/kit';

export async function load({ params }) {
	const mentorId = params.mentorId;

	try {
		const mentor = await pb.collection(Collections.Users).getOne<UsersResponse>(mentorId);

		// Only show the page if the user is actually a mentor
		if (!mentor.isMentor) {
			throw error(404, 'Mentor not found');
		}

		// Load available slots (not booked and in the future)
		const slots = await pb.collection(Collections.Slots).getFullList<SlotsResponse>({
			filter: `mentor = "${mentorId}" && booked = false && start >= "${new Date().toISOString()}"`,
			sort: 'start'
		});

		return {
			mentor,
			slots
		};
	} catch (err) {
		console.error('Failed to load mentor profile:', err);
		throw error(404, 'Mentor not found');
	}
}
