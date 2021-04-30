import {useDispatch, useSelector} from 'react-redux';
import {add, selectTodo} from './todoSlice';

export function TodoList() {
  const todo = useSelector(selectTodo);

  return (
      <div>
        <ul>
          {todo.tasks.map((i) => (
              <li key={i.id}>{i.name}</li>
          ))}
        </ul>
      </div>
  );
}