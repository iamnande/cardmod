package models

// Card is the domain interface.
type Card interface {
	Name() string
	Level() int32
}
