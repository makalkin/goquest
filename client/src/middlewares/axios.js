/**
 * This middleware checks if dispatched action is http request and if so fills headers if none specified.
 * Exists for authorization sake and might be obsolete.
 *
 * Looks bizarre but actually is useful.
 *
 * @param store
 */
export const axiosHeaders = store => next => action => {
  if (action.payload && action.payload.request) {
    action.payload.request.headers = Object.assign({}, action.payload.request.headers, store.getState().auth.headers);
  }
  return next(action)
};
