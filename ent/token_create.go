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
	"github.com/paycrest/protocol/ent/lockpaymentorder"
	"github.com/paycrest/protocol/ent/network"
	"github.com/paycrest/protocol/ent/paymentorder"
	"github.com/paycrest/protocol/ent/token"
)

// TokenCreate is the builder for creating a Token entity.
type TokenCreate struct {
	config
	mutation *TokenMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (tc *TokenCreate) SetCreatedAt(t time.Time) *TokenCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TokenCreate) SetNillableCreatedAt(t *time.Time) *TokenCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TokenCreate) SetUpdatedAt(t time.Time) *TokenCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TokenCreate) SetNillableUpdatedAt(t *time.Time) *TokenCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetSymbol sets the "symbol" field.
func (tc *TokenCreate) SetSymbol(s string) *TokenCreate {
	tc.mutation.SetSymbol(s)
	return tc
}

// SetContractAddress sets the "contract_address" field.
func (tc *TokenCreate) SetContractAddress(s string) *TokenCreate {
	tc.mutation.SetContractAddress(s)
	return tc
}

// SetDecimals sets the "decimals" field.
func (tc *TokenCreate) SetDecimals(i int8) *TokenCreate {
	tc.mutation.SetDecimals(i)
	return tc
}

// SetIsEnabled sets the "is_enabled" field.
func (tc *TokenCreate) SetIsEnabled(b bool) *TokenCreate {
	tc.mutation.SetIsEnabled(b)
	return tc
}

// SetNillableIsEnabled sets the "is_enabled" field if the given value is not nil.
func (tc *TokenCreate) SetNillableIsEnabled(b *bool) *TokenCreate {
	if b != nil {
		tc.SetIsEnabled(*b)
	}
	return tc
}

// SetNetworkID sets the "network" edge to the Network entity by ID.
func (tc *TokenCreate) SetNetworkID(id int) *TokenCreate {
	tc.mutation.SetNetworkID(id)
	return tc
}

// SetNetwork sets the "network" edge to the Network entity.
func (tc *TokenCreate) SetNetwork(n *Network) *TokenCreate {
	return tc.SetNetworkID(n.ID)
}

// AddPaymentOrderIDs adds the "payment_orders" edge to the PaymentOrder entity by IDs.
func (tc *TokenCreate) AddPaymentOrderIDs(ids ...uuid.UUID) *TokenCreate {
	tc.mutation.AddPaymentOrderIDs(ids...)
	return tc
}

// AddPaymentOrders adds the "payment_orders" edges to the PaymentOrder entity.
func (tc *TokenCreate) AddPaymentOrders(p ...*PaymentOrder) *TokenCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tc.AddPaymentOrderIDs(ids...)
}

// AddLockPaymentOrderIDs adds the "lock_payment_orders" edge to the LockPaymentOrder entity by IDs.
func (tc *TokenCreate) AddLockPaymentOrderIDs(ids ...uuid.UUID) *TokenCreate {
	tc.mutation.AddLockPaymentOrderIDs(ids...)
	return tc
}

// AddLockPaymentOrders adds the "lock_payment_orders" edges to the LockPaymentOrder entity.
func (tc *TokenCreate) AddLockPaymentOrders(l ...*LockPaymentOrder) *TokenCreate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return tc.AddLockPaymentOrderIDs(ids...)
}

// Mutation returns the TokenMutation object of the builder.
func (tc *TokenCreate) Mutation() *TokenMutation {
	return tc.mutation
}

// Save creates the Token in the database.
func (tc *TokenCreate) Save(ctx context.Context) (*Token, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TokenCreate) SaveX(ctx context.Context) *Token {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TokenCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TokenCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TokenCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := token.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := token.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tc.mutation.IsEnabled(); !ok {
		v := token.DefaultIsEnabled
		tc.mutation.SetIsEnabled(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TokenCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Token.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Token.updated_at"`)}
	}
	if _, ok := tc.mutation.Symbol(); !ok {
		return &ValidationError{Name: "symbol", err: errors.New(`ent: missing required field "Token.symbol"`)}
	}
	if v, ok := tc.mutation.Symbol(); ok {
		if err := token.SymbolValidator(v); err != nil {
			return &ValidationError{Name: "symbol", err: fmt.Errorf(`ent: validator failed for field "Token.symbol": %w`, err)}
		}
	}
	if _, ok := tc.mutation.ContractAddress(); !ok {
		return &ValidationError{Name: "contract_address", err: errors.New(`ent: missing required field "Token.contract_address"`)}
	}
	if v, ok := tc.mutation.ContractAddress(); ok {
		if err := token.ContractAddressValidator(v); err != nil {
			return &ValidationError{Name: "contract_address", err: fmt.Errorf(`ent: validator failed for field "Token.contract_address": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Decimals(); !ok {
		return &ValidationError{Name: "decimals", err: errors.New(`ent: missing required field "Token.decimals"`)}
	}
	if _, ok := tc.mutation.IsEnabled(); !ok {
		return &ValidationError{Name: "is_enabled", err: errors.New(`ent: missing required field "Token.is_enabled"`)}
	}
	if _, ok := tc.mutation.NetworkID(); !ok {
		return &ValidationError{Name: "network", err: errors.New(`ent: missing required edge "Token.network"`)}
	}
	return nil
}

func (tc *TokenCreate) sqlSave(ctx context.Context) (*Token, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TokenCreate) createSpec() (*Token, *sqlgraph.CreateSpec) {
	var (
		_node = &Token{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(token.Table, sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(token.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(token.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tc.mutation.Symbol(); ok {
		_spec.SetField(token.FieldSymbol, field.TypeString, value)
		_node.Symbol = value
	}
	if value, ok := tc.mutation.ContractAddress(); ok {
		_spec.SetField(token.FieldContractAddress, field.TypeString, value)
		_node.ContractAddress = value
	}
	if value, ok := tc.mutation.Decimals(); ok {
		_spec.SetField(token.FieldDecimals, field.TypeInt8, value)
		_node.Decimals = value
	}
	if value, ok := tc.mutation.IsEnabled(); ok {
		_spec.SetField(token.FieldIsEnabled, field.TypeBool, value)
		_node.IsEnabled = value
	}
	if nodes := tc.mutation.NetworkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.NetworkTable,
			Columns: []string{token.NetworkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(network.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.network_tokens = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.PaymentOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   token.PaymentOrdersTable,
			Columns: []string{token.PaymentOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymentorder.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.LockPaymentOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   token.LockPaymentOrdersTable,
			Columns: []string{token.LockPaymentOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lockpaymentorder.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TokenCreateBulk is the builder for creating many Token entities in bulk.
type TokenCreateBulk struct {
	config
	builders []*TokenCreate
}

// Save creates the Token entities in the database.
func (tcb *TokenCreateBulk) Save(ctx context.Context) ([]*Token, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Token, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TokenMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TokenCreateBulk) SaveX(ctx context.Context) []*Token {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TokenCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TokenCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
