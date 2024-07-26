package models

// CustomError defines a custom error type with status code and message
type CustomError struct {
	StatusCode int
	Message    string
}

// Error implements the error interface for CustomError
func (e *CustomError) Error() string {
	return e.Message
}
