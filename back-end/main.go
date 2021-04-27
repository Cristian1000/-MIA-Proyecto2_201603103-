package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/godror/godror"
)

func main() {
	//Oracle 12c
	db, err := sql.Open("godror", "cris/1234@localhost:1521/ORCL18")
	//Oracle 18c
	// db, err := sql.Open("godror", "user/password@localhost:1521/ORCL18.localdomain")
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
	for rows.Next() {

		rows.Scan(&x, &id)

		fmt.Printf(strconv.Itoa(x)+" %s", id)
		fmt.Printf("\n")
	}
}
