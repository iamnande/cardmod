package models

// Refinement is the domain interface.
//go:generate mockgen -source refinement.go -destination=./mocks/refinement.go -package mocks
type Refinement interface {
	Source() string
	Target() string
	Numerator() int32
	Denominator() int32
}
