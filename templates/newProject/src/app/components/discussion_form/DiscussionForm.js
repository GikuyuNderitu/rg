import React, { Component } from 'react';

class DiscussionForm extends Component {
    state = {};
    handleSubmit = () => this.setState({title:'', content:''});
    handleChange = (e, {name, value}) => this.setState({[name]: value});

    render() {
        return(
            <Form onSubmit={this.handleSubmit}>
                <Form.Input label="Title" type="text" onChange={this.handleChange}/>
                <Form.TextArea label="Enter discussion Content" onChange={this.handleChange} />
                <Button type="submit">Submit</Button>
            </Form>
        )
    }
}

export default DiscussionForm;