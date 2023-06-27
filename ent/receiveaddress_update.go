// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/paycrest/paycrest-protocol/ent/predicate"
	"github.com/paycrest/paycrest-protocol/ent/receiveaddress"
)

// ReceiveAddressUpdate is the builder for updating ReceiveAddress entities.
type ReceiveAddressUpdate struct {
	config
	hooks    []Hook
	mutation *ReceiveAddressMutation
}

// Where appends a list predicates to the ReceiveAddressUpdate builder.
func (rau *ReceiveAddressUpdate) Where(ps ...predicate.ReceiveAddress) *ReceiveAddressUpdate {
	rau.mutation.Where(ps...)
	return rau
}

// SetAddress sets the "address" field.
func (rau *ReceiveAddressUpdate) SetAddress(s string) *ReceiveAddressUpdate {
	rau.mutation.SetAddress(s)
	return rau
}

// SetAccountIndex sets the "accountIndex" field.
func (rau *ReceiveAddressUpdate) SetAccountIndex(i int) *ReceiveAddressUpdate {
	rau.mutation.ResetAccountIndex()
	rau.mutation.SetAccountIndex(i)
	return rau
}

// AddAccountIndex adds i to the "accountIndex" field.
func (rau *ReceiveAddressUpdate) AddAccountIndex(i int) *ReceiveAddressUpdate {
	rau.mutation.AddAccountIndex(i)
	return rau
}

// SetStatus sets the "status" field.
func (rau *ReceiveAddressUpdate) SetStatus(r receiveaddress.Status) *ReceiveAddressUpdate {
	rau.mutation.SetStatus(r)
	return rau
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rau *ReceiveAddressUpdate) SetNillableStatus(r *receiveaddress.Status) *ReceiveAddressUpdate {
	if r != nil {
		rau.SetStatus(*r)
	}
	return rau
}

// Mutation returns the ReceiveAddressMutation object of the builder.
func (rau *ReceiveAddressUpdate) Mutation() *ReceiveAddressMutation {
	return rau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rau *ReceiveAddressUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rau.sqlSave, rau.mutation, rau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rau *ReceiveAddressUpdate) SaveX(ctx context.Context) int {
	affected, err := rau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rau *ReceiveAddressUpdate) Exec(ctx context.Context) error {
	_, err := rau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rau *ReceiveAddressUpdate) ExecX(ctx context.Context) {
	if err := rau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rau *ReceiveAddressUpdate) check() error {
	if v, ok := rau.mutation.Status(); ok {
		if err := receiveaddress.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ReceiveAddress.status": %w`, err)}
		}
	}
	return nil
}

func (rau *ReceiveAddressUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(receiveaddress.Table, receiveaddress.Columns, sqlgraph.NewFieldSpec(receiveaddress.FieldID, field.TypeInt))
	if ps := rau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rau.mutation.Address(); ok {
		_spec.SetField(receiveaddress.FieldAddress, field.TypeString, value)
	}
	if value, ok := rau.mutation.AccountIndex(); ok {
		_spec.SetField(receiveaddress.FieldAccountIndex, field.TypeInt, value)
	}
	if value, ok := rau.mutation.AddedAccountIndex(); ok {
		_spec.AddField(receiveaddress.FieldAccountIndex, field.TypeInt, value)
	}
	if value, ok := rau.mutation.Status(); ok {
		_spec.SetField(receiveaddress.FieldStatus, field.TypeEnum, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{receiveaddress.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rau.mutation.done = true
	return n, nil
}

// ReceiveAddressUpdateOne is the builder for updating a single ReceiveAddress entity.
type ReceiveAddressUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReceiveAddressMutation
}

// SetAddress sets the "address" field.
func (rauo *ReceiveAddressUpdateOne) SetAddress(s string) *ReceiveAddressUpdateOne {
	rauo.mutation.SetAddress(s)
	return rauo
}

// SetAccountIndex sets the "accountIndex" field.
func (rauo *ReceiveAddressUpdateOne) SetAccountIndex(i int) *ReceiveAddressUpdateOne {
	rauo.mutation.ResetAccountIndex()
	rauo.mutation.SetAccountIndex(i)
	return rauo
}

// AddAccountIndex adds i to the "accountIndex" field.
func (rauo *ReceiveAddressUpdateOne) AddAccountIndex(i int) *ReceiveAddressUpdateOne {
	rauo.mutation.AddAccountIndex(i)
	return rauo
}

// SetStatus sets the "status" field.
func (rauo *ReceiveAddressUpdateOne) SetStatus(r receiveaddress.Status) *ReceiveAddressUpdateOne {
	rauo.mutation.SetStatus(r)
	return rauo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rauo *ReceiveAddressUpdateOne) SetNillableStatus(r *receiveaddress.Status) *ReceiveAddressUpdateOne {
	if r != nil {
		rauo.SetStatus(*r)
	}
	return rauo
}

// Mutation returns the ReceiveAddressMutation object of the builder.
func (rauo *ReceiveAddressUpdateOne) Mutation() *ReceiveAddressMutation {
	return rauo.mutation
}

// Where appends a list predicates to the ReceiveAddressUpdate builder.
func (rauo *ReceiveAddressUpdateOne) Where(ps ...predicate.ReceiveAddress) *ReceiveAddressUpdateOne {
	rauo.mutation.Where(ps...)
	return rauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rauo *ReceiveAddressUpdateOne) Select(field string, fields ...string) *ReceiveAddressUpdateOne {
	rauo.fields = append([]string{field}, fields...)
	return rauo
}

// Save executes the query and returns the updated ReceiveAddress entity.
func (rauo *ReceiveAddressUpdateOne) Save(ctx context.Context) (*ReceiveAddress, error) {
	return withHooks(ctx, rauo.sqlSave, rauo.mutation, rauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rauo *ReceiveAddressUpdateOne) SaveX(ctx context.Context) *ReceiveAddress {
	node, err := rauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rauo *ReceiveAddressUpdateOne) Exec(ctx context.Context) error {
	_, err := rauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rauo *ReceiveAddressUpdateOne) ExecX(ctx context.Context) {
	if err := rauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rauo *ReceiveAddressUpdateOne) check() error {
	if v, ok := rauo.mutation.Status(); ok {
		if err := receiveaddress.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ReceiveAddress.status": %w`, err)}
		}
	}
	return nil
}

func (rauo *ReceiveAddressUpdateOne) sqlSave(ctx context.Context) (_node *ReceiveAddress, err error) {
	if err := rauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(receiveaddress.Table, receiveaddress.Columns, sqlgraph.NewFieldSpec(receiveaddress.FieldID, field.TypeInt))
	id, ok := rauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ReceiveAddress.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, receiveaddress.FieldID)
		for _, f := range fields {
			if !receiveaddress.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != receiveaddress.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rauo.mutation.Address(); ok {
		_spec.SetField(receiveaddress.FieldAddress, field.TypeString, value)
	}
	if value, ok := rauo.mutation.AccountIndex(); ok {
		_spec.SetField(receiveaddress.FieldAccountIndex, field.TypeInt, value)
	}
	if value, ok := rauo.mutation.AddedAccountIndex(); ok {
		_spec.AddField(receiveaddress.FieldAccountIndex, field.TypeInt, value)
	}
	if value, ok := rauo.mutation.Status(); ok {
		_spec.SetField(receiveaddress.FieldStatus, field.TypeEnum, value)
	}
	_node = &ReceiveAddress{config: rauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{receiveaddress.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rauo.mutation.done = true
	return _node, nil
}
