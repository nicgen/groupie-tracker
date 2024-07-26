package main

import (
	"handlers"
	"log"
	"net/http"
	"time"
)

func main() {
	// // Create a new ServeMux
	// mux := http.NewServeMux()

	// // Serve static files
	// mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// // Set up routes
	// mux.HandleFunc("/", handlers.IndexHandler)
	// mux.HandleFunc("/about", handlers.AboutHandler)
	// mux.HandleFunc("/force500", handlers.Force500Handler)

	// // Wrap the ServeMux with our error handling middleware
	// handler := handlers.WithErrorHandling(mux)

	// // Create a new http.Server with security settings
	// server := &http.Server{
	// 	Addr:              "localhost:8080",  // restricts the server to listen only on the localhost interface
	// 	Handler:           handler,           // assign existing handler to the Handler
	// 	ReadHeaderTimeout: 10 * time.Second,  // timeout for reading the request headers, helps prevent slow-loris type attacks
	// 	WriteTimeout:      10 * time.Second,  // timeout for writing the response, helps prevent resource exhaustion from slow clients
	// 	IdleTimeout:       120 * time.Second, // timeout for keep-alive connections, helps manage idle connections and frees up resources
	// 	MaxHeaderBytes:    1 << 20,           // limits the maximum header size to 1 MB, prevent denial-of-service attacks via large headers
	// }
	// log.Printf("Server starting on http://%s...\n", server.Addr)

	// if err := server.ListenAndServe(); err != nil {
	// 	log.Fatalf("Error starting server: %v", err)
	// }
	mux := setupMux()

	server := setupServer(mux)
	log.Printf("Server starting on http://%s...\n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func setupMux() *http.ServeMux {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Serve static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Set up routes
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/about", handlers.AboutHandler)
	mux.HandleFunc("/force500", handlers.Force500Handler) // for testing purpose only

	return mux
}

func setupServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              "localhost:8080",
		Handler:           handlers.WithErrorHandling(handler),
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
}
