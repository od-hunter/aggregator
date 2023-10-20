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
	"github.com/paycrest/paycrest-protocol/ent/fiatcurrency"
	"github.com/paycrest/paycrest-protocol/ent/lockpaymentorder"
	"github.com/paycrest/paycrest-protocol/ent/provideravailability"
	"github.com/paycrest/paycrest-protocol/ent/providerordertoken"
	"github.com/paycrest/paycrest-protocol/ent/providerprofile"
	"github.com/paycrest/paycrest-protocol/ent/providerrating"
	"github.com/paycrest/paycrest-protocol/ent/provisionbucket"
	"github.com/paycrest/paycrest-protocol/ent/user"
)

// ProviderProfileCreate is the builder for creating a ProviderProfile entity.
type ProviderProfileCreate struct {
	config
	mutation *ProviderProfileMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ppc *ProviderProfileCreate) SetCreatedAt(t time.Time) *ProviderProfileCreate {
	ppc.mutation.SetCreatedAt(t)
	return ppc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ppc *ProviderProfileCreate) SetNillableCreatedAt(t *time.Time) *ProviderProfileCreate {
	if t != nil {
		ppc.SetCreatedAt(*t)
	}
	return ppc
}

// SetUpdatedAt sets the "updated_at" field.
func (ppc *ProviderProfileCreate) SetUpdatedAt(t time.Time) *ProviderProfileCreate {
	ppc.mutation.SetUpdatedAt(t)
	return ppc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ppc *ProviderProfileCreate) SetNillableUpdatedAt(t *time.Time) *ProviderProfileCreate {
	if t != nil {
		ppc.SetUpdatedAt(*t)
	}
	return ppc
}

// SetTradingName sets the "trading_name" field.
func (ppc *ProviderProfileCreate) SetTradingName(s string) *ProviderProfileCreate {
	ppc.mutation.SetTradingName(s)
	return ppc
}

// SetHostIdentifier sets the "host_identifier" field.
func (ppc *ProviderProfileCreate) SetHostIdentifier(s string) *ProviderProfileCreate {
	ppc.mutation.SetHostIdentifier(s)
	return ppc
}

// SetNillableHostIdentifier sets the "host_identifier" field if the given value is not nil.
func (ppc *ProviderProfileCreate) SetNillableHostIdentifier(s *string) *ProviderProfileCreate {
	if s != nil {
		ppc.SetHostIdentifier(*s)
	}
	return ppc
}

// SetProvisionMode sets the "provision_mode" field.
func (ppc *ProviderProfileCreate) SetProvisionMode(pm providerprofile.ProvisionMode) *ProviderProfileCreate {
	ppc.mutation.SetProvisionMode(pm)
	return ppc
}

// SetNillableProvisionMode sets the "provision_mode" field if the given value is not nil.
func (ppc *ProviderProfileCreate) SetNillableProvisionMode(pm *providerprofile.ProvisionMode) *ProviderProfileCreate {
	if pm != nil {
		ppc.SetProvisionMode(*pm)
	}
	return ppc
}

// SetIsPartner sets the "is_partner" field.
func (ppc *ProviderProfileCreate) SetIsPartner(b bool) *ProviderProfileCreate {
	ppc.mutation.SetIsPartner(b)
	return ppc
}

// SetNillableIsPartner sets the "is_partner" field if the given value is not nil.
func (ppc *ProviderProfileCreate) SetNillableIsPartner(b *bool) *ProviderProfileCreate {
	if b != nil {
		ppc.SetIsPartner(*b)
	}
	return ppc
}

// SetID sets the "id" field.
func (ppc *ProviderProfileCreate) SetID(s string) *ProviderProfileCreate {
	ppc.mutation.SetID(s)
	return ppc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ppc *ProviderProfileCreate) SetNillableID(s *string) *ProviderProfileCreate {
	if s != nil {
		ppc.SetID(*s)
	}
	return ppc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ppc *ProviderProfileCreate) SetUserID(id uuid.UUID) *ProviderProfileCreate {
	ppc.mutation.SetUserID(id)
	return ppc
}

// SetUser sets the "user" edge to the User entity.
func (ppc *ProviderProfileCreate) SetUser(u *User) *ProviderProfileCreate {
	return ppc.SetUserID(u.ID)
}

// SetCurrencyID sets the "currency" edge to the FiatCurrency entity by ID.
func (ppc *ProviderProfileCreate) SetCurrencyID(id uuid.UUID) *ProviderProfileCreate {
	ppc.mutation.SetCurrencyID(id)
	return ppc
}

// SetCurrency sets the "currency" edge to the FiatCurrency entity.
func (ppc *ProviderProfileCreate) SetCurrency(f *FiatCurrency) *ProviderProfileCreate {
	return ppc.SetCurrencyID(f.ID)
}

// AddProvisionBucketIDs adds the "provision_buckets" edge to the ProvisionBucket entity by IDs.
func (ppc *ProviderProfileCreate) AddProvisionBucketIDs(ids ...int) *ProviderProfileCreate {
	ppc.mutation.AddProvisionBucketIDs(ids...)
	return ppc
}

// AddProvisionBuckets adds the "provision_buckets" edges to the ProvisionBucket entity.
func (ppc *ProviderProfileCreate) AddProvisionBuckets(p ...*ProvisionBucket) *ProviderProfileCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ppc.AddProvisionBucketIDs(ids...)
}

// AddOrderTokenIDs adds the "order_tokens" edge to the ProviderOrderToken entity by IDs.
func (ppc *ProviderProfileCreate) AddOrderTokenIDs(ids ...int) *ProviderProfileCreate {
	ppc.mutation.AddOrderTokenIDs(ids...)
	return ppc
}

// AddOrderTokens adds the "order_tokens" edges to the ProviderOrderToken entity.
func (ppc *ProviderProfileCreate) AddOrderTokens(p ...*ProviderOrderToken) *ProviderProfileCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ppc.AddOrderTokenIDs(ids...)
}

// SetAvailabilityID sets the "availability" edge to the ProviderAvailability entity by ID.
func (ppc *ProviderProfileCreate) SetAvailabilityID(id int) *ProviderProfileCreate {
	ppc.mutation.SetAvailabilityID(id)
	return ppc
}

// SetNillableAvailabilityID sets the "availability" edge to the ProviderAvailability entity by ID if the given value is not nil.
func (ppc *ProviderProfileCreate) SetNillableAvailabilityID(id *int) *ProviderProfileCreate {
	if id != nil {
		ppc = ppc.SetAvailabilityID(*id)
	}
	return ppc
}

// SetAvailability sets the "availability" edge to the ProviderAvailability entity.
func (ppc *ProviderProfileCreate) SetAvailability(p *ProviderAvailability) *ProviderProfileCreate {
	return ppc.SetAvailabilityID(p.ID)
}

// SetProviderRatingID sets the "provider_rating" edge to the ProviderRating entity by ID.
func (ppc *ProviderProfileCreate) SetProviderRatingID(id int) *ProviderProfileCreate {
	ppc.mutation.SetProviderRatingID(id)
	return ppc
}

// SetNillableProviderRatingID sets the "provider_rating" edge to the ProviderRating entity by ID if the given value is not nil.
func (ppc *ProviderProfileCreate) SetNillableProviderRatingID(id *int) *ProviderProfileCreate {
	if id != nil {
		ppc = ppc.SetProviderRatingID(*id)
	}
	return ppc
}

// SetProviderRating sets the "provider_rating" edge to the ProviderRating entity.
func (ppc *ProviderProfileCreate) SetProviderRating(p *ProviderRating) *ProviderProfileCreate {
	return ppc.SetProviderRatingID(p.ID)
}

// AddAssignedOrderIDs adds the "assigned_orders" edge to the LockPaymentOrder entity by IDs.
func (ppc *ProviderProfileCreate) AddAssignedOrderIDs(ids ...uuid.UUID) *ProviderProfileCreate {
	ppc.mutation.AddAssignedOrderIDs(ids...)
	return ppc
}

// AddAssignedOrders adds the "assigned_orders" edges to the LockPaymentOrder entity.
func (ppc *ProviderProfileCreate) AddAssignedOrders(l ...*LockPaymentOrder) *ProviderProfileCreate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ppc.AddAssignedOrderIDs(ids...)
}

// Mutation returns the ProviderProfileMutation object of the builder.
func (ppc *ProviderProfileCreate) Mutation() *ProviderProfileMutation {
	return ppc.mutation
}

// Save creates the ProviderProfile in the database.
func (ppc *ProviderProfileCreate) Save(ctx context.Context) (*ProviderProfile, error) {
	ppc.defaults()
	return withHooks(ctx, ppc.sqlSave, ppc.mutation, ppc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ppc *ProviderProfileCreate) SaveX(ctx context.Context) *ProviderProfile {
	v, err := ppc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ppc *ProviderProfileCreate) Exec(ctx context.Context) error {
	_, err := ppc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ppc *ProviderProfileCreate) ExecX(ctx context.Context) {
	if err := ppc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ppc *ProviderProfileCreate) defaults() {
	if _, ok := ppc.mutation.CreatedAt(); !ok {
		v := providerprofile.DefaultCreatedAt()
		ppc.mutation.SetCreatedAt(v)
	}
	if _, ok := ppc.mutation.UpdatedAt(); !ok {
		v := providerprofile.DefaultUpdatedAt()
		ppc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ppc.mutation.ProvisionMode(); !ok {
		v := providerprofile.DefaultProvisionMode
		ppc.mutation.SetProvisionMode(v)
	}
	if _, ok := ppc.mutation.IsPartner(); !ok {
		v := providerprofile.DefaultIsPartner
		ppc.mutation.SetIsPartner(v)
	}
	if _, ok := ppc.mutation.ID(); !ok {
		v := providerprofile.DefaultID()
		ppc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ppc *ProviderProfileCreate) check() error {
	if _, ok := ppc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ProviderProfile.created_at"`)}
	}
	if _, ok := ppc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ProviderProfile.updated_at"`)}
	}
	if _, ok := ppc.mutation.TradingName(); !ok {
		return &ValidationError{Name: "trading_name", err: errors.New(`ent: missing required field "ProviderProfile.trading_name"`)}
	}
	if v, ok := ppc.mutation.TradingName(); ok {
		if err := providerprofile.TradingNameValidator(v); err != nil {
			return &ValidationError{Name: "trading_name", err: fmt.Errorf(`ent: validator failed for field "ProviderProfile.trading_name": %w`, err)}
		}
	}
	if _, ok := ppc.mutation.ProvisionMode(); !ok {
		return &ValidationError{Name: "provision_mode", err: errors.New(`ent: missing required field "ProviderProfile.provision_mode"`)}
	}
	if v, ok := ppc.mutation.ProvisionMode(); ok {
		if err := providerprofile.ProvisionModeValidator(v); err != nil {
			return &ValidationError{Name: "provision_mode", err: fmt.Errorf(`ent: validator failed for field "ProviderProfile.provision_mode": %w`, err)}
		}
	}
	if _, ok := ppc.mutation.IsPartner(); !ok {
		return &ValidationError{Name: "is_partner", err: errors.New(`ent: missing required field "ProviderProfile.is_partner"`)}
	}
	if _, ok := ppc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "ProviderProfile.user"`)}
	}
	if _, ok := ppc.mutation.CurrencyID(); !ok {
		return &ValidationError{Name: "currency", err: errors.New(`ent: missing required edge "ProviderProfile.currency"`)}
	}
	return nil
}

func (ppc *ProviderProfileCreate) sqlSave(ctx context.Context) (*ProviderProfile, error) {
	if err := ppc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ppc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ppc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected ProviderProfile.ID type: %T", _spec.ID.Value)
		}
	}
	ppc.mutation.id = &_node.ID
	ppc.mutation.done = true
	return _node, nil
}

func (ppc *ProviderProfileCreate) createSpec() (*ProviderProfile, *sqlgraph.CreateSpec) {
	var (
		_node = &ProviderProfile{config: ppc.config}
		_spec = sqlgraph.NewCreateSpec(providerprofile.Table, sqlgraph.NewFieldSpec(providerprofile.FieldID, field.TypeString))
	)
	if id, ok := ppc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ppc.mutation.CreatedAt(); ok {
		_spec.SetField(providerprofile.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ppc.mutation.UpdatedAt(); ok {
		_spec.SetField(providerprofile.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ppc.mutation.TradingName(); ok {
		_spec.SetField(providerprofile.FieldTradingName, field.TypeString, value)
		_node.TradingName = value
	}
	if value, ok := ppc.mutation.HostIdentifier(); ok {
		_spec.SetField(providerprofile.FieldHostIdentifier, field.TypeString, value)
		_node.HostIdentifier = value
	}
	if value, ok := ppc.mutation.ProvisionMode(); ok {
		_spec.SetField(providerprofile.FieldProvisionMode, field.TypeEnum, value)
		_node.ProvisionMode = value
	}
	if value, ok := ppc.mutation.IsPartner(); ok {
		_spec.SetField(providerprofile.FieldIsPartner, field.TypeBool, value)
		_node.IsPartner = value
	}
	if nodes := ppc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   providerprofile.UserTable,
			Columns: []string{providerprofile.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_provider_profile = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ppc.mutation.CurrencyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   providerprofile.CurrencyTable,
			Columns: []string{providerprofile.CurrencyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fiatcurrency.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.fiat_currency_provider = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ppc.mutation.ProvisionBucketsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   providerprofile.ProvisionBucketsTable,
			Columns: providerprofile.ProvisionBucketsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(provisionbucket.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ppc.mutation.OrderTokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   providerprofile.OrderTokensTable,
			Columns: []string{providerprofile.OrderTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(providerordertoken.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ppc.mutation.AvailabilityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   providerprofile.AvailabilityTable,
			Columns: []string{providerprofile.AvailabilityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(provideravailability.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ppc.mutation.ProviderRatingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   providerprofile.ProviderRatingTable,
			Columns: []string{providerprofile.ProviderRatingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(providerrating.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ppc.mutation.AssignedOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   providerprofile.AssignedOrdersTable,
			Columns: []string{providerprofile.AssignedOrdersColumn},
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

// ProviderProfileCreateBulk is the builder for creating many ProviderProfile entities in bulk.
type ProviderProfileCreateBulk struct {
	config
	builders []*ProviderProfileCreate
}

// Save creates the ProviderProfile entities in the database.
func (ppcb *ProviderProfileCreateBulk) Save(ctx context.Context) ([]*ProviderProfile, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ppcb.builders))
	nodes := make([]*ProviderProfile, len(ppcb.builders))
	mutators := make([]Mutator, len(ppcb.builders))
	for i := range ppcb.builders {
		func(i int, root context.Context) {
			builder := ppcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProviderProfileMutation)
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
					_, err = mutators[i+1].Mutate(root, ppcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ppcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ppcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ppcb *ProviderProfileCreateBulk) SaveX(ctx context.Context) []*ProviderProfile {
	v, err := ppcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ppcb *ProviderProfileCreateBulk) Exec(ctx context.Context) error {
	_, err := ppcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ppcb *ProviderProfileCreateBulk) ExecX(ctx context.Context) {
	if err := ppcb.Exec(ctx); err != nil {
		panic(err)
	}
}
