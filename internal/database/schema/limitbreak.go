package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// LimitBreak holds the schema definition for the LimitBreak entity.
type LimitBreak struct {
	ent.Schema
}

// Fields of the LimitBreak.
func (LimitBreak) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().NotEmpty(),
	}
}

// Edges of the LimitBreak.
func (LimitBreak) Edges() []ent.Edge {
	return nil
}
