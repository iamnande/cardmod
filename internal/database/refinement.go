// Code generated by entc, DO NOT EDIT.

package database

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/iamnande/cardmod/internal/database/refinement"
)

// Refinement is the model entity for the Refinement schema.
type Refinement struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Source holds the value of the "source" field.
	Source string `json:"source,omitempty"`
	// Target holds the value of the "target" field.
	Target string `json:"target,omitempty"`
	// Numerator holds the value of the "numerator" field.
	Numerator int32 `json:"numerator,omitempty"`
	// Denominator holds the value of the "denominator" field.
	Denominator int32 `json:"denominator,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Refinement) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case refinement.FieldID, refinement.FieldNumerator, refinement.FieldDenominator:
			values[i] = new(sql.NullInt64)
		case refinement.FieldSource, refinement.FieldTarget:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Refinement", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Refinement fields.
func (r *Refinement) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case refinement.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case refinement.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				r.Source = value.String
			}
		case refinement.FieldTarget:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field target", values[i])
			} else if value.Valid {
				r.Target = value.String
			}
		case refinement.FieldNumerator:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field numerator", values[i])
			} else if value.Valid {
				r.Numerator = int32(value.Int64)
			}
		case refinement.FieldDenominator:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field denominator", values[i])
			} else if value.Valid {
				r.Denominator = int32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Refinement.
// Note that you need to call Refinement.Unwrap() before calling this method if this Refinement
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Refinement) Update() *RefinementUpdateOne {
	return (&RefinementClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Refinement entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Refinement) Unwrap() *Refinement {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("database: Refinement is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Refinement) String() string {
	var builder strings.Builder
	builder.WriteString("Refinement(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", source=")
	builder.WriteString(r.Source)
	builder.WriteString(", target=")
	builder.WriteString(r.Target)
	builder.WriteString(", numerator=")
	builder.WriteString(fmt.Sprintf("%v", r.Numerator))
	builder.WriteString(", denominator=")
	builder.WriteString(fmt.Sprintf("%v", r.Denominator))
	builder.WriteByte(')')
	return builder.String()
}

// Refinements is a parsable slice of Refinement.
type Refinements []*Refinement

func (r Refinements) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}