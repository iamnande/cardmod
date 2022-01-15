package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Calculation holds the schema definition for the Calculation entity.
type Calculation struct {
	ent.Schema
}

// Fields of the Calculation.
func (Calculation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("card_id", uuid.UUID{}),
		field.UUID("magic_id", uuid.UUID{}),
		field.Int32("card_ratio").NonNegative(),
		field.Int32("magic_ratio").NonNegative(),
	}
}

// Edges of the Calculation.
// NOTE: when https://github.com/ent/ent/issues/46 is complete, specifically edge fields (metadata on edges), migrate
// 		 to native edges instead of the junction table.
func (Calculation) Edges() []ent.Edge {
	return nil
}
