import React from 'react'
import {BrowserRouter, Switch, Route} from 'react-router-dom'
import Login from '../paginas/login.js'
import Admin from '../paginas/Admin.js'
import CargaMasiva from '../paginas/Carga.js'

function Router() {
  return(
    <BrowserRouter>
      <Switch>
        <Route exact path="/" component={Login} />
        <Route exact path="/Admin" component={Admin} />
        <Route exact path="/CargaMasiva" component={CargaMasiva} />
      </Switch>
    </BrowserRouter>
  );
}

export default Router;
