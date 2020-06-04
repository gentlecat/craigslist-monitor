import React, { Component } from 'react';
import { css } from '@emotion/core';
import axios from 'axios';
import { connect } from 'react-redux';
import { loadListings, hideListing } from '../actions';
import { ListItem } from './ListItem';
import { SystemState } from '../reducers';

interface State {
  data: Listing[] | undefined;
  loadingState: LoadingState;
}

export interface Listing {
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

const changeStatus = (id: string, isHidden: boolean) => {
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

class Listings extends Component<any, State> {
  public state = {
    data: undefined,
    loadingState: LoadingState.Loading,
  };

  public componentDidMount = async () => {
    const data = await loadList();

    this.props.loadListings(data);

    this.setState({ data, loadingState: LoadingState.Loaded });
  };

  private hideListing = (listingID: string) => {
    this.props.hideListing(listingID);
    changeStatus(listingID, true);
  };

  private renderList = (listings: Listing[]) => {
    let items = [];
    listings.forEach(l => {
      items.push(
        <ListItem
          key={l.id}
          listing={l}
          onHideListing={() => this.hideListing(l.id)}
        />
      );
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
        return this.renderList(this.props.listings);
      case LoadingState.Error:
      default:
        return <div>Error occurred :(</div>;
    }
  };
}

export default connect(
  (state: SystemState) => {
    return { listings: state.listings };
  },
  { loadListings, hideListing }
)(Listings);
