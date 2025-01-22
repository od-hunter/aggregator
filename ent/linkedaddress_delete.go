// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/paycrest/aggregator/ent/linkedaddress"
	"github.com/paycrest/aggregator/ent/predicate"
)

// LinkedAddressDelete is the builder for deleting a LinkedAddress entity.
type LinkedAddressDelete struct {
	config
	hooks    []Hook
	mutation *LinkedAddressMutation
}

// Where appends a list predicates to the LinkedAddressDelete builder.
func (lad *LinkedAddressDelete) Where(ps ...predicate.LinkedAddress) *LinkedAddressDelete {
	lad.mutation.Where(ps...)
	return lad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lad *LinkedAddressDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, lad.sqlExec, lad.mutation, lad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (lad *LinkedAddressDelete) ExecX(ctx context.Context) int {
	n, err := lad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lad *LinkedAddressDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(linkedaddress.Table, sqlgraph.NewFieldSpec(linkedaddress.FieldID, field.TypeInt))
	if ps := lad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, lad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	lad.mutation.done = true
	return affected, err
}

// LinkedAddressDeleteOne is the builder for deleting a single LinkedAddress entity.
type LinkedAddressDeleteOne struct {
	lad *LinkedAddressDelete
}

// Where appends a list predicates to the LinkedAddressDelete builder.
func (lado *LinkedAddressDeleteOne) Where(ps ...predicate.LinkedAddress) *LinkedAddressDeleteOne {
	lado.lad.mutation.Where(ps...)
	return lado
}

// Exec executes the deletion query.
func (lado *LinkedAddressDeleteOne) Exec(ctx context.Context) error {
	n, err := lado.lad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{linkedaddress.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lado *LinkedAddressDeleteOne) ExecX(ctx context.Context) {
	if err := lado.Exec(ctx); err != nil {
		panic(err)
	}
}
