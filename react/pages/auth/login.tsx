import Link from "next/link";
import Layout from "../../components/layout";
import DocsLink from "../../components/docsLink";
import "../../setting";
import React from "react";

export default function Auth() {
  return (
    <>
      <Layout>
        <DocsLink
          title="Log in"
          href="https://mdbootstrap.com/docs/react/navigation/navbar/"
        />
        <h2>
          <Link href="/">
            <a>Back to home</a>
          </Link>
        </h2>
      </Layout>
    </>
  );
}
