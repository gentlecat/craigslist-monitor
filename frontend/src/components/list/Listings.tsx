import React, { Component } from 'react';
import { css } from '@emotion/core';
import axios from 'axios';
import { ListItem } from './ListItem';

interface State {
  data: ListItem[] | undefined;
  loadingState: LoadingState;
}

export interface ListItemInt {
  id: string;
  title: string;
  url: string;
  prices: Price[];
  images: string[];
  isHidden: boolean;
}

export interface Price {
  price: number;
}

enum LoadingState {
  Loading,
  Loaded,
  Error,
}

const loadList = () =>
  axios
    .get('/api/list')
    .then(response => {
      console.log(response.data);
      return response.data;
    })
    .catch(error => {
      console.error(error);
    });

export class Listings extends Component<any, State> {
  public state = {
    data: undefined,
    loadingState: LoadingState.Loading,
  };

  public componentDidMount = async () => {
    const data = await loadList();
    this.setState({ data, loadingState: LoadingState.Loaded });
  };

  private renderList = (games: ListItemInt[]) => {
    let items = [];
    games.forEach(g => {
      items.push(<ListItem key={g.id} listing={g} />);
    });
    return (
      <div
        css={css`
          display: flex;
          flex-direction: column;
        `}
      >
        {items}
      </div>
    );
  };

  public render = () => {
    switch (this.state.loadingState) {
      case LoadingState.Loading:
        return <div>Loading...</div>;
      case LoadingState.Loaded:
        return this.renderList(this.state.data);
      case LoadingState.Error:
      default:
        return <div>Error occurred :(</div>;
    }
  };
}
