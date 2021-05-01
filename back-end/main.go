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

type Hola struct {
	Texto string `json:"Texto"`
}

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
		fmt.Println("ID:", key)
		fmt.Print("\t")
		fmt.Println("Usuario:", element.User)
		fmt.Print("\t")
		fmt.Println("Clave:", element.Pass)
		fmt.Print("\t")
		fmt.Println("Nombre:", element.Nombre)
		fmt.Print("\t")
		fmt.Println("Apellido:", element.Apellido)
		for _, element := range element.Resultado_ {
			fmt.Print("\t")
			fmt.Println("Resultados:")
			fmt.Print("\t\t")
			fmt.Println("Temporada:", element.Temporada)
			fmt.Print("\t\t")
			fmt.Println("Membresia:", element.Tier)
			for _, element := range element.Jornadas {
				fmt.Print("\t\t")
				fmt.Println("Jornadas:")
				fmt.Print("\t\t\t")
				fmt.Println("Jornada:", element.Jornada)
				for _, element := range element.Evento {
					fmt.Print("\t\t\t")
					fmt.Println("Predicciones:")
					fmt.Print("\t\t\t\t")
					fmt.Println("Deporte:", element.Deporte)
					fmt.Print("\t\t\t\t")
					fmt.Println("Local:", element.Local)
					fmt.Print("\t\t\t\t")
					fmt.Println("Visitante:", element.Visitante)
					fmt.Print("\t\t\t\t")
					fmt.Println("Fecha:", element.Fecha)
					fmt.Print("\t\t\t\t")
					fmt.Println("Prediccion:")
					fmt.Print("\t\t\t\t\t")
					fmt.Println("P Local:", element.Prediccion.Local)
					fmt.Print("\t\t\t\t\t")
					fmt.Println("P Visita:", element.Prediccion.Visitante)
					fmt.Print("\t\t\t\t")
					fmt.Println("Resultado:")
					fmt.Print("\t\t\t\t\t")
					fmt.Println("R Local:", element.Resultado.Local)
					fmt.Print("\t\t\t\t\t")
					fmt.Println("R Visita:", element.Resultado.Visitante)

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
