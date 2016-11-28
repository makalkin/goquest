import {loginRequest, logout} from '../auth';
import {AUTH} from '../../constants/actionTypes';

describe("Auth related actions", ()=> {
	test("login request", () => {
		let email = "test@test.com",
			password = "thePassword";

		let expectedAction = {
			types: [AUTH.LOGIN_USER_REQUEST, AUTH.LOGIN_USER_SUCCESS, AUTH.LOGIN_USER_FAIL],
			payload: {
				request: {
					url: '/api/v1/Auth/Signin',
					method: 'post',
					data: {
						email,
						password
					}
				}
			}
		};

		expect(loginRequest(email, password)).toEqual(expectedAction);
	});

	test("logout", () => {
		let expectedAction = {
			type: AUTH.LOGOUT
		};

		expect(logout()).toEqual(expectedAction);
	});

});