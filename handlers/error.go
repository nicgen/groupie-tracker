package handlers

import (
	"log"
	"net"
	"net/http"
	"runtime/debug"
	"web-starter/models"
)

// HandleError sends an error response with the given status code and message
func HandleError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	data := struct {
		Code    int
		Message string
	}{
		Code:    status,
		Message: message,
	}
	renderTemplate(w, "error", data)
	// renderTemplate(w, "layout", data)
}

// WithErrorHandling middleware that handles all errors and panics
func WithErrorHandling(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Log the error and stack trace
				log.Printf("Panic: %v\n%s", err, debug.Stack())

				// Determine the status code and message
				statusCode := http.StatusInternalServerError
				message := "Internal Server Error"

				switch e := err.(type) {
				case *models.CustomError:
					statusCode = e.StatusCode
					message = e.Message
				case *net.OpError:
					statusCode = http.StatusInternalServerError
					message = "A network error occurred"
				default:
					switch err.(string) {
					case "bad request":
						statusCode = http.StatusBadRequest
						message = "Bad Request"
					case "not found":
						statusCode = http.StatusNotFound
						message = "Not Found"
					default:
						statusCode = http.StatusInternalServerError
						message = "Internal Server Error"
					}
				}

				// Render the error page
				HandleError(w, statusCode, message)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// Force500Handler forces a 500 error for testing purposes
func Force500Handler(w http.ResponseWriter, r *http.Request) {
	panic("This is a forced panic to test 500 error handling")
}
