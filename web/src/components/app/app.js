import React from 'react';
import logo from '../../images/logo.svg';
import './app.css';
import {taskActions} from '../../_actions';
import {connect} from 'react-redux';

class App extends React.Component {
  render() {
    return (
        <div className="App">
          <header className="App-header">
            <img src={logo} className="App-logo" alt="logo"/>
            <p>
              Edit <code>src/App.js</code> and save to reload.
            </p>
            <a
                className="App-link"
                href="https://reactjs.org"
                target="_blank"
                rel="noopener noreferrer"
            >
              Learn React
            </a>
          </header>
        </div>
    );
  }
}

export default App;

function mapStateToProps(state) {
  const {tasks} = state.data;
  return {tasks};
}

const actionCreators = {
  list: taskActions.list,
};

const connectedApp = connect(mapStateToProps, actionCreators)(App);
export {connectedApp as App};
