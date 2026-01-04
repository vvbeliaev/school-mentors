import { pb, type UsersResponse, type UserExpand, Collections } from '$lib';

import type { UserTags } from './models';

export async function globalUserLoad() {
	console.log('globalUserLoad', pb.authStore.isValid);
	let user: UsersResponse<UserTags, UserExpand> | null = null;

	try {
		const res = await pb.collection(Collections.Users).authRefresh({ expand: 'slots_via_mentor' });
		user = res.record as UsersResponse<UserTags, UserExpand>;

		return { user };
	} catch (error) {
		console.error(error);
		pb.authStore.clear();
		return { user: null };
	}
}
