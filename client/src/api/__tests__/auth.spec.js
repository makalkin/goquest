/*
	This is the first working test example for async APIs.

	STEP 1: import all the stuff (for store and http mocking, APi itself, actions, consts, etc.)
 */

import configureMockStore  from 'redux-mock-store';
import MockAdapter from 'axios-mock-adapter';
import axiosMiddleware from 'redux-axios-middleware';
import thunk from 'redux-thunk'

import {axiosHeaders} from '../../middlewares/axios.js';
import {auth} from '../index';
import axios from '../../helpers/axios';
import {BASE_URL} from '../../constants/common'
import {loginRequest} from '../../actions/auth';
import {AUTH} from '../../constants/actionTypes';

let mock = new MockAdapter(axios);
let middlewares = [thunk, axiosHeaders, axiosMiddleware(axios)];
let mockStore = configureMockStore(middlewares);

// STEP 2: define and setup general test suit.
describe("Auth async actions", ()=> {
	afterEach(() => {
		mock.reset();
	});

	// STEP 3: create individual test cases for API methods.
	it("Should login", () => {
		let token = "TOKEN",
			email = "test@test.com",
			password = "password",
			response = {
				access_token: token,
				expires_in: '300'
			};
		// Here we mock API endpoint.
		mock.onPost('/api/v1/Auth/Signin')
			.reply(200, response);

		// Here we mock store and actions.
		const store = mockStore();
		let expectedActions = [{
			...loginRequest(email, password),
			type: AUTH.LOGIN_USER_REQUEST
		}, {
			type: AUTH.LOGIN_USER_SUCCESS,
			payload: response,
			meta: {previousAction: loginRequest(email, password)}
		}];

		return store.dispatch(auth.login(email, password)).then((action) => {
			// Here we check if the API dispatches all the actions we want it to dispatch.
			expect(store.getActions()).toEqual(expectedActions)
		});

	})
});