import {auth, initialState} from '../auth';
import {AUTH} from '../../constants/actionTypes';


describe('Auth reducer', ()=> {
	it('should return initial state', () => {
		expect(auth(undefined, {})).toEqual(initialState);
	});

	it(`should handle ${AUTH.LOGIN_USER_REQUEST}`, () => {
		expect(auth(initialState, {type: AUTH.LOGIN_USER_REQUEST, payload: {}})).toEqual(
			Object.assign({}, initialState, {isAuthenticated: false, isLoading: true})
		);
	})

});