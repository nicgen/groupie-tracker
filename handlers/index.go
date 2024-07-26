package handlers

import (
	"groupie_tracker/models"
	"net/http"
)

// IndexHandler handles requests to the root URL
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		err := &models.CustomError{
			StatusCode: http.StatusNotFound,
			Message:    "Page Not Found",
		}
		// Use HandleError to send the error response
		HandleError(w, err.StatusCode, err.Message)
		return
	}

	data := models.PageData{
		Title:  "Home",
		Header: "Welcome to My Web App",
		Content: map[string]interface{}{
			"Message": "This is the home page.",
		},
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	renderTemplate(w, "layout", data)
}
