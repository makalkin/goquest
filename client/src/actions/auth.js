/**
 * This module holds authentication actions that are to be dispatched.
 */

import {AUTH} from '../constants/actionTypes';

/**
 * Begin the auth process.
 * @returns {{type}}
 */
export const loginRequest = (email, password) => {
	return {
		types: [AUTH.LOGIN_USER_REQUEST, AUTH.LOGIN_USER_SUCCESS, AUTH.LOGIN_USER_FAIL],
		payload: {
			request: {
				url: '/api/v1/auth/signin',
				method: 'post',
				data: {
					email,
					password
				}
			}
		}
	}
};

/**
 * Refresh token action creator.
 * @returns {{types: *[], payload: {request: {url: string, method: string}}}}
 */
export const refreshToken = () => {
	return {
		types: [AUTH.REFRESH_TOKEN, AUTH.LOGIN_USER_SUCCESS, AUTH.LOGOUT],
    payload: {
			request: {
				url: '/api/v1/auth/refresh-token',
				method: 'post'
			}
		}
	}
};

/**
 * Logout.
 * @returns {{type}}
 */
export const logout = () => {
	return {
		type: AUTH.LOGOUT
	}
};
