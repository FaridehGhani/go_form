package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/FaridehGhani/go_form/form"
)

func ContactForm(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/contact.html", nil)
}

func SubmitMessage(w http.ResponseWriter, r *http.Request) {
	// validate form parameters
	msg := &form.Contact{
		Email:   r.PostFormValue("email"),
		Content: r.PostFormValue("content"),
	}

	if msg.Validate() == false {
		render(w, "templates/contact.html", msg)

		return
	}

	// redirect to submitted page
	http.Redirect(w, r, "/submitted", http.StatusSeeOther)
}

func SubmittedMessage(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/submitted.html", nil)
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		log.Print(err)
		http.Error(w, "parsing template error", http.StatusInternalServerError)

		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Print(err)
		http.Error(w, "executing template error", http.StatusInternalServerError)
	}
}
