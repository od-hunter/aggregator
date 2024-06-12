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
	"github.com/paycrest/protocol/ent/fiatcurrency"
	"github.com/paycrest/protocol/ent/institution"
	"github.com/paycrest/protocol/ent/providerprofile"
	"github.com/paycrest/protocol/ent/provisionbucket"
	"github.com/shopspring/decimal"
)

// FiatCurrencyCreate is the builder for creating a FiatCurrency entity.
type FiatCurrencyCreate struct {
	config
	mutation *FiatCurrencyMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (fcc *FiatCurrencyCreate) SetCreatedAt(t time.Time) *FiatCurrencyCreate {
	fcc.mutation.SetCreatedAt(t)
	return fcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fcc *FiatCurrencyCreate) SetNillableCreatedAt(t *time.Time) *FiatCurrencyCreate {
	if t != nil {
		fcc.SetCreatedAt(*t)
	}
	return fcc
}

// SetUpdatedAt sets the "updated_at" field.
func (fcc *FiatCurrencyCreate) SetUpdatedAt(t time.Time) *FiatCurrencyCreate {
	fcc.mutation.SetUpdatedAt(t)
	return fcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fcc *FiatCurrencyCreate) SetNillableUpdatedAt(t *time.Time) *FiatCurrencyCreate {
	if t != nil {
		fcc.SetUpdatedAt(*t)
	}
	return fcc
}

// SetCode sets the "code" field.
func (fcc *FiatCurrencyCreate) SetCode(s string) *FiatCurrencyCreate {
	fcc.mutation.SetCode(s)
	return fcc
}

// SetShortName sets the "short_name" field.
func (fcc *FiatCurrencyCreate) SetShortName(s string) *FiatCurrencyCreate {
	fcc.mutation.SetShortName(s)
	return fcc
}

// SetDecimals sets the "decimals" field.
func (fcc *FiatCurrencyCreate) SetDecimals(i int) *FiatCurrencyCreate {
	fcc.mutation.SetDecimals(i)
	return fcc
}

// SetNillableDecimals sets the "decimals" field if the given value is not nil.
func (fcc *FiatCurrencyCreate) SetNillableDecimals(i *int) *FiatCurrencyCreate {
	if i != nil {
		fcc.SetDecimals(*i)
	}
	return fcc
}

// SetSymbol sets the "symbol" field.
func (fcc *FiatCurrencyCreate) SetSymbol(s string) *FiatCurrencyCreate {
	fcc.mutation.SetSymbol(s)
	return fcc
}

// SetName sets the "name" field.
func (fcc *FiatCurrencyCreate) SetName(s string) *FiatCurrencyCreate {
	fcc.mutation.SetName(s)
	return fcc
}

// SetMarketRate sets the "market_rate" field.
func (fcc *FiatCurrencyCreate) SetMarketRate(d decimal.Decimal) *FiatCurrencyCreate {
	fcc.mutation.SetMarketRate(d)
	return fcc
}

// SetIsEnabled sets the "is_enabled" field.
func (fcc *FiatCurrencyCreate) SetIsEnabled(b bool) *FiatCurrencyCreate {
	fcc.mutation.SetIsEnabled(b)
	return fcc
}

// SetNillableIsEnabled sets the "is_enabled" field if the given value is not nil.
func (fcc *FiatCurrencyCreate) SetNillableIsEnabled(b *bool) *FiatCurrencyCreate {
	if b != nil {
		fcc.SetIsEnabled(*b)
	}
	return fcc
}

// SetID sets the "id" field.
func (fcc *FiatCurrencyCreate) SetID(u uuid.UUID) *FiatCurrencyCreate {
	fcc.mutation.SetID(u)
	return fcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fcc *FiatCurrencyCreate) SetNillableID(u *uuid.UUID) *FiatCurrencyCreate {
	if u != nil {
		fcc.SetID(*u)
	}
	return fcc
}

// AddProviderIDs adds the "providers" edge to the ProviderProfile entity by IDs.
func (fcc *FiatCurrencyCreate) AddProviderIDs(ids ...string) *FiatCurrencyCreate {
	fcc.mutation.AddProviderIDs(ids...)
	return fcc
}

// AddProviders adds the "providers" edges to the ProviderProfile entity.
func (fcc *FiatCurrencyCreate) AddProviders(p ...*ProviderProfile) *FiatCurrencyCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return fcc.AddProviderIDs(ids...)
}

// AddProvisionBucketIDs adds the "provision_buckets" edge to the ProvisionBucket entity by IDs.
func (fcc *FiatCurrencyCreate) AddProvisionBucketIDs(ids ...int) *FiatCurrencyCreate {
	fcc.mutation.AddProvisionBucketIDs(ids...)
	return fcc
}

// AddProvisionBuckets adds the "provision_buckets" edges to the ProvisionBucket entity.
func (fcc *FiatCurrencyCreate) AddProvisionBuckets(p ...*ProvisionBucket) *FiatCurrencyCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return fcc.AddProvisionBucketIDs(ids...)
}

// AddInstitutionIDs adds the "institutions" edge to the Institution entity by IDs.
func (fcc *FiatCurrencyCreate) AddInstitutionIDs(ids ...int) *FiatCurrencyCreate {
	fcc.mutation.AddInstitutionIDs(ids...)
	return fcc
}

// AddInstitutions adds the "institutions" edges to the Institution entity.
func (fcc *FiatCurrencyCreate) AddInstitutions(i ...*Institution) *FiatCurrencyCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return fcc.AddInstitutionIDs(ids...)
}

// Mutation returns the FiatCurrencyMutation object of the builder.
func (fcc *FiatCurrencyCreate) Mutation() *FiatCurrencyMutation {
	return fcc.mutation
}

// Save creates the FiatCurrency in the database.
func (fcc *FiatCurrencyCreate) Save(ctx context.Context) (*FiatCurrency, error) {
	fcc.defaults()
	return withHooks(ctx, fcc.sqlSave, fcc.mutation, fcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fcc *FiatCurrencyCreate) SaveX(ctx context.Context) *FiatCurrency {
	v, err := fcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcc *FiatCurrencyCreate) Exec(ctx context.Context) error {
	_, err := fcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcc *FiatCurrencyCreate) ExecX(ctx context.Context) {
	if err := fcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fcc *FiatCurrencyCreate) defaults() {
	if _, ok := fcc.mutation.CreatedAt(); !ok {
		v := fiatcurrency.DefaultCreatedAt()
		fcc.mutation.SetCreatedAt(v)
	}
	if _, ok := fcc.mutation.UpdatedAt(); !ok {
		v := fiatcurrency.DefaultUpdatedAt()
		fcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fcc.mutation.Decimals(); !ok {
		v := fiatcurrency.DefaultDecimals
		fcc.mutation.SetDecimals(v)
	}
	if _, ok := fcc.mutation.IsEnabled(); !ok {
		v := fiatcurrency.DefaultIsEnabled
		fcc.mutation.SetIsEnabled(v)
	}
	if _, ok := fcc.mutation.ID(); !ok {
		v := fiatcurrency.DefaultID()
		fcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fcc *FiatCurrencyCreate) check() error {
	if _, ok := fcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "FiatCurrency.created_at"`)}
	}
	if _, ok := fcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "FiatCurrency.updated_at"`)}
	}
	if _, ok := fcc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "FiatCurrency.code"`)}
	}
	if _, ok := fcc.mutation.ShortName(); !ok {
		return &ValidationError{Name: "short_name", err: errors.New(`ent: missing required field "FiatCurrency.short_name"`)}
	}
	if _, ok := fcc.mutation.Decimals(); !ok {
		return &ValidationError{Name: "decimals", err: errors.New(`ent: missing required field "FiatCurrency.decimals"`)}
	}
	if _, ok := fcc.mutation.Symbol(); !ok {
		return &ValidationError{Name: "symbol", err: errors.New(`ent: missing required field "FiatCurrency.symbol"`)}
	}
	if _, ok := fcc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "FiatCurrency.name"`)}
	}
	if _, ok := fcc.mutation.MarketRate(); !ok {
		return &ValidationError{Name: "market_rate", err: errors.New(`ent: missing required field "FiatCurrency.market_rate"`)}
	}
	if _, ok := fcc.mutation.IsEnabled(); !ok {
		return &ValidationError{Name: "is_enabled", err: errors.New(`ent: missing required field "FiatCurrency.is_enabled"`)}
	}
	return nil
}

func (fcc *FiatCurrencyCreate) sqlSave(ctx context.Context) (*FiatCurrency, error) {
	if err := fcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fcc.driver, _spec); err != nil {
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
	fcc.mutation.id = &_node.ID
	fcc.mutation.done = true
	return _node, nil
}

func (fcc *FiatCurrencyCreate) createSpec() (*FiatCurrency, *sqlgraph.CreateSpec) {
	var (
		_node = &FiatCurrency{config: fcc.config}
		_spec = sqlgraph.NewCreateSpec(fiatcurrency.Table, sqlgraph.NewFieldSpec(fiatcurrency.FieldID, field.TypeUUID))
	)
	if id, ok := fcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := fcc.mutation.CreatedAt(); ok {
		_spec.SetField(fiatcurrency.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := fcc.mutation.UpdatedAt(); ok {
		_spec.SetField(fiatcurrency.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := fcc.mutation.Code(); ok {
		_spec.SetField(fiatcurrency.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := fcc.mutation.ShortName(); ok {
		_spec.SetField(fiatcurrency.FieldShortName, field.TypeString, value)
		_node.ShortName = value
	}
	if value, ok := fcc.mutation.Decimals(); ok {
		_spec.SetField(fiatcurrency.FieldDecimals, field.TypeInt, value)
		_node.Decimals = value
	}
	if value, ok := fcc.mutation.Symbol(); ok {
		_spec.SetField(fiatcurrency.FieldSymbol, field.TypeString, value)
		_node.Symbol = value
	}
	if value, ok := fcc.mutation.Name(); ok {
		_spec.SetField(fiatcurrency.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := fcc.mutation.MarketRate(); ok {
		_spec.SetField(fiatcurrency.FieldMarketRate, field.TypeFloat64, value)
		_node.MarketRate = value
	}
	if value, ok := fcc.mutation.IsEnabled(); ok {
		_spec.SetField(fiatcurrency.FieldIsEnabled, field.TypeBool, value)
		_node.IsEnabled = value
	}
	if nodes := fcc.mutation.ProvidersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fiatcurrency.ProvidersTable,
			Columns: []string{fiatcurrency.ProvidersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(providerprofile.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fcc.mutation.ProvisionBucketsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fiatcurrency.ProvisionBucketsTable,
			Columns: []string{fiatcurrency.ProvisionBucketsColumn},
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
	if nodes := fcc.mutation.InstitutionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fiatcurrency.InstitutionsTable,
			Columns: []string{fiatcurrency.InstitutionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(institution.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FiatCurrencyCreateBulk is the builder for creating many FiatCurrency entities in bulk.
type FiatCurrencyCreateBulk struct {
	config
	builders []*FiatCurrencyCreate
}

// Save creates the FiatCurrency entities in the database.
func (fccb *FiatCurrencyCreateBulk) Save(ctx context.Context) ([]*FiatCurrency, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fccb.builders))
	nodes := make([]*FiatCurrency, len(fccb.builders))
	mutators := make([]Mutator, len(fccb.builders))
	for i := range fccb.builders {
		func(i int, root context.Context) {
			builder := fccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FiatCurrencyMutation)
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
					_, err = mutators[i+1].Mutate(root, fccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fccb *FiatCurrencyCreateBulk) SaveX(ctx context.Context) []*FiatCurrency {
	v, err := fccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fccb *FiatCurrencyCreateBulk) Exec(ctx context.Context) error {
	_, err := fccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fccb *FiatCurrencyCreateBulk) ExecX(ctx context.Context) {
	if err := fccb.Exec(ctx); err != nil {
		panic(err)
	}
}
