package errors

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
)

// APIError is the base API error model.
type APIError struct {

	// code is the response code.
	code codes.Code `json:"code"`

	// message is the client message.
	message string `json:"message"`

	// baseError is the underlying error of what went wrong.
	// NOTE: this should not be returned to callers.
	baseError error `json:"error"`
}

// Code gets the error code.
func (err *APIError) Code() codes.Code {
	return err.code
}

// Message gets the error message.
func (err *APIError) Message() string {
	return err.message
}

// BaseError gets the error's base error.
func (err *APIError) BaseError() error {
	return err.baseError
}

// Error is a stringer method on *APIError to satisfy the generic error type.
func (err *APIError) Error() string {
	return fmt.Sprintf("%s: %s", err.message, err.baseError)
}

// NewAPIError initializes a new API error.
func NewAPIError(code codes.Code, msg string, err error) error {
	return &APIError{
		code:      code,
		message:   msg,
		baseError: err,
	}
}

// IsAPIError is a helper method to determine whether or not the error is of a type *APIError.
func IsAPIError(err error) bool {
	if err != nil {
		apiErr := new(APIError)
		return errors.As(err, &apiErr)
	}
	return false
}
