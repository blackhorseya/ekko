import React from 'react';
import './App.css';
import {TodoList} from './components/todo/TodoList';
import {AddTodo} from './components/todo/AddTodo';

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
