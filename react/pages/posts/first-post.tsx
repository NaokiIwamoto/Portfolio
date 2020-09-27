import Link from "next/link";
import { useEffect, useState } from "react";
import Layout from "../../components/layout";
import DocsLink from "../../components/docsLink";
import "../../setting";
import axios from "axios";
import React from "react";
import { MDBInput } from "mdbreact";

interface Article {
  id?: number;
  title: string;
  article: string;
}

const defaultArticle: Article = {
  title: "",
  article: "",
};

export default function FirstPost() {
  const [data, setData] = useState(defaultArticle);
  useEffect(() => {
    const fetchData = async () => {
      axios
        .get("/ping")
        .then((res) => {
          console.log("成功");
        })
        .catch((err) => {
          console.log(err);
        });
    };

    fetchData();
  }, []);

  const handleTitleChange = (event: React.FormEvent<HTMLInputElement>) => {
    setData({ title: event.currentTarget.value, article: data.article });
    console.info(data);
  };

  const handleArticleChange = (event: React.FormEvent<HTMLInputElement>) => {
    setData({ title: data.title, article: event.currentTarget.value });
    console.info(data);
  };

  return (
    <>
      <Layout>
        <DocsLink
          title="New Article"
          href="https://mdbootstrap.com/docs/react/navigation/navbar/"
        />
        <MDBInput label="Title" onChange={handleTitleChange} />

        <MDBInput
          type="textarea"
          label="Body of Article"
          rows="14"
          onChange={handleArticleChange}
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
