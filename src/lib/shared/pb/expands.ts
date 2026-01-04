import type { BookingsResponse, SlotsResponse } from './pocketbase-types';

export type UserExpand =
	| {
			slots_via_mentor: SlotsResponse[] | undefined;
	  }
	| undefined;

export type SlotExpand =
	| {
			bookings_via_slot: BookingsResponse[] | undefined;
	  }
	| undefined;
