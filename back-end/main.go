package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/godror/godror"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Resultados struct {
	Temporada string    `json:"temporada"`
	Tier      string    `json:"tier"`
	Jornadas  []Jornada `Json:"jornadas"`
}

type Jornada struct {
	Jornada string   `json:"jornada"`
	Evento  []Evento `json:"predicciones"`
}

type Evento struct {
	Deporte    string     `json:"deporte"`
	Fecha      string     `json:"fecha"`
	Visitante  string     `json:"visitante"`
	Local      string     `json:"local"`
	Prediccion Prediccion `json:"prediccion"`
	Resultado  Resultado  `json:"resultado"`
}

type Prediccion struct {
	Visitante int `json:"visitante"`
	Local     int `json:"local"`
}

type Resultado struct {
	Visitante int `json:"visitante"`
	Local     int `json:"local"`
}

type Usuario struct {
	Nombre     string       `json:"nombre"`
	Apellido   string       `json:"apellido"`
	Pass       string       `json:"password"`
	User       string       `json:"username"`
	Resultado_ []Resultados `json:"resultados"`
}

type Info map[string]Usuario

type Evento2 struct {
	NombreL   string `json:"nombreL"`
	NombreV   string `json:"nombreV"`
	Fecha     string `json:"fecha"`
	IdJornada string `json:"idJornada"`
	Deporte   string `json:"deporte"`
}

type Retorno struct {
	Eventos []Dentro `json:"Eventos"`
}
type Dentro struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Start string `json:"start"`
}

type Contiene map[string]Retorno

func indexRoute(w http.ResponseWriter, r *http.Request) {
	//retornar_Evento("hola", "adios", "03/02/2019 11:29")
	//valida_Usuario("Cris10")
	//validar_Deporte("golf")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Inicio("jpook0@army.mil", "Mvjtqy"))

}

func CargaMasiva(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var Carga Info
	json.NewDecoder(r.Body).Decode(&Carga)

	for key, element := range Carga {
		var jornada string
		var cliente string
		var temporada string
		var idJornada string
		var idTemporada string
		//var resultado string
		fmt.Println(key)
		cliente = element.User
		resultado := valida_Usuario(cliente)
		if resultado != cliente {
			ingresar_Usuario(element.Nombre, element.Apellido, element.Pass, element.User)
		}
		for _, element := range element.Resultado_ {
			temporada = element.Temporada
			resultado_Temp := validar_Temporada(temporada)
			if resultado_Temp == "" {
				ingresar_Temorada(element.Temporada)
				idTemporada = id_Temporada(temporada)
			}
			resultado_Tier := validar_Tier(element.Tier)
			if resultado_Tier != element.Tier {
				Ingresar_Membresia(element.Tier)
			}
			ingrrsar_Membresia_Temp(cliente, element.Tier, element.Temporada)
			for _, element := range element.Jornadas {
				jornada = element.Jornada
				retornarJornada := validar_Jornada(jornada, temporada)
				if retornarJornada == "" {
					ingresar_Jornada(element.Jornada, "03/02/2019 11:29", "03/02/2019 11:29", temporada, "Finalizada")
					idJornada = id_jornada(jornada, temporada)
				}
				for _, element := range element.Evento {

					//var evento string
					resultado_Dep := validar_Deporte(element.Deporte)
					if resultado_Dep == "" {
						ingresar_Deporte(element.Deporte)
					}
					retornoTem := retornar_Temporada(temporada)
					validEvento := validar_Evento(element.Local, element.Visitante, element.Fecha)
					if validEvento == "" {
						ingresar_Evento(element.Local, element.Visitante, strconv.Itoa(element.Resultado.Local), strconv.Itoa(element.Resultado.Visitante), element.Fecha, element.Deporte, jornada, retornoTem)
					}
					evento := retornar_Evento(element.Local, element.Visitante, element.Fecha)
					validPrediccion := validar_Prediccion(cliente, element.Fecha)
					if validPrediccion == "" {
						ingresar_Prediccion(strconv.Itoa(element.Prediccion.Local), strconv.Itoa(element.Prediccion.Visitante), cliente, evento)
					}

					Actualizar_Jornada(idJornada)
					Actualizar_Temporada(idTemporada)
				}
			}
		}
	}

}

