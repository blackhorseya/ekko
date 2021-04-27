import {useSelector} from 'react-redux';
import {selectTodo} from './todoSlice';

export function Todo() {
  const todo = useSelector(selectTodo);

  return (
      <ul>
        {todo.tasks.map((i) => (
            <li key={i.id}>{i.name}</li>
        ))}
      </ul>
  );
}