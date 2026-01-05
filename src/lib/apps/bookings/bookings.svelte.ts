import { pb, Collections, type BookingsResponse } from '$lib';

class BookingsStore {
	private userId: string | null = null;

	bookings: BookingsResponse[] = $state([]);

	mentorBookings = $derived(this.bookings.filter((b) => b.expand?.slot?.mentor === this.userId));
	menteeBookings = $derived(this.bookings.filter((b) => b.mentee === this.userId));

	page: number = $state(1);
	perPage: number = $state(20);
	totalItems: number = $state(0);

	set(bookings: BookingsResponse[], totalItems: number) {
		this.bookings = bookings;
		this.totalItems = totalItems;
	}

	async load(userId: string, page = 1) {
		this.userId = userId;
		this.page = page;

		const result = await pb.collection(Collections.Bookings).getList(page, this.perPage, {
			filter: `mentee = "${userId}" || slot.mentor = "${userId}"`,
			sort: '-created',
			expand: 'slot,slot.mentor,mentee'
		});

		this.bookings = result.items as BookingsResponse[];
		this.totalItems = result.totalItems;

		return result;
	}

	subscribe() {
		if (!this.userId) return;

		return pb.collection(Collections.Bookings).subscribe(
			'*',
			async (e) => {
				const booking = e.record as BookingsResponse;

				// Re-fetch with expand if it's a create/update
				let expandedBooking = booking;
				if (e.action === 'create' || e.action === 'update') {
					expandedBooking = await pb.collection(Collections.Bookings).getOne(booking.id, {
						expand: 'slot,slot.mentor,mentee'
					});
				}

				switch (e.action) {
					case 'create':
						// Only add if it belongs to user
						if (
							expandedBooking.mentee === this.userId ||
							(expandedBooking.expand?.slot as any)?.mentor === this.userId
						) {
							this.bookings = [expandedBooking, ...this.bookings];
							this.totalItems += 1;
						}
						break;
					case 'update':
						this.bookings = this.bookings.map((b) =>
							b.id === expandedBooking.id ? expandedBooking : b
						);
						break;
					case 'delete':
						this.bookings = this.bookings.filter((b) => b.id !== booking.id);
						this.totalItems -= 1;
						break;
				}
			},
			{
				expand: 'slot,slot.mentor,mentee'
			}
		);
	}

	unsubscribe() {
		pb.collection(Collections.Bookings).unsubscribe();
	}

	clear() {
		this.userId = null;
		this.bookings = [];
		this.page = 1;
		this.totalItems = 0;
	}
}

export const bookingsStore = new BookingsStore();
