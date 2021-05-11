import React, { Component } from 'react'
import 'bootstrap/dist/css/bootstrap.min.css'
//import {Dropdown, DropdownItem, DropdownMenu, DropdownToggle} from 'reactstrap'
import axios from 'axios'
import DataTable from 'react-data-table-component'


export default class Temporada extends Component {

    actualizar = () =>{
        var resultado = {
            id_tempo:this.state.Solicitud.temporada,
            id_cliente:this.state.Solicitud.usuario
        }
        console.log(resultado)

        axios.post("http://localhost:3003/TemporadaUsuario", JSON.stringify(resultado))
        .then(response=>{
            this.setState({
                tabla:response.data.datos
            })
        })
        .catch(error=>{
            console.log(error);
        })

    }
    

    state = {
        temoradas:[],
        usuarios:[],
        Solicitud:{
            temporada: '',
            usuario: ''
        },
        tabla:[]
    }

    handleChange = e => {
        this.setState({
            Solicitud:{
               ...this.state.Solicitud,
               [e.target.name]: e.target.value
           }
       })
       console.log(this.state.Solicitud)
   }

    componentDidMount() {
        axios.get("http://localhost:3003/Temporadas")
        .then((response) =>{
            this.setState({
                temoradas:response.data.temoradas
            })
        } )
        .catch((error) => {console.log(error)})

        axios.get("http://localhost:3003/Usuarios")
        .then((response) =>{
            this.setState({
                usuarios:response.data.usuario
            })
        } )
        .catch((error) => {console.log(error)})
    }

    render() {
        const columnas = [
            {
                name: 'Deporte',
                selector:'deporte',
                sortable: true,
            },
            {
                name: 'Local',
                selector:'local',
                sortable: true,
            },
            {
                name: 'Visitante',
                selector:'visitante',
                sortable: true,
            },
            {
                name: 'Prediccion',
                selector:'prediccion',
                sortable: true,
            },
            {
                name: 'Resultado',
                selector:'resultado',
                sortable: true,
            },
            {
                name: 'Puntos',
                selector:'puntos',
                sortable: true,
            },
            {
                name: 'Fecha',
                selector:'fecha',
                sortable: true,
            },
        ]

        

        return (
            <div>
                <div id="barra">
                <nav className="navbar navbar-expand-lg navbar-light bg-light">
                    <a className="navbar-brand" href="#">Quinela</a>
                    <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
                        <span className="navbar-toggler-icon"></span>
                    </button>
                    <div className="collapse navbar-collapse" id="navbarNav">
                        <ul className="navbar-nav">
                        <li className="nav-item active">
                            <a className="nav-link" href="#">Carga Masiva <span className="sr-only">(current)</span></a>
                        </li>
                        <li className="nav-item active">
                            <a className="nav-link" href="#"  >Jornada </a>
                        </li>
                        <li className="nav-item active">
                            <a className="nav-link" href="#">Temporada </a>
                        </li>
                        <li className="nav-item active">
                            <a className="nav-link " href="#">Recompensas </a>
                        </li>
                        <li className="nav-item active">
                            <a className="nav-link " href="#">Deportes </a>
                        </li>
                        <li className="nav-item active">
                            <a className="nav-link " href="https://www.google.com/">Reportes </a>
                        </li>
                        </ul>
                    </div>
                    </nav>
                </div>
                <br/>
                <div>
                    <button type="button" className="btn btn-primary btn-lg" onClick={this.actualizar}>Actualizar</button>
                </div>
                <br/>
                <div>
                    <select name="temporada" className="form-control" onChange={this.handleChange}>
                        {this.state.temoradas.map(elemento =>(
                            <option key={elemento.id} value={elemento.id} name="temporada" >{elemento.nombre}</option>
                        ))

                        }
                    </select>
                </div>
                <br/>
                <div>
                    <select name="usuario" className="form-control" onChange={this.handleChange}>
                        {this.state.usuarios.map(elemento =>(
                            <option key={elemento.id} value={elemento.id} name="usuario" >{elemento.usuario}</option>
                        ))

                        }
                    </select>
                </div>
                <br/>
                <div>
                    <DataTable
                    columns = {columnas}
                    data = {this.state.tabla}
                    title = "Resultados"
                    />
                </div>
            </div>
        )
    }
}
