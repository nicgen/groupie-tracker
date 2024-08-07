package handlers

import (
	"fmt"
	"groupie-tracker/models"
	"net/http"
)

// IndexHandler handles requests to the root URL
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Debug: Print the full URL path
	fmt.Printf("Received request for path: %s\n", r.URL.Path)

	if r.URL.Path != "/" {
		// * generate your error message
		// err := &models.CustomError{
		// 	StatusCode: http.StatusNotFound,
		// 	Message:    "Page Not Found",
		// }
		// Use HandleError to send the error response
		// HandleError(w, err.StatusCode, err.Message)
		// return
		// * alt. use the auto-generated error code & message
		HandleError(w, http.StatusNotFound, "Page Not Found")
		return
	}

	artists, err := models.FetchArtists()
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "Error fetching artists")
		return
	}

	data := models.PageData{
		Title:  "Home",
		Header: "Welcome to My Web App",
		Content: map[string]interface{}{
			"Message":   "This is the home page.",
			"Paragraph": "This is a new paragraph",
			"Artists":   artists,
		},
		IsError: false,
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	renderTemplate(w, "index", data)
}
