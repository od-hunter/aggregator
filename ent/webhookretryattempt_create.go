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
	"github.com/paycrest/aggregator/ent/webhookretryattempt"
)

// WebhookRetryAttemptCreate is the builder for creating a WebhookRetryAttempt entity.
type WebhookRetryAttemptCreate struct {
	config
	mutation *WebhookRetryAttemptMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (wrac *WebhookRetryAttemptCreate) SetCreatedAt(t time.Time) *WebhookRetryAttemptCreate {
	wrac.mutation.SetCreatedAt(t)
	return wrac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wrac *WebhookRetryAttemptCreate) SetNillableCreatedAt(t *time.Time) *WebhookRetryAttemptCreate {
	if t != nil {
		wrac.SetCreatedAt(*t)
	}
	return wrac
}

// SetUpdatedAt sets the "updated_at" field.
func (wrac *WebhookRetryAttemptCreate) SetUpdatedAt(t time.Time) *WebhookRetryAttemptCreate {
	wrac.mutation.SetUpdatedAt(t)
	return wrac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wrac *WebhookRetryAttemptCreate) SetNillableUpdatedAt(t *time.Time) *WebhookRetryAttemptCreate {
	if t != nil {
		wrac.SetUpdatedAt(*t)
	}
	return wrac
}

// SetAttemptNumber sets the "attempt_number" field.
func (wrac *WebhookRetryAttemptCreate) SetAttemptNumber(i int) *WebhookRetryAttemptCreate {
	wrac.mutation.SetAttemptNumber(i)
	return wrac
}

// SetNextRetryTime sets the "next_retry_time" field.
func (wrac *WebhookRetryAttemptCreate) SetNextRetryTime(t time.Time) *WebhookRetryAttemptCreate {
	wrac.mutation.SetNextRetryTime(t)
	return wrac
}

// SetNillableNextRetryTime sets the "next_retry_time" field if the given value is not nil.
func (wrac *WebhookRetryAttemptCreate) SetNillableNextRetryTime(t *time.Time) *WebhookRetryAttemptCreate {
	if t != nil {
		wrac.SetNextRetryTime(*t)
	}
	return wrac
}

// SetPayload sets the "payload" field.
func (wrac *WebhookRetryAttemptCreate) SetPayload(m map[string]interface{}) *WebhookRetryAttemptCreate {
	wrac.mutation.SetPayload(m)
	return wrac
}

// SetSignature sets the "signature" field.
func (wrac *WebhookRetryAttemptCreate) SetSignature(s string) *WebhookRetryAttemptCreate {
	wrac.mutation.SetSignature(s)
	return wrac
}

// SetNillableSignature sets the "signature" field if the given value is not nil.
func (wrac *WebhookRetryAttemptCreate) SetNillableSignature(s *string) *WebhookRetryAttemptCreate {
	if s != nil {
		wrac.SetSignature(*s)
	}
	return wrac
}

// SetWebhookURL sets the "webhook_url" field.
func (wrac *WebhookRetryAttemptCreate) SetWebhookURL(s string) *WebhookRetryAttemptCreate {
	wrac.mutation.SetWebhookURL(s)
	return wrac
}

// SetStatus sets the "status" field.
func (wrac *WebhookRetryAttemptCreate) SetStatus(w webhookretryattempt.Status) *WebhookRetryAttemptCreate {
	wrac.mutation.SetStatus(w)
	return wrac
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wrac *WebhookRetryAttemptCreate) SetNillableStatus(w *webhookretryattempt.Status) *WebhookRetryAttemptCreate {
	if w != nil {
		wrac.SetStatus(*w)
	}
	return wrac
}

// Mutation returns the WebhookRetryAttemptMutation object of the builder.
func (wrac *WebhookRetryAttemptCreate) Mutation() *WebhookRetryAttemptMutation {
	return wrac.mutation
}

// Save creates the WebhookRetryAttempt in the database.
func (wrac *WebhookRetryAttemptCreate) Save(ctx context.Context) (*WebhookRetryAttempt, error) {
	wrac.defaults()
	return withHooks(ctx, wrac.sqlSave, wrac.mutation, wrac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wrac *WebhookRetryAttemptCreate) SaveX(ctx context.Context) *WebhookRetryAttempt {
	v, err := wrac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wrac *WebhookRetryAttemptCreate) Exec(ctx context.Context) error {
	_, err := wrac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wrac *WebhookRetryAttemptCreate) ExecX(ctx context.Context) {
	if err := wrac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wrac *WebhookRetryAttemptCreate) defaults() {
	if _, ok := wrac.mutation.CreatedAt(); !ok {
		v := webhookretryattempt.DefaultCreatedAt()
		wrac.mutation.SetCreatedAt(v)
	}
	if _, ok := wrac.mutation.UpdatedAt(); !ok {
		v := webhookretryattempt.DefaultUpdatedAt()
		wrac.mutation.SetUpdatedAt(v)
	}
	if _, ok := wrac.mutation.NextRetryTime(); !ok {
		v := webhookretryattempt.DefaultNextRetryTime()
		wrac.mutation.SetNextRetryTime(v)
	}
	if _, ok := wrac.mutation.Status(); !ok {
		v := webhookretryattempt.DefaultStatus
		wrac.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wrac *WebhookRetryAttemptCreate) check() error {
	if _, ok := wrac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "WebhookRetryAttempt.created_at"`)}
	}
	if _, ok := wrac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "WebhookRetryAttempt.updated_at"`)}
	}
	if _, ok := wrac.mutation.AttemptNumber(); !ok {
		return &ValidationError{Name: "attempt_number", err: errors.New(`ent: missing required field "WebhookRetryAttempt.attempt_number"`)}
	}
	if _, ok := wrac.mutation.NextRetryTime(); !ok {
		return &ValidationError{Name: "next_retry_time", err: errors.New(`ent: missing required field "WebhookRetryAttempt.next_retry_time"`)}
	}
	if _, ok := wrac.mutation.Payload(); !ok {
		return &ValidationError{Name: "payload", err: errors.New(`ent: missing required field "WebhookRetryAttempt.payload"`)}
	}
	if _, ok := wrac.mutation.WebhookURL(); !ok {
		return &ValidationError{Name: "webhook_url", err: errors.New(`ent: missing required field "WebhookRetryAttempt.webhook_url"`)}
	}
	if _, ok := wrac.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "WebhookRetryAttempt.status"`)}
	}
	if v, ok := wrac.mutation.Status(); ok {
		if err := webhookretryattempt.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "WebhookRetryAttempt.status": %w`, err)}
		}
	}
	return nil
}

