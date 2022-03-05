package models

// LimitBreak is the domain interface.
//go:generate mockgen -source limitbreak.go -destination=./mocks/limitbreak.go -package mocks
type LimitBreak interface {
	Name() string
}
