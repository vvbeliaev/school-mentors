import { slotsStore } from '$lib/apps/slots';
import { bookingsStore } from '$lib/apps/bookings';
import { pb, type UsersResponse, type UserExpand } from '$lib';

import type { UserTags } from './models';
import { userStore } from './user.svelte';

export async function globalUserLoad() {
	console.log('globalUserLoad', pb.authStore.isValid);
	let user: UsersResponse<UserTags, UserExpand> | null = null;

	try {
		user = await userStore.load();
		const [slots, bookings] = await Promise.all([
			slotsStore.load(user.id),
			bookingsStore.load(user.id)
		]);

		return { user, slots, bookings };
	} catch (error) {
		console.error(error);
		pb.authStore.clear();
		return { user: null, slots: [], bookings: null };
	}
}