func (wrac *WebhookRetryAttemptCreate) sqlSave(ctx context.Context) (*WebhookRetryAttempt, error) {
	if err := wrac.check(); err != nil {
		return nil, err
	}
	_node, _spec := wrac.createSpec()
	if err := sqlgraph.CreateNode(ctx, wrac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	wrac.mutation.id = &_node.ID
	wrac.mutation.done = true
	return _node, nil
}

func (wrac *WebhookRetryAttemptCreate) createSpec() (*WebhookRetryAttempt, *sqlgraph.CreateSpec) {
	var (
		_node = &WebhookRetryAttempt{config: wrac.config}
		_spec = sqlgraph.NewCreateSpec(webhookretryattempt.Table, sqlgraph.NewFieldSpec(webhookretryattempt.FieldID, field.TypeInt))
	)
	_spec.OnConflict = wrac.conflict
	if value, ok := wrac.mutation.CreatedAt(); ok {
		_spec.SetField(webhookretryattempt.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := wrac.mutation.UpdatedAt(); ok {
		_spec.SetField(webhookretryattempt.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := wrac.mutation.AttemptNumber(); ok {
		_spec.SetField(webhookretryattempt.FieldAttemptNumber, field.TypeInt, value)
		_node.AttemptNumber = value
	}
	if value, ok := wrac.mutation.NextRetryTime(); ok {
		_spec.SetField(webhookretryattempt.FieldNextRetryTime, field.TypeTime, value)
		_node.NextRetryTime = value
	}
	if value, ok := wrac.mutation.Payload(); ok {
		_spec.SetField(webhookretryattempt.FieldPayload, field.TypeJSON, value)
		_node.Payload = value
	}
	if value, ok := wrac.mutation.Signature(); ok {
		_spec.SetField(webhookretryattempt.FieldSignature, field.TypeString, value)
		_node.Signature = value
	}
	if value, ok := wrac.mutation.WebhookURL(); ok {
		_spec.SetField(webhookretryattempt.FieldWebhookURL, field.TypeString, value)
		_node.WebhookURL = value
	}
	if value, ok := wrac.mutation.Status(); ok {
		_spec.SetField(webhookretryattempt.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.WebhookRetryAttempt.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.WebhookRetryAttemptUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (wrac *WebhookRetryAttemptCreate) OnConflict(opts ...sql.ConflictOption) *WebhookRetryAttemptUpsertOne {
	wrac.conflict = opts
	return &WebhookRetryAttemptUpsertOne{
		create: wrac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.WebhookRetryAttempt.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (wrac *WebhookRetryAttemptCreate) OnConflictColumns(columns ...string) *WebhookRetryAttemptUpsertOne {
	wrac.conflict = append(wrac.conflict, sql.ConflictColumns(columns...))
	return &WebhookRetryAttemptUpsertOne{
		create: wrac,
	}
}

type (
	// WebhookRetryAttemptUpsertOne is the builder for "upsert"-ing
	//  one WebhookRetryAttempt node.
	WebhookRetryAttemptUpsertOne struct {
		create *WebhookRetryAttemptCreate
	}

	// WebhookRetryAttemptUpsert is the "OnConflict" setter.
	WebhookRetryAttemptUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *WebhookRetryAttemptUpsert) SetUpdatedAt(v time.Time) *WebhookRetryAttemptUpsert {
	u.Set(webhookretryattempt.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *WebhookRetryAttemptUpsert) UpdateUpdatedAt() *WebhookRetryAttemptUpsert {
	u.SetExcluded(webhookretryattempt.FieldUpdatedAt)
	return u
}

// SetAttemptNumber sets the "attempt_number" field.
func (u *WebhookRetryAttemptUpsert) SetAttemptNumber(v int) *WebhookRetryAttemptUpsert {
	u.Set(webhookretryattempt.FieldAttemptNumber, v)
	return u
}

// UpdateAttemptNumber sets the "attempt_number" field to the value that was provided on create.
func (u *WebhookRetryAttemptUpsert) UpdateAttemptNumber() *WebhookRetryAttemptUpsert {
	u.SetExcluded(webhookretryattempt.FieldAttemptNumber)
	return u
}

// AddAttemptNumber adds v to the "attempt_number" field.
func (u *WebhookRetryAttemptUpsert) AddAttemptNumber(v int) *WebhookRetryAttemptUpsert {
	u.Add(webhookretryattempt.FieldAttemptNumber, v)
	return u
}

// SetNextRetryTime sets the "next_retry_time" field.
func (u *WebhookRetryAttemptUpsert) SetNextRetryTime(v time.Time) *WebhookRetryAttemptUpsert {
	u.Set(webhookretryattempt.FieldNextRetryTime, v)
	return u
}

// UpdateNextRetryTime sets the "next_retry_time" field to the value that was provided on create.
func (u *WebhookRetryAttemptUpsert) UpdateNextRetryTime() *WebhookRetryAttemptUpsert {
	u.SetExcluded(webhookretryattempt.FieldNextRetryTime)
	return u
}

// SetPayload sets the "payload" field.
func (u *WebhookRetryAttemptUpsert) SetPayload(v map[string]interface{}) *WebhookRetryAttemptUpsert {
	u.Set(webhookretryattempt.FieldPayload, v)
	return u
}

// UpdatePayload sets the "payload" field to the value that was provided on create.
func (u *WebhookRetryAttemptUpsert) UpdatePayload() *WebhookRetryAttemptUpsert {
	u.SetExcluded(webhookretryattempt.FieldPayload)
	return u
}

// SetSignature sets the "signature" field.
func (u *WebhookRetryAttemptUpsert) SetSignature(v string) *WebhookRetryAttemptUpsert {
	u.Set(webhookretryattempt.FieldSignature, v)
	return u
}

// UpdateSignature sets the "signature" field to the value that was provided on create.
func (u *WebhookRetryAttemptUpsert) UpdateSignature() *WebhookRetryAttemptUpsert {
	u.SetExcluded(webhookretryattempt.FieldSignature)
	return u
}

// ClearSignature clears the value of the "signature" field.
func (u *WebhookRetryAttemptUpsert) ClearSignature() *WebhookRetryAttemptUpsert {
	u.SetNull(webhookretryattempt.FieldSignature)
	return u
}

// SetWebhookURL sets the "webhook_url" field.
func (u *WebhookRetryAttemptUpsert) SetWebhookURL(v string) *WebhookRetryAttemptUpsert {
	u.Set(webhookretryattempt.FieldWebhookURL, v)
	return u
}

// UpdateWebhookURL sets the "webhook_url" field to the value that was provided on create.
func (u *WebhookRetryAttemptUpsert) UpdateWebhookURL() *WebhookRetryAttemptUpsert {
	u.SetExcluded(webhookretryattempt.FieldWebhookURL)
	return u
}

// SetStatus sets the "status" field.
func (u *WebhookRetryAttemptUpsert) SetStatus(v webhookretryattempt.Status) *WebhookRetryAttemptUpsert {
	u.Set(webhookretryattempt.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *WebhookRetryAttemptUpsert) UpdateStatus() *WebhookRetryAttemptUpsert {
	u.SetExcluded(webhookretryattempt.FieldStatus)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.WebhookRetryAttempt.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *WebhookRetryAttemptUpsertOne) UpdateNewValues() *WebhookRetryAttemptUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(webhookretryattempt.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.WebhookRetryAttempt.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *WebhookRetryAttemptUpsertOne) Ignore() *WebhookRetryAttemptUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *WebhookRetryAttemptUpsertOne) DoNothing() *WebhookRetryAttemptUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the WebhookRetryAttemptCreate.OnConflict
// documentation for more info.
func (u *WebhookRetryAttemptUpsertOne) Update(set func(*WebhookRetryAttemptUpsert)) *WebhookRetryAttemptUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&WebhookRetryAttemptUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *WebhookRetryAttemptUpsertOne) SetUpdatedAt(v time.Time) *WebhookRetryAttemptUpsertOne {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *WebhookRetryAttemptUpsertOne) UpdateUpdatedAt() *WebhookRetryAttemptUpsertOne {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetAttemptNumber sets the "attempt_number" field.
func (u *WebhookRetryAttemptUpsertOne) SetAttemptNumber(v int) *WebhookRetryAttemptUpsertOne {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.SetAttemptNumber(v)
	})
}

// AddAttemptNumber adds v to the "attempt_number" field.
func (u *WebhookRetryAttemptUpsertOne) AddAttemptNumber(v int) *WebhookRetryAttemptUpsertOne {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.AddAttemptNumber(v)
	})
}

// UpdateAttemptNumber sets the "attempt_number" field to the value that was provided on create.
func (u *WebhookRetryAttemptUpsertOne) UpdateAttemptNumber() *WebhookRetryAttemptUpsertOne {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.UpdateAttemptNumber()
	})
}

// SetNextRetryTime sets the "next_retry_time" field.
func (u *WebhookRetryAttemptUpsertOne) SetNextRetryTime(v time.Time) *WebhookRetryAttemptUpsertOne {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.SetNextRetryTime(v)
	})
}

// UpdateNextRetryTime sets the "next_retry_time" field to the value that was provided on create.
func (u *WebhookRetryAttemptUpsertOne) UpdateNextRetryTime() *WebhookRetryAttemptUpsertOne {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.UpdateNextRetryTime()
	})
}

// SetPayload sets the "payload" field.
func (u *WebhookRetryAttemptUpsertOne) SetPayload(v map[string]interface{}) *WebhookRetryAttemptUpsertOne {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.SetPayload(v)
	})
}

