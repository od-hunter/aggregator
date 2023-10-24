// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/paycrest/paycrest-protocol/ent/apikey"
	"github.com/paycrest/paycrest-protocol/ent/paymentorder"
	"github.com/paycrest/paycrest-protocol/ent/predicate"
)

// APIKeyUpdate is the builder for updating APIKey entities.
type APIKeyUpdate struct {
	config
	hooks    []Hook
	mutation *APIKeyMutation
}

// Where appends a list predicates to the APIKeyUpdate builder.
func (aku *APIKeyUpdate) Where(ps ...predicate.APIKey) *APIKeyUpdate {
	aku.mutation.Where(ps...)
	return aku
}

// SetSecret sets the "secret" field.
func (aku *APIKeyUpdate) SetSecret(s string) *APIKeyUpdate {
	aku.mutation.SetSecret(s)
	return aku
}

// AddPaymentOrderIDs adds the "payment_orders" edge to the PaymentOrder entity by IDs.
func (aku *APIKeyUpdate) AddPaymentOrderIDs(ids ...uuid.UUID) *APIKeyUpdate {
	aku.mutation.AddPaymentOrderIDs(ids...)
	return aku
}

// AddPaymentOrders adds the "payment_orders" edges to the PaymentOrder entity.
func (aku *APIKeyUpdate) AddPaymentOrders(p ...*PaymentOrder) *APIKeyUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return aku.AddPaymentOrderIDs(ids...)
}

// Mutation returns the APIKeyMutation object of the builder.
func (aku *APIKeyUpdate) Mutation() *APIKeyMutation {
	return aku.mutation
}

// ClearPaymentOrders clears all "payment_orders" edges to the PaymentOrder entity.
func (aku *APIKeyUpdate) ClearPaymentOrders() *APIKeyUpdate {
	aku.mutation.ClearPaymentOrders()
	return aku
}

// RemovePaymentOrderIDs removes the "payment_orders" edge to PaymentOrder entities by IDs.
func (aku *APIKeyUpdate) RemovePaymentOrderIDs(ids ...uuid.UUID) *APIKeyUpdate {
	aku.mutation.RemovePaymentOrderIDs(ids...)
	return aku
}

