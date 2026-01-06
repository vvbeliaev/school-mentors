import { error } from '@sveltejs/kit';

import { mentorsStore } from '$lib/apps/user';
import { slotsStore } from '$lib/apps/slots';

export async function load({ params }) {
	const mentorId = params.mentorId;

	try {
		const mentor = await mentorsStore.getOne(mentorId);

		if (!mentor.isMentor) {
			throw error(404, 'Mentor not found');
		}

		const slots = await slotsStore.loadMentorSlots(mentorId);

		return {
			mentor,
			slots
		};
	} catch (err) {
		console.error('Failed to load mentor profile:', err);
		throw error(404, 'Mentor not found');
	}
}
