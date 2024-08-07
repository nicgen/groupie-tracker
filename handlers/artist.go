package handlers

import (
	"fmt"
	"groupie-tracker/models"
	"net/http"
	"strconv"
	"strings"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	// Debug: Print the full URL path
	fmt.Printf("Received request for path: %s\n", r.URL.Path)

	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	fmt.Printf("Extracted ID string: %s\n", idStr) // Debug: Print extracted ID string

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Printf("Error converting ID to int: %v\n", err) // Debug: Print conversion error
		HandleError(w, http.StatusBadRequest, "Invalid artist ID")
		return
	}

	fmt.Printf("Looking up artist with ID: %d\n", id) // Debug: Print the ID we're looking up

	artist, ok := models.GetArtistByID(id)
	if !ok {
		fmt.Printf("Artist with ID %d not found\n", id) // Debug: Print when artist is not found
		HandleError(w, http.StatusNotFound, "Artist not found")
		return
	}

	data := models.PageData{
		Title:   artist.Name,
		Header:  artist.Name,
		Content: artist,
		IsError: false,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	renderTemplate(w, "artist", data)
}
