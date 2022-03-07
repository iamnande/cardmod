package cards

// card is the concrete implementation of the Card interface.
type card struct {
	name  string
	level int32
}

// Name returns the card name.
func (c *card) Name() string {
	return c.name
}

// Level returns the card level.
func (c *card) Level() int32 {
	return c.level
}
