import React, { Component } from 'react'
import FullCalendar from '@fullcalendar/react'
import dayGridPlugin from '@fullcalendar/daygrid'
import interactionPlugin from "@fullcalendar/interaction";
import listPlugin from '@fullcalendar/list';
import time from '@fullcalendar/timegrid';
import moment from 'moment';
import '../css/Admin.css';
import axios from 'axios';

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
        calendarWeekends: true,
        calendarEvents: [
            {title: "Research and Development vs Business Development", start: "2018/04/26 09:32"},
            { title: 'event 2', date: '2021-04-02' }
        ]
      };
  

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
                <div>
                <button type="button" class="btn btn-primary btn-lg">Crear Temporada</button>
                <button type="button" class="btn btn-primary btn-lg">Crear Jornada</button> 
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
                
            </div>
        )
    }

    handleDateSelect = (selectInfo) => {
        axios.get("http://localhost:3003/ConsultarJornada").then(respuesta => {
            if (respuesta.data.fecha == "Activa") {
                let title = prompt('Ingrese Equipo Local')
                let title2 = prompt('Ingrese Equipo Visitante')
                let deporte2 = prompt('Ingrese el Deporte')
                let calendarApi = selectInfo.view.calendar

                var Evento = {
                    nombreL:title,
                    nombreV:title2,
                    fecha:selectInfo.startStr,
                    idJornada:respuesta.data.id,
                    deporte:deporte2
                }

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
                    if (res.data.fecha == "Activa") {
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
        if (confirm(`Are you sure you want to delete the event '${clickInfo.event.title}'`)) {
          clickInfo.event.remove()
        }
      }
/*
      handleEvents = (events) => {
        this.setState({
            calendarEvents: events
        })
      }*/
      
}

function renderEventContent(eventInfo) {
    return (
      <>
        <b>{eventInfo.timeText}</b>
        <i>{eventInfo.event.title}</i>
      </>
    )
  }


