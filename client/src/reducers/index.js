import { combineReducers } from 'redux';
import { routerReducer } from 'react-router-redux';
import { reducer as formReducer } from 'redux-form';
import {userReducer} from './user'
import {auth} from './auth';
import {accountReducer} from './account';


export const rootReducer = combineReducers({
	routing: routerReducer,
	form: formReducer,
  account: accountReducer,
	auth,
	user: userReducer
});
