import type { BookingsResponse, SlotsResponse, UsersResponse } from './pocketbase-types';

export type UserExpand =
	| {
			slots_via_mentor: SlotsResponse<SlotExpand>[] | undefined;
	  }
	| undefined;

export type SlotExpand =
	| {
			mentor: UsersResponse;
			bookings_via_slot: BookingsResponse<unknown, BookingsExpand>[] | undefined;
	  }
	| undefined;

export type BookingsExpand =
	| {
			slot: SlotsResponse<SlotExpand>;
			mentee: UsersResponse;
	  }
	| undefined;
