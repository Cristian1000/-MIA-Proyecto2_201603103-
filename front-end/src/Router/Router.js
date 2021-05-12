import React from 'react'
import {BrowserRouter, Switch, Route} from 'react-router-dom'
import Login from '../paginas/login.js'
import Admin from '../paginas/Admin.js'
import CargaMasiva from '../paginas/Carga.js'
import Usuario from '../paginas/Usuario.js'
import Temporada from '../paginas/Temporada.js'
import Registro from '../paginas/Registro.js'
import Deporte from '../paginas/Deporte.js'
import Jornada from '../paginas/Jornada'

function Router() {
  return(
    <BrowserRouter>
      <Switch>
        <Route exact path="/" component={Login} />
        <Route exact path="/Admin" component={Admin} />
        <Route exact path="/CargaMasiva" component={CargaMasiva} />
        <Route exact path="/Usuario" component={Usuario} />
        <Route exact path="/Temporada" component={Temporada} />
        <Route exact path="/Registro" component={Registro} />
        <Route exact path="/Deporte" component={Deporte} />
        <Route exact path="/Jornada" component={Jornada} />
      </Switch>
    </BrowserRouter>
  );
}

export default Router;
