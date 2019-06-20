package main

import (
	"fmt"

	"../models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "mstranger"
	password = "password"
	dbname   = "webapp_dev"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}

	defer us.Close()

	us.DestructiveReset()

	user := models.User{
		Name:  "Michael Scott",
		Email: "michaelscott@mail.com",
	}
	if err := us.Create(&user); err != nil {
		panic(err)
	}

	// user.Email = "blabla@mail.com"
	// if err := us.Update(&user); err != nil {
	// 	panic(err)
	// }

	if err := us.Delete(user.ID); err != nil {
		panic(err)
	}

	userByID, err := us.ByID(user.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println(userByID)

}
