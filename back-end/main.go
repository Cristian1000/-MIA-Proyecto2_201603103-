package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	Title string `json:"title"`
	Start string `json:"start"`
}

type Contiene map[string]Retorno

func indexRoute(w http.ResponseWriter, r *http.Request) {
	//retornar_Evento("hola", "adios", "03/02/2019 11:29")
	//valida_Usuario("Cris10")
	//validar_Deporte("golf")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Consultar_Temporada())

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
			if resultado_Temp != temporada {
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
	fmt.Println(newEvento)
	Crear_Evento(newEvento.NombreL, newEvento.NombreV, newEvento.Fecha, newEvento.IdJornada, newEvento.Deporte)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/Carga", CargaMasiva).Methods("POST")
	router.HandleFunc("/Eventos", Enviar_Evento).Methods("GET")
	router.HandleFunc("/ConsultarJornada", Estado_Jornada).Methods("GET")
	router.HandleFunc("/ConsultarTemporada", Estado_Temporada).Methods("GET")
	router.HandleFunc("/AgregarEvento", Agregar_Evento).Methods("POST")

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

	rows, err := db.Query(`SELECT NOMBRE_LOCAL, NOMBRE_VISITANTE, TO_CHAR(FECHA,'YYYY-MM-DD HH24:MI') FROM EVENTO`)

	if err != nil {
		fmt.Println("Validar jornada")
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()
	var retorno Retorno
	var local string
	var visitante string
	var fecha string
	//var captura string
	//layout := "2006-01-02T15:04:05.000Z"
	for rows.Next() {
		rows.Scan(&local, &visitante, &fecha)
		//captura := `{ "title":"` + local + ` vs ` + visitante + `", "start":` + fecha + `}`

		var dento Dentro
		dento.Title = local + " vs " + visitante

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
	Id    string `json:"id"`
	Fecha string `json:"fecha"`
}

func Consultar_Jornada() (consulta Jornada2) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Query(`SELECT * FROM (SELECT ID, TO_CHAR(JORNADA.FECHA_FIN,'YYYY-MM-DD HH24:MI') FROM JORNADA
	ORDER BY TO_CHAR(JORNADA.FECHA_FIN,'YYYY-MM-DD HH24:MI') ASC)
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
	var jornada Jornada2
	for rows.Next() {
		rows.Scan(&id, &fecha)
		jornada.Id = strconv.Itoa(id)
		if fecha == "" {
			jornada.Fecha = "Activa"
		} else {
			jornada.Fecha = fecha
		}
	}
	return jornada
}

type Temporada2 struct {
	Id    string `json:"id"`
	Fecha string `json:"fecha"`
}

func Consultar_Temporada() (consulta Temporada2) {
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Query(`SELECT * FROM (SELECT ID, TO_CHAR(TEMPORADA.FECHA_FIN,'YYYY-MM-DD HH24:MI') FROM TEMPORADA
	ORDER BY TO_CHAR(TEMPORADA.FECHA_FIN,'YYYY-MM-DD HH24:MI') DESC)
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
	var jornada Temporada2
	for rows.Next() {
		rows.Scan(&id, &fecha)
		jornada.Id = strconv.Itoa(id)
		if fecha == "" {
			jornada.Fecha = "Activa"
		} else {
			jornada.Fecha = fecha
		}
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
