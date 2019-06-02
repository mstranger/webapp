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

	type User struct {
		ID    int
		Name  string
		Email string
	}

	var users []User

	// var id int
	// row := db.QueryRow(`
	// 	INSERT INTO users(name, email)
	// 	VALUES($1, $2)
	// 	RETURNING id`,
	// 	"Mike Smith", "mike@mail.com")

	// var name, email string
	// row := db.QueryRow(`
	// 	SELECT id, name, email
	// 	FROM users
	// 	WHERE id=$1`, 3)

	rows, err := db.Query(`
		SELECT id, name, email
		FROM users`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	if rows.Err() != nil {
		// handle the err
	}

	fmt.Println(users)

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Successfully connected!")
	// db.Close()
}
