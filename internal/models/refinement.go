package models

// Refinement is the domain interface.
type Refinement interface {
	Source() string
	Target() string
	Numerator() int32
	Denominator() int32
}
