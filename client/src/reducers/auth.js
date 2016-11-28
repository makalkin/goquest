/**
 * Auth store holds everything related to authentication and authorization of user.
 */
import {AUTH} from '../constants/actionTypes';
import * as Cookies from "js-cookie";
import {AUTH_COOKIE_KEY, JWT_PAYLOAD_ROLES_KEY} from '../constants/common';
// import {KJUR, b64utoutf8} from 'jsrsasign';

export const initialState = {
	isAuthenticated: false,
	headers: {},
  roles: [],
	expiresOn: null,
  validForRefreshUntil: null,
	isLoading: false,
	error: {},
  isFetchingRefreshToken: false,
  lastRequest: {}
};

let cookieState = Cookies.get(AUTH_COOKIE_KEY);

if (cookieState) {
	cookieState = JSON.parse(cookieState)
}

export const auth = (state = cookieState || initialState, action) => {
	switch (action.type) {
		case AUTH.LOGIN_USER_REQUEST:
			return Object.assign({}, state, {
				isAuthenticated: false,
				isLoading: true
			});
    case AUTH.REFRESH_TOKEN:
			return Object.assign({}, state, {
        isFetchingRefreshToken: true
			});
    case AUTH.LOGIN_USER_SUCCESS:
      // Here we extract payload encoded into JWT payload, and make sure they are saved as an array.
      // let payload = KJUR.jws.JWS.readSafeJSONString(b64utoutf8(action.payload.access_token.split(".")[1]));
      // let roles = payload[JWT_PAYLOAD_ROLES_KEY];
      // if (typeof roles === 'string' || !Array.isArray(roles)) {
      //   roles = [roles];
      // }
		// 	let authState = Object.assign({}, initialState, {
		// 		isAuthenticated: true,
		// 		headers: {Authorization: `Bearer ${action.payload.access_token}`},
      //   expiresOn: payload['exp'],
      //   validForRefreshUntil: payload['refresh_expiration'],
      //   roles: roles
		// 	});
      // Cookies.set(AUTH_COOKIE_KEY, authState, {expires: (payload['refresh_expiration'] - (new Date()) / 1000) / 86400});  // All cookies must die!
		// 	return authState;
		case AUTH.LOGIN_USER_FAIL:
			return Object.assign({}, initialState, {
				error: action.payload
			});
		case AUTH.LOGOUT:
			// Cookies.remove(AUTH_COOKIE_KEY);
			// return initialState;
		default:
		  if (action.payload && action.payload.request) {
		    let lastRequest = Object.assign({}, action);
		    if (lastRequest.types) {
		      delete lastRequest.type;
        }
		    return Object.assign({}, state, {lastRequest: lastRequest});
      }
			return state;
	}
};
