import React, { Component } from 'react'
import axios from 'axios'

export default class Jornada extends Component {


    state = {
        Jornada:{
            jornada:''
        },
        jornadas:[],
        fases:[],
        Estados:{
            estado:''
        }
    }

    componentDidMount() {
        axios.get("http://localhost:3003/Jornadas")
        .then((response) =>{
            this.setState({
                jornadas:response.data.jornada
            })
        } )
        .catch((error) => {console.log(error)})

        axios.get("http://localhost:3003/Fase")
        .then((response) =>{
            this.setState({
                fases:response.data.estado
            })
        } )
        .catch((error) => {console.log(error)})
    }

    Moddificar = ()=>{
        var nuevo = {
            idJ:this.state.Jornada.jornada,
            idF:this.state.Estados.estado
        }
        axios.post("http://localhost:3003/ModificarFase", JSON.stringify(nuevo))
        .then(res =>{
            alert("Modificado")
        }).catch(err => {console.log(err)})
    }

    handleChange = e => {
        this.setState({
            Jornada:{
               ...this.state.Jornada,
               [e.target.name]: e.target.value
           },
           Estados:{
            ...this.state.Estados,
            [e.target.name]: e.target.value
        }
       })
       console.log(this.state.Estados.estado)
       console.log(this.state.Jornada.jornada)
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
      Deporte = ()=>{
        window.location.href ="/Deporte"
      }

      Admin = () =>{
        window.location.href ="/Admin"
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
                            <a className="nav-link" href="#">Jornada </a>
                        </li>
                        <li className="nav-item active">
                            <a className="nav-link" onClick={this.Temporada}>Temporada </a>
                        </li>
                        <li className="nav-item active">
                            <a className="nav-link " href="#">Recompensas </a>
                        </li>
                        <li className="nav-item active">
                            <a className="nav-link " onClick={this.Deporte} >Deportes </a>
                        </li>
                        <li className="nav-item active">
                            <a className="nav-link " onClick={this.salir}>Salir </a>
                        </li>
                        </ul>
                    </div>
                    </nav>
                </div>
                <div>
                    <select name="jornada" className="form-control" onChange={this.handleChange}>
                        {this.state.jornadas.map(elemento =>(
                            <option key={elemento.id} value={elemento.id} name="jornada" >{elemento.Nombre +" - " +elemento.estado}</option>
                        ))

                        }
                    </select>
                    <br/>
                    <select name="estado" className="form-control" onChange={this.handleChange}>
                        {this.state.fases.map(elemento =>(
                            <option key={elemento.id} value={elemento.id} name="estado" >{elemento.nombre}</option>
                        ))

                        }
                    </select>
                </div>
                <div>
                    <button type="button" className="btn btn-primary btn-lg" onClick={this.Moddificar}>Actualizar</button>
                </div>
            </div>
        )
    }
}
