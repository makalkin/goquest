import {Navbar, Nav, NavDropdown} from "react-bootstrap";
import React from 'react';
import {connect} from 'react-redux'

class AppNavbar extends React.Component {
    render() {
        return(
            <Navbar>
                <Navbar.Header>
                    <Navbar.Brand>
                        <a href="/">{this.props.brand}</a>
                    </Navbar.Brand>
                    <Navbar.Toggle />
                </Navbar.Header>
                <Navbar.Collapse>
                    <Nav pullRight>

                        <NavDropdown title="Dropdown" id="ayy">

                        </NavDropdown>
                    </Nav>
                </Navbar.Collapse>
            </Navbar>
        )
    }
}

export default connect()(AppNavbar)