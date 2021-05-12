import React, { Component } from 'react'
import FullCalendar from '@fullcalendar/react'
import dayGridPlugin from '@fullcalendar/daygrid'
import interactionPlugin from "@fullcalendar/interaction";
import listPlugin from '@fullcalendar/list';
import time from '@fullcalendar/timegrid';
import moment from 'moment';
import '../css/Admin.css';
import axios from 'axios';
import {Modal, ModalHeader, ModalBody, ModalFooter, FormGroup, Input, Label, Button, FormText } from 'reactstrap'

export default class Admin extends Component {
    

    async componentDidMount() {
        await axios
        .get("http://localhost:3003/Eventos")
        .then(response => {
        response.data.forEach(element => {
          //this.fixEvent(element.id, element.local, element.visita, element.m_local, element.m_visita, element.fecha_inicio, element.fecha_final);
          this.state2.Events.push(element)
          
          
        });
        //console.log(this.state2.Events)
        this.setState({
          calendarEvents: this.state2.Events
        })
      });
    }

    state2 = {
        Events: []
      }

    state = {
        Temporada:{
            nombre_Temp:'',
            fecha_iTemp:'',
            fecha_fTemp:''
        },
        Jornada:{
            nombre_Jor:'',
            fecha_iJor:'',
            fecha_fJor:''
        },
        abierto: false,
        abierto2: false,
        calendarWeekends: true,
        calendarEvents: [ ]
      };
    
    abrirJornada = () =>{
        this.setState({abierto: !this.state.abierto});
    }
    
    abrirTemporada = () =>{
        this.setState({abierto2: !this.state.abierto2});
    }

    handleChange = e => {
         this.setState({
            Temporada:{
                ...this.state.Temporada,
                [e.target.name]: e.target.value
            },
            Jornada:{
                ...this.state.Jornada,
                [e.target.name]: e.target.value
            }
        })

        console.log(this.state.Jornada);
        console.log(this.state.Temporada);
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

      Jornada = () =>{
        window.location.href ="/Jornada"
      }

    render() {

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
                            <a className="nav-link" onClick={this.Carga_Masiva} >Carga Masiva <span className="sr-only">(current)</span></a>
                        </li>
                        <li className="nav-item active">
                            <a className="nav-link" onClick={this.Jornada}>Jornada </a>
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
                    <br/>
                <button type="button" className="btn btn-primary btn-lg" onClick = {this.abrirTemporada}>Crear Temporada</button>
                <button type="button" className="btn btn-primary btn-lg" onClick ={this.abrirJornada} >Crear Jornada</button> 
                    <br/>
                </div>
                <div>
                <br/>
                <button type="button" className="btn btn-primary btn-lg">Terminar Temporada</button>
                    <br/>
                </div>
                <div id="Calendario">
                    <br/> <br/>
                    <FullCalendar
                        plugins={[ dayGridPlugin, listPlugin, time, interactionPlugin]}
                        initialView="dayGridMonth"
                        headerToolbar={{
                            left: 'dayGridMonth,timeGridWeek,listYear',
                            center: 'title,today',
                            right: 'prevYear,prev,next,nextYear'
                        }}
                        editable={true}
                        selectable={true}
                        selectMirror={true}
                        dayMaxEvents={true}
                        locale = {'es'}
                        //initialEvents = {this.Ingresar_eventos()}
                        events={this.state.calendarEvents}
                        select={this.handleDateSelect}
                        eventContent={renderEventContent} // custom render function
                        eventClick={this.handleEventClick}
                    />
                </div>
                
                <Modal isOpen={this.state.abierto}>
                    <ModalHeader>
                        Crear Jornada
                    </ModalHeader>

                    <ModalBody>
                        <FormGroup>
                            <Label for="nombre-jornada">Nombre</Label>
                            <Input type="text" id="nombre-jornada" name="nombre_Jor" onChange={this.handleChange}></Input>
                        </FormGroup>
                        <FormGroup>
                            <Label for="fecha_iJ">Fecha Inicio</Label>
                            <Input type="datetime-local" id="fecha_iJ" name="fecha_iJor" onChange={this.handleChange}></Input>
                        </FormGroup>
                        <FormGroup>
                            <Label for="fecha_fJ">Fecha Fin</Label>
                            <Input type="datetime-local" id="fecha_fJ" name="fecha_fJor" onChange={this.handleChange}></Input>
                        </FormGroup>
                    </ModalBody>

                    <ModalFooter>
                        <Button onClick={this.Agregar_Jornada}>Crear</Button>
                        <Button onClick={this.abrirJornada}>Cancelar</Button>
                    </ModalFooter>
                </Modal>

                <Modal isOpen={this.state.abierto2}>
                    <ModalHeader>
                        Crear Temporada
                    </ModalHeader>
                    
                    <ModalBody>
                        <FormGroup>
                            <Label for="nombre-temporada">Nombre</Label>
                            <Input type="text" id="nombre-temporada" name="nombre_Temp" onChange={this.handleChange}></Input>
                        </FormGroup>
                        <FormGroup>
                            <Label for="fecha_iT">Fecha Inicio</Label>
                            <Input type="datetime-local" id="fecha_iT" name="fecha_iTemp" onChange={this.handleChange}></Input>
                        </FormGroup>
                        <FormGroup>
                            <Label for="fecha_fT">Fecha Fin</Label>
                            <Input type="datetime-local" id="fecha_fT" name="fecha_fTemp" onChange={this.handleChange}></Input>
                        </FormGroup>
                    </ModalBody>

                    <ModalFooter>
                        <Button onClick={this.Agregar_Temporada}>Crear</Button>
                        <Button onClick={this.abrirTemporada}>Cancelar</Button>
                    </ModalFooter>
                </Modal>

            </div>

        )
    }

    handleDateSelect = (selectInfo) => {
        axios.get("http://localhost:3003/ConsultarJornada").then(respuesta => {
            if (respuesta.data.fase == "Activa") {
                let title = prompt('Ingrese Equipo Local')
                let title2 = prompt('Ingrese Equipo Visitante')
                let deporte2 = prompt('Ingrese el Deporte')
                let calendarApi = selectInfo.view.calendar
                var fecha_E = selectInfo.startStr
                var fecha_S = ""
                for (let i = 0; i < 16; i++) {
                    fecha_S += fecha_E.charAt(i)
                    console.log(fecha_S)           
                }
                var Evento = {
                    nombreL:title,
                    nombreV:title2,
                    fecha:fecha_S,
                    idJornada:respuesta.data.id,
                    deporte:deporte2
                }

                console.log(Evento)
                title += " vs " + title2;

                axios.post("http://localhost:3003/AgregarEvento", JSON.stringify(Evento)).then(
                    result => {
                        console.log("Se envio la informacion");
                        //console.log(dato)
                    }
                ).catch(console.log)
            
                calendarApi.unselect() // clear date selection
            
                if (title) {
                calendarApi.addEvent({
                    title,
                    start: selectInfo.startStr,
                })
                }
            } else {
                axios.get("http://localhost:3003/ConsultarTemporada").then(res =>{
                    if (res.data.fase == "Activa") {
                        alert("No hay Jornada activa, Por favor crear una nueva")
                    }else{
                        alert("No hay temporada Activa, Por vafor cree una Temporada y Jornada que esten activas")
                    }
                })
            }
        }
        )
        
      }

      handleEventClick = (clickInfo) => {
        clickInfo.event.id
        let local2 = prompt("Local")
        let Visitante = prompt("Visitante")
        var Resultado = {
            id:clickInfo.event.id,
            local:local2,
            visitante:Visitante
        }
        console.log(Resultado)
        axios.post("http://localhost:3003/AgregarResultado", JSON.stringify(Resultado)).then(res =>{
            alert(res.data)
        }).catch()

      }
/*
      handleEvents = (events) => {
        this.setState({
            calendarEvents: events
        })
      }*/

      Agregar_Jornada = () =>{
        axios.get("http://localhost:3003/ConsultarJornada").then(evento =>{
            if (evento.data.fase != "Activa") {
                axios.get("http://localhost:3003/ConsultarTemporada").then(evento2 =>{
                    if (evento2.data.fase == "Activa") {
                        var jornada ={
                            nombre:this.state.Jornada.nombre_Jor,
                            fecha_i:this.state.Jornada.fecha_iJor,
                            fecha_f:this.state.Jornada.fecha_fJor,
                            temporada:evento2.data.id,
                            fase:"Activa"
                        }
                        console.log(jornada)
                        axios.post("http://localhost:3003/AgregarJornada", JSON.stringify(jornada)).then(event =>{
                            alert(event.data)
                        }

                        ).catch();

                    } else {
                        alert("No hay temporada Activa")
                    }
                })

                
            }else{
                alert("Tiene que terminar la Jornada Actual para crear otra")
            }
        })
      }


      Agregar_Temporada = () =>{
        axios.get("http://localhost:3003/ConsultarTemporada").then(evento =>{
            if (evento.data.fase != "Activa") {
                var temporada = {
                    nombre:this.state.Temporada.nombre_Temp,
                    fecha:this.state.Temporada.fecha_iTemp,
                    fechaf:this.state.Temporada.fecha_fTemp,
                    fase:"Activa"
                }

                console.log(temporada)

                axios.post("http://localhost:3003/AgregarTemporada", JSON.stringify(temporada)).then(event =>{
                            alert(event.data)
                        }

                        ).catch();
                
            }else{
                alert("Tiene que terminar la Temporada Actual para crear otra")
            }
        })
      }
      
}

function renderEventContent(eventInfo) {
    return (
      <>
        <b>{eventInfo.timeText}</b>
        <i>{eventInfo.event.title}</i>
      </>
    )
  }

/*
function Agregar_Jornada() {
    axios.get("http://localhost:3003/ConsultarJornada").then(evento =>{
        if (evento.data.fecha != "Activa") {
            jornada ={
                nombre:this.state
            }
        }else{
            alert("Tiene que terminar la Jornada Actual para crear otra")
        }
    })
}


function Agregar_Temporada() {
    axios.get("http://localhost:3003/ConsultarTemorada").then(evento =>{
        if (evento.data.fecha != "Activa") {
            let nombre = prompt('Ingrese un nombre')
            
        }else{
            alert("Tiene que terminar la Temporada Actual para crear otra")
        }
    })
}*/