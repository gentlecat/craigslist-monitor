import { css } from '@emotion/core';
import React from 'react';
import axios from 'axios';

const changeStatus = (id: number, isHidden: boolean) => {
  return axios
    .post(isHidden ? '/api/hide' : '/api/unhide', {
      listingId: id,
    })
    .then(response => {
      console.log(response);
    })
    .catch(error => {
      console.error(error);
    });
};

export const ListItem = ({ listing }) => {
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
          src={listing['images'][0]}
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
        <div
          css={css`
            font-weight: bold;
          `}
        >
          ${listing['prices'][0]['price']}
        </div>
        <div>
          <a href={listing['url']}>{listing['title']}</a>
        </div>
        <div>
          <button
            onClick={() => changeStatus(listing['id'], !listing['isHidden'])}
          >
            {listing['isHidden'] ? 'Unhide' : 'Hide'}
          </button>
        </div>
        <div>
          <textarea rows="3" cols="33" placeholder="Notes..." />
        </div>
      </div>
    </div>
  );
};
