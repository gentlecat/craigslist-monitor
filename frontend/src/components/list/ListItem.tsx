import { css } from '@emotion/core';
import React from 'react';
import dayjs from 'dayjs';
import relativeTime from 'dayjs/plugin/relativeTime';
import { Listing } from './Listings';

dayjs.extend(relativeTime);

type ListItemProps = {
  listing: Listing;
  onHideListing: Function;
};

export const ListItem = ({ listing, onHideListing }: ListItemProps) => {
  return (
    <div
      css={css`
        flex: 1;
        display: flex;
        padding-top: 20px;
        padding-bottom: 20px;
        border-bottom: 1px solid #7d7d7d1c;
        ${listing.isHidden ? 'opacity: 0.4;' : ''}
      `}
    >
      <div
        css={css`
          flex: 1;
          max-width: 230px;
        `}
      >
        <img
          src={listing.images[0]}
          css={css`
            max-width: 200px;
            max-height: 200px;
          `}
        />
      </div>

      <div
        css={css`
          flex: 1;
        `}
      >
        <div>
          <span
            css={css`
              font-weight: bold;
            `}
          >
            $
            {listing.prices[0].price}
          </span>
          {' '}
          – 
          {' '}
          <a href={listing.url}>{listing.title}</a>
        </div>
        <div>
          Posted
          {' '}
          <span title={listing.postedAt.toString()}>
            {dayjs(listing.postedAt).fromNow()}
          </span>
        </div>
        <div>
          Last updated
          {' '}
          <span title={listing.updatedAt.toString()}>
            {dayjs(listing.updatedAt).fromNow()}
          </span>
        </div>
        <div>
          <button onClick={() => onHideListing()}>
            {listing.isHidden ? 'Unhide' : 'Hide'}
          </button>
        </div>
        <div>
          <textarea rows="3" cols="33" placeholder="Notes..." />
        </div>
      </div>
    </div>
  );
};
