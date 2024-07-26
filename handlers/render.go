package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// Load all templates
var templates = template.Must(template.ParseGlob(filepath.Join("templates", "*.html")))

// var errorTemplate = template.Must(template.ParseGlob(filepath.Join("templates", "*.html")))

var errorTemplate = template.Must(template.ParseFiles(filepath.Join("templates", "error.html")))

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	// fmt.Println("DATA:", data)
	var err error
	if tmpl == "error" {
		// err = templates.ExecuteTemplate(w, tmpl, data)
		err = errorTemplate.Execute(w, data)
	} else {
		err = templates.ExecuteTemplate(w, tmpl, data)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
