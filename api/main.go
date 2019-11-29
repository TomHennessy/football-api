package main

import (
	"fmt"
	// "io/ioutil"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	// db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/football")
	db, err := sql.Open("mysql", "root:@/football")

	if err != nil {
		log.Fatal(err)
	}
	// err = db.Ping()
	var (
		ID int
		FirstName string
	)
	rows, err := db.Query("SELECT ID, FirstName FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ID, &FirstName)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(ID, FirstName)
		// fmt.Printf(ID, FirstName)
	}
	err = rows.Err()
	if err != nil {
		// log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
		fmt.Fprintln(w, "This is a print test")
	})
	go log.Fatal(http.ListenAndServe(":8080", nil))
	select {}


	// fmt.Printf("This is a print test")



	defer db.Close()
}
