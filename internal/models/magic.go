package models

// Magic is the domain interface.
//go:generate mockgen -source magic.go -destination=./mocks/magic.go -package mocks
type Magic interface {
	Name() string
	Purpose() string
}
