package main

import (
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/jinzhu/gorm"
	"github.com/qor/activity"

	_ "github.com/mattn/go-sqlite3"

	"net/http"
	"fmt"
)

type Order struct {
	gorm.Model

	Name string
}

func main() {
	DB, _ := gorm.Open("sqlite3", "demo.db")
	defer DB.Close()


	DB.AutoMigrate(&Order{})

	Admin := admin.New(&qor.Config{DB: DB})

	order := Admin.AddResource(&Order{})
	activity.Register(order)

	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)

	port := ":9090"
	fmt.Println("will listen on port ", port)
	http.ListenAndServe(port, mux)
}
