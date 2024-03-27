package errors_utils

import "fmt"

type badRequestError struct {
	message    string
	statusCode int
}

func NewBadRequestError(customMessage string) Error {

	if customMessage == "" {
		customMessage = "Bad Request"
	} else {
		customMessage = "Bad Request: " + customMessage
	}

	return badRequestError{
		message:    customMessage,
		statusCode: 400,
	}
}

func (e badRequestError) Error() string {
	return fmt.Sprintf("%d: %s", e.statusCode, e.message)
}
