// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/paycrest/protocol/ent/predicate"
	"github.com/paycrest/protocol/ent/verificationtoken"
)

// VerificationTokenUpdate is the builder for updating VerificationToken entities.
type VerificationTokenUpdate struct {
	config
	hooks    []Hook
	mutation *VerificationTokenMutation
}

// Where appends a list predicates to the VerificationTokenUpdate builder.
func (vtu *VerificationTokenUpdate) Where(ps ...predicate.VerificationToken) *VerificationTokenUpdate {
	vtu.mutation.Where(ps...)
	return vtu
}

// SetUpdatedAt sets the "updated_at" field.
func (vtu *VerificationTokenUpdate) SetUpdatedAt(t time.Time) *VerificationTokenUpdate {
	vtu.mutation.SetUpdatedAt(t)
	return vtu
}

// SetScope sets the "scope" field.
func (vtu *VerificationTokenUpdate) SetScope(v verificationtoken.Scope) *VerificationTokenUpdate {
	vtu.mutation.SetScope(v)
	return vtu
}

// SetNillableScope sets the "scope" field if the given value is not nil.
func (vtu *VerificationTokenUpdate) SetNillableScope(v *verificationtoken.Scope) *VerificationTokenUpdate {
	if v != nil {
		vtu.SetScope(*v)
	}
	return vtu
}

// Mutation returns the VerificationTokenMutation object of the builder.
func (vtu *VerificationTokenUpdate) Mutation() *VerificationTokenMutation {
	return vtu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vtu *VerificationTokenUpdate) Save(ctx context.Context) (int, error) {
	if err := vtu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, vtu.sqlSave, vtu.mutation, vtu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vtu *VerificationTokenUpdate) SaveX(ctx context.Context) int {
	affected, err := vtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vtu *VerificationTokenUpdate) Exec(ctx context.Context) error {
	_, err := vtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vtu *VerificationTokenUpdate) ExecX(ctx context.Context) {
	if err := vtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vtu *VerificationTokenUpdate) defaults() error {
	if _, ok := vtu.mutation.UpdatedAt(); !ok {
		if verificationtoken.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized verificationtoken.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := verificationtoken.UpdateDefaultUpdatedAt()
		vtu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (vtu *VerificationTokenUpdate) check() error {
	if v, ok := vtu.mutation.Scope(); ok {
		if err := verificationtoken.ScopeValidator(v); err != nil {
			return &ValidationError{Name: "scope", err: fmt.Errorf(`ent: validator failed for field "VerificationToken.scope": %w`, err)}
		}
	}
	if vtu.mutation.OwnerCleared() && len(vtu.mutation.OwnerIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "VerificationToken.owner"`)
	}
	return nil
}

func (vtu *VerificationTokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := vtu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(verificationtoken.Table, verificationtoken.Columns, sqlgraph.NewFieldSpec(verificationtoken.FieldID, field.TypeUUID))
	if ps := vtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vtu.mutation.UpdatedAt(); ok {
		_spec.SetField(verificationtoken.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := vtu.mutation.Scope(); ok {
		_spec.SetField(verificationtoken.FieldScope, field.TypeEnum, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{verificationtoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vtu.mutation.done = true
	return n, nil
}

// VerificationTokenUpdateOne is the builder for updating a single VerificationToken entity.
type VerificationTokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VerificationTokenMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (vtuo *VerificationTokenUpdateOne) SetUpdatedAt(t time.Time) *VerificationTokenUpdateOne {
	vtuo.mutation.SetUpdatedAt(t)
	return vtuo
}

// SetScope sets the "scope" field.
func (vtuo *VerificationTokenUpdateOne) SetScope(v verificationtoken.Scope) *VerificationTokenUpdateOne {
	vtuo.mutation.SetScope(v)
	return vtuo
}

// SetNillableScope sets the "scope" field if the given value is not nil.
func (vtuo *VerificationTokenUpdateOne) SetNillableScope(v *verificationtoken.Scope) *VerificationTokenUpdateOne {
	if v != nil {
		vtuo.SetScope(*v)
	}
	return vtuo
}

// Mutation returns the VerificationTokenMutation object of the builder.
func (vtuo *VerificationTokenUpdateOne) Mutation() *VerificationTokenMutation {
	return vtuo.mutation
}

// Where appends a list predicates to the VerificationTokenUpdate builder.
func (vtuo *VerificationTokenUpdateOne) Where(ps ...predicate.VerificationToken) *VerificationTokenUpdateOne {
	vtuo.mutation.Where(ps...)
	return vtuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vtuo *VerificationTokenUpdateOne) Select(field string, fields ...string) *VerificationTokenUpdateOne {
	vtuo.fields = append([]string{field}, fields...)
	return vtuo
}

// Save executes the query and returns the updated VerificationToken entity.
func (vtuo *VerificationTokenUpdateOne) Save(ctx context.Context) (*VerificationToken, error) {
	if err := vtuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, vtuo.sqlSave, vtuo.mutation, vtuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vtuo *VerificationTokenUpdateOne) SaveX(ctx context.Context) *VerificationToken {
	node, err := vtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vtuo *VerificationTokenUpdateOne) Exec(ctx context.Context) error {
	_, err := vtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vtuo *VerificationTokenUpdateOne) ExecX(ctx context.Context) {
	if err := vtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vtuo *VerificationTokenUpdateOne) defaults() error {
	if _, ok := vtuo.mutation.UpdatedAt(); !ok {
		if verificationtoken.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized verificationtoken.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := verificationtoken.UpdateDefaultUpdatedAt()
		vtuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (vtuo *VerificationTokenUpdateOne) check() error {
	if v, ok := vtuo.mutation.Scope(); ok {
		if err := verificationtoken.ScopeValidator(v); err != nil {
			return &ValidationError{Name: "scope", err: fmt.Errorf(`ent: validator failed for field "VerificationToken.scope": %w`, err)}
		}
	}
	if vtuo.mutation.OwnerCleared() && len(vtuo.mutation.OwnerIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "VerificationToken.owner"`)
	}
	return nil
}

func (vtuo *VerificationTokenUpdateOne) sqlSave(ctx context.Context) (_node *VerificationToken, err error) {
	if err := vtuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(verificationtoken.Table, verificationtoken.Columns, sqlgraph.NewFieldSpec(verificationtoken.FieldID, field.TypeUUID))
	id, ok := vtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "VerificationToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, verificationtoken.FieldID)
		for _, f := range fields {
			if !verificationtoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != verificationtoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vtuo.mutation.UpdatedAt(); ok {
		_spec.SetField(verificationtoken.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := vtuo.mutation.Scope(); ok {
		_spec.SetField(verificationtoken.FieldScope, field.TypeEnum, value)
	}
	_node = &VerificationToken{config: vtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{verificationtoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vtuo.mutation.done = true
	return _node, nil
}
