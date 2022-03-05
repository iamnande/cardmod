package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Refinement holds the schema definition for the Refinement entity.
type Refinement struct {
	ent.Schema
}

// Fields of the Refinement.
func (Refinement) Fields() []ent.Field {
	return []ent.Field{
		field.String("source").NotEmpty(),
		field.String("target").NotEmpty(),
		field.Int32("numerator").Positive(),
		field.Int32("denominator").Positive(),
	}
}

// Edges of the Refinement.
func (Refinement) Edges() []ent.Edge {
	return nil
}
