package models

// Item is the domain interface.
//go:generate mockgen -source item.go -destination=./mocks/item.go -package mocks
type Item interface {
	Name() string
}
