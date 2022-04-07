package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	// const (
	// 	host     = "localhost"
	// 	port     = 5432
	// 	user     = "pavi"
	// 	password = "pavi"
	// 	dbname   = "pavidb"
	// )
	table := flag.String("t", "", "enter tablename")
	username := flag.String("u", "", "enter username")
	password := flag.String("p", "", "enter password")
	host := flag.String("h", "localhost", "enter host")
	dbname := flag.String("db", "", "enter dbname")
	port := flag.Int("port", 5432, "enter port")
	action := flag.String("a", "", "enter an action")

	flag.Parse()
	t := *action
	z := t + " " + *table
	fmt.Println(z)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", *host, *port, *username, *password, *dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	rows, err := db.Query(z)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	println(rows)
	for rows.Next() {
		var first_name, last_name, email, gender string
		var id int
		if err := rows.Scan(&id, &first_name, &last_name, &email, &gender); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %d\nfirst_name: %s\nlast_name: %s\nemail: %s\ngender: %s\n---\n\n", id, first_name, last_name, email, gender)
	}
}

//   ./main testcase -u=username -p=password -db=databasename -t=tablename -a="select * from"
