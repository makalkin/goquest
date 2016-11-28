import {ACCOUNT} from '../constants/actionTypes';

/**
 * Action creator for account roles fetching.
 *
 * @returns FluxAlmostStandardAction
 */
export const getAccountRolesRequest = () => {
  return {
    types: [ACCOUNT.GET_ACCOUNT_ROLES_REQUEST, ACCOUNT.GET_ACCOUNT_ROLES_SUCCESS, ACCOUNT.GET_ACCOUNT_ROLES_FAIL],
    payload: {
      request: {
        url: '/api/v1/accounts/roles',
        method: 'get'
      }
    }
  }
};


/**
 * Action creator for account creation.
 *
 * @param account
 * @returns FluxAlmostStandardAction
 */
export const createAccountRequest = (account) => {
  return {
    types: [ACCOUNT.CREATE_ACCOUNT_REQUEST, ACCOUNT.CREATE_ACCOUNT_SUCCESS, ACCOUNT.CREATE_ACCOUNT_FAIL],
    payload: {
      request: {
        url: '/api/v1/accounts',
        method: 'post',
        data: {
          ...account
        }
      }
    }
  }
};

export const getAccounts = (searchQuery) => {
  return {
    types: [ACCOUNT.GET_ACCOUNTS_REQUEST, ACCOUNT.GET_ACCOUNTS_SUCCESS, ACCOUNT.GET_ACCOUNTS_FAIL],
    payload: {
      request: {
        url: "/api/v1/accounts" + (searchQuery || ''),
        method: "get"
      }
    }
  }
};
