import {configureStore} from '@reduxjs/toolkit';
import counterReducer from '../features/counter/counterSlice';
import rootReducer from '../_reducers';

export const store = configureStore({
  reducer: {
    root: rootReducer,
    counter: counterReducer,
  },
});
