import React, { lazy } from "react";
const Home = lazy(() => import("../view/home"));
const Manage = lazy(() => import('../view/manage'))
export default [
  {
    path: '/appdl',
    component: Home,
    exact: true
  },
  {
    path: '/app/manage',
    component: Manage,
    exact: true
  }
]
