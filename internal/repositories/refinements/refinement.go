package refinements

// refinement is the concrete implementation of the Refinement interface.
type refinement struct {
	source      string
	target      string
	numerator   int32
	denominator int32
}

// Source returns the refinement source.
func (c *refinement) Source() string {
	return c.source
}

// Target returns the refinement target.
func (c *refinement) Target() string {
	return c.target
}

// Numerator returns the refinement numerator.
func (c *refinement) Numerator() int32 {
	return c.numerator
}

// Denominator returns the refinement denominator.
func (c *refinement) Denominator() int32 {
	return c.denominator
}
