import React from 'react';
import { css } from '@emotion/core';

export const Header = () => {
  return (
    <nav>
      <h1
        css={css`
          font-size: 2em;
        `}
      >
        Craigslist Monitor
      </h1>
    </nav>
  );
};
