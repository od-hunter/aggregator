// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/paycrest/paycrest-protocol/ent/senderprofile"
	"github.com/paycrest/paycrest-protocol/ent/user"
)

// SenderProfileCreate is the builder for creating a SenderProfile entity.
type SenderProfileCreate struct {
	config
	mutation *SenderProfileMutation
	hooks    []Hook
}

// SetWebhookURL sets the "webhook_url" field.
func (spc *SenderProfileCreate) SetWebhookURL(s string) *SenderProfileCreate {
	spc.mutation.SetWebhookURL(s)
	return spc
}

// SetNillableWebhookURL sets the "webhook_url" field if the given value is not nil.
func (spc *SenderProfileCreate) SetNillableWebhookURL(s *string) *SenderProfileCreate {
	if s != nil {
		spc.SetWebhookURL(*s)
	}
	return spc
}

// SetDomainWhitelist sets the "domain_whitelist" field.
func (spc *SenderProfileCreate) SetDomainWhitelist(s []string) *SenderProfileCreate {
	spc.mutation.SetDomainWhitelist(s)
	return spc
}

// SetID sets the "id" field.
func (spc *SenderProfileCreate) SetID(u uuid.UUID) *SenderProfileCreate {
	spc.mutation.SetID(u)
	return spc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (spc *SenderProfileCreate) SetNillableID(u *uuid.UUID) *SenderProfileCreate {
	if u != nil {
		spc.SetID(*u)
	}
	return spc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (spc *SenderProfileCreate) SetUserID(id uuid.UUID) *SenderProfileCreate {
	spc.mutation.SetUserID(id)
	return spc
}

// SetUser sets the "user" edge to the User entity.
func (spc *SenderProfileCreate) SetUser(u *User) *SenderProfileCreate {
	return spc.SetUserID(u.ID)
}

// Mutation returns the SenderProfileMutation object of the builder.
func (spc *SenderProfileCreate) Mutation() *SenderProfileMutation {
	return spc.mutation
}

// Save creates the SenderProfile in the database.
func (spc *SenderProfileCreate) Save(ctx context.Context) (*SenderProfile, error) {
	spc.defaults()
	return withHooks(ctx, spc.sqlSave, spc.mutation, spc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (spc *SenderProfileCreate) SaveX(ctx context.Context) *SenderProfile {
	v, err := spc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spc *SenderProfileCreate) Exec(ctx context.Context) error {
	_, err := spc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spc *SenderProfileCreate) ExecX(ctx context.Context) {
	if err := spc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (spc *SenderProfileCreate) defaults() {
	if _, ok := spc.mutation.DomainWhitelist(); !ok {
		v := senderprofile.DefaultDomainWhitelist
		spc.mutation.SetDomainWhitelist(v)
	}
	if _, ok := spc.mutation.ID(); !ok {
		v := senderprofile.DefaultID()
		spc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (spc *SenderProfileCreate) check() error {
	if _, ok := spc.mutation.DomainWhitelist(); !ok {
		return &ValidationError{Name: "domain_whitelist", err: errors.New(`ent: missing required field "SenderProfile.domain_whitelist"`)}
	}
	if _, ok := spc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "SenderProfile.user"`)}
	}
	return nil
}

func (spc *SenderProfileCreate) sqlSave(ctx context.Context) (*SenderProfile, error) {
	if err := spc.check(); err != nil {
		return nil, err
	}
	_node, _spec := spc.createSpec()
	if err := sqlgraph.CreateNode(ctx, spc.driver, _spec); err != nil {
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
	spc.mutation.id = &_node.ID
	spc.mutation.done = true
	return _node, nil
}

func (spc *SenderProfileCreate) createSpec() (*SenderProfile, *sqlgraph.CreateSpec) {
	var (
		_node = &SenderProfile{config: spc.config}
		_spec = sqlgraph.NewCreateSpec(senderprofile.Table, sqlgraph.NewFieldSpec(senderprofile.FieldID, field.TypeUUID))
	)
	if id, ok := spc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := spc.mutation.WebhookURL(); ok {
		_spec.SetField(senderprofile.FieldWebhookURL, field.TypeString, value)
		_node.WebhookURL = value
	}
	if value, ok := spc.mutation.DomainWhitelist(); ok {
		_spec.SetField(senderprofile.FieldDomainWhitelist, field.TypeJSON, value)
		_node.DomainWhitelist = value
	}
	if nodes := spc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   senderprofile.UserTable,
			Columns: []string{senderprofile.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_sender_profile = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SenderProfileCreateBulk is the builder for creating many SenderProfile entities in bulk.
type SenderProfileCreateBulk struct {
	config
	builders []*SenderProfileCreate
}

// Save creates the SenderProfile entities in the database.
func (spcb *SenderProfileCreateBulk) Save(ctx context.Context) ([]*SenderProfile, error) {
	specs := make([]*sqlgraph.CreateSpec, len(spcb.builders))
	nodes := make([]*SenderProfile, len(spcb.builders))
	mutators := make([]Mutator, len(spcb.builders))
	for i := range spcb.builders {
		func(i int, root context.Context) {
			builder := spcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SenderProfileMutation)
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
					_, err = mutators[i+1].Mutate(root, spcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, spcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, spcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (spcb *SenderProfileCreateBulk) SaveX(ctx context.Context) []*SenderProfile {
	v, err := spcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spcb *SenderProfileCreateBulk) Exec(ctx context.Context) error {
	_, err := spcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spcb *SenderProfileCreateBulk) ExecX(ctx context.Context) {
	if err := spcb.Exec(ctx); err != nil {
		panic(err)
	}
}
