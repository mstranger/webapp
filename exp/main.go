package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "mstranger"
	password = "password"
	dbname   = "webapp_dev"
)

func main() {
	// var psqlInfo string

	// if password == "" {

	// } else {

	// }

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	_, err = db.Exec(`
		INSERT INTO users(name, email)
		VALUES($1, $2)`, "John Doe", "john@mail.com")
	if err != nil {
		panic(err)
	}

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Successfully connected!")
	// db.Close()
}
