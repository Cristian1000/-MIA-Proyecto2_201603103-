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

func indexRoute(w http.ResponseWriter, r *http.Request) {
	//retornar_Evento("hola", "adios", "03/02/2019 11:29")
	//valida_Usuario("Cris10")
	//validar_Deporte("golf")
	fmt.Println(retornar_Evento("Kermit Garbar", "Sarge Rowlings", "05/03/2018 11:58"))
	fmt.Fprint(w, conexion())
}

/*
func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}*/

func CargaMasiva(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var Carga Info
	json.NewDecoder(r.Body).Decode(&Carga)

	for key, element := range Carga {
		var jornada string
		var cliente string
		var temporada string
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
			}
			resultado_Tier := validar_Tier(element.Tier)
			if resultado_Tier != element.Tier {
				Ingresar_Membresia(element.Tier)
			}
			ingrrsar_Membresia_Temp(cliente, element.Tier, element.Temporada)
			for _, element := range element.Jornadas {
				jornada = element.Jornada
				ingresar_Jornada(element.Jornada, "03/02/2019 11:29", "03/02/2019 11:29", temporada, "Finalizada")
				for _, element := range element.Evento {

					//var evento string
					resultado_Dep := validar_Deporte(element.Deporte)
					if resultado_Dep != element.Deporte {
						ingresar_Deporte(element.Deporte)
					}
					retornoTem := retornar_Temporada(temporada)
					ingresar_Evento(element.Local, element.Visitante, strconv.Itoa(element.Resultado.Local), strconv.Itoa(element.Resultado.Visitante), element.Fecha, element.Deporte, jornada, retornoTem)
					evento := retornar_Evento(element.Local, element.Visitante, element.Fecha)
					ingresar_Prediccion(strconv.Itoa(element.Prediccion.Local), strconv.Itoa(element.Prediccion.Visitante), cliente, evento)
				}
			}
		}
	}

	/*(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	var newTask Usuario
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task Data")
	}

	json.Unmarshal(reqBody, &newTask)

	fmt.Println(newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)*/
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/Carga", CargaMasiva).Methods("POST")

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

//Peticiones a Oracle

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

	rows, err := db.Query(`SELECT NOMBRE from DEPORTE WHERE NOMBRE = :1`,
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
