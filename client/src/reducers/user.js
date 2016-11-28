import {USER} from '../constants/actionTypes';

const initialState = {
  name: null,
  isAuthenticated: false,
  users: [],
  resultSet: {},
  error: {}
};
export const userReducer = (state = initialState, action) => {
  let newstate = Object.assign({}, state);
  switch (action.type) {

    case USER.GET_USERS_SUCCESS:
      return Object.assign({}, newstate,
        {
          users: action.payload.results,
          resultSet: action.payload.metadata.resultset
        });

    case USER.GET_USERS_FAILED:
      return Object.assign({}, newstate, {
        error: action.payload
      });

    //case LOGIN_SUCCESS:
    //  return {...state, name: action.payload.name, isAuthenticated: action.payload.isAuthenticated}
    //
    //case LOGIN_FAIL:
    //  // TODO
    //  return state
    //
    //case LOGOUT_SUCCESS:
    //  // TODO
    //  return state

    default:
      return newstate
  }
};
