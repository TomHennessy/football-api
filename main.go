package main

import (
	// "fmt"
	// "io/ioutil"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
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

	// defer db.Close()
}