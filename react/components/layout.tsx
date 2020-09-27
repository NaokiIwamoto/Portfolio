import React from "react";
import { MDBContainer } from "mdbreact";
import FooterPage from "./footer";
import Body from "./body";
import NavbarCommon from "./navbar";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <>
      <NavbarCommon></NavbarCommon>
      <Body>{children}</Body>
      <FooterPage></FooterPage>
    </>
  );
}
