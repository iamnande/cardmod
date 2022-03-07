package items

// item is the concrete implementation of the Item interface.
type item struct {
	name    string
	purpose string
}

// Name returns the item name.
func (c *item) Name() string {
	return c.name
}

// Purpose returns the item purpose.
func (c *item) Purpose() string {
	return c.purpose
}
