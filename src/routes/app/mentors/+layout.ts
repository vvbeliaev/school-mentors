import { mentorsStore } from '$lib/apps/user';
import type { MentorFilters } from '$lib/apps/user/models';

export async function load({ url, depends }) {
	depends('app:mentors');

	const filters: MentorFilters = {
		q: url.searchParams.get('q') || '',
		university: url.searchParams.get('university') || '',
		tags: url.searchParams.get('tags')?.split(',').filter(Boolean) || [],
		minPrice: url.searchParams.get('minPrice')
			? parseInt(url.searchParams.get('minPrice')!)
			: undefined,
		maxPrice: url.searchParams.get('maxPrice')
			? parseInt(url.searchParams.get('maxPrice')!)
			: undefined,
		sort: url.searchParams.get('sort') || '-created'
	};

	const mentorsPromise = mentorsStore.load(filters);
	return { mentorsPromise, filters };
}
