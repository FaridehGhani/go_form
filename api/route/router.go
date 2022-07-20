package route

import (
	"github.com/FaridehGhani/go_form/controller/form"
	"net/http"

	"github.com/bmizerany/pat"
)

func GetRouter() *pat.PatternServeMux {
	route := pat.New()

	route.Get("/", http.HandlerFunc(form.ContactForm))
	route.Post("/", http.HandlerFunc(form.SubmitContactForm))
	route.Get("/submitted", http.HandlerFunc(form.SubmittedContactForm))

	return route
}
