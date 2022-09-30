import { validSessionGuard } from "../common/common";
import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";

const Welcome = () => {
  const navigate = useNavigate();
  useEffect(() => {
    validSessionGuard(navigate);
  }, [navigate]);
  const user = window.sessionStorage.getItem("username");

  return <div>Hello {user}! :)</div>;
};

export default Welcome;
