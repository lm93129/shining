import React, { lazy } from "react";
const Home = lazy(() => import("../view/Home"));
export default [
  {
    path: "/appdl",
    component: Home,
    exact: true
  }
];
