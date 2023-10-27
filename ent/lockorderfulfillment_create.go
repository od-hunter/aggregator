// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/paycrest/protocol/ent/lockorderfulfillment"
	"github.com/paycrest/protocol/ent/lockpaymentorder"
	"github.com/paycrest/protocol/ent/validatorprofile"
)

// LockOrderFulfillmentCreate is the builder for creating a LockOrderFulfillment entity.
type LockOrderFulfillmentCreate struct {
	config
	mutation *LockOrderFulfillmentMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (lofc *LockOrderFulfillmentCreate) SetCreatedAt(t time.Time) *LockOrderFulfillmentCreate {
	lofc.mutation.SetCreatedAt(t)
	return lofc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lofc *LockOrderFulfillmentCreate) SetNillableCreatedAt(t *time.Time) *LockOrderFulfillmentCreate {
	if t != nil {
		lofc.SetCreatedAt(*t)
	}
	return lofc
}

// SetUpdatedAt sets the "updated_at" field.
func (lofc *LockOrderFulfillmentCreate) SetUpdatedAt(t time.Time) *LockOrderFulfillmentCreate {
	lofc.mutation.SetUpdatedAt(t)
	return lofc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lofc *LockOrderFulfillmentCreate) SetNillableUpdatedAt(t *time.Time) *LockOrderFulfillmentCreate {
	if t != nil {
		lofc.SetUpdatedAt(*t)
	}
	return lofc
}

// SetTxID sets the "tx_id" field.
func (lofc *LockOrderFulfillmentCreate) SetTxID(s string) *LockOrderFulfillmentCreate {
	lofc.mutation.SetTxID(s)
	return lofc
}

// SetTxReceiptImage sets the "tx_receipt_image" field.
func (lofc *LockOrderFulfillmentCreate) SetTxReceiptImage(s string) *LockOrderFulfillmentCreate {
	lofc.mutation.SetTxReceiptImage(s)
	return lofc
}

// SetConfirmations sets the "confirmations" field.
func (lofc *LockOrderFulfillmentCreate) SetConfirmations(i int) *LockOrderFulfillmentCreate {
	lofc.mutation.SetConfirmations(i)
	return lofc
}

// SetNillableConfirmations sets the "confirmations" field if the given value is not nil.
func (lofc *LockOrderFulfillmentCreate) SetNillableConfirmations(i *int) *LockOrderFulfillmentCreate {
	if i != nil {
		lofc.SetConfirmations(*i)
	}
	return lofc
}

// SetValidationErrors sets the "validation_errors" field.
func (lofc *LockOrderFulfillmentCreate) SetValidationErrors(s []string) *LockOrderFulfillmentCreate {
	lofc.mutation.SetValidationErrors(s)
	return lofc
}

// SetID sets the "id" field.
func (lofc *LockOrderFulfillmentCreate) SetID(u uuid.UUID) *LockOrderFulfillmentCreate {
	lofc.mutation.SetID(u)
	return lofc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (lofc *LockOrderFulfillmentCreate) SetNillableID(u *uuid.UUID) *LockOrderFulfillmentCreate {
	if u != nil {
		lofc.SetID(*u)
	}
	return lofc
}

// SetOrderID sets the "order" edge to the LockPaymentOrder entity by ID.
func (lofc *LockOrderFulfillmentCreate) SetOrderID(id uuid.UUID) *LockOrderFulfillmentCreate {
	lofc.mutation.SetOrderID(id)
	return lofc
}

// SetOrder sets the "order" edge to the LockPaymentOrder entity.
func (lofc *LockOrderFulfillmentCreate) SetOrder(l *LockPaymentOrder) *LockOrderFulfillmentCreate {
	return lofc.SetOrderID(l.ID)
}

// AddValidatorIDs adds the "validators" edge to the ValidatorProfile entity by IDs.
func (lofc *LockOrderFulfillmentCreate) AddValidatorIDs(ids ...uuid.UUID) *LockOrderFulfillmentCreate {
	lofc.mutation.AddValidatorIDs(ids...)
	return lofc
}

// AddValidators adds the "validators" edges to the ValidatorProfile entity.
func (lofc *LockOrderFulfillmentCreate) AddValidators(v ...*ValidatorProfile) *LockOrderFulfillmentCreate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return lofc.AddValidatorIDs(ids...)
}

// Mutation returns the LockOrderFulfillmentMutation object of the builder.
func (lofc *LockOrderFulfillmentCreate) Mutation() *LockOrderFulfillmentMutation {
	return lofc.mutation
}

// Save creates the LockOrderFulfillment in the database.
func (lofc *LockOrderFulfillmentCreate) Save(ctx context.Context) (*LockOrderFulfillment, error) {
	lofc.defaults()
	return withHooks(ctx, lofc.sqlSave, lofc.mutation, lofc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lofc *LockOrderFulfillmentCreate) SaveX(ctx context.Context) *LockOrderFulfillment {
	v, err := lofc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lofc *LockOrderFulfillmentCreate) Exec(ctx context.Context) error {
	_, err := lofc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lofc *LockOrderFulfillmentCreate) ExecX(ctx context.Context) {
	if err := lofc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lofc *LockOrderFulfillmentCreate) defaults() {
	if _, ok := lofc.mutation.CreatedAt(); !ok {
		v := lockorderfulfillment.DefaultCreatedAt()
		lofc.mutation.SetCreatedAt(v)
	}
	if _, ok := lofc.mutation.UpdatedAt(); !ok {
		v := lockorderfulfillment.DefaultUpdatedAt()
		lofc.mutation.SetUpdatedAt(v)
	}
	if _, ok := lofc.mutation.Confirmations(); !ok {
		v := lockorderfulfillment.DefaultConfirmations
		lofc.mutation.SetConfirmations(v)
	}
	if _, ok := lofc.mutation.ValidationErrors(); !ok {
		v := lockorderfulfillment.DefaultValidationErrors
		lofc.mutation.SetValidationErrors(v)
	}
	if _, ok := lofc.mutation.ID(); !ok {
		v := lockorderfulfillment.DefaultID()
		lofc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lofc *LockOrderFulfillmentCreate) check() error {
	if _, ok := lofc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "LockOrderFulfillment.created_at"`)}
	}
	if _, ok := lofc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "LockOrderFulfillment.updated_at"`)}
	}
	if _, ok := lofc.mutation.TxID(); !ok {
		return &ValidationError{Name: "tx_id", err: errors.New(`ent: missing required field "LockOrderFulfillment.tx_id"`)}
	}
	if _, ok := lofc.mutation.TxReceiptImage(); !ok {
		return &ValidationError{Name: "tx_receipt_image", err: errors.New(`ent: missing required field "LockOrderFulfillment.tx_receipt_image"`)}
	}
	if _, ok := lofc.mutation.Confirmations(); !ok {
		return &ValidationError{Name: "confirmations", err: errors.New(`ent: missing required field "LockOrderFulfillment.confirmations"`)}
	}
	if _, ok := lofc.mutation.ValidationErrors(); !ok {
		return &ValidationError{Name: "validation_errors", err: errors.New(`ent: missing required field "LockOrderFulfillment.validation_errors"`)}
	}
	if _, ok := lofc.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required edge "LockOrderFulfillment.order"`)}
	}
	return nil
}

func (lofc *LockOrderFulfillmentCreate) sqlSave(ctx context.Context) (*LockOrderFulfillment, error) {
	if err := lofc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lofc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lofc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	lofc.mutation.id = &_node.ID
	lofc.mutation.done = true
	return _node, nil
}

func (lofc *LockOrderFulfillmentCreate) createSpec() (*LockOrderFulfillment, *sqlgraph.CreateSpec) {
	var (
		_node = &LockOrderFulfillment{config: lofc.config}
		_spec = sqlgraph.NewCreateSpec(lockorderfulfillment.Table, sqlgraph.NewFieldSpec(lockorderfulfillment.FieldID, field.TypeUUID))
	)
	if id, ok := lofc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := lofc.mutation.CreatedAt(); ok {
		_spec.SetField(lockorderfulfillment.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lofc.mutation.UpdatedAt(); ok {
		_spec.SetField(lockorderfulfillment.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := lofc.mutation.TxID(); ok {
		_spec.SetField(lockorderfulfillment.FieldTxID, field.TypeString, value)
		_node.TxID = value
	}
	if value, ok := lofc.mutation.TxReceiptImage(); ok {
		_spec.SetField(lockorderfulfillment.FieldTxReceiptImage, field.TypeString, value)
		_node.TxReceiptImage = value
	}
	if value, ok := lofc.mutation.Confirmations(); ok {
		_spec.SetField(lockorderfulfillment.FieldConfirmations, field.TypeInt, value)
		_node.Confirmations = value
	}
	if value, ok := lofc.mutation.ValidationErrors(); ok {
		_spec.SetField(lockorderfulfillment.FieldValidationErrors, field.TypeJSON, value)
		_node.ValidationErrors = value
	}
	if nodes := lofc.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   lockorderfulfillment.OrderTable,
			Columns: []string{lockorderfulfillment.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lockpaymentorder.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.lock_payment_order_fulfillment = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lofc.mutation.ValidatorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   lockorderfulfillment.ValidatorsTable,
			Columns: lockorderfulfillment.ValidatorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(validatorprofile.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LockOrderFulfillmentCreateBulk is the builder for creating many LockOrderFulfillment entities in bulk.
type LockOrderFulfillmentCreateBulk struct {
	config
	builders []*LockOrderFulfillmentCreate
}

// Save creates the LockOrderFulfillment entities in the database.
func (lofcb *LockOrderFulfillmentCreateBulk) Save(ctx context.Context) ([]*LockOrderFulfillment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lofcb.builders))
	nodes := make([]*LockOrderFulfillment, len(lofcb.builders))
	mutators := make([]Mutator, len(lofcb.builders))
	for i := range lofcb.builders {
		func(i int, root context.Context) {
			builder := lofcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LockOrderFulfillmentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lofcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lofcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lofcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lofcb *LockOrderFulfillmentCreateBulk) SaveX(ctx context.Context) []*LockOrderFulfillment {
	v, err := lofcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lofcb *LockOrderFulfillmentCreateBulk) Exec(ctx context.Context) error {
	_, err := lofcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lofcb *LockOrderFulfillmentCreateBulk) ExecX(ctx context.Context) {
	if err := lofcb.Exec(ctx); err != nil {
		panic(err)
	}
}
