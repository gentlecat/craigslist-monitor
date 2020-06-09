import React from 'react';
import dayjs from 'dayjs';
import { css } from '@emotion/core';
import { Listing } from './Listings';

type MetadataProps = {
  listing: Listing;
};

export const Metadata = React.memo(({ listing }: MetadataProps) => {
  return (
    <div>
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

      <div
        css={css`
          font-size: 0.85rem;
          margin-top: 6px;
          color: #6a6a6a;
        `}
      >
        Posted
        {' '}
        <span title={listing.postedAt.toString()}>
          {dayjs(listing.postedAt).fromNow()}
        </span>
        <br />
        Last updated
        {' '}
        <span title={listing.updatedAt.toString()}>
          {dayjs(listing.updatedAt).fromNow()}
        </span>
      </div>
    </div>
  );
});
