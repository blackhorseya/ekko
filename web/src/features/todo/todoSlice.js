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
      state.tasks.push({
        id: state.tasks[state.tasks.length - 1].id + 1,
        name: action.payload,
      });
    },
  },
});

export const {add} = todoSlice.actions;

export const selectTodo = (state) => state.todo;

export default todoSlice.reducer;
