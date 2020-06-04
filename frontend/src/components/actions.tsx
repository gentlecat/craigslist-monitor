import { Listing } from './list/Listings';

export enum Action {
  LoadListings = 'LOAD_LISTINGS',
  HideListing = 'HIDE_LISTING',
}

export const loadListings = (listings: Listing[]) => ({
  type: Action.LoadListings,
  listings,
});

export const hideListing = (listingID: string) => ({
  type: Action.HideListing,
  listingID,
});
