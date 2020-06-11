import React from 'react';
import dayjs from 'dayjs';
import { css } from '@emotion/core';
import { Listing } from './Listings';
import { Price } from './Price';

type MetadataProps = {
  listing: Listing;
};

export const Metadata = React.memo(({ listing }: MetadataProps) => {
  const prices: number[] = [];
  const priceElements = [];

  listing.prices.forEach((p) => {
    let isDown;
    if (prices.length > 0) {
      const previous = prices[prices.length - 1];

      if (previous === p.price) return;

      if (p.price < previous) {
        isDown = true;
      } else if (p.price > previous) {
        isDown = false;
      }
    }

    prices.push(p.price);
    priceElements.push(<Price price={p.price} isDown={isDown} />);
  });

  return (
    <div>
      <div
        css={css`
          display: flex;
          grid-gap: 6px;
          align-items: baseline;
        `}
      >
        {
          // Reversing since we want to show current price
          priceElements.reverse()
        }
      </div>

      <div>
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
