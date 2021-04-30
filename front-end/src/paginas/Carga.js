import React, { Component } from 'react'
import {useState} from 'react'
import 'bootstrap/dist/css/bootstrap.min.css'
import axios from 'axios'

export default class Carga extends Component {

    render() {

        var ArchivoYaml = "";

        var openFile = function(evt) { 
            let status = [];
            const fileObj = evt.target.files[0];
            const reader = new FileReader(); 
            let fileloaded = e => {
                ArchivoYaml = e.target.result;
              console.log(e.target.result);
            }
            fileloaded = fileloaded.bind(this);
            reader.onload = fileloaded;
            reader.readAsText(fileObj);
        };

        var Enviar = function(){
            var url = 'http://localhost:3003/Carga';
            var dato = JSON.stringify({Texto:ArchivoYaml})
            axios.post(url, dato).then(
                result => {
                    console.log("Se envio la informacion");
                }
            ).catch(console.log)
        }

        

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
                        <li className="nav-item">
                            <a className="nav-link" href="#">Jornada </a>
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
                <div id='Carga'>
                    <br/><br/>
                    <input type='file' name='Archivo' onChange= {evt => openFile(evt)}/>
                    <br/><br/>
                    <button className="btn btn-primary" onClick={Enviar}  >Cargar Archivo</button>
                </div>

                
            </div>
        )
    }

    
}


  