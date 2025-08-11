package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/index.html",
	))
	tmpl.ExecuteTemplate(w, "base", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Chorely!</h1>")
}
