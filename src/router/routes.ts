import React, { lazy } from "react";
const Home = lazy(() => import("../view/Home"));
export default [
  {
    path: "/download",
    component: Home,
    exact: true
  }
];
