import React, { useState } from "react";
import {
  Navbar,
  NavbarBrand,
  NavbarToggler,
  Collapse,
  Nav,
  NavItem,
  NavbarText,
} from "reactstrap";

function Navigation({ author, children }) {
  const [isOpen, setIsOpen] = useState(false);
  const toggle = () => setIsOpen(!isOpen);
  return (
    <Navbar color="dark" light={false} dark>
      <NavbarBrand href="/">OKR Generator</NavbarBrand>
      <NavbarToggler onClick={toggle} />
      <Collapse isOpen={isOpen} navbar>
        <NavbarText>{author}</NavbarText>
        <Nav navbar>
          <NavItem>{children}</NavItem>
        </Nav>
      </Collapse>
    </Navbar>
  );
}

export default Navigation;
