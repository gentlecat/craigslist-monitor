import React, { Component } from 'react';
import { css } from '@emotion/core';
import axios from 'axios';
import _ from 'lodash';
import { connect } from 'react-redux';
import { loadListings, hideListing, updateNote } from '../actions';
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
  postedAt: Date;
  updatedAt: Date;
  note: string;
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
    .then((response) => {
      console.log(response.data);
      return response.data;
    })
    .catch((error) => {
      console.error(error);
    });

const changeStatus = (id: string, isHidden: boolean) => {
  return axios
    .post(isHidden ? '/api/hide' : '/api/unhide', {
      listingId: id,
    })
    .then((response) => {
      console.log(response);
    })
    .catch((error) => {
      console.error(error);
    });
};

const setNote = _.debounce(async (id: string, note: string) => {
  return axios
    .post('/api/note', {
      listingId: id,
      note,
    })
    .then((response) => {
      console.log(response);
    })
    .catch((error) => {
      console.error(error);
    });
}, 1200);

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

  private updateNote = (listingID: string, newNote: string) => {
    this.props.updateNote(listingID, newNote);
    setNote(listingID, newNote);
  };

  private renderList = (listings: Listing[]) => {
    const items = [];
    listings.forEach((l) => {
      items.push(
        <ListItem
          key={l.id}
          listing={l}
          onHideListing={this.hideListing}
          onNoteUpdate={this.updateNote}
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
        return (
          <div
            css={css`
              text-align: center;
              margin-top: 100px;
              color: #6a6a6a;
            `}
          >
            Loading...
          </div>
        );
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
  { loadListings, hideListing, updateNote }
)(Listings);
