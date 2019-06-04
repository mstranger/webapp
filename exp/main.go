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

	// var u User
	// newdb := db.Where("id = ? AND color = ?", 4, "red")
	// newdb.First(&u)
	// db.Where("color = ?", "red").
	// 	Where("id > ?", 4).
	// 	First(&u)
	// db.First(&u, "color = ?", "red")
	// fmt.Println(u)

	var u User
	// newDB := db.Where("email = ?", "blabla@bla.com")
	// newDB = newDB.Or("color = ?", "red")
	// newDB = newDB.First(&u)

	// db = db.Where("email = ? ", "blabla@bla.com").First(&u)
	if err := db.Where("email = ?", "bla@bla.com").First(&u).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			fmt.Println("No user found!")
		case gorm.ErrInvalidSQL:
			fmt.Println("Invalid query!")
		default:
			panic(err)
		}
	}
	fmt.Println(u)
	// errors := db.GetErrors()
	// if db.RecordNotFound() {
	// 	fmt.Println("No user found!")
	// } else if db.Error != nil {
	// 	panic(db.Error)
	// } else {
	// 	fmt.Println(u)
	// }

	// if len(errors) > 0 {
	// 	fmt.Println(errors)
	// 	os.Exit(1)
	// }

	// if db.Error != nil {
	// panic(db.Error)
	// }

	// fmt.Println(u)

	// var users []User
	// db.Find(&users)
	// fmt.Println(len(users))
	// fmt.Println(users)
}
