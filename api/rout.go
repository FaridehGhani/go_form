package api

import (
	"net/http"

	"github.com/bmizerany/pat"

	"github.com/FaridehGhani/go_form/handler"
)

func GetRouter() *pat.PatternServeMux {
	route := pat.New()

	route.Get("/", http.HandlerFunc(handler.ContactForm))
	route.Post("/", http.HandlerFunc(handler.SubmitMessage))
	route.Get("/submitted", http.HandlerFunc(handler.SubmittedMessage))

	return route
}
