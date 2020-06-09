import { Listing } from './list/Listings';

export enum Action {
  LoadListings = 'LOAD_LISTINGS',
  HideListing = 'HIDE_LISTING',
  UpdateNote = 'UPDATE_NOTE',
}

export const loadListings = (listings: Listing[]) => ({
  type: Action.LoadListings,
  listings,
});

export const hideListing = (listingID: string) => ({
  type: Action.HideListing,
  listingID,
});

export const updateNote = (listingID: string, note: string) => ({
  type: Action.UpdateNote,
  listingID,
  note,
});
