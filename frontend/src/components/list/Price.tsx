import React from 'react';
import { css } from '@emotion/core';
import { is } from 'immer/dist/internal';

type PriceProps = {
  price: number;
  isDown: boolean | undefined;
};

export const Price = ({ price, isDown }: PriceProps) => {
  return (
    <div
      css={css`
        &:first-child {
          font-size: 1.25rem;
        }
      `}
    >
      <span
        css={css`
          ${isDown !== undefined ? `color: ${isDown ? 'green' : 'red'}` : ''}
        `}
      >
        $
        {price}
      </span>
    </div>
  );
};

export default React.memo(Price);
