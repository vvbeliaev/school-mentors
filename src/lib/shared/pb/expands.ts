import type { SlotsResponse } from './pocketbase-types';

export type UserExpand =
	| {
			slots_via_mentor: SlotsResponse[] | undefined;
	  }
	| undefined;
