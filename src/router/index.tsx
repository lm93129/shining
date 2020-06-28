import React from "react";
import { BrowserRouter as Router, Route, Switch, Redirect } from "react-router-dom";
import frontRoutes from './frontRoutes'
import consoleRoutes from './consoleRoutes'

const routes = [...frontRoutes, ...consoleRoutes]
function RouteWithSubRoutes(route: any) {
  return (
    <Route
      path={route.path}
      render={props => (
        <route.component {...props} routes={route.routes} />
      )}
    />
  );
}
function RouteConfig() {
  return (
    <Switch>
      <Redirect from="/" to="/appdl" exact />
      {routes.map((route, i) => (
        <RouteWithSubRoutes key={i} {...route} />
      ))}
    </Switch>
  )
}

export default RouteConfig;