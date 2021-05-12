import React, { Component } from 'react'
import axios from "axios";
import '../css/Login.css'
import 'bootstrap/dist/css/bootstrap.min.css'

export default class Deporte extends Component {

    state = {
        Deporte:{
            id:'',
        nombre:''
        },
        Deportes:[]
    }

    handleChange = e => {
        this.setState({
            Deporte:{
               ...this.state.Deporte,
               [e.target.name]: e.target.value
           },
       })
       console.log(this.state.Deporte)
   }

   componentDidMount() {
    axios.get("http://localhost:3003/Deportes")
    .then((response) =>{
        this.setState({
            Deportes:response.data.deporte
        })
    } )
    .catch((error) => {console.log(error)})
    }

    Carga_Masiva = () =>{
        window.location.href ="/CargaMasiva"
      }

    Temporada = () =>{
        window.location.href ="/Temporada"
      }

      salir = ()=>{
        window.location.href ="/"
      }

      Admin = ()=>{
        window.location.href ="/Admin"
      }

      Jornada = () =>{
        window.location.href ="/Jornada"
      }

      agregar = () =>{
        var nuevo = {
            id:this.state.Deporte.id,
            nombre:this.state.Deporte.nombre
        }
        console.log(nuevo)

        axios.post("http://localhost:3003/AgregarDeporte", JSON.stringify(nuevo))
      .then(response=>{
          alert("Deporte Agregado")
      })
      .catch(error=>{
          console.log(error);
      })
      }
      modificar = () =>{
        var nuevo = {
            id:this.state.Deporte.id,
            nombre:this.state.Deporte.nombre
        }

        console.log(nuevo)
        axios.post("http://localhost:3003/ModificarDeporte", JSON.stringify(nuevo))
        .then(response=>{
            alert("Deporte Modificado")
        })
        .catch(error=>{
            console.log(error);
        })
      }
      eliminar = () =>{
        var nuevo = {
            id:this.state.Deporte.id,
            nombre:this.state.Deporte.nombre
        }

        console.log(nuevo)
        axios.post("http://localhost:3003/EliminarDeporte", JSON.stringify(nuevo))
        .then(response=>{
            alert("Deporte Eliminado")
        })
        .catch(error=>{
            console.log(error);
        })
      }

    render() {
        return (
            <div>
                <div id="barra">
                <nav className="navbar navbar-expand-lg navbar-light bg-light">
                    <a className="navbar-brand" onClick={this.Admin}>Quinela</a>
                    <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
                        <span className="navbar-toggler-icon"></span>
                    </button>
                    <div className="collapse navbar-collapse" id="navbarNav">
                        <ul className="navbar-nav">
                        <li className="nav-item active">
                            <a className="nav-link" onClick={this.Carga_Masiva} >Carga Masiva <span className="sr-only">(current)</span></a>
                        </li>
                        <li className="nav-item active">
                            <a className="nav-link" onClick = {this.Jornada}>Jornada </a>
                        </li>
                        <li className="nav-item active">
                            <a className="nav-link" onClick={this.Temporada}>Temporada </a>
                        </li>
                        <li className="nav-item active">
                            <a className="nav-link " href="#">Recompensas </a>
                        </li>
                        <li className="nav-item active">
                            <a className="nav-link " href="#">Deportes </a>
                        </li>
                        <li className="nav-item active">
                            <a className="nav-link " onClick={this.salir}>Salir </a>
                        </li>
                        </ul>
                    </div>
                    </nav>
                </div>
                <div>
                    <select name="id" className="form-control" onChange={this.handleChange}>
                        {this.state.Deportes.map(elemento =>(
                            <option key={elemento.id} value={elemento.id} name="id" >{elemento.nombre}</option>
                        ))
                        }
                    </select>
                </div>
                <div>
                    <label>Nombre: </label>
                    <br />
                    <input
                    type="text"
                    className="form-control"
                    name="nombre"
                    onChange={this.handleChange}
                    />
                    <br />
                </div>
                <div>
                    <br/>
                    <button type="button" className="btn btn-primary btn-lg" onClick = {this.agregar}>Crear Deporte</button>
                    <button type="button" className="btn btn-primary btn-lg" onClick ={this.modificar} >Modificar Deporte</button> 
                    <button type="button" className="btn btn-primary btn-lg" onClick ={this.eliminar} >Eliminar Deporte</button> 
                    <br/>
                </div>
            </div>
        )
    }
}
