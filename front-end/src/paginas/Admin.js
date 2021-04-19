import React, { Component } from 'react'
import FullCalendar from '@fullcalendar/react'
import dayGridPlugin from '@fullcalendar/daygrid'
import '../css/Admin.css'

export default class Admin extends Component {

    handleDateClick = (arg) => { // bind with an arrow function
        alert(arg.dateStr)
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
                
               

                <div id="Calendario">
                    <button type="button" class="btn btn-outline-secondary">Evento</button>
                    <FullCalendar
                        plugins={[ dayGridPlugin ]}
                        initialView="dayGridMonth"
                        events={[
                            { title: 'event 1', date: '2021-04-01' },
                            { title: 'event 2', date: '2021-04-02' }
                          ]}
                    />
                </div>
            </div>
        )
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