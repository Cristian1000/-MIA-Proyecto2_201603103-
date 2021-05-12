import React, {Component} from 'react'
import '../css/Login.css'
import 'bootstrap/dist/css/bootstrap.min.css'
import axios from 'axios'
import Cookies from 'universal-cookie';

const cookies = new Cookies();


class login extends Component{

    state = {
        from:{
            usua:'',
            contrasena:''
        }
    }

    handleChange = e => {
      this.setState({
         from:{
             ...this.state.from,
             [e.target.name]: e.target.value
         }
     })

     console.log(this.state.from);
 }

    iniciarSesion=()=>{
      if (this.state.from.contrasena == "admin" && this.state.from.usua == "admin") {
          window.location.href="./Admin";
      }else{

      var inicio ={
        usuario:this.state.from.usua,
        pass:this.state.from.contrasena
      }
      console.log(inicio)
      axios.post("http://localhost:3003/InicioSesion", JSON.stringify(inicio))
      .then(response=>{
          if(response.data.id != ""){
              console.log(response.data)
              cookies.set('id', response.data.id, {path: "/"});
              alert(`Bienvenido`);
              window.location.href="./Usuario";
          }else{
              alert('El usuario o la contraseña no son correctos');
          }
      })
      .catch(error=>{
          console.log(error);
      })
    }
  }

  regristrar = () =>{
    window.location.href ="/Registro"
  }

  componentDidMount() {
    if(cookies.get('id')){
        window.location.href="./Usuario";
    }
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
              name="usua"
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
            <button className="btn btn-primary" onClick={this.iniciarSesion}>Iniciar Sesión</button>
            <br />
            <br />
            <button className="btn btn-primary" onClick={this.regristrar}>Registro</button>
          </div>
        </div>
      </div>
        );
    }
}

export default login