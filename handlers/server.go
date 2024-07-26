package handlers

import (
	"net/http"
	"time"
)

func SetupMux() *http.ServeMux {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Serve static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Set up routes
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/about", AboutHandler)
	mux.HandleFunc("/force500", Force500Handler) // for testing purpose only

	return mux
}

func SetupServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              "localhost:8080",
		Handler:           WithErrorHandling(handler),
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
}
