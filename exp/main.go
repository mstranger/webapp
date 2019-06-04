package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "mstranger"
	password = "password"
	dbname   = "webapp_dev"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
	Color string
}

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.LogMode(true)
	db.AutoMigrate(&User{})

	var u User
	// newdb := db.Where("id = ? AND color = ?", 4, "red")
	// newdb.First(&u)
	db.Where("color = ?", "red").
		Where("id > ?", 4).
		First(&u)
	// db.First(&u, "color = ?", "red")
	fmt.Println(u)

	var users []User
	db.Find(&users)
	fmt.Println(len(users))
	fmt.Println(users)
}
