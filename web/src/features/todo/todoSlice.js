import {createAsyncThunk, createSlice} from '@reduxjs/toolkit';

const initialState = {
  tasks: [],
  loading: false,
  error: false,
};

export const listTasks = createAsyncThunk(
    'tasks',
    async (start, end) => {
      const resp = await fetch(`http://localhost:8080/api/v1/tasks`);
      return resp.data;
    },
);

export const todoSlice = createSlice({
  name: 'todo',
  initialState: initialState,
  reducers: {
    list: (state, action) => {
      state.tasks = action.payload;
      state.loading = true;
      state.error = false;
    },
    add: (state, action) => {
      state.tasks.push({
        id: state.tasks[state.tasks.length - 1].id + 1,
        name: action.payload,
      });
      state.loading = false;
    },
  },
  extraReducers: (builder) => {
    builder.addCase(listTasks.pending, (state) => {
      state.loading = true;
    }).addCase(listTasks.fulfilled, (state, action) => {
      state.tasks.push(action.payload);
      state.loading = false;
    });
  },
});

export const {list, add} = todoSlice.actions;

export const selectTodo = (state) => state.todo;

export default todoSlice.reducer;
