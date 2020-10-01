import React, { useState, useEffect } from "react";
import { Button, Form } from "react-bootstrap";
import axios from "axios";

interface loginUserInfo {
  mailAddress: string;
  password: string;
}

const defaultLoginUserInfo: loginUserInfo = {
  mailAddress: "",
  password: "",
};

export default function Login() {
  const [data, setData] = useState(defaultLoginUserInfo);
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

  function handleMailAddressChange(event: React.ChangeEvent<HTMLInputElement>) {
    setData({
      mailAddress: event.currentTarget.value,
      password: data.password,
    });
  }

  function handlePasswordChange(event: React.ChangeEvent<HTMLInputElement>) {
    setData({
      mailAddress: data.mailAddress,
      password: event.currentTarget.value,
    });
  }

  function submitLoginUserInfo() {
    axios
      .post("/auth/login", {})
      .then((res) => {
        console.info(res);
      })
      .catch((err) => {
        console.log(err);
      });
  }

  return (
    <>
      <Form>
        <Form.Group controlId="formBasicEmail">
          <Form.Label>Email address</Form.Label>
          <Form.Control
            type="email"
            placeholder="Enter email"
            onChange={handleMailAddressChange}
          />
          <Form.Text className="text-muted">
            We'll never share your email with anyone else.
          </Form.Text>
        </Form.Group>

        <Form.Group controlId="formBasicPassword">
          <Form.Label>Password</Form.Label>
          <Form.Control
            type="password"
            placeholder="Password"
            onChange={handlePasswordChange}
          />
        </Form.Group>
        <Form.Group controlId="formBasicCheckbox">
          <Form.Check type="checkbox" label="Check me out" />
        </Form.Group>
        <Button variant="primary" type="submit" onChange={submitLoginUserInfo}>
          Submit
        </Button>
      </Form>
    </>
  );
}
