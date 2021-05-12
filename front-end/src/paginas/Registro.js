import React, {Component} from 'react'
import '../css/Login.css'
import 'bootstrap/dist/css/bootstrap.min.css'
import axios from 'axios'
import Cookies from 'universal-cookie';

const cookies = new Cookies();


class Registro extends Component{

    state = {
        from:{
            nombre:'',
            apellido:'',
            pass:'',
            usuario:'',
            fecha_nacimiento:'',
            fecha:'',
            correo:''
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
      var inicio ={
            nombre:this.state.from.nombre,
            apellido:this.state.from.apellido,
            pass:this.state.from.pass,
            usuario:this.state.from.usuario,
            fecha:this.state.from.fecha_nacimiento,
            correo:this.state.from.correo
      }
      console.log(inicio)
      axios.post("http://localhost:3003/CrearUsuario", JSON.stringify(inicio))
      .then(response=>{
          if(response.data != ""){
              console.log(response.data)
              cookies.set('id', response.data, {path: "/"});
              alert(`Bienvenido`);
              window.location.href="./Usuario";
          }else{
              alert('El usuario ya Existe');
          }
      })
      .catch(error=>{
          console.log(error);
      })

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
            <label>Nombre: </label>
            <br />
            <input
              type="text"
              className="form-control"
              name="nombre"
              onChange={this.handleChange}
            />
            <br />
            <label>Apellido: </label>
            <br />
            <input
              type="text"
              className="form-control"
              name="apellido"
              onChange={this.handleChange}
            />
            <br />
            <label>Contraseña: </label>
            <br />
            <input
              type="password"
              className="form-control"
              name="pass"
              onChange={this.handleChange}
            />
            <br />
            <label>Usuario: </label>
            <br />
            <input
              type="text"
              className="form-control"
              name="usuario"
              onChange={this.handleChange}
            />
            <br />
            <label>Fecha Nacimiento: </label>
            <br />
            <input
              type="date"
              className="form-control"
              name="fecha_nacimiento"
              onChange={this.handleChange}
            />
            <br />
            <label>Correo: </label>
            <br />
            <input
              type="email"
              className="form-control"
              name="correo"
              onChange={this.handleChange}
            />
            <br />
            <button className="btn btn-primary" onClick={this.iniciarSesion} >Registro</button>
            <br />
            <br />
            <button className="btn btn-primary" >Iniciar Sesión</button>
          </div>
        </div>
      </div>
        );
    }

}

export default Registro
