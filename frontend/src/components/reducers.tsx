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

export const reducer = (state = initialState, action: any) => {
  switch (action.type) {
    case Action.LoadListings:
      return {
        isLoading: false,
        listings: action.listings,
      };

    case Action.HideListing:
      const updatedListings = [];
      for (const l of state.listings) {
        if (l.id !== action.listingID) {
          updatedListings.push(l);
        }
      }
      return { ...state, listings: updatedListings };

    default:
      return state;
  }
};
