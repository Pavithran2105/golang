package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "pavi"
// 	password = "pavi"
// 	dbname   = "pavidb"
// )

type sandbox struct {
	id        int
	Firstname string
	Lastname  string
	Age       int
}

func init() {

	connStr := "postgres://pavi:pavi@localhost/pavidb?sslmode=disable"
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	// this will be printed in the terminal, notifying that we successfully connected to our database
	fmt.Println("You are now connected to the database.")
}

func main() {
	http.HandleFunc("/retrieve", retrieveRecord)
	http.ListenAndServe(":8080", nil)
}

func retrieveRecord(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Cool, your connected to the browser \n")

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * FROM tblretrieveTest")

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	snbs := make([]sandbox, 0)

	for rows.Next() {
		snb := sandbox{}
		err := rows.Scan(&snb.id, &snb.Firstname, &snb.Lastname, &snb.Age)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		snbs = append(snbs, snb)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, snb := range snbs {
		fmt.Fprintf(w, "%d %s %s %d\n", snb.id, snb.Firstname, snb.Lastname, snb.Age)
	}

}
