import React from 'react'
import {BrowserRouter, Switch, Route} from 'react-router-dom'
import Login from '../paginas/login.js'
import Admin from '../paginas/Admin.js'

function Router() {
  return(
    <BrowserRouter>
      <switch>
        <Route exact path="/" component={Login} />
        <Route exact path="/Admin" component={Admin} />
      </switch>
    </BrowserRouter>
  );
}

export default Router;
