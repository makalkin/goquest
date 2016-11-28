/**
 * This module contains configurations for axios client which is used as base http client for XHR.
 */

import axios from 'axios';
import {BASE_URL} from '../constants/common';
import Qs from 'qs';
import {store} from '../index';
import {auth} from '../api';

let client = axios.create({
	baseURL: BASE_URL
});

/**
 * The default url param serializer. For all those GET methods with filters.
 * @param params
 * @returns {*|string}
 */
axios.paramsSerializer = (params) => {
	return Qs.stringify(params, {arrayFormat: 'brackets'})
};

// TODO: might be useful might be not, needs review.
client.interceptors.request.use(function (config) {
	return config;
});

/*
  What happens here is very important. In case of request error we check if it's auth related, then
  if it wasn't about token refresh we have to ask for token and try that failed request again. In this
  case an empty Promise is returned so axios-middleware is not confused which action should be send. In
  other cases we just reject the error.
 */
client.interceptors.response.use(response => {
  return response.data;
}, error => {
  return new Promise((resolve, reject) => {
    let state = store.getState();
    let isInvalidToken = false;
    try {
      isInvalidToken = error.response.headers['www-authenticate'].includes('invalid_token');
    } catch (e) {}
    if (error.response.status === 401 && isInvalidToken && !state.auth.isFetchingRefreshToken) {
      store.dispatch(auth.refreshToken()).then(() => {
        store.dispatch(state.auth.lastRequest);
      }, (err) => {
        return err;
      });
    } else {
      return reject(error);
    }
  });
});

let options = {
  returnRejectedPromiseOnError: true
};

export default {
  client,
  options
}
