import {useDispatch, useSelector} from 'react-redux';
import {listTasks, selectTodo} from './todoSlice';

export function TodoList() {
  const todo = useSelector(selectTodo);
  const dispatch = useDispatch();

  return (
      <div>
        <button onClick={() => dispatch(listTasks())}>list</button>
        <ul>
          {todo.tasks.map((i) => (
              <li key={i.id}>{i.name}</li>
          ))}
        </ul>
      </div>
  );
}