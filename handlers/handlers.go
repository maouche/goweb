package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	if err != nil {
		fmt.Fprintf(w, err.Error(), http.StatusInternalServerError)
	}

	t.Execute(w, nil)
}
