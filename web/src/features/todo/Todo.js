import {useDispatch, useSelector} from 'react-redux';
import {add, selectTodo} from './todoSlice';

export function Todo() {
  const todo = useSelector(selectTodo);
  const dispatch = useDispatch();

  const handleAdd = () => {
    dispatch(
        add('a new task'),
    );
  };

  return (
      <div>
        <button onClick={handleAdd}>add todo</button>
        <ul>
          {todo.tasks.map((i) => (
              <li key={i.id}>{i.name}</li>
          ))}
        </ul>
      </div>
  );
}