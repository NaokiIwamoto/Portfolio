import React, { useState } from "react";
import {
  MDBCollapse,
  MDBContainer,
  MDBDropdown,
  MDBDropdownItem,
  MDBDropdownMenu,
  MDBDropdownToggle,
  MDBIcon,
  MDBNavbar,
  MDBNavbarBrand,
  MDBNavbarNav,
  MDBNavbarToggler,
  MDBNavItem,
  MDBNavLink,
} from "mdbreact";
import Link from "next/link";

interface navbarStruct {
  collapseID: string;
}

const initialNavbarState: navbarStruct = {
  collapseID: "",
};

export default function NavbarCommon() {
  const [data, setData] = useState(initialNavbarState);
  function toggleCollapse(collapseID: string): void {
    setData({ collapseID: data.collapseID == collapseID ? "" : collapseID });
  }
  return (
    <>
      <MDBNavbar color="unique-color-dark" dark expand="md">
        <MDBNavbarBrand>
          <strong className="white-text">Port</strong>
        </MDBNavbarBrand>
        <MDBNavbarToggler onClick={() => toggleCollapse("navbarCollapse3")} />
        <MDBCollapse id="navbarCollapse3" isOpen={data.collapseID} navbar>
          <MDBNavbarNav left>
            <MDBNavItem>
              <Link href={"#!"}>
                <a className="nav-link text-white">Home</a>
              </Link>
            </MDBNavItem>
            <MDBNavItem>
              <Link href={"#!"}>
                <a className="nav-link text-white">Features</a>
              </Link>
            </MDBNavItem>
            <MDBNavItem>
              <Link href={"#!"}>
                <a className="nav-link text-white">Features</a>
              </Link>
            </MDBNavItem>
            <MDBNavItem>
              <MDBDropdown>
                <MDBDropdownToggle nav caret>
                  <div className="d-none d-md-inline">MDBDropdown</div>
                </MDBDropdownToggle>
                <MDBDropdownMenu className="dropdown-default" right>
                  <MDBDropdownItem href="#!">Action</MDBDropdownItem>
                  <MDBDropdownItem href="#!">Another Action</MDBDropdownItem>
                  <MDBDropdownItem href="#!">
                    Something else here
                  </MDBDropdownItem>
                  <MDBDropdownItem href="#!">
                    Something else here
                  </MDBDropdownItem>
                </MDBDropdownMenu>
              </MDBDropdown>
            </MDBNavItem>
          </MDBNavbarNav>
          <MDBNavbarNav right>
            {/* <MDBNavItem>
                <MDBNavLink
                  className="waves-effect waves-light d-flex align-items-center"
                  to="#!"
                >
                  1
                  <MDBIcon icon="envelope" className="ml-1" />
                </MDBNavLink>
              </MDBNavItem> */}
            <MDBNavItem>
              <MDBDropdown>
                <MDBDropdownToggle className="dopdown-toggle" nav>
                  <img
                    src="https://mdbootstrap.com/img/Photos/Avatars/avatar-2.jpg"
                    className="rounded-circle z-depth-0"
                    style={{ height: "35px", padding: 0 }}
                    alt=""
                  />
                </MDBDropdownToggle>
                <MDBDropdownMenu className="dropdown-default" right>
                  <MDBDropdownItem>
                    <Link href={"/auth/login"}>Log in</Link>
                  </MDBDropdownItem>
                </MDBDropdownMenu>
              </MDBDropdown>
            </MDBNavItem>
          </MDBNavbarNav>
        </MDBCollapse>
      </MDBNavbar>
    </>
  );
}
