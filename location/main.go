package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/location"
	"github.com/jinzhu/configor"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"net/http"
	"fmt"
)

type Store struct {
	gorm.Model

	Name string
	location.Location
}

type Config struct {
	GoogleAPIKey string
}


func main(){
	db, _ := gorm.Open("sqlite3", "demo.db")
	defer db.Close()

	db.AutoMigrate(&Store{})

	// before your run the project, create the config/app.yml file and add the following content:
	// googleapikey: xxxxxx
	config := Config{}
	configor.Load(&config, "config/app.yml")

	// you should set the google api key
	location.GoogleAPIKey = config.GoogleAPIKey

	Admin := admin.New(&qor.Config{DB:db})
	Admin.AddResource(&Store{})

	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)

	port := ":9090"
	fmt.Println("server listen on ", port)
	http.ListenAndServe(port, mux)

}
