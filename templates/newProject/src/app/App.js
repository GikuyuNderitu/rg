import React, { Component } from 'react';
import { connect } from 'react-redux';
import {BrowserRouter as Router, Route} from 'react-router-dom';

class App extends Component {
    componentDidMount() {
        console.log(this.props);
    }
    render() {
        return (
            <Router>
                <div>
                    <Header user={this.props.user}/>
                    <Route exact path="/" component={Home}/>
                    <Route path="/discussion" component={Discussion}/>
                    <Route path="/statistics" component={Statistics}/>
                </div>
            </Router>
        )
    }
}

const mapStateToProps = state => {
    return {user: state.authReducer.user}
}

export default connect(mapStateToProps)(App);