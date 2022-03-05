package cerrors

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
)

// APIError is the base API error model.
type APIError struct {

	// Code is the response code.
	Code codes.Code `json:"code"`

	// Message is the client message.
	Message string `json:"message"`

	// BaseError is the underlying error of what went wrong.
	// NOTE: this should not be returned to callers.
	BaseError error `json:"error"`
}

// Error is a stringer method on *APIError to satisfy the generic error type.
func (err *APIError) Error() string {
	return fmt.Sprintf("%s: %s", err.Message, err.BaseError)
}

// IsAPIError is a helper method to determine whether or not the error is of a type *APIError.
func IsAPIError(err error) bool {
	if err != nil {
		apiErr := new(APIError)
		return errors.As(err, &apiErr)
	}
	return false
}
