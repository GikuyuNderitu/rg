import React, { Component } from 'react';
import { connect } from 'react-redux';

import { Container, Header, Feed, Loader, Dimmer, Segment } from 'semantic-ui-react';


import { RECEIVED_DISCUSSIONS, LOADING_DISCUSSIONS, FAILED_DISCUSSIONS } from '../../state/types';
import { getDiscussions } from '../../state/actions';

class Discussion extends Component {
    constructor(props) {
        super(props);
        this.state={}
    }

    getDiscussions() {
        this.props.getDiscussions();
    }

    componentDidMount() {
        this.getDiscussions()
    }

    componentWillReceiveProps({loading}){
        console.log(loading);
    }

    render() {
        return(
            <Container>
                <Header as="h1">Discussion</Header>

                <Segment>
                    <Feed>
                    
                        <Dimmer active={this.props.loading}>
                            <Loader />
                        </Dimmer>

                        {this.props.discussions.map((val, idx) => (
                            <Feed.Event key={idx}>
                                <Feed.Content>
                                    <Feed.Summary>{val.summary}</Feed.Summary>
                                    <Feed.Date>{val.date}</Feed.Date>
                                </Feed.Content>
                            </Feed.Event>
                        ))}
                        
                    </Feed>

                    {this.props.loading ?
                        <div>
                        <p>
                            LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR 
                        </p>

                        <p>
                            LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR 
                        </p>
                        <p>
                            LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR 
                        </p>
                        <p>
                            LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR LOREM IPSUM COLOR DOLOR 
                        </p> 
                        </div>:
                        null
                    }
                </Segment>
                

            </Container>
        )
    }
}

const mapStatetoProps = ({discussReducer}) => (
    {
        discussions: discussReducer.events,
        loading: discussReducer.loading
    }
)

const mapDispatchToProps = dispatch => (
    {
        getDiscussions() {
            dispatch({type: LOADING_DISCUSSIONS});
            getDiscussions()
            .then( events => dispatch({type: RECEIVED_DISCUSSIONS, payload:events}))
            .catch( err => dispatch({type: FAILED_DISCUSSIONS, payload: err}));
        }
    }
)

export default connect(mapStatetoProps, mapDispatchToProps)(Discussion);