package magics

// magic is the concrete implementation of the Magic interface.
type magic struct {
	name    string
	purpose string
}

// Name returns the magic name.
func (c *magic) Name() string {
	return c.name
}

// Purpose returns the magic purpose.
func (c *magic) Purpose() string {
	return c.purpose
}
