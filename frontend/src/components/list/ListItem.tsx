import { css } from '@emotion/core';
import React from 'react';
import dayjs from 'dayjs';
import relativeTime from 'dayjs/plugin/relativeTime';
import { Note } from './Note';
import { Listing } from './Listings';
import { Metadata } from './Metadata';

dayjs.extend(relativeTime);

type ListItemProps = {
  listing: Listing;
  onHideListing: Function;
  onNoteUpdate: Function;
};

export const ListItem = React.memo(
  ({ listing, onHideListing, onNoteUpdate }: ListItemProps) => {
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
            text-align: center;
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
            margin-left: 20px;
          `}
        >
          <Metadata listing={listing} />

          <div
            css={css`
              margin-top: 20px;
            `}
          >
            <Note
              note={listing.note}
              onNoteUpdate={(newNote: string) =>
                onNoteUpdate(listing.id, newNote)}
            />
          </div>
        </div>

        <div
          css={css`
            flex: 1;
            max-width: 60px;
            margin-left: 20px;
          `}
        >
          <button onClick={() => onHideListing(listing.id)}>
            {listing.isHidden ? 'Unhide' : '🗑'}
          </button>
        </div>
      </div>
    );
  }
);
