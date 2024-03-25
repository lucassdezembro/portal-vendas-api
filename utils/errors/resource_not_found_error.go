package errors_utils

import "fmt"

type resourceNotFoundError struct {
	message    string
	statusCode int
}

func NewResourceNotFoundError(customMessage string) Error {

	if customMessage == "" {
		customMessage = "Resource Not Found"
	} else {
		customMessage = "Resource Not Found: " + customMessage
	}

	return resourceNotFoundError{
		message:    customMessage,
		statusCode: 404,
	}
}

func (e resourceNotFoundError) Error() string {
	return fmt.Sprintf("%d: %s", e.statusCode, e.message)
}
