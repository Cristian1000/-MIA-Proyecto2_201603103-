import React, {Component} from 'react'
import '../css/Login.css'
import 'bootstrap/dist/css/bootstrap.min.css'
import axios from 'axios'


class login extends Component{

    state = {
        from:{
            usuario:"",
            contrasena:""
        }
    }

    handleChange = async e => {
        await this.setState({
            from:{
                ...this.state.form,
                [e.target.name]: e.target.value
            }
        })

        console.log(this.state.from);
    }

    render(){
        return(
            <div className="containerPrincipal">
            <div className="containerSecundario">
            <div className="form-group">
            <label>Usuario: </label>
            <br />
            <input
              type="text"
              className="form-control"
              name="usuario"
              onChange={this.handleChange}
            />
            <br />
            <label>Contraseña: </label>
            <br />
            <input
              type="password"
              className="form-control"
              name="contrasena"
              onChange={this.handleChange}
            />
            <br />
            <button className="btn btn-primary" onClick={()=> this.iniciarSesion()}>Iniciar Sesión</button>
          </div>
        </div>
      </div>
        );
    }
}

export default login