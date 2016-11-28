import { createStore, applyMiddleware, compose } from 'redux';
import thunkMiddleware from 'redux-thunk';
import { rootReducer } from '../reducers';
import { redirect } from '../middlewares/redirect';
import axiosMiddleware from 'redux-axios-middleware';
import axios from '../helpers/axios';
import {axiosHeaders} from '../middlewares/axios';

export default function configureStore() {
  const middlewares = [thunkMiddleware, redirect, axiosHeaders, axiosMiddleware(axios.client, axios.options)];
  const isDev = process.env.NODE_ENV === 'development';
  if (isDev && process.env.LOGGER) {
    const logger = require('redux-logger')();
    middlewares.push(logger);
  }

  const store = compose(
    applyMiddleware(...middlewares),
    (isDev && window.devToolsExtension) ? window.devToolsExtension() : f => f
  )(createStore)(rootReducer);

  if (module.hot) {
    module.hot.accept('../reducers', () => {
      const nextRootReducer = require('../reducers').rootReducer;
      store.replaceReducer(nextRootReducer)
    });
  }

  return store
}
