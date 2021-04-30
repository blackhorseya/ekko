import {useDispatch} from 'react-redux';
import {add} from './todoSlice';

export function AddTodo() {
  const dispatch = useDispatch();

  const handleAdd = () => {
    dispatch(
        add('a new task'),
    );
  };

  return (
      <div>
        <button onClick={handleAdd}>add todo</button>
      </div>
  );
}