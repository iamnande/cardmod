package models

// Refinement is the domain interface.
type Refinement interface {
	Source() string
	Target() string
	Numerator() int32
	Denominator() int32
}

// RefinementFilter is used to filter refinement collection results.
type RefinementFilter interface {
	Source() string
	Target() string
}
