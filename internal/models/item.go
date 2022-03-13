package models

// Item is the domain interface.
type Item interface {
	Name() string
	Purpose() string
}
