package models

// Card is the domain interface.
//go:generate mockgen -source card.go -destination=./mocks/card.go -package mocks
type Card interface {
	Name() string
	Level() int32
}
