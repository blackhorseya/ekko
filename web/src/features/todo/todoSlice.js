import {createSlice} from '@reduxjs/toolkit';

const initialState = {
  tasks: [
    {id: 1, name: 'first todo on redux'},
    {id: 2, name: 'second todo in list'},
  ],
};

export const todoSlice = createSlice({
  name: 'todo',
  initialState: initialState,
  reducers: {
    add: (state, action) => {
      state.tasks.push(action.payload);
    },
  },
});

export const selectTodo = (state) => state.todo;

export default todoSlice.reducer;