// RemovePaymentOrders removes "payment_orders" edges to PaymentOrder entities.
func (aku *APIKeyUpdate) RemovePaymentOrders(p ...*PaymentOrder) *APIKeyUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return aku.RemovePaymentOrderIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aku *APIKeyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, aku.sqlSave, aku.mutation, aku.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aku *APIKeyUpdate) SaveX(ctx context.Context) int {
	affected, err := aku.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aku *APIKeyUpdate) Exec(ctx context.Context) error {
	_, err := aku.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aku *APIKeyUpdate) ExecX(ctx context.Context) {
	if err := aku.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aku *APIKeyUpdate) check() error {
	if v, ok := aku.mutation.Secret(); ok {
		if err := apikey.SecretValidator(v); err != nil {
			return &ValidationError{Name: "secret", err: fmt.Errorf(`ent: validator failed for field "APIKey.secret": %w`, err)}
		}
	}
	if _, ok := aku.mutation.SenderProfileID(); aku.mutation.SenderProfileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "APIKey.sender_profile"`)
	}
	if _, ok := aku.mutation.ProviderProfileID(); aku.mutation.ProviderProfileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "APIKey.provider_profile"`)
	}
	if _, ok := aku.mutation.ValidatorProfileID(); aku.mutation.ValidatorProfileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "APIKey.validator_profile"`)
	}
	return nil
}

func (aku *APIKeyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := aku.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(apikey.Table, apikey.Columns, sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeUUID))
	if ps := aku.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aku.mutation.Secret(); ok {
		_spec.SetField(apikey.FieldSecret, field.TypeString, value)
	}
	if aku.mutation.PaymentOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   apikey.PaymentOrdersTable,
			Columns: []string{apikey.PaymentOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymentorder.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aku.mutation.RemovedPaymentOrdersIDs(); len(nodes) > 0 && !aku.mutation.PaymentOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   apikey.PaymentOrdersTable,
			Columns: []string{apikey.PaymentOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymentorder.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aku.mutation.PaymentOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   apikey.PaymentOrdersTable,
			Columns: []string{apikey.PaymentOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymentorder.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aku.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apikey.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aku.mutation.done = true
	return n, nil
}

// APIKeyUpdateOne is the builder for updating a single APIKey entity.
type APIKeyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *APIKeyMutation
}

// SetSecret sets the "secret" field.
func (akuo *APIKeyUpdateOne) SetSecret(s string) *APIKeyUpdateOne {
	akuo.mutation.SetSecret(s)
	return akuo
}

// AddPaymentOrderIDs adds the "payment_orders" edge to the PaymentOrder entity by IDs.
func (akuo *APIKeyUpdateOne) AddPaymentOrderIDs(ids ...uuid.UUID) *APIKeyUpdateOne {
	akuo.mutation.AddPaymentOrderIDs(ids...)
	return akuo
}

// AddPaymentOrders adds the "payment_orders" edges to the PaymentOrder entity.
func (akuo *APIKeyUpdateOne) AddPaymentOrders(p ...*PaymentOrder) *APIKeyUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return akuo.AddPaymentOrderIDs(ids...)
}

// Mutation returns the APIKeyMutation object of the builder.
func (akuo *APIKeyUpdateOne) Mutation() *APIKeyMutation {
	return akuo.mutation
}

// ClearPaymentOrders clears all "payment_orders" edges to the PaymentOrder entity.
func (akuo *APIKeyUpdateOne) ClearPaymentOrders() *APIKeyUpdateOne {
	akuo.mutation.ClearPaymentOrders()
	return akuo
}

// RemovePaymentOrderIDs removes the "payment_orders" edge to PaymentOrder entities by IDs.
func (akuo *APIKeyUpdateOne) RemovePaymentOrderIDs(ids ...uuid.UUID) *APIKeyUpdateOne {
	akuo.mutation.RemovePaymentOrderIDs(ids...)
	return akuo
}

// RemovePaymentOrders removes "payment_orders" edges to PaymentOrder entities.
func (akuo *APIKeyUpdateOne) RemovePaymentOrders(p ...*PaymentOrder) *APIKeyUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return akuo.RemovePaymentOrderIDs(ids...)
}

// Where appends a list predicates to the APIKeyUpdate builder.
func (akuo *APIKeyUpdateOne) Where(ps ...predicate.APIKey) *APIKeyUpdateOne {
	akuo.mutation.Where(ps...)
	return akuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (akuo *APIKeyUpdateOne) Select(field string, fields ...string) *APIKeyUpdateOne {
	akuo.fields = append([]string{field}, fields...)
	return akuo
}

// Save executes the query and returns the updated APIKey entity.
func (akuo *APIKeyUpdateOne) Save(ctx context.Context) (*APIKey, error) {
	return withHooks(ctx, akuo.sqlSave, akuo.mutation, akuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (akuo *APIKeyUpdateOne) SaveX(ctx context.Context) *APIKey {
	node, err := akuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (akuo *APIKeyUpdateOne) Exec(ctx context.Context) error {
	_, err := akuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (akuo *APIKeyUpdateOne) ExecX(ctx context.Context) {
	if err := akuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (akuo *APIKeyUpdateOne) check() error {
	if v, ok := akuo.mutation.Secret(); ok {
		if err := apikey.SecretValidator(v); err != nil {
			return &ValidationError{Name: "secret", err: fmt.Errorf(`ent: validator failed for field "APIKey.secret": %w`, err)}
		}
	}
	if _, ok := akuo.mutation.SenderProfileID(); akuo.mutation.SenderProfileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "APIKey.sender_profile"`)
	}
	if _, ok := akuo.mutation.ProviderProfileID(); akuo.mutation.ProviderProfileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "APIKey.provider_profile"`)
	}
	if _, ok := akuo.mutation.ValidatorProfileID(); akuo.mutation.ValidatorProfileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "APIKey.validator_profile"`)
	}
	return nil
}

func (akuo *APIKeyUpdateOne) sqlSave(ctx context.Context) (_node *APIKey, err error) {
	if err := akuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(apikey.Table, apikey.Columns, sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeUUID))
	id, ok := akuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "APIKey.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := akuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apikey.FieldID)
		for _, f := range fields {
			if !apikey.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != apikey.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := akuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := akuo.mutation.Secret(); ok {
		_spec.SetField(apikey.FieldSecret, field.TypeString, value)
	}
	if akuo.mutation.PaymentOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   apikey.PaymentOrdersTable,
			Columns: []string{apikey.PaymentOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymentorder.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := akuo.mutation.RemovedPaymentOrdersIDs(); len(nodes) > 0 && !akuo.mutation.PaymentOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   apikey.PaymentOrdersTable,
			Columns: []string{apikey.PaymentOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymentorder.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := akuo.mutation.PaymentOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   apikey.PaymentOrdersTable,
			Columns: []string{apikey.PaymentOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymentorder.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &APIKey{config: akuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, akuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apikey.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	akuo.mutation.done = true
	return _node, nil
}
