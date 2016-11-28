import {getAccountRolesRequest, createAccountRequest, getAccounts} from '../actions/account';

class AccountAPI {
  getAccountRoles() {
    return dispatch => {
      return dispatch(getAccountRolesRequest()).then(response => {
        return response;
      }).catch((response) => {
        return response;
      });
    }
  }

  createAccount(account) {
    return dispatch => {
      return dispatch(createAccountRequest(account)).then(response => {
        return response;
      }).catch((response) => {
        return response;
      });
    }
  }

  getAccounts(searchQuery) {
    return dispatch => {
      return dispatch(getAccounts(searchQuery)).then(response => {
        return response;
      }).catch((response) => {
        return response;
      });
    }
  }

}

export default AccountAPI;
