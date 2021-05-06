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
                <div id="Calendario">
                    <br/> <br/>
                    <FullCalendar
                        plugins={[ dayGridPlugin, listPlugin, time]}
                        initialView="dayGridMonth"
                        headerToolbar={{
                            left: 'dayGridMonth,timeGridWeek,listYear',
                            center: 'title,today',
                            right: 'prevYear,prev,next,nextYear'
                        }}
                        locale = {'es'}
                        //initialEvents = {this.Ingresar_eventos()}
                        events={this.state.calendarEvents}
                    />
                </div>
                
            </div>
        )
    }
}


