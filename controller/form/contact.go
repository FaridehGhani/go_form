package form

import (
	"github.com/FaridehGhani/go_form/model/form"
	"html/template"
	"log"
	"net/http"
)

func ContactForm(w http.ResponseWriter, r *http.Request) {
	render(w, "view/templates/contact.html", nil)
}

func SubmitContactForm(w http.ResponseWriter, r *http.Request) {
	// validate form parameters
	contact := &form.Contact{
		Email:   r.PostFormValue("email"),
		Content: r.PostFormValue("content"),
	}

	if contact.Validate() == false {
		render(w, "view/templates/contact.html", contact)

		return
	}

	_, err := contact.CreateContent()
	if err != nil {
		contact.Errors = err.Error()
		render(w, "view/templates/contact.html", contact)

		return
	}

	// redirect to submitted page
	http.Redirect(w, r, "/submitted", http.StatusSeeOther)
}

func SubmittedContactForm(w http.ResponseWriter, r *http.Request) {
	render(w, "view/templates/submitted.html", nil)
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
