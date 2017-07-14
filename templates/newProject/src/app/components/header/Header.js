import React, { Component } from 'react';

import { Link, withRouter } from 'react-router-dom';

import { 
    Container,
    Divider,
    Dropdown,
    Header,
    Menu } from 'semantic-ui-react';

import styles from './Header.sass';


const links = [
    {path: "/", root: "home", title: "Home"},
    {path: "/discussion", root: "discussion", title: "Discussion", color: "blue"},
    {path: "/statistics", root: "statistics", title: "Statistics", color: "blue"},
]
class MyHeader extends Component {
    constructor(props) {
        super(props);

        console.log(props);

        this.state={
            active:"home"
        };
        this.handleMenuItem = this.handleMenuItem.bind(this);
    }

    handleMenuItem(e, {name}) {
        this.setState({active: name})
    }

    render() {
        const { pathname } = this.props.location; 

        const active = pathname === '/' ? 'home' : pathname.substring(1);

        return(
            <header className={styles.headerContainer}>
                <Header className={styles.logo} size="huge">Health Nut</Header>

                <Container>
                    <Menu color="teal" pointing secondary>
                        {links.map((link, idx) => (
                            <Menu.Item
                                key={idx}
                                as={Link}
                                to={link.path}
                                onClick={this.handleMenuItem}
                                active={active ===  link.root}
                                name={link.root}
                                color={link.color || "black"}>
                                {link.title}
                            </Menu.Item>
                        ))}
                        <Menu.Menu position="right">
                            <Dropdown text={this.props.user.name}>
                                <Dropdown.Menu>
                                    <Dropdown.Item>
                                        Profile
                                    </Dropdown.Item>
                                </Dropdown.Menu>
                            </Dropdown>
                        </Menu.Menu>
                    </Menu>                
                </Container>
            </header>
        )
    }
}

export default withRouter(MyHeader);