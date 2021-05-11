import React, { Component } from 'react'
import FullCalendar from '@fullcalendar/react'
import dayGridPlugin from '@fullcalendar/daygrid'
import interactionPlugin from "@fullcalendar/interaction";
import listPlugin from '@fullcalendar/list';
import time from '@fullcalendar/timegrid';
import Cookies from 'universal-cookie';
import axios from 'axios';
const cookies = new Cookies();

export default class Usuario extends Component {

    ShowSelected()
        {
        /* Para obtener el valor */
        var cod = document.getElementById("OS").value;
        alert(cod);
        
        /* Para obtener el texto */
        var combo = document.getElementById("producto");
        var selected = combo.options[combo.selectedIndex].text;
        alert(selected);
    }

    async componentDidMount() {
        if(!cookies.get('id')){
            window.location.href="./";
        }
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

    state ={
        calendarWeekends: true,
        calendarEvents: []
    }


    cerrarSesion=()=>{
        cookies.remove('id', {path: "/"});
        window.location.href='./';
    }

    render() {
        console.log('id: '+ cookies.get('id'));
        return (
            
            <div>
                <div id="barra">
                <nav className="navbar navbar-expand-lg navbar-light bg-light">
                    <a className="navbar-brand" href="#">Usuario</a>
                    <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
                        <span className="navbar-toggler-icon"></span>
                    </button>
                    <div className="collapse navbar-collapse" id="navbarNav">
                        <ul className="navbar-nav">
                        <li className="nav-item active">
                            <a className="nav-link" href="#">Comprar Tier <span className="sr-only">(current)</span></a>
                        </li>
                        <li className="nav-item active">
                            <a className="nav-link" href="#" onClick={()=>this.cerrarSesion()} >Cerrar Sesion </a>
                        </li>
                        </ul>
                    </div>
                    </nav>
                </div>

                <div>
                    <select name="OS" onClick={this.ShowSelected}>
                        <option value="1">Windows Vista</option> 
                        <option value="2">Windows 7</option> 
                        <option value="3">Windows XP</option>
                        <option value="10">Fedora</option> 
                        <option value="11">Debian</option> 
                        <option value="12">Suse</option> 
                    </select>
                </div>

                <div>
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
                        events={this.state.calendarEvents}
                        eventContent={renderEventContent}
                        eventClick={this.handleEventClick}
                    />
                </div>

            </div>
        )
    }

    handleEventClick = (clickInfo) => {
        clickInfo.event.id
        console.log(cookies.get('id'))
        let local2 = prompt("Local")
        let Visitante = prompt("Visitante")
        var Resultado = {
            id_evento:clickInfo.event.id,
            local:local2,
            visitante:Visitante,
            id_cliente:cookies.get('id')
        }
        console.log(Resultado)
        axios.post("http://localhost:3003/IngresarPrediccion", JSON.stringify(Resultado)).then(res =>{
            alert(res.data)
        }).catch()

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