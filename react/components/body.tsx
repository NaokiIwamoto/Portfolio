import React from "react";
import { MDBContainer } from "mdbreact";

export default function Body({
  children,
}: {
  children: React.ReactNode;
  home?: boolean;
}) {
  return (
    <>
      <MDBContainer>
        <main>{children}</main>
      </MDBContainer>
    </>
  );
}
