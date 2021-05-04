import React from 'react';
import {connect} from 'react-redux';

class AddTodo extends React.Component {
  render() {
    return (
        <div>
          <h1>Add Todo</h1>
        </div>
    );
  }
}

function mapStateToProps(state) {
  return {};
}

const actionCreators = {};

const connectedAddTodo = connect(mapStateToProps, actionCreators)(AddTodo);
export {connectedAddTodo as AddTodo};