import {Navbar, Nav, NavDropdown} from "react-bootstrap";


const AppNavbar = ({brand}) => (
    <Navbar>
        <Navbar.Header>
            <Navbar.Brand>
                <a href="/">Goquest</a>
            </Navbar.Brand>
            <Navbar.Toogle />
        </Navbar.Header>
        <Navbar.Collapse>
            <Nav>
                <NavDropdown title="">

                </NavDropdown>
            </Nav>
        </Navbar.Collapse>
    </Navbar>
);

export default AppNavbar;