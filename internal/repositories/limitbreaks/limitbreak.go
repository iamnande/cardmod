package limitbreaks

// limitbreak is the concrete implementation of the LimitBreak interface.
type limitbreak struct {
	name string
}

// Name returns the limitbreak name.
func (c *limitbreak) Name() string {
	return c.name
}