func Enviar_Evento(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var nuevoRetorno Retorno

	json.NewEncoder(w).Encode(Buscar_evento().Eventos)
}

func Estado_Jornada(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var nuevoRetorno Retorno

	json.NewEncoder(w).Encode(Consultar_Jornada())
}

func Estado_Temporada(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var nuevoRetorno Retorno

	json.NewEncoder(w).Encode(Consultar_Temporada())
}

func Agregar_Evento(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var nuevoRetorno Retorno
	var newEvento Evento2
	json.NewDecoder(r.Body).Decode(&newEvento)
	newEvento.Fecha = strings.Replace(newEvento.Fecha, "T", " ", -1)
	fmt.Println(newEvento)
	//fmt.Println(newEvento)
	Crear_Evento(newEvento.NombreL, newEvento.NombreV, newEvento.Fecha, newEvento.IdJornada, newEvento.Deporte)
}

func Agregar_Jornada(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newJornada Jornada2
	json.NewDecoder(r.Body).Decode(&newJornada)
	newJornada.Fecha = strings.Replace(newJornada.Fecha, "T", " ", -1)
	newJornada.Fecha2 = strings.Replace(newJornada.Fecha2, "T", " ", -1)
	tempo := validar_Temporada2(newJornada.Temporada)
	jor := validar_Jornada(newJornada.Nombre, tempo)
	fmt.Println(newJornada)
	if jor == "" {
		Crear_Jornada(newJornada.Nombre, newJornada.Fecha, newJornada.Fecha2, newJornada.Temporada, newJornada.Fase)
		json.NewEncoder(w).Encode("Jornada Creada")
	} else {
		json.NewEncoder(w).Encode("Jornada no Creada")
	}
}

func Agregar_Temporada(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newTemporada Temporada2
	json.NewDecoder(r.Body).Decode(&newTemporada)
	newTemporada.Fecha = strings.Replace(newTemporada.Fecha, "T", " ", -1)
	newTemporada.Fecha2 = strings.Replace(newTemporada.Fecha2, "T", " ", -1)
	fmt.Println(newTemporada)
	tempo := validar_Temporada(newTemporada.Nombre)
	if tempo == "" {
		Crear_Temporada(newTemporada.Nombre, newTemporada.Fecha, newTemporada.Fecha2, newTemporada.Fase)
		json.NewEncoder(w).Encode("Temporada Creada")
	} else {
		json.NewEncoder(w).Encode("Temporada no Creada")
	}
}

type Resultado2 struct {
	Id        string `json:"id"`
	Local     string `json:"local"`
	Visitante string `json:"visitante"`
}

func Actualizar_Resultado(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newResultado Resultado2
	json.NewDecoder(r.Body).Decode(&newResultado)

	Agregar_Resultado(newResultado.Local, newResultado.Visitante, newResultado.Id)

}

type Usu struct {
	Usuario string `json:"usuario"`
	Pass    string `json:"pass"`
	Id      string `json:"id"`
}

func Inicio_Sesion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var nuevoInicio Usu
	json.NewDecoder(r.Body).Decode(&nuevoInicio)
	nuevoInicio.Id = Inicio(nuevoInicio.Usuario, nuevoInicio.Pass)
	json.NewEncoder(w).Encode(nuevoInicio)
}

type Predic struct {
	Local     string `json:"local"`
	Visitante string `json:"visitante"`
	IdCliente string `json:"id_cliente"`
	IdEvento  string `json:"id_evento"`
}

func Crear_Prediccion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var nuevaPrediccion Predic
	json.NewDecoder(r.Body).Decode(&nuevaPrediccion)
	fmt.Println(nuevaPrediccion)
	salida := Agregar_Prediccion(nuevaPrediccion.Local, nuevaPrediccion.Visitante, nuevaPrediccion.IdCliente, nuevaPrediccion.IdEvento)
	fmt.Println(salida)
	json.NewEncoder(w).Encode(salida)
}

func Retornar_Usuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Usuarios())
}

func retornar_Tempo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Temporadas())
}