// UpdatePayload sets the "payload" field to the value that was provided on create.
func (u *WebhookRetryAttemptUpsertOne) UpdatePayload() *WebhookRetryAttemptUpsertOne {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.UpdatePayload()
	})
}

// SetSignature sets the "signature" field.
func (u *WebhookRetryAttemptUpsertOne) SetSignature(v string) *WebhookRetryAttemptUpsertOne {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.SetSignature(v)
	})
}

// UpdateSignature sets the "signature" field to the value that was provided on create.
func (u *WebhookRetryAttemptUpsertOne) UpdateSignature() *WebhookRetryAttemptUpsertOne {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.UpdateSignature()
	})
}

// ClearSignature clears the value of the "signature" field.
func (u *WebhookRetryAttemptUpsertOne) ClearSignature() *WebhookRetryAttemptUpsertOne {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.ClearSignature()
	})
}

// SetWebhookURL sets the "webhook_url" field.
func (u *WebhookRetryAttemptUpsertOne) SetWebhookURL(v string) *WebhookRetryAttemptUpsertOne {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.SetWebhookURL(v)
	})
}

// UpdateWebhookURL sets the "webhook_url" field to the value that was provided on create.
func (u *WebhookRetryAttemptUpsertOne) UpdateWebhookURL() *WebhookRetryAttemptUpsertOne {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.UpdateWebhookURL()
	})
}

