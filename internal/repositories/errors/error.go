package errors

// RepositoryError is a repository error.
type RepositoryError struct {
	message string
}

// Error is the stringer method on the RepositoryError.
func (e *RepositoryError) Error() string {
	return e.message
}

// NewRepositoryError initializes a new repository error.
func NewRepositoryError(msg string) error {
	return &RepositoryError{
		message: msg,
	}
}
