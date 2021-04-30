package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/godror/godror"
	"github.com/gorilla/mux"
)

type Hola struct {
	Texto string `json:"Texto" json`
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, conexion())

}

func CargaMasiva(w http.ResponseWriter, r *http.Request) {
	var newTask Hola
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task Data")
	}

	json.Unmarshal(reqBody, &newTask)

	fmt.Println(newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/Carga", CargaMasiva).Methods("POST")
	log.Fatal(http.ListenAndServe(":3003", router))
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
