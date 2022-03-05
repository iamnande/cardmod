package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Magic holds the schema definition for the Magic entity.
type Magic struct {
	ent.Schema
}

// Fields of the Magic.
func (Magic) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().NotEmpty(),
	}
}

// Edges of the Magic.
func (Magic) Edges() []ent.Edge {
	return nil
}
