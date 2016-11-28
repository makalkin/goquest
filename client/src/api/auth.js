import {loginRequest, logout, refreshToken} from '../actions/auth';

/**
 * This API provides methods to work with user authentication in application.
 */
class AuthAPI {
	/**
	 * Authenticate user with given credentials. On success acquire and save auth payload.
	 * @param email
	 * @param password
	 */
	login(email, password) {
		return dispatch => {
			return dispatch(loginRequest(email, password)).then(response => {
				return response;
			}).catch((response) => {
				return response;
			});
		}
	}

  refreshToken() {
    return dispatch => {
      return dispatch(refreshToken());
    }
  }

	/**
	 * Logout user from the application. Cleanup auth data.
	 */
	logout() {
		return dispatch => {
			return dispatch(logout());
		}
	}
}

export default AuthAPI;
