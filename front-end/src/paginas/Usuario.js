import React, { Component, useState } from 'react'
import FullCalendar from '@fullcalendar/react'
import dayGridPlugin from '@fullcalendar/daygrid'
import interactionPlugin from "@fullcalendar/interaction";
import listPlugin from '@fullcalendar/list';
import time from '@fullcalendar/timegrid';
import Cookies from 'universal-cookie';
import axios from 'axios';
import {Modal, ModalHeader, ModalBody, ModalFooter, FormGroup, Input, Label, Button, ButtonDropdown, DropdownToggle, DropdownMenu, DropdownItem } from 'reactstrap'


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

      axios.get("http://localhost:3003/Membresia")
        .then((response) =>{
            this.setState({
                membresias:response.data.datos
            })
        } )
        .catch((error) => {console.log(error)})
    }

    state2 = {
        Events: []
      }

    state ={
        calendarWeekends: true,
        calendarEvents: [],
        membresias:[],
        Membre:{
            membresia:''
        }
    }

    handleChange = e => {
        this.setState({
            Membre:{
               ...this.state.Membre,
               [e.target.name]: e.target.value
           }
       })
  
       console.log(this.state.Membre);
   }

    cerrarSesion=()=>{
        cookies.remove('id', {path: "/"});
        window.location.href='./';
    }

    render() {

        return (
            
            <div>
                <div id="barra">
                <nav className="navbar navbar-expand-lg navbar-light bg-light">
                    <a className="navbar-brand" href="#">{cookies.get('id')}</a>
                    <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
                        <span className="navbar-toggler-icon"></span>
                    </button>
                    <div className="collapse navbar-collapse" id="navbarNav">
                        <ul className="navbar-nav">
                        <li className="nav-item active">
                            <a className="nav-link" href="#" onClick ={this.Compra}>Comprar Tier <span className="sr-only">(current)</span></a>
                        </li>
                        <li className="nav-item active">
                            <a className="nav-link" href="#" onClick={()=>this.cerrarSesion()} >Cerrar Sesion </a>
                        </li>
                        </ul>
                    </div>
                    </nav>
                </div>

                <div>
                <select name="membresia" className="form-control" onChange={this.handleChange}>
                        {this.state.membresias.map(elemento =>(
                            <option key={elemento.id} value={elemento.id} name="membresia" >{elemento.nombre}</option>
                        ))

                        }
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
        axios.get("http://localhost:3003/ConsultarTemporada").then(evento =>{
            if (evento.data.fase == "Activa") {

                var actual ={
                    idC:cookies.get('id'),
                    idT:evento.data.id
                }

                axios.post("http://localhost:3003/MembresiaActual", JSON.stringify(actual))
                .then(response=>{
                    if(response.data.id != ""){
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
                    }else{
                        alert('No tiene Membresia Vigente');
                    }
                })
                .catch(error=>{
                    console.log(error);
                })
                
            }else{
                alert("No hay temporada Activa")
            }
        })
        

      }

      Compra = () =>{
        axios.get("http://localhost:3003/ConsultarTemporada").then(evento =>{
            if (evento.data.fase == "Activa") {
                var temporada = {
                    idC:cookies.get('id'),
                    idM:this.state.Membre.membresia,
                    idT:evento.data.id
                }

                console.log(temporada)

                axios.post("http://localhost:3003/CompraMembresia", JSON.stringify(temporada)).then(event =>{
                            alert(event.data)
                        }

                        ).catch();
                
            }else{
                alert("No hay temporada Activa")
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