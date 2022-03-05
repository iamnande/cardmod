package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().NotEmpty(),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return nil
}