type Tempo_Usu struct {
	Id_tempo   string `json:"id_tempo"`
	Id_cliente string `json:"id_cliente"`
}

func Tabla_TU(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var nuevoTu Tempo_Usu
	json.NewDecoder(r.Body).Decode(&nuevoTu)
	fmt.Println(nuevoTu)
	json.NewEncoder(w).Encode(Tabal_Temporada(nuevoTu.Id_cliente, nuevoTu.Id_tempo))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/Carga", CargaMasiva).Methods("POST")
	router.HandleFunc("/Eventos", Enviar_Evento).Methods("GET")
	router.HandleFunc("/ConsultarJornada", Estado_Jornada).Methods("GET")
	router.HandleFunc("/ConsultarTemporada", Estado_Temporada).Methods("GET")
	router.HandleFunc("/AgregarEvento", Agregar_Evento).Methods("POST")
	router.HandleFunc("/AgregarJornada", Agregar_Jornada).Methods("POST")
	router.HandleFunc("/AgregarTemporada", Agregar_Temporada).Methods("POST")
	router.HandleFunc("/AgregarResultado", Actualizar_Resultado).Methods("POST")
	router.HandleFunc("/InicioSesion", Inicio_Sesion).Methods("POST")
	router.HandleFunc("/IngresarPrediccion", Crear_Prediccion).Methods("POST")
	router.HandleFunc("/Usuarios", Retornar_Usuario).Methods("GET")
	router.HandleFunc("/Temporadas", retornar_Tempo).Methods("GET")
	router.HandleFunc("/TemporadaUsuario", Tabla_TU).Methods("POST")

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":3003", handler))
}

/*
import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/godror/godror"
)*/

func conexion() (consulta string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Query("SELECT ID, NOMBRE FROM cliente")

	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var x int
	var id string
	var cunsulta string
	for rows.Next() {
		rows.Scan(&x, &id)
		cunsulta = fmt.Sprintf("%s%s%s", strconv.Itoa(x), ",", id)
	}

	return cunsulta
}

func ingresar_Usuario(nombre string, apellido string, pass string, usuario string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Exec(`CALL Ingresar_Cliente(:1,:2,:3,:4)`,
		nombre, apellido, pass, usuario)

	if err != nil {
		fmt.Println("ingresar usuario")
		fmt.Println("Error running query")
		fmt.Println(err)
		fmt.Println(rows)
		return
	}

	//defer rows.Close()
}

func ingresar_Deporte(nombre string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Exec(`CALL Ingresar_Deporte(:1)`,
		nombre)

	if err != nil {
		fmt.Println("ingresar deporte")
		fmt.Println("Error running query")
		fmt.Println(err)
		fmt.Println(rows)
		return
	}
	//defer rows.Close()
}

func ingresar_Evento(nombreL string, nombreV string, resultadoL string, resultadoV string, fecha string, deporte string, jornada string, temporada string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Exec(`CALL Ingresar_Evento(:1, :2, :3, :4, :5, :6, :7, :8)`,
		nombreL, nombreV, resultadoL, resultadoV, fecha, deporte, jornada, temporada)

	if err != nil {
		fmt.Println("ingresar evento")
		fmt.Println("Error running query")
		fmt.Println(err)
		fmt.Println(rows)
		return
	}
}

func ingresar_Jornada(nombre string, fecha_i string, fecha_f string, temporada string, fase string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Exec(`CALL Ingresar_Jornada(:1, :2, :3, :4, :5)`,
		nombre, fecha_i, fecha_f, temporada, fase)

	if err != nil {
		fmt.Println("ingresar jornada")
		fmt.Println("Error running query")
		fmt.Println(err)
		fmt.Println(rows)
		return
	}
}

func Ingresar_Membresia(nombre string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Exec(`CALL Ingresar_Membresia(:1)`,
		nombre)

	if err != nil {
		fmt.Println("ingresar membresia")
		fmt.Println("Error running query")
		fmt.Println(err)
		fmt.Println(rows)
		return
	}
}

func ingrrsar_Membresia_Temp(cliente string, membresia string, temporada string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Exec(`CALL Ingresar_Membresia_Temp(:1, :2, :3)`,
		cliente, membresia, temporada)

	if err != nil {
		fmt.Println("ingresar Membresia temporal")
		fmt.Println("Error running query")
		fmt.Println(err)
		fmt.Println(rows)
		return
	}
}

