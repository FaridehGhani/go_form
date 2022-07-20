package main

import (
	"github.com/FaridehGhani/go_form/model/form"
	"log"
	"net/http"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"

	"github.com/FaridehGhani/go_form/api/route"
	"github.com/FaridehGhani/go_form/infra/mysql"
)

func init() {
	// load project env variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %s", err)
	}
}

func main() {
	// connect to database
	db := mysql.Connect()
	log.Println("mysql connects successfully ...")
	db.Debug().AutoMigrate(&form.Contact{})
	defer db.Close()

	// load routes
	routes := route.GetRouter()
	log.Print("listening ...")
	err := http.ListenAndServe(os.Getenv("server_port"), routes)
	if err != nil {
		log.Fatal(err)
	}
}
