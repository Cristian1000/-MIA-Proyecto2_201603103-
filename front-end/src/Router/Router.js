import React from 'react'
import {BrowserRouter, Switch, Route} from 'react-router-dom'
import Login from '../paginas/login.js'
import Admin from '../paginas/Admin.js'
import CargaMasiva from '../paginas/Carga.js'
import Usuario from '../paginas/Usuario.js'
import Temporada from '../paginas/Temporada.js'

function Router() {
  return(
    <BrowserRouter>
      <Switch>
        <Route exact path="/" component={Login} />
        <Route exact path="/Admin" component={Admin} />
        <Route exact path="/CargaMasiva" component={CargaMasiva} />
        <Route exact path="/Usuario" component={Usuario} />
        <Route exact path="/Temporada" component={Temporada} />
      </Switch>
    </BrowserRouter>
  );
}

export default Router;