func ingresar_Prediccion(prediccionL string, prediccionV string, cliente string, evento string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Exec(`CALL Ingresar_Prediccion(:1, :2, :3, :4)`,
		prediccionL, prediccionV, cliente, evento)

	if err != nil {
		fmt.Println("ingresar prediccion")
		fmt.Println("Error running query")
		fmt.Println(err)
		fmt.Println(rows)
		return
	}
}

func ingresar_Temorada(nombre string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Exec(`CALL Ingresar_Temporada(:1)`,
		nombre)

	if err != nil {
		fmt.Print("ingresar temorada")
		fmt.Println("Error running query")
		fmt.Println(err)
		fmt.Println(rows)
		return
	}
}

func retornar_Evento(nombreL string, nombreV string, fecha string) (consulta string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Query(`SELECT ID FROM EVENTO where NOMBRE_LOCAL = :1 and NOMBRE_VISITANTE = :2 and TO_CHAR(FECHA,'DD/MM/YYYY HH24:MI') = :3`,
		nombreL, nombreV, fecha)

	if err != nil {
		fmt.Println("retorno de evento")
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var x int
	for rows.Next() {
		rows.Scan(&x)
		consulta = strconv.Itoa(x)
	}
	fmt.Println(strconv.Itoa(x))
	return consulta
}

func retornar_Temporada(nombreL string) (consulta string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Query(`SELECT ID FROM TEMPORADA WHERE NOMBRE = :1`,
		nombreL)

	if err != nil {
		fmt.Println("retorno de evento")
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var x int
	for rows.Next() {
		rows.Scan(&x)
		consulta = strconv.Itoa(x)
	}
	return consulta
}

func valida_Usuario(nombre string) (consulta string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Query(`SELECT USUARIO FROM CLIENTE WHERE USUARIO = :1`,
		nombre)

	if err != nil {
		fmt.Println("Validar usuario")
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var x string
	for rows.Next() {
		rows.Scan(&x)
		consulta = x
	}
	fmt.Println(x)
	return consulta
}

func validar_Temporada(nombre string) (consulta string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Query(`SELECT NOMBRE FROM TEMPORADA WHERE NOMBRE = :1`,
		nombre)

	if err != nil {
		fmt.Println("validar temporada")
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var x string
	for rows.Next() {
		rows.Scan(&x)
		consulta = x
	}
	fmt.Println(x)
	return consulta
}

func validar_Temporada2(nombre string) (consulta string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Query(`SELECT NOMBRE FROM TEMPORADA WHERE ID = :1`,
		nombre)

	if err != nil {
		fmt.Println("validar temporada2")
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var x string
	for rows.Next() {
		rows.Scan(&x)
		consulta = x
	}
	fmt.Println(x)
	return consulta
}

func validar_Tier(nombre string) (consulta string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Query(`SELECT NOMBRE FROM MEMBRESIA WHERE NOMBRE = :1`,
		nombre)

	if err != nil {
		fmt.Println("validar tier")
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var x string
	for rows.Next() {
		rows.Scan(&x)
		consulta = x
	}
	fmt.Println(x)
	return consulta
}

func validar_Deporte(nombre string) (consulta string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Query(`SELECT ID from DEPORTE WHERE NOMBRE = :1`,
		nombre)

	if err != nil {
		fmt.Println("Validar deporte")
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var x string
	for rows.Next() {
		rows.Scan(&x)
		consulta = x
	}
	fmt.Println(x)
	return consulta
}

func validar_Jornada(nombre string, temporada string) (consulta string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Query(`SELECT jornada.NOMBRE FROM JORNADA, TEMPORADA WHERE jornada.NOMBRE = :1 and TEMPORADA.ID = JORNADA.ID_TEMPORADA and TEMPORADA.NOMBRE = :2`,
		nombre, temporada)

	if err != nil {
		fmt.Println("Validar jornada")
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var x string
	for rows.Next() {
		rows.Scan(&x)
		consulta = x
	}
	fmt.Println(x)
	return consulta
}

func validar_Evento(nombreL string, nombreV string, fecha string) (consulta string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Query(`SELECT NOMBRE_LOCAL FROM EVENTO WHERE NOMBRE_LOCAL = :1 and NOMBRE_VISITANTE = :2 and to_char(fecha,'DD/MM/YYYY HH24:MI') = :3`,
		nombreL, nombreV, fecha)

	if err != nil {
		fmt.Println("Validar jornada")
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var x string
	for rows.Next() {
		rows.Scan(&x)
		consulta = x
	}
	fmt.Println(x)
	return consulta
}

func validar_Prediccion(nombre string, fecha string) (consulta string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Query(`SELECT PREDICCION.ID FROM PREDICCION, CLIENTE, EVENTO WHERE CLIENTE.ID = PREDICCION.ID_CLIENTE and EVENTO.ID = PREDICCION.ID_CLIENTE and CLIENTE.USUARIO = :1 and TO_CHAR(evento.FECHA,'DD/MM/YYYY HH24:MI') = :2`,
		nombre, fecha)

	if err != nil {
		fmt.Println("Validar jornada")
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var x string
	for rows.Next() {
		rows.Scan(&x)
		consulta = x
	}
	fmt.Println(x)
	return consulta
}

type Eventos []Dentro

func Buscar_evento() (consulta Retorno) {
	//var EventosCargados = Eventos{}
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Query(`SELECT ID, NOMBRE_LOCAL, NOMBRE_VISITANTE, TO_CHAR(FECHA,'YYYY-MM-DD HH24:MI'), R_LOCAL, R_VISITANTE FROM EVENTO`)

	if err != nil {
		fmt.Println("Validar jornada")
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()
	var retorno Retorno
	var id int
	var local string
	var visitante string
	var fecha string
	var r_local int
	var r_visitante int
	//var captura string
	//layout := "2006-01-02T15:04:05.000Z"
	for rows.Next() {
		rows.Scan(&id, &local, &visitante, &fecha, &r_local, &r_visitante)
		//captura := `{ "title":"` + local + ` vs ` + visitante + `", "start":` + fecha + `}`

		var dento Dentro
		dento.Id = strconv.Itoa(id)
		dento.Title = local + " vs " + visitante + " " + strconv.Itoa(r_local) + " - " + strconv.Itoa(r_visitante)

		//fmt.Println(t)
		dento.Start = fecha

		//json.Unmarshal([]byte(captura), &dento)

		retorno.Eventos = append(retorno.Eventos, dento)
	}

	return retorno
}

func id_jornada(nombre string, temporada string) (consulta string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Query(`select jornada.ID from Jornada, Temporada where jornada.nombre = :1 and jornada.id_temporada = temporada.id and temporada.nombre = :2`,
		nombre, temporada)

	if err != nil {
		fmt.Println("Validar jornada")
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var x string
	for rows.Next() {
		rows.Scan(&x)
		consulta = x
	}
	fmt.Println(x)
	return consulta
}

func id_Temporada(nombre string) (consulta string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Query(`select TEMPORADA.ID FROM Temporada WHERE Temporada.NOMBRE = :1`,
		nombre)

	if err != nil {
		fmt.Println("Validar jornada")
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var x string
	for rows.Next() {
		rows.Scan(&x)
		consulta = x
	}
	fmt.Println(x)
	return consulta
}

func Actualizar_Jornada(id string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Exec(`CALL Actualizar_Jornada(:1)`,
		id)

	if err != nil {
		fmt.Println("Actualizar Jornada")
		fmt.Println("Error running query")
		fmt.Println(err)
		fmt.Println(rows)
		return
	}
}

func Actualizar_Temporada(id string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Exec(`CALL Actualizar_Temporada(:1)`,
		id)

	if err != nil {
		fmt.Println("Actualizar Jornada")
		fmt.Println("Error running query")
		fmt.Println(err)
		fmt.Println(rows)
		return
	}
}

type Jornada2 struct {
	Id        string `json:"id"`
	Nombre    string `json:"nombre"`
	Fecha     string `json:"fecha_i"`
	Fecha2    string `json:"fecha_f"`
	Temporada string `json:"temporada"`
	Fase      string `json:"fase"`
}

func Consultar_Jornada() (consulta Jornada2) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Query(`SELECT * FROM (SELECT JORNADA.ID, FASE.NOMBRE, TO_CHAR(JORNADA.FECHA_FIN,'YYYY-MM-DD HH24:MI') FROM JORNADA, FASE WHERE JORNADA.ID_FASE = FASE.ID
	ORDER BY TO_CHAR(JORNADA.FECHA_FIN,'YYYY-MM-DD HH24:MI') DESC)
	WHERE ROWNUM = 1`)

	if err != nil {
		fmt.Println("Validar jornada")
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var id int
	var fecha string
	var fase string
	var jornada Jornada2
	for rows.Next() {
		rows.Scan(&id, &fase, &fecha)
		jornada.Id = strconv.Itoa(id)
		jornada.Fase = fase
		jornada.Fecha2 = fecha
	}
	return jornada
}

type Temporada2 struct {
	Id     string `json:"id"`
	Nombre string `json:"nombre"`
	Fecha  string `json:"fecha"`
	Fecha2 string `json:"fechaf"`
	Fase   string `json:"fase"`
}

func Consultar_Temporada() (consulta Temporada2) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Query(`SELECT * FROM (SELECT TEMPORADA.ID, FASE.NOMBRE, TO_CHAR(TEMPORADA.FECHA_FIN,'YYYY-MM-DD HH24:MI') FROM TEMPORADA, FASE WHERE TEMPORADA.FASE = FASE.ID
	ORDER BY TO_CHAR(TEMPORADA.FECHA_FIN,'YYYY-MM-DD HH24:MI') ASC)
	WHERE ROWNUM = 1`)

	if err != nil {
		fmt.Println("Validar jornada")
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var id int
	var fecha string
	var fase string
	var jornada Temporada2
	for rows.Next() {
		rows.Scan(&id, &fase, &fecha)
		jornada.Id = strconv.Itoa(id)
		jornada.Fase = fase
		jornada.Fecha2 = fecha
	}
	return jornada
}

func Crear_Evento(nombreL string, nombreV string, fecha string, idJornada string, deporte string) {

	idDeporte := validar_Deporte(deporte)
	if idDeporte == "" {
		ingresar_Deporte(deporte)
		idDeporte = validar_Deporte(deporte)
	}

	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Exec(`CALL Crear_Evento(:1, :2, :3, :4, :5)`,
		nombreL, nombreV, fecha, idJornada, idDeporte)

	if err != nil {
		fmt.Println("Crear Evento")
		fmt.Println("Error running query")
		fmt.Println(err)
		fmt.Println(rows)
		return
	}
}

func Crear_Temporada(nombre string, fecha string, fechaf string, fase string) {

	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Exec(`CALL Creat_Temporada(:1, :2, :3, :4)`,
		nombre, fecha, fechaf, fase)

	if err != nil {
		fmt.Println("Crear Temporada")
		fmt.Println("Error running query")
		fmt.Println(err)
		fmt.Println(rows)
		return
	}
}

func Crear_Jornada(nombre string, fechai string, fechaf string, temorada string, fase string) {

	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Exec(`CALL Crear_Jornada(:1, :2, :3, :4, :5)`,
		nombre, fechai, fechaf, temorada, fase)

	if err != nil {
		fmt.Println("Crear Jornada")
		fmt.Println("Error running query")
		fmt.Println(err)
		fmt.Println(rows)
		return
	}
}

func Agregar_Resultado(resultadoL string, resultadoV string, id string) {

	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Exec(`CALL Agregar_Resultado(:1, :2, :3)`,
		resultadoL, resultadoV, id)

	if err != nil {
		fmt.Println("Agregar Resultado")
		fmt.Println("Error running query")
		fmt.Println(err)
		fmt.Println(rows)
		return
	}
}

func Inicio(nombre string, pass string) (consulta string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()
	var salida string
	rows, err := db.Exec(`CALL Inicio_sesion(:1, :2, :3)`,
		sql.Out{Dest: &salida}, nombre, pass)

	if err != nil {
		fmt.Println("Inicio Secion")
		fmt.Println(rows)
		fmt.Println(err)
		return
	}
	return salida
}

func Agregar_Prediccion(local string, visitante string, idC string, idE string) (consulta string) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()
	var salida string
	rows, err := db.Exec(`CALL Agregar_Prediccion(:1, :2, :3, :4, :5)`,
		local, visitante, idC, idE, sql.Out{Dest: &salida})

	if err != nil {
		fmt.Println("Ingresar Prediccion")
		fmt.Println(rows)
		fmt.Println(err)
		return
	}
	return salida
}

type Contenedor_usuario struct {
	Usuarios []Usu `json:"usuario"`
}

func Usuarios() (consulta Contenedor_usuario) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT ID, USUARIO FROM cliente")

	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var usuarios Contenedor_usuario
	var id int
	var nombre string
	for rows.Next() {
		var usuario Usu
		rows.Scan(&id, &nombre)
		usuario.Id = strconv.Itoa(id)
		usuario.Usuario = nombre
		usuarios.Usuarios = append(usuarios.Usuarios, usuario)
	}

	return usuarios
}

type Lista_Temporada struct {
	Id     string `json:"id"`
	Nombre string `json:"nombre"`
}
type Contenedor_Temporada struct {
	Temoradas []Lista_Temporada `json:"temoradas"`
}

func Temporadas() (consulta Contenedor_Temporada) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Query("SELECT ID, NOMBRE FROM Temporada")

	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var temporadas Contenedor_Temporada
	var id int
	var nombre string
	for rows.Next() {
		var tempo Lista_Temporada
		rows.Scan(&id, &nombre)
		tempo.Id = strconv.Itoa(id)
		tempo.Nombre = nombre
		temporadas.Temoradas = append(temporadas.Temoradas, tempo)
	}

	return temporadas
}

type Contenedor_TU struct {
	Datos []Temporada_Usuario `json:"datos"`
}

type Temporada_Usuario struct {
	Deporte    string `json:"deporte"`
	Local      string `json:"local"`
	Visitante  string `json:"visitante"`
	Prediccion string `json:"prediccion"`
	Resultado  string `json:"resultado"`
	Puntos     string `json:"puntos"`
	Fecha      string `json:"fecha"`
}

func Tabal_Temporada(idC string, idT string) (consulta Contenedor_TU) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()
	fmt.Println(idC + " " + idT)
	rows, err := db.Query(`select Deporte.NOMBRE, Evento.NOMBRE_LOCAL, EVENTO.NOMBRE_VISITANTE, PREDICCION.PUNTOD_LOCAL, PREDICCION.PUNTOS_VISITANTE, EVENTO.R_LOCAL, EVENTO.R_VISITANTE, PREDICCION.PUNTOS_OBTENIDOS, to_char(EVENTO.FECHA, 'YYYY-MM-DD HH24:MI')
	from Deporte, Cliente, Evento, Temporada, Jornada, Prediccion
	WHERE TEMPORADA.ID = JORNADA.ID_TEMPORADA and EVENTO.ID_JORNADA = JORNADA.ID and DEPORTE.ID = EVENTO.ID_DEPORTE and EVENTO.ID = PREDICCION.ID_EVENTO and CLIENTE.ID = PREDICCION.ID_CLIENTE and CLIENTE.ID = :1 and TEMPORADA.ID = :2`,
		idC, idT)

	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var contenedor Contenedor_TU
	var deporte string
	var local string
	var visita string
	var p_local int
	var p_visita int
	var r_local int
	var r_visita int
	var puntos int
	var fecha string
	for rows.Next() {
		var usuario Temporada_Usuario
		rows.Scan(&deporte, &local, &visita, &p_local, &p_visita, &r_local, &r_visita, &puntos, &fecha)
		usuario.Deporte = deporte
		usuario.Local = local
		usuario.Visitante = visita
		usuario.Prediccion = strconv.Itoa(p_local) + " - " + strconv.Itoa(p_visita)
		usuario.Resultado = strconv.Itoa(r_local) + " - " + strconv.Itoa(r_visita)
		usuario.Puntos = strconv.Itoa(puntos)
		usuario.Fecha = fecha
		contenedor.Datos = append(contenedor.Datos, usuario)
	}

	return contenedor
}
