import React from 'react';
import './App.css';
import {TodoList} from './features/todo/TodoList';
import {AddTodo} from './features/todo/AddTodo';

function App() {
  return (
      <div className="App">
        <header className="App-header">
          <AddTodo/>
          <TodoList/>
        </header>
      </div>
  );
}

export default App;
