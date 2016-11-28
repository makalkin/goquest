import {USER} from '../constants/actionTypes';

export const getUsers = (searchQuery) => {
  return {
    types: [USER.GET_USERS_REQUEST, USER.GET_USERS_SUCCESS, USER.GET_USERS_FAILED],
    payload: {
      request: {
        url: "/api/v1/users" + (searchQuery || ''),
        method: "get"
      }
    }
  }
};