// SetStatus sets the "status" field.
func (u *WebhookRetryAttemptUpsertOne) SetStatus(v webhookretryattempt.Status) *WebhookRetryAttemptUpsertOne {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *WebhookRetryAttemptUpsertOne) UpdateStatus() *WebhookRetryAttemptUpsertOne {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *WebhookRetryAttemptUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for WebhookRetryAttemptCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *WebhookRetryAttemptUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *WebhookRetryAttemptUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *WebhookRetryAttemptUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// WebhookRetryAttemptCreateBulk is the builder for creating many WebhookRetryAttempt entities in bulk.
type WebhookRetryAttemptCreateBulk struct {
	config
	err      error
	builders []*WebhookRetryAttemptCreate
	conflict []sql.ConflictOption
}

// Save creates the WebhookRetryAttempt entities in the database.
func (wracb *WebhookRetryAttemptCreateBulk) Save(ctx context.Context) ([]*WebhookRetryAttempt, error) {
	if wracb.err != nil {
		return nil, wracb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wracb.builders))
	nodes := make([]*WebhookRetryAttempt, len(wracb.builders))
	mutators := make([]Mutator, len(wracb.builders))
	for i := range wracb.builders {
		func(i int, root context.Context) {
			builder := wracb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WebhookRetryAttemptMutation)
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
					_, err = mutators[i+1].Mutate(root, wracb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = wracb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wracb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wracb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wracb *WebhookRetryAttemptCreateBulk) SaveX(ctx context.Context) []*WebhookRetryAttempt {
	v, err := wracb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wracb *WebhookRetryAttemptCreateBulk) Exec(ctx context.Context) error {
	_, err := wracb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wracb *WebhookRetryAttemptCreateBulk) ExecX(ctx context.Context) {
	if err := wracb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.WebhookRetryAttempt.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.WebhookRetryAttemptUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (wracb *WebhookRetryAttemptCreateBulk) OnConflict(opts ...sql.ConflictOption) *WebhookRetryAttemptUpsertBulk {
	wracb.conflict = opts
	return &WebhookRetryAttemptUpsertBulk{
		create: wracb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.WebhookRetryAttempt.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (wracb *WebhookRetryAttemptCreateBulk) OnConflictColumns(columns ...string) *WebhookRetryAttemptUpsertBulk {
	wracb.conflict = append(wracb.conflict, sql.ConflictColumns(columns...))
	return &WebhookRetryAttemptUpsertBulk{
		create: wracb,
	}
}

// WebhookRetryAttemptUpsertBulk is the builder for "upsert"-ing
// a bulk of WebhookRetryAttempt nodes.
type WebhookRetryAttemptUpsertBulk struct {
	create *WebhookRetryAttemptCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.WebhookRetryAttempt.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *WebhookRetryAttemptUpsertBulk) UpdateNewValues() *WebhookRetryAttemptUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(webhookretryattempt.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.WebhookRetryAttempt.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *WebhookRetryAttemptUpsertBulk) Ignore() *WebhookRetryAttemptUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *WebhookRetryAttemptUpsertBulk) DoNothing() *WebhookRetryAttemptUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the WebhookRetryAttemptCreateBulk.OnConflict
// documentation for more info.
func (u *WebhookRetryAttemptUpsertBulk) Update(set func(*WebhookRetryAttemptUpsert)) *WebhookRetryAttemptUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&WebhookRetryAttemptUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *WebhookRetryAttemptUpsertBulk) SetUpdatedAt(v time.Time) *WebhookRetryAttemptUpsertBulk {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *WebhookRetryAttemptUpsertBulk) UpdateUpdatedAt() *WebhookRetryAttemptUpsertBulk {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetAttemptNumber sets the "attempt_number" field.
func (u *WebhookRetryAttemptUpsertBulk) SetAttemptNumber(v int) *WebhookRetryAttemptUpsertBulk {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.SetAttemptNumber(v)
	})
}

// AddAttemptNumber adds v to the "attempt_number" field.
func (u *WebhookRetryAttemptUpsertBulk) AddAttemptNumber(v int) *WebhookRetryAttemptUpsertBulk {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.AddAttemptNumber(v)
	})
}

