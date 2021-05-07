import React from 'react';
import {connect} from 'react-redux';
import {todoActions} from '../../_actions';
import {Button, Grid, Paper, TextField} from '@material-ui/core';

class AddTodo extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      title: '',
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(e) {
    const {name, value} = e.target;
    this.setState({[name]: value});
  }

  handleSubmit(e) {
    e.preventDefault();

    const {title} = this.state;
    if (title) {
      this.props.add({title: title});
    }
    this.setState({'title': ''});
  }

  render() {
    const {title} = this.state;

    return (
        <Grid container spacing={2}>
          <Grid item xs={12}>
            <Paper style={{padding: 20}}>
              <Grid container spacing={2}>
                <Grid item>
                  <TextField label="new task" variant="outlined" name={'title'}
                             size={'small'} value={title}
                             onChange={this.handleChange}/>
                </Grid>
                <Grid item>
                  <Button variant="contained" color="primary"
                          onClick={this.handleSubmit}>add</Button>
                </Grid>
              </Grid>
            </Paper>
          </Grid>
        </Grid>
    );
  }
}

function mapStateToProps(state) {
  const {todo} = state;
  return {todo};
}

const actionCreators = {
  add: todoActions.add,
};

const connectedAddTodo = connect(mapStateToProps, actionCreators)(AddTodo);
export {connectedAddTodo as AddTodo};