import React from 'react';
import ReactDOM from 'react-dom';
import { Main } from './components/Main';

import { createStore } from 'redux';
import { Provider } from 'react-redux';
import { reducer } from './components/reducers';

const reduxStore = createStore(
  reducer,
  window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__()
);

window.onload = function() {
  ReactDOM.render(
    <Provider store={reduxStore}>
      <Main />
    </Provider>,
    document.getElementById('app-container')
  );
};