// UpdateAttemptNumber sets the "attempt_number" field to the value that was provided on create.
func (u *WebhookRetryAttemptUpsertBulk) UpdateAttemptNumber() *WebhookRetryAttemptUpsertBulk {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.UpdateAttemptNumber()
	})
}

// SetNextRetryTime sets the "next_retry_time" field.
func (u *WebhookRetryAttemptUpsertBulk) SetNextRetryTime(v time.Time) *WebhookRetryAttemptUpsertBulk {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.SetNextRetryTime(v)
	})
}

// UpdateNextRetryTime sets the "next_retry_time" field to the value that was provided on create.
func (u *WebhookRetryAttemptUpsertBulk) UpdateNextRetryTime() *WebhookRetryAttemptUpsertBulk {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.UpdateNextRetryTime()
	})
}

// SetPayload sets the "payload" field.
func (u *WebhookRetryAttemptUpsertBulk) SetPayload(v map[string]interface{}) *WebhookRetryAttemptUpsertBulk {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.SetPayload(v)
	})
}

// UpdatePayload sets the "payload" field to the value that was provided on create.
func (u *WebhookRetryAttemptUpsertBulk) UpdatePayload() *WebhookRetryAttemptUpsertBulk {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.UpdatePayload()
	})
}

// SetSignature sets the "signature" field.
func (u *WebhookRetryAttemptUpsertBulk) SetSignature(v string) *WebhookRetryAttemptUpsertBulk {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.SetSignature(v)
	})
}

