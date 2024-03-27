package errors_utils

import "fmt"

type internalServerError struct {
	message    string
	statusCode int
}

func NewInternalServerError(customMessage string) Error {

	if customMessage == "" {
		customMessage = "Internal Server Error"
	} else {
		customMessage = "Internal Server Error: " + customMessage
	}

	return internalServerError{
		message:    customMessage,
		statusCode: 500,
	}
}

func (e internalServerError) Error() string {
	return fmt.Sprintf("%d: %s", e.statusCode, e.message)
}
