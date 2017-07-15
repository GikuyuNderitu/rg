import React, { Component } from 'react';

import styles from './App.sass'

class App extends Component {
    render() {
        return (
            <div className={styles.container}>
                <h1>Hello {this.props.title}</h1>
            </div>
        )
    }
}

export default App;