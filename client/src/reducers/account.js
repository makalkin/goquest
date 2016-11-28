import {ACCOUNT} from '../constants/actionTypes';

export const initialState = {
  roles: [],
  defaultRoleId: null,
  isLoading: false,
  accounts: [],
  resultSet: {},
  error: {}
};


export const accountReducer = (state = initialState, action) => {
  let newstate = Object.assign({}, state);

  switch (action.type) {
    case ACCOUNT.GET_ACCOUNT_ROLES_SUCCESS:
      return Object.assign({}, state,
        {
          roles: action.payload,
          defaultRoleId: action.payload && action.payload.length ? action.payload[0].id : null
        });

    case ACCOUNT.CREATE_ACCOUNT_SUCCESS:
      return Object.assign({}, state, {roles: action.payload});

    case ACCOUNT.GET_ACCOUNTS_SUCCESS:
      return Object.assign({}, newstate,
        {
          accounts: action.payload.results,
          resultSet: action.payload.metadata.resultset
        });

    case ACCOUNT.GET_ACCOUNTS_FAIL:
      return Object.assign({}, newstate, {
        error: action.payload
      });

    default:
      return newstate;
  }
};
