import {combineReducers} from '@reduxjs/toolkit';
import {todos} from './todos.reducer';

const rootReducer = combineReducers({
  todos,
});

export default rootReducer;