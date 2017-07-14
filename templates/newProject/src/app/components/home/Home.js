import React, { Component } from 'react';

import { Container, Header } from 'semantic-ui-react';

class Home extends Component {
    constructor(props) {
        super(props);
        this.state={}
    }

    render() {
        return(
            <Container>
                <Header as="h1">Home</Header>
            </Container>
        )
    }
}

export default Home;