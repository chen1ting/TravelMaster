import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { validateToken } from "../api/api";

const Welcome = () => {
  const navigate = useNavigate();
  useEffect(() => {
    const ok = validateToken(window.sessionStorage.getItem("session_token"));
    if (!ok) {
      console.log("not ok");
      navigate("/");
    }
  }, [navigate]);
  const user = window.sessionStorage.getItem("username");

  return <div>Hello {user}! :)</div>;
};

export default Welcome;
