package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/validations"
)

type User struct {
	gorm.Model
	Age uint
}

func (user User) Validate(db *gorm.DB) {
	fmt.Println("start validation", user.Age)
	if user.Age <= 18 {
		db.AddError(validations.NewError(user, "Age", "age need to be 18+"))
	}
}

func main() {
	db, _ := gorm.Open("sqlite3", "demo.db")
	defer db.Close()

	validations.RegisterCallbacks(db)

	db.AutoMigrate(&User{})

	user := User{
		Age: 15,
	}
	errs := db.Create(&user).GetErrors()
	for _, err := range errs {
		if e, ok := err.(*validations.Error); ok {
			fmt.Println(e.Label() + " " + e.Error())
		} else {
			fmt.Println(err)
		}
	}

}
