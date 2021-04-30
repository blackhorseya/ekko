import {useSelector} from 'react-redux';
import {selectTodo} from './todoSlice';

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