import React from "react";
import Company from "./companies/Company";
import { useAppSelector } from "./hooks";
import { Navigate } from "react-router-dom";

function App() {
  const isLoggedIn = useAppSelector((state) => state.login.isLoggedIn);
  if (!isLoggedIn) {
    return <Navigate replace to="/login" />;
  }

  return <Company />;
}

export default App;
