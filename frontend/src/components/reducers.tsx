import produce from 'immer';
import { Action } from './actions';
import { Listing } from './list/Listings';

export interface SystemState {
  isLoading: boolean;
  listings: Listing[];
}

const initialState: SystemState = {
  isLoading: true,
  listings: [],
};

export const reducer = produce((draft, action) => {
  switch (action.type) {
    case Action.LoadListings:
      draft.isLoading = false;
      draft.listings = action.listings;
      break;

    case Action.HideListing:
      const updatedListings = [];
      for (const l of draft.listings) {
        if (l.id !== action.listingID) {
          updatedListings.push(l);
        }
      }
      draft.listings = updatedListings;

    case Action.UpdateNote:
      for (const l of draft.listings) {
        if (l.id === action.listingID) {
          l.note = action.note;
          break;
        }
      }
      break;
  }
}, initialState);
