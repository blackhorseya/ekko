import React from 'react';
import {taskActions} from '../../_actions';
import {connect} from 'react-redux';

class ListTodo extends React.Component {
  componentDidMount() {
    this.props.list();
  }

  render() {
    const {tasks} = this.props;

    return (
        <div>
          <h1>Todos</h1>
          <ul>
            {
              tasks.item && tasks.item.map((task, index) =>
                  <li key={index}>
                    {task.title}
                  </li>,
              )
            }
          </ul>
        </div>
    );
  }
}

function mapStateToProps(state) {
  const {tasks} = state;
  return {tasks};
}

const actionCreators = {
  list: taskActions.list,
};

const connectedListTodo = connect(mapStateToProps, actionCreators)(ListTodo);
export {connectedListTodo as ListTodo};
