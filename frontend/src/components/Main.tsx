import React from 'react';
import { css, Global, jsx } from '@emotion/core';
import emotionNormalize from 'emotion-normalize';
import { charterFont } from 'fonts';
import { Header } from '../components/Header';
import { Listings } from './list/Listings';

export const Main = () => {
  return (
    <div>
      <Global
        styles={css`
          ${emotionNormalize}
          ${charterFont}

          html,
          body {
            padding: 0;
            margin: 0;
            background: white;
            min-height: 100%;
            font-family: 'Charter';
            font-size: 16px;
            color: black;
          }
        `}
      />

      <div
        css={css`
          padding: 20px;
          max-width: 1200px;
          margin: 0 auto;
        `}
      >
        <Header />
        <Listings />
      </div>
    </div>
  );
};
