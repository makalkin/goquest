import {getUsers} from '../actions/user';

class UserAPI {
  getUsers(searchQuery) {
    return dispatch => {
      return dispatch(getUsers(searchQuery)).then(response => {
        return response;
      }).catch((response) => {
        return response;
      });
    }
  }
}

export default UserAPI;
