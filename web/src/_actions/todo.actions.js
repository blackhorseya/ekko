import {todoConstants} from '../_constants';

function list() {
  return dispatch => {
    dispatch(request());

    // todo: 2021-05-04|20:51|doggy|implement caller api via services
  };

  function request() {
    return {type: todoConstants.LIST_REQUEST};
  }

  function success(tasks) {
    return {type: todoConstants.LIST_SUCCESS, tasks};
  }

  function failure(error) {
    return {type: todoConstants.LIST_FAILURE, error};
  }
}