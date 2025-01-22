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
	"github.com/google/uuid"
	"github.com/paycrest/aggregator/ent/linkedaddress"
	"github.com/paycrest/aggregator/ent/paymentorder"
)

// LinkedAddressCreate is the builder for creating a LinkedAddress entity.
type LinkedAddressCreate struct {
	config
	mutation *LinkedAddressMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (lac *LinkedAddressCreate) SetCreatedAt(t time.Time) *LinkedAddressCreate {
	lac.mutation.SetCreatedAt(t)
	return lac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lac *LinkedAddressCreate) SetNillableCreatedAt(t *time.Time) *LinkedAddressCreate {
	if t != nil {
		lac.SetCreatedAt(*t)
	}
	return lac
}

// SetUpdatedAt sets the "updated_at" field.
func (lac *LinkedAddressCreate) SetUpdatedAt(t time.Time) *LinkedAddressCreate {
	lac.mutation.SetUpdatedAt(t)
	return lac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lac *LinkedAddressCreate) SetNillableUpdatedAt(t *time.Time) *LinkedAddressCreate {
	if t != nil {
		lac.SetUpdatedAt(*t)
	}
	return lac
}

// SetAddress sets the "address" field.
func (lac *LinkedAddressCreate) SetAddress(s string) *LinkedAddressCreate {
	lac.mutation.SetAddress(s)
	return lac
}

// SetSalt sets the "salt" field.
func (lac *LinkedAddressCreate) SetSalt(b []byte) *LinkedAddressCreate {
	lac.mutation.SetSalt(b)
	return lac
}

// SetInstitution sets the "institution" field.
func (lac *LinkedAddressCreate) SetInstitution(s string) *LinkedAddressCreate {
	lac.mutation.SetInstitution(s)
	return lac
}

// SetAccountIdentifier sets the "account_identifier" field.
func (lac *LinkedAddressCreate) SetAccountIdentifier(s string) *LinkedAddressCreate {
	lac.mutation.SetAccountIdentifier(s)
	return lac
}

// SetAccountName sets the "account_name" field.
func (lac *LinkedAddressCreate) SetAccountName(s string) *LinkedAddressCreate {
	lac.mutation.SetAccountName(s)
	return lac
}

// SetOwnerAddress sets the "owner_address" field.
func (lac *LinkedAddressCreate) SetOwnerAddress(s string) *LinkedAddressCreate {
	lac.mutation.SetOwnerAddress(s)
	return lac
}

// SetLastIndexedBlock sets the "last_indexed_block" field.
func (lac *LinkedAddressCreate) SetLastIndexedBlock(i int64) *LinkedAddressCreate {
	lac.mutation.SetLastIndexedBlock(i)
	return lac
}

// SetNillableLastIndexedBlock sets the "last_indexed_block" field if the given value is not nil.
func (lac *LinkedAddressCreate) SetNillableLastIndexedBlock(i *int64) *LinkedAddressCreate {
	if i != nil {
		lac.SetLastIndexedBlock(*i)
	}
	return lac
}

// SetTxHash sets the "tx_hash" field.
func (lac *LinkedAddressCreate) SetTxHash(s string) *LinkedAddressCreate {
	lac.mutation.SetTxHash(s)
	return lac
}

// SetNillableTxHash sets the "tx_hash" field if the given value is not nil.
func (lac *LinkedAddressCreate) SetNillableTxHash(s *string) *LinkedAddressCreate {
	if s != nil {
		lac.SetTxHash(*s)
	}
	return lac
}

// AddPaymentOrderIDs adds the "payment_orders" edge to the PaymentOrder entity by IDs.
func (lac *LinkedAddressCreate) AddPaymentOrderIDs(ids ...uuid.UUID) *LinkedAddressCreate {
	lac.mutation.AddPaymentOrderIDs(ids...)
	return lac
}

// AddPaymentOrders adds the "payment_orders" edges to the PaymentOrder entity.
func (lac *LinkedAddressCreate) AddPaymentOrders(p ...*PaymentOrder) *LinkedAddressCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return lac.AddPaymentOrderIDs(ids...)
}

// Mutation returns the LinkedAddressMutation object of the builder.
func (lac *LinkedAddressCreate) Mutation() *LinkedAddressMutation {
	return lac.mutation
}

// Save creates the LinkedAddress in the database.
func (lac *LinkedAddressCreate) Save(ctx context.Context) (*LinkedAddress, error) {
	lac.defaults()
	return withHooks(ctx, lac.sqlSave, lac.mutation, lac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lac *LinkedAddressCreate) SaveX(ctx context.Context) *LinkedAddress {
	v, err := lac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lac *LinkedAddressCreate) Exec(ctx context.Context) error {
	_, err := lac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lac *LinkedAddressCreate) ExecX(ctx context.Context) {
	if err := lac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lac *LinkedAddressCreate) defaults() {
	if _, ok := lac.mutation.CreatedAt(); !ok {
		v := linkedaddress.DefaultCreatedAt()
		lac.mutation.SetCreatedAt(v)
	}
	if _, ok := lac.mutation.UpdatedAt(); !ok {
		v := linkedaddress.DefaultUpdatedAt()
		lac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lac *LinkedAddressCreate) check() error {
	if _, ok := lac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "LinkedAddress.created_at"`)}
	}
	if _, ok := lac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "LinkedAddress.updated_at"`)}
	}
	if _, ok := lac.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "LinkedAddress.address"`)}
	}
	if _, ok := lac.mutation.Salt(); !ok {
		return &ValidationError{Name: "salt", err: errors.New(`ent: missing required field "LinkedAddress.salt"`)}
	}
	if _, ok := lac.mutation.Institution(); !ok {
		return &ValidationError{Name: "institution", err: errors.New(`ent: missing required field "LinkedAddress.institution"`)}
	}
	if _, ok := lac.mutation.AccountIdentifier(); !ok {
		return &ValidationError{Name: "account_identifier", err: errors.New(`ent: missing required field "LinkedAddress.account_identifier"`)}
	}
	if _, ok := lac.mutation.AccountName(); !ok {
		return &ValidationError{Name: "account_name", err: errors.New(`ent: missing required field "LinkedAddress.account_name"`)}
	}
	if _, ok := lac.mutation.OwnerAddress(); !ok {
		return &ValidationError{Name: "owner_address", err: errors.New(`ent: missing required field "LinkedAddress.owner_address"`)}
	}
	if v, ok := lac.mutation.TxHash(); ok {
		if err := linkedaddress.TxHashValidator(v); err != nil {
			return &ValidationError{Name: "tx_hash", err: fmt.Errorf(`ent: validator failed for field "LinkedAddress.tx_hash": %w`, err)}
		}
	}
	return nil
}

func (lac *LinkedAddressCreate) sqlSave(ctx context.Context) (*LinkedAddress, error) {
	if err := lac.check(); err != nil {
		return nil, err
	}
	_node, _spec := lac.createSpec()
	if err := sqlgraph.CreateNode(ctx, lac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	lac.mutation.id = &_node.ID
	lac.mutation.done = true
	return _node, nil
}

func (lac *LinkedAddressCreate) createSpec() (*LinkedAddress, *sqlgraph.CreateSpec) {
	var (
		_node = &LinkedAddress{config: lac.config}
		_spec = sqlgraph.NewCreateSpec(linkedaddress.Table, sqlgraph.NewFieldSpec(linkedaddress.FieldID, field.TypeInt))
	)
	_spec.OnConflict = lac.conflict
	if value, ok := lac.mutation.CreatedAt(); ok {
		_spec.SetField(linkedaddress.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lac.mutation.UpdatedAt(); ok {
		_spec.SetField(linkedaddress.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := lac.mutation.Address(); ok {
		_spec.SetField(linkedaddress.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := lac.mutation.Salt(); ok {
		_spec.SetField(linkedaddress.FieldSalt, field.TypeBytes, value)
		_node.Salt = value
	}
	if value, ok := lac.mutation.Institution(); ok {
		_spec.SetField(linkedaddress.FieldInstitution, field.TypeString, value)
		_node.Institution = value
	}
	if value, ok := lac.mutation.AccountIdentifier(); ok {
		_spec.SetField(linkedaddress.FieldAccountIdentifier, field.TypeString, value)
		_node.AccountIdentifier = value
	}
	if value, ok := lac.mutation.AccountName(); ok {
		_spec.SetField(linkedaddress.FieldAccountName, field.TypeString, value)
		_node.AccountName = value
	}
	if value, ok := lac.mutation.OwnerAddress(); ok {
		_spec.SetField(linkedaddress.FieldOwnerAddress, field.TypeString, value)
		_node.OwnerAddress = value
	}
	if value, ok := lac.mutation.LastIndexedBlock(); ok {
		_spec.SetField(linkedaddress.FieldLastIndexedBlock, field.TypeInt64, value)
		_node.LastIndexedBlock = value
	}
	if value, ok := lac.mutation.TxHash(); ok {
		_spec.SetField(linkedaddress.FieldTxHash, field.TypeString, value)
		_node.TxHash = value
	}
	if nodes := lac.mutation.PaymentOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   linkedaddress.PaymentOrdersTable,
			Columns: []string{linkedaddress.PaymentOrdersColumn},
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
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.LinkedAddress.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LinkedAddressUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (lac *LinkedAddressCreate) OnConflict(opts ...sql.ConflictOption) *LinkedAddressUpsertOne {
	lac.conflict = opts
	return &LinkedAddressUpsertOne{
		create: lac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.LinkedAddress.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (lac *LinkedAddressCreate) OnConflictColumns(columns ...string) *LinkedAddressUpsertOne {
	lac.conflict = append(lac.conflict, sql.ConflictColumns(columns...))
	return &LinkedAddressUpsertOne{
		create: lac,
	}
}

type (
	// LinkedAddressUpsertOne is the builder for "upsert"-ing
	//  one LinkedAddress node.
	LinkedAddressUpsertOne struct {
		create *LinkedAddressCreate
	}

	// LinkedAddressUpsert is the "OnConflict" setter.
	LinkedAddressUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *LinkedAddressUpsert) SetUpdatedAt(v time.Time) *LinkedAddressUpsert {
	u.Set(linkedaddress.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *LinkedAddressUpsert) UpdateUpdatedAt() *LinkedAddressUpsert {
	u.SetExcluded(linkedaddress.FieldUpdatedAt)
	return u
}

// SetAddress sets the "address" field.
func (u *LinkedAddressUpsert) SetAddress(v string) *LinkedAddressUpsert {
	u.Set(linkedaddress.FieldAddress, v)
	return u
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *LinkedAddressUpsert) UpdateAddress() *LinkedAddressUpsert {
	u.SetExcluded(linkedaddress.FieldAddress)
	return u
}

// SetInstitution sets the "institution" field.
func (u *LinkedAddressUpsert) SetInstitution(v string) *LinkedAddressUpsert {
	u.Set(linkedaddress.FieldInstitution, v)
	return u
}

// UpdateInstitution sets the "institution" field to the value that was provided on create.
func (u *LinkedAddressUpsert) UpdateInstitution() *LinkedAddressUpsert {
	u.SetExcluded(linkedaddress.FieldInstitution)
	return u
}

// SetAccountIdentifier sets the "account_identifier" field.
func (u *LinkedAddressUpsert) SetAccountIdentifier(v string) *LinkedAddressUpsert {
	u.Set(linkedaddress.FieldAccountIdentifier, v)
	return u
}

// UpdateAccountIdentifier sets the "account_identifier" field to the value that was provided on create.
func (u *LinkedAddressUpsert) UpdateAccountIdentifier() *LinkedAddressUpsert {
	u.SetExcluded(linkedaddress.FieldAccountIdentifier)
	return u
}

// SetAccountName sets the "account_name" field.
func (u *LinkedAddressUpsert) SetAccountName(v string) *LinkedAddressUpsert {
	u.Set(linkedaddress.FieldAccountName, v)
	return u
}

// UpdateAccountName sets the "account_name" field to the value that was provided on create.
func (u *LinkedAddressUpsert) UpdateAccountName() *LinkedAddressUpsert {
	u.SetExcluded(linkedaddress.FieldAccountName)
	return u
}

// SetOwnerAddress sets the "owner_address" field.
func (u *LinkedAddressUpsert) SetOwnerAddress(v string) *LinkedAddressUpsert {
	u.Set(linkedaddress.FieldOwnerAddress, v)
	return u
}

// UpdateOwnerAddress sets the "owner_address" field to the value that was provided on create.
func (u *LinkedAddressUpsert) UpdateOwnerAddress() *LinkedAddressUpsert {
	u.SetExcluded(linkedaddress.FieldOwnerAddress)
	return u
}

// SetLastIndexedBlock sets the "last_indexed_block" field.
func (u *LinkedAddressUpsert) SetLastIndexedBlock(v int64) *LinkedAddressUpsert {
	u.Set(linkedaddress.FieldLastIndexedBlock, v)
	return u
}

// UpdateLastIndexedBlock sets the "last_indexed_block" field to the value that was provided on create.
func (u *LinkedAddressUpsert) UpdateLastIndexedBlock() *LinkedAddressUpsert {
	u.SetExcluded(linkedaddress.FieldLastIndexedBlock)
	return u
}

// AddLastIndexedBlock adds v to the "last_indexed_block" field.
func (u *LinkedAddressUpsert) AddLastIndexedBlock(v int64) *LinkedAddressUpsert {
	u.Add(linkedaddress.FieldLastIndexedBlock, v)
	return u
}

// ClearLastIndexedBlock clears the value of the "last_indexed_block" field.
func (u *LinkedAddressUpsert) ClearLastIndexedBlock() *LinkedAddressUpsert {
	u.SetNull(linkedaddress.FieldLastIndexedBlock)
	return u
}

// SetTxHash sets the "tx_hash" field.
func (u *LinkedAddressUpsert) SetTxHash(v string) *LinkedAddressUpsert {
	u.Set(linkedaddress.FieldTxHash, v)
	return u
}

// UpdateTxHash sets the "tx_hash" field to the value that was provided on create.
func (u *LinkedAddressUpsert) UpdateTxHash() *LinkedAddressUpsert {
	u.SetExcluded(linkedaddress.FieldTxHash)
	return u
}

// ClearTxHash clears the value of the "tx_hash" field.
func (u *LinkedAddressUpsert) ClearTxHash() *LinkedAddressUpsert {
	u.SetNull(linkedaddress.FieldTxHash)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.LinkedAddress.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *LinkedAddressUpsertOne) UpdateNewValues() *LinkedAddressUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(linkedaddress.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.Salt(); exists {
			s.SetIgnore(linkedaddress.FieldSalt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.LinkedAddress.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *LinkedAddressUpsertOne) Ignore() *LinkedAddressUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LinkedAddressUpsertOne) DoNothing() *LinkedAddressUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LinkedAddressCreate.OnConflict
// documentation for more info.
func (u *LinkedAddressUpsertOne) Update(set func(*LinkedAddressUpsert)) *LinkedAddressUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LinkedAddressUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *LinkedAddressUpsertOne) SetUpdatedAt(v time.Time) *LinkedAddressUpsertOne {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *LinkedAddressUpsertOne) UpdateUpdatedAt() *LinkedAddressUpsertOne {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetAddress sets the "address" field.
func (u *LinkedAddressUpsertOne) SetAddress(v string) *LinkedAddressUpsertOne {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *LinkedAddressUpsertOne) UpdateAddress() *LinkedAddressUpsertOne {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.UpdateAddress()
	})
}

// SetInstitution sets the "institution" field.
func (u *LinkedAddressUpsertOne) SetInstitution(v string) *LinkedAddressUpsertOne {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.SetInstitution(v)
	})
}

// UpdateInstitution sets the "institution" field to the value that was provided on create.
func (u *LinkedAddressUpsertOne) UpdateInstitution() *LinkedAddressUpsertOne {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.UpdateInstitution()
	})
}

// SetAccountIdentifier sets the "account_identifier" field.
func (u *LinkedAddressUpsertOne) SetAccountIdentifier(v string) *LinkedAddressUpsertOne {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.SetAccountIdentifier(v)
	})
}

// UpdateAccountIdentifier sets the "account_identifier" field to the value that was provided on create.
func (u *LinkedAddressUpsertOne) UpdateAccountIdentifier() *LinkedAddressUpsertOne {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.UpdateAccountIdentifier()
	})
}

// SetAccountName sets the "account_name" field.
func (u *LinkedAddressUpsertOne) SetAccountName(v string) *LinkedAddressUpsertOne {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.SetAccountName(v)
	})
}

// UpdateAccountName sets the "account_name" field to the value that was provided on create.
func (u *LinkedAddressUpsertOne) UpdateAccountName() *LinkedAddressUpsertOne {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.UpdateAccountName()
	})
}

// SetOwnerAddress sets the "owner_address" field.
func (u *LinkedAddressUpsertOne) SetOwnerAddress(v string) *LinkedAddressUpsertOne {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.SetOwnerAddress(v)
	})
}

// UpdateOwnerAddress sets the "owner_address" field to the value that was provided on create.
func (u *LinkedAddressUpsertOne) UpdateOwnerAddress() *LinkedAddressUpsertOne {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.UpdateOwnerAddress()
	})
}

// SetLastIndexedBlock sets the "last_indexed_block" field.
func (u *LinkedAddressUpsertOne) SetLastIndexedBlock(v int64) *LinkedAddressUpsertOne {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.SetLastIndexedBlock(v)
	})
}

// AddLastIndexedBlock adds v to the "last_indexed_block" field.
func (u *LinkedAddressUpsertOne) AddLastIndexedBlock(v int64) *LinkedAddressUpsertOne {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.AddLastIndexedBlock(v)
	})
}

// UpdateLastIndexedBlock sets the "last_indexed_block" field to the value that was provided on create.
func (u *LinkedAddressUpsertOne) UpdateLastIndexedBlock() *LinkedAddressUpsertOne {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.UpdateLastIndexedBlock()
	})
}

// ClearLastIndexedBlock clears the value of the "last_indexed_block" field.
func (u *LinkedAddressUpsertOne) ClearLastIndexedBlock() *LinkedAddressUpsertOne {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.ClearLastIndexedBlock()
	})
}

// SetTxHash sets the "tx_hash" field.
func (u *LinkedAddressUpsertOne) SetTxHash(v string) *LinkedAddressUpsertOne {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.SetTxHash(v)
	})
}

// UpdateTxHash sets the "tx_hash" field to the value that was provided on create.
func (u *LinkedAddressUpsertOne) UpdateTxHash() *LinkedAddressUpsertOne {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.UpdateTxHash()
	})
}

// ClearTxHash clears the value of the "tx_hash" field.
func (u *LinkedAddressUpsertOne) ClearTxHash() *LinkedAddressUpsertOne {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.ClearTxHash()
	})
}

// Exec executes the query.
func (u *LinkedAddressUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LinkedAddressCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LinkedAddressUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *LinkedAddressUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *LinkedAddressUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// LinkedAddressCreateBulk is the builder for creating many LinkedAddress entities in bulk.
type LinkedAddressCreateBulk struct {
	config
	err      error
	builders []*LinkedAddressCreate
	conflict []sql.ConflictOption
}

// Save creates the LinkedAddress entities in the database.
func (lacb *LinkedAddressCreateBulk) Save(ctx context.Context) ([]*LinkedAddress, error) {
	if lacb.err != nil {
		return nil, lacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(lacb.builders))
	nodes := make([]*LinkedAddress, len(lacb.builders))
	mutators := make([]Mutator, len(lacb.builders))
	for i := range lacb.builders {
		func(i int, root context.Context) {
			builder := lacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LinkedAddressMutation)
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
					_, err = mutators[i+1].Mutate(root, lacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = lacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lacb *LinkedAddressCreateBulk) SaveX(ctx context.Context) []*LinkedAddress {
	v, err := lacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lacb *LinkedAddressCreateBulk) Exec(ctx context.Context) error {
	_, err := lacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lacb *LinkedAddressCreateBulk) ExecX(ctx context.Context) {
	if err := lacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.LinkedAddress.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LinkedAddressUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (lacb *LinkedAddressCreateBulk) OnConflict(opts ...sql.ConflictOption) *LinkedAddressUpsertBulk {
	lacb.conflict = opts
	return &LinkedAddressUpsertBulk{
		create: lacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.LinkedAddress.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (lacb *LinkedAddressCreateBulk) OnConflictColumns(columns ...string) *LinkedAddressUpsertBulk {
	lacb.conflict = append(lacb.conflict, sql.ConflictColumns(columns...))
	return &LinkedAddressUpsertBulk{
		create: lacb,
	}
}

// LinkedAddressUpsertBulk is the builder for "upsert"-ing
// a bulk of LinkedAddress nodes.
type LinkedAddressUpsertBulk struct {
	create *LinkedAddressCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.LinkedAddress.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *LinkedAddressUpsertBulk) UpdateNewValues() *LinkedAddressUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(linkedaddress.FieldCreatedAt)
			}
			if _, exists := b.mutation.Salt(); exists {
				s.SetIgnore(linkedaddress.FieldSalt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.LinkedAddress.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *LinkedAddressUpsertBulk) Ignore() *LinkedAddressUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LinkedAddressUpsertBulk) DoNothing() *LinkedAddressUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LinkedAddressCreateBulk.OnConflict
// documentation for more info.
func (u *LinkedAddressUpsertBulk) Update(set func(*LinkedAddressUpsert)) *LinkedAddressUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LinkedAddressUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *LinkedAddressUpsertBulk) SetUpdatedAt(v time.Time) *LinkedAddressUpsertBulk {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *LinkedAddressUpsertBulk) UpdateUpdatedAt() *LinkedAddressUpsertBulk {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetAddress sets the "address" field.
func (u *LinkedAddressUpsertBulk) SetAddress(v string) *LinkedAddressUpsertBulk {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *LinkedAddressUpsertBulk) UpdateAddress() *LinkedAddressUpsertBulk {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.UpdateAddress()
	})
}

// SetInstitution sets the "institution" field.
func (u *LinkedAddressUpsertBulk) SetInstitution(v string) *LinkedAddressUpsertBulk {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.SetInstitution(v)
	})
}

// UpdateInstitution sets the "institution" field to the value that was provided on create.
func (u *LinkedAddressUpsertBulk) UpdateInstitution() *LinkedAddressUpsertBulk {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.UpdateInstitution()
	})
}

// SetAccountIdentifier sets the "account_identifier" field.
func (u *LinkedAddressUpsertBulk) SetAccountIdentifier(v string) *LinkedAddressUpsertBulk {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.SetAccountIdentifier(v)
	})
}

// UpdateAccountIdentifier sets the "account_identifier" field to the value that was provided on create.
func (u *LinkedAddressUpsertBulk) UpdateAccountIdentifier() *LinkedAddressUpsertBulk {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.UpdateAccountIdentifier()
	})
}

// SetAccountName sets the "account_name" field.
func (u *LinkedAddressUpsertBulk) SetAccountName(v string) *LinkedAddressUpsertBulk {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.SetAccountName(v)
	})
}

// UpdateAccountName sets the "account_name" field to the value that was provided on create.
func (u *LinkedAddressUpsertBulk) UpdateAccountName() *LinkedAddressUpsertBulk {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.UpdateAccountName()
	})
}

// SetOwnerAddress sets the "owner_address" field.
func (u *LinkedAddressUpsertBulk) SetOwnerAddress(v string) *LinkedAddressUpsertBulk {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.SetOwnerAddress(v)
	})
}

// UpdateOwnerAddress sets the "owner_address" field to the value that was provided on create.
func (u *LinkedAddressUpsertBulk) UpdateOwnerAddress() *LinkedAddressUpsertBulk {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.UpdateOwnerAddress()
	})
}

// SetLastIndexedBlock sets the "last_indexed_block" field.
func (u *LinkedAddressUpsertBulk) SetLastIndexedBlock(v int64) *LinkedAddressUpsertBulk {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.SetLastIndexedBlock(v)
	})
}

// AddLastIndexedBlock adds v to the "last_indexed_block" field.
func (u *LinkedAddressUpsertBulk) AddLastIndexedBlock(v int64) *LinkedAddressUpsertBulk {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.AddLastIndexedBlock(v)
	})
}

// UpdateLastIndexedBlock sets the "last_indexed_block" field to the value that was provided on create.
func (u *LinkedAddressUpsertBulk) UpdateLastIndexedBlock() *LinkedAddressUpsertBulk {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.UpdateLastIndexedBlock()
	})
}

// ClearLastIndexedBlock clears the value of the "last_indexed_block" field.
func (u *LinkedAddressUpsertBulk) ClearLastIndexedBlock() *LinkedAddressUpsertBulk {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.ClearLastIndexedBlock()
	})
}

// SetTxHash sets the "tx_hash" field.
func (u *LinkedAddressUpsertBulk) SetTxHash(v string) *LinkedAddressUpsertBulk {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.SetTxHash(v)
	})
}

// UpdateTxHash sets the "tx_hash" field to the value that was provided on create.
func (u *LinkedAddressUpsertBulk) UpdateTxHash() *LinkedAddressUpsertBulk {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.UpdateTxHash()
	})
}

// ClearTxHash clears the value of the "tx_hash" field.
func (u *LinkedAddressUpsertBulk) ClearTxHash() *LinkedAddressUpsertBulk {
	return u.Update(func(s *LinkedAddressUpsert) {
		s.ClearTxHash()
	})
}

// Exec executes the query.
func (u *LinkedAddressUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the LinkedAddressCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LinkedAddressCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LinkedAddressUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
