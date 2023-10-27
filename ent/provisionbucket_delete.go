// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/paycrest/protocol/ent/predicate"
	"github.com/paycrest/protocol/ent/provisionbucket"
)

// ProvisionBucketDelete is the builder for deleting a ProvisionBucket entity.
type ProvisionBucketDelete struct {
	config
	hooks    []Hook
	mutation *ProvisionBucketMutation
}

// Where appends a list predicates to the ProvisionBucketDelete builder.
func (pbd *ProvisionBucketDelete) Where(ps ...predicate.ProvisionBucket) *ProvisionBucketDelete {
	pbd.mutation.Where(ps...)
	return pbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pbd *ProvisionBucketDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pbd.sqlExec, pbd.mutation, pbd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pbd *ProvisionBucketDelete) ExecX(ctx context.Context) int {
	n, err := pbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pbd *ProvisionBucketDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(provisionbucket.Table, sqlgraph.NewFieldSpec(provisionbucket.FieldID, field.TypeInt))
	if ps := pbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pbd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pbd.mutation.done = true
	return affected, err
}

// ProvisionBucketDeleteOne is the builder for deleting a single ProvisionBucket entity.
type ProvisionBucketDeleteOne struct {
	pbd *ProvisionBucketDelete
}

// Where appends a list predicates to the ProvisionBucketDelete builder.
func (pbdo *ProvisionBucketDeleteOne) Where(ps ...predicate.ProvisionBucket) *ProvisionBucketDeleteOne {
	pbdo.pbd.mutation.Where(ps...)
	return pbdo
}

// Exec executes the deletion query.
func (pbdo *ProvisionBucketDeleteOne) Exec(ctx context.Context) error {
	n, err := pbdo.pbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{provisionbucket.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pbdo *ProvisionBucketDeleteOne) ExecX(ctx context.Context) {
	if err := pbdo.Exec(ctx); err != nil {
		panic(err)
	}
}
