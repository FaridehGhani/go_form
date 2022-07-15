package main

import (
	"log"
	"net/http"

	"github.com/FaridehGhani/go_form/api"
)

func main() {
	// load routes
	routes := api.GetRouter()

	log.Print("Listening on localhost:8080 ...")
	err := http.ListenAndServe(":8080", routes)
	if err != nil {
		log.Fatal(err)
	}
}
