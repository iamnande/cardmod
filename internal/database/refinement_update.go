// Code generated by entc, DO NOT EDIT.

package database

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/iamnande/cardmod/internal/database/predicate"
	"github.com/iamnande/cardmod/internal/database/refinement"
)

// RefinementUpdate is the builder for updating Refinement entities.
type RefinementUpdate struct {
	config
	hooks    []Hook
	mutation *RefinementMutation
}

// Where appends a list predicates to the RefinementUpdate builder.
func (ru *RefinementUpdate) Where(ps ...predicate.Refinement) *RefinementUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetSource sets the "source" field.
func (ru *RefinementUpdate) SetSource(s string) *RefinementUpdate {
	ru.mutation.SetSource(s)
	return ru
}

// SetTarget sets the "target" field.
func (ru *RefinementUpdate) SetTarget(s string) *RefinementUpdate {
	ru.mutation.SetTarget(s)
	return ru
}

// SetNumerator sets the "numerator" field.
func (ru *RefinementUpdate) SetNumerator(i int32) *RefinementUpdate {
	ru.mutation.ResetNumerator()
	ru.mutation.SetNumerator(i)
	return ru
}

// AddNumerator adds i to the "numerator" field.
func (ru *RefinementUpdate) AddNumerator(i int32) *RefinementUpdate {
	ru.mutation.AddNumerator(i)
	return ru
}

// SetDenominator sets the "denominator" field.
func (ru *RefinementUpdate) SetDenominator(i int32) *RefinementUpdate {
	ru.mutation.ResetDenominator()
	ru.mutation.SetDenominator(i)
	return ru
}

// AddDenominator adds i to the "denominator" field.
func (ru *RefinementUpdate) AddDenominator(i int32) *RefinementUpdate {
	ru.mutation.AddDenominator(i)
	return ru
}

// Mutation returns the RefinementMutation object of the builder.
func (ru *RefinementUpdate) Mutation() *RefinementMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RefinementUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RefinementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("database: uninitialized hook (forgotten import database/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RefinementUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RefinementUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RefinementUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RefinementUpdate) check() error {
	if v, ok := ru.mutation.Source(); ok {
		if err := refinement.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf("database: validator failed for field \"source\": %w", err)}
		}
	}
	if v, ok := ru.mutation.Target(); ok {
		if err := refinement.TargetValidator(v); err != nil {
			return &ValidationError{Name: "target", err: fmt.Errorf("database: validator failed for field \"target\": %w", err)}
		}
	}
	if v, ok := ru.mutation.Numerator(); ok {
		if err := refinement.NumeratorValidator(v); err != nil {
			return &ValidationError{Name: "numerator", err: fmt.Errorf("database: validator failed for field \"numerator\": %w", err)}
		}
	}
	if v, ok := ru.mutation.Denominator(); ok {
		if err := refinement.DenominatorValidator(v); err != nil {
			return &ValidationError{Name: "denominator", err: fmt.Errorf("database: validator failed for field \"denominator\": %w", err)}
		}
	}
	return nil
}

func (ru *RefinementUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   refinement.Table,
			Columns: refinement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: refinement.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Source(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: refinement.FieldSource,
		})
	}
	if value, ok := ru.mutation.Target(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: refinement.FieldTarget,
		})
	}
	if value, ok := ru.mutation.Numerator(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: refinement.FieldNumerator,
		})
	}
	if value, ok := ru.mutation.AddedNumerator(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: refinement.FieldNumerator,
		})
	}
	if value, ok := ru.mutation.Denominator(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: refinement.FieldDenominator,
		})
	}
	if value, ok := ru.mutation.AddedDenominator(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: refinement.FieldDenominator,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{refinement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RefinementUpdateOne is the builder for updating a single Refinement entity.
type RefinementUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RefinementMutation
}

// SetSource sets the "source" field.
func (ruo *RefinementUpdateOne) SetSource(s string) *RefinementUpdateOne {
	ruo.mutation.SetSource(s)
	return ruo
}

// SetTarget sets the "target" field.
func (ruo *RefinementUpdateOne) SetTarget(s string) *RefinementUpdateOne {
	ruo.mutation.SetTarget(s)
	return ruo
}

// SetNumerator sets the "numerator" field.
func (ruo *RefinementUpdateOne) SetNumerator(i int32) *RefinementUpdateOne {
	ruo.mutation.ResetNumerator()
	ruo.mutation.SetNumerator(i)
	return ruo
}

// AddNumerator adds i to the "numerator" field.
func (ruo *RefinementUpdateOne) AddNumerator(i int32) *RefinementUpdateOne {
	ruo.mutation.AddNumerator(i)
	return ruo
}

// SetDenominator sets the "denominator" field.
func (ruo *RefinementUpdateOne) SetDenominator(i int32) *RefinementUpdateOne {
	ruo.mutation.ResetDenominator()
	ruo.mutation.SetDenominator(i)
	return ruo
}

// AddDenominator adds i to the "denominator" field.
func (ruo *RefinementUpdateOne) AddDenominator(i int32) *RefinementUpdateOne {
	ruo.mutation.AddDenominator(i)
	return ruo
}

// Mutation returns the RefinementMutation object of the builder.
func (ruo *RefinementUpdateOne) Mutation() *RefinementMutation {
	return ruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RefinementUpdateOne) Select(field string, fields ...string) *RefinementUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Refinement entity.
func (ruo *RefinementUpdateOne) Save(ctx context.Context) (*Refinement, error) {
	var (
		err  error
		node *Refinement
	)
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RefinementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("database: uninitialized hook (forgotten import database/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RefinementUpdateOne) SaveX(ctx context.Context) *Refinement {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RefinementUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RefinementUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RefinementUpdateOne) check() error {
	if v, ok := ruo.mutation.Source(); ok {
		if err := refinement.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf("database: validator failed for field \"source\": %w", err)}
		}
	}
	if v, ok := ruo.mutation.Target(); ok {
		if err := refinement.TargetValidator(v); err != nil {
			return &ValidationError{Name: "target", err: fmt.Errorf("database: validator failed for field \"target\": %w", err)}
		}
	}
	if v, ok := ruo.mutation.Numerator(); ok {
		if err := refinement.NumeratorValidator(v); err != nil {
			return &ValidationError{Name: "numerator", err: fmt.Errorf("database: validator failed for field \"numerator\": %w", err)}
		}
	}
	if v, ok := ruo.mutation.Denominator(); ok {
		if err := refinement.DenominatorValidator(v); err != nil {
			return &ValidationError{Name: "denominator", err: fmt.Errorf("database: validator failed for field \"denominator\": %w", err)}
		}
	}
	return nil
}

func (ruo *RefinementUpdateOne) sqlSave(ctx context.Context) (_node *Refinement, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   refinement.Table,
			Columns: refinement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: refinement.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Refinement.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, refinement.FieldID)
		for _, f := range fields {
			if !refinement.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("database: invalid field %q for query", f)}
			}
			if f != refinement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Source(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: refinement.FieldSource,
		})
	}
	if value, ok := ruo.mutation.Target(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: refinement.FieldTarget,
		})
	}
	if value, ok := ruo.mutation.Numerator(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: refinement.FieldNumerator,
		})
	}
	if value, ok := ruo.mutation.AddedNumerator(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: refinement.FieldNumerator,
		})
	}
	if value, ok := ruo.mutation.Denominator(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: refinement.FieldDenominator,
		})
	}
	if value, ok := ruo.mutation.AddedDenominator(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: refinement.FieldDenominator,
		})
	}
	_node = &Refinement{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{refinement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}