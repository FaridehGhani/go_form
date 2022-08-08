package route

import (
	"net/http"

	"github.com/bmizerany/pat"

	"github.com/FaridehGhani/go_form/controller/form"
)

func GetRouter() *pat.PatternServeMux {
	route := pat.New()

	route.Get("/", http.HandlerFunc(form.ContactForm))
	route.Post("/", http.HandlerFunc(form.SubmitContactForm))
	route.Get("/submitted", http.HandlerFunc(form.SubmittedContactForm))

	return route
}