// UpdateSignature sets the "signature" field to the value that was provided on create.
func (u *WebhookRetryAttemptUpsertBulk) UpdateSignature() *WebhookRetryAttemptUpsertBulk {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.UpdateSignature()
	})
}

// ClearSignature clears the value of the "signature" field.
func (u *WebhookRetryAttemptUpsertBulk) ClearSignature() *WebhookRetryAttemptUpsertBulk {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.ClearSignature()
	})
}

// SetWebhookURL sets the "webhook_url" field.
func (u *WebhookRetryAttemptUpsertBulk) SetWebhookURL(v string) *WebhookRetryAttemptUpsertBulk {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.SetWebhookURL(v)
	})
}

// UpdateWebhookURL sets the "webhook_url" field to the value that was provided on create.
func (u *WebhookRetryAttemptUpsertBulk) UpdateWebhookURL() *WebhookRetryAttemptUpsertBulk {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.UpdateWebhookURL()
	})
}

// SetStatus sets the "status" field.
func (u *WebhookRetryAttemptUpsertBulk) SetStatus(v webhookretryattempt.Status) *WebhookRetryAttemptUpsertBulk {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *WebhookRetryAttemptUpsertBulk) UpdateStatus() *WebhookRetryAttemptUpsertBulk {
	return u.Update(func(s *WebhookRetryAttemptUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *WebhookRetryAttemptUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the WebhookRetryAttemptCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for WebhookRetryAttemptCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *WebhookRetryAttemptUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
