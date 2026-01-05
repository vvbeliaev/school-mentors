import { Collections, pb, type UsersResponse } from '$lib';
import type { MentorFilters, UserTags } from './models';

const PAGE_SIZE = 20;

export type MentorResponse = UsersResponse<UserTags> & {
	rating?: number; // Mock rating for now as it's not in the schema yet
};

class MentorsStore {
	page = $state(1);
	totalPages = $state(0);
	totalItems = $state(0);
	loading = $state(true);
	isFallback = $state(false);

	private _mentors: MentorResponse[] = $state([]);
	private userId: string | null = null;
	private currentFilters: MentorFilters = {};

	mentors = $derived(this._mentors);

	set(mentors: MentorResponse[], page: number, totalPages: number, totalItems: number) {
		this.loading = false;
		this._mentors = mentors.map((m) => ({
			...m,
			rating: 4.5
		}));
		this.page = page;
		this.totalPages = totalPages;
		this.totalItems = totalItems;
	}

	async load(filters: MentorFilters = {}) {
		this.loading = true;
		this.currentFilters = filters;
		this.isFallback = false;

		const filterQuery = this.buildFilterQuery(filters);
		const sort = filters.sort || '-created';

		try {
			let res = await pb.collection(Collections.Users).getList<MentorResponse>(1, PAGE_SIZE, {
				filter: filterQuery,
				sort: sort,
				requestKey: JSON.stringify(filters)
			});

			// Fallback logic: if no results and filters were applied, try a broader search
			const hasFilters =
				filters.q ||
				filters.university ||
				(filters.tags && filters.tags.length > 0) ||
				filters.minPrice ||
				filters.maxPrice;

			if (res.items.length === 0 && hasFilters) {
				this.isFallback = true;
				res = await pb.collection(Collections.Users).getList<MentorResponse>(1, PAGE_SIZE, {
					filter: 'isMentor = true && isVerified = true',
					sort: '-created'
				});
			}

			this.set(res.items, res.page, res.totalPages, res.totalItems);
			return res;
		} catch (error) {
			console.error('Failed to load mentors:', error);
			this.loading = false;
			throw error;
		}
	}

	private buildFilterQuery(filters: MentorFilters): string {
		const parts = ['isMentor = true', 'isVerified = true'];

		if (filters.q) {
			parts.push(`(name ~ "${filters.q}" || bio ~ "${filters.q}" || university ~ "${filters.q}")`);
		}

		if (filters.university) {
			parts.push(`university = "${filters.university}"`);
		}

		if (filters.tags && filters.tags.length > 0) {
			// PocketBase JSON field filter for tags
			const tagFilters = filters.tags.map((tag) => `tags ~ "${tag}"`).join(' || ');
			parts.push(`(${tagFilters})`);
		}

		if (filters.minPrice) {
			parts.push(`hourRateCents >= ${filters.minPrice}`);
		}

		if (filters.maxPrice) {
			parts.push(`hourRateCents <= ${filters.maxPrice}`);
		}

		return parts.join(' && ');
	}

	async loadNextPage() {
		if (this.page >= this.totalPages || this.isFallback) return;

		this.loading = true;
		const filterQuery = this.buildFilterQuery(this.currentFilters);
		const sort = this.currentFilters.sort || '-created';

		const res = await pb
			.collection(Collections.Users)
			.getList<MentorResponse>(this.page + 1, PAGE_SIZE, {
				filter: filterQuery,
				sort: sort
			});

		this._mentors = [...this._mentors, ...res.items];
		this.page = res.page;
		this.totalPages = res.totalPages;
		this.totalItems = res.totalItems;
		this.loading = false;
	}

	clear() {
		this._mentors = [];
		this.page = 1;
		this.totalPages = 0;
		this.totalItems = 0;
		this.loading = true;
		this.isFallback = false;
		this.currentFilters = {};
	}
}

export const mentorsStore = new MentorsStore();
