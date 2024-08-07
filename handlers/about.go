package handlers

import (
	"net/http"
	"web-starter/models"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	data := models.PageData{
		Title:  "Title About",
		Header: "Header About",
		Content: map[string]interface{}{
			"Message": "This is the about page.",
			"other":   "other string",
		},
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	renderTemplate(w, "layout", data)
}
