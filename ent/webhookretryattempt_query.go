// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/paycrest/protocol/ent/predicate"
	"github.com/paycrest/protocol/ent/webhookretryattempt"
)

// WebhookRetryAttemptQuery is the builder for querying WebhookRetryAttempt entities.
type WebhookRetryAttemptQuery struct {
	config
	ctx        *QueryContext
	order      []webhookretryattempt.OrderOption
	inters     []Interceptor
	predicates []predicate.WebhookRetryAttempt
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WebhookRetryAttemptQuery builder.
func (wraq *WebhookRetryAttemptQuery) Where(ps ...predicate.WebhookRetryAttempt) *WebhookRetryAttemptQuery {
	wraq.predicates = append(wraq.predicates, ps...)
	return wraq
}

// Limit the number of records to be returned by this query.
func (wraq *WebhookRetryAttemptQuery) Limit(limit int) *WebhookRetryAttemptQuery {
	wraq.ctx.Limit = &limit
	return wraq
}

// Offset to start from.
func (wraq *WebhookRetryAttemptQuery) Offset(offset int) *WebhookRetryAttemptQuery {
	wraq.ctx.Offset = &offset
	return wraq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wraq *WebhookRetryAttemptQuery) Unique(unique bool) *WebhookRetryAttemptQuery {
	wraq.ctx.Unique = &unique
	return wraq
}

// Order specifies how the records should be ordered.
func (wraq *WebhookRetryAttemptQuery) Order(o ...webhookretryattempt.OrderOption) *WebhookRetryAttemptQuery {
	wraq.order = append(wraq.order, o...)
	return wraq
}

// First returns the first WebhookRetryAttempt entity from the query.
// Returns a *NotFoundError when no WebhookRetryAttempt was found.
func (wraq *WebhookRetryAttemptQuery) First(ctx context.Context) (*WebhookRetryAttempt, error) {
	nodes, err := wraq.Limit(1).All(setContextOp(ctx, wraq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{webhookretryattempt.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wraq *WebhookRetryAttemptQuery) FirstX(ctx context.Context) *WebhookRetryAttempt {
	node, err := wraq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WebhookRetryAttempt ID from the query.
// Returns a *NotFoundError when no WebhookRetryAttempt ID was found.
func (wraq *WebhookRetryAttemptQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wraq.Limit(1).IDs(setContextOp(ctx, wraq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{webhookretryattempt.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wraq *WebhookRetryAttemptQuery) FirstIDX(ctx context.Context) int {
	id, err := wraq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WebhookRetryAttempt entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WebhookRetryAttempt entity is found.
// Returns a *NotFoundError when no WebhookRetryAttempt entities are found.
func (wraq *WebhookRetryAttemptQuery) Only(ctx context.Context) (*WebhookRetryAttempt, error) {
	nodes, err := wraq.Limit(2).All(setContextOp(ctx, wraq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{webhookretryattempt.Label}
	default:
		return nil, &NotSingularError{webhookretryattempt.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wraq *WebhookRetryAttemptQuery) OnlyX(ctx context.Context) *WebhookRetryAttempt {
	node, err := wraq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WebhookRetryAttempt ID in the query.
// Returns a *NotSingularError when more than one WebhookRetryAttempt ID is found.
// Returns a *NotFoundError when no entities are found.
func (wraq *WebhookRetryAttemptQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wraq.Limit(2).IDs(setContextOp(ctx, wraq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{webhookretryattempt.Label}
	default:
		err = &NotSingularError{webhookretryattempt.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wraq *WebhookRetryAttemptQuery) OnlyIDX(ctx context.Context) int {
	id, err := wraq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WebhookRetryAttempts.
func (wraq *WebhookRetryAttemptQuery) All(ctx context.Context) ([]*WebhookRetryAttempt, error) {
	ctx = setContextOp(ctx, wraq.ctx, ent.OpQueryAll)
	if err := wraq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WebhookRetryAttempt, *WebhookRetryAttemptQuery]()
	return withInterceptors[[]*WebhookRetryAttempt](ctx, wraq, qr, wraq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wraq *WebhookRetryAttemptQuery) AllX(ctx context.Context) []*WebhookRetryAttempt {
	nodes, err := wraq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WebhookRetryAttempt IDs.
func (wraq *WebhookRetryAttemptQuery) IDs(ctx context.Context) (ids []int, err error) {
	if wraq.ctx.Unique == nil && wraq.path != nil {
		wraq.Unique(true)
	}
	ctx = setContextOp(ctx, wraq.ctx, ent.OpQueryIDs)
	if err = wraq.Select(webhookretryattempt.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wraq *WebhookRetryAttemptQuery) IDsX(ctx context.Context) []int {
	ids, err := wraq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wraq *WebhookRetryAttemptQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wraq.ctx, ent.OpQueryCount)
	if err := wraq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wraq, querierCount[*WebhookRetryAttemptQuery](), wraq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wraq *WebhookRetryAttemptQuery) CountX(ctx context.Context) int {
	count, err := wraq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wraq *WebhookRetryAttemptQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wraq.ctx, ent.OpQueryExist)
	switch _, err := wraq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wraq *WebhookRetryAttemptQuery) ExistX(ctx context.Context) bool {
	exist, err := wraq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WebhookRetryAttemptQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wraq *WebhookRetryAttemptQuery) Clone() *WebhookRetryAttemptQuery {
	if wraq == nil {
		return nil
	}
	return &WebhookRetryAttemptQuery{
		config:     wraq.config,
		ctx:        wraq.ctx.Clone(),
		order:      append([]webhookretryattempt.OrderOption{}, wraq.order...),
		inters:     append([]Interceptor{}, wraq.inters...),
		predicates: append([]predicate.WebhookRetryAttempt{}, wraq.predicates...),
		// clone intermediate query.
		sql:  wraq.sql.Clone(),
		path: wraq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WebhookRetryAttempt.Query().
//		GroupBy(webhookretryattempt.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wraq *WebhookRetryAttemptQuery) GroupBy(field string, fields ...string) *WebhookRetryAttemptGroupBy {
	wraq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WebhookRetryAttemptGroupBy{build: wraq}
	grbuild.flds = &wraq.ctx.Fields
	grbuild.label = webhookretryattempt.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.WebhookRetryAttempt.Query().
//		Select(webhookretryattempt.FieldCreatedAt).
//		Scan(ctx, &v)
func (wraq *WebhookRetryAttemptQuery) Select(fields ...string) *WebhookRetryAttemptSelect {
	wraq.ctx.Fields = append(wraq.ctx.Fields, fields...)
	sbuild := &WebhookRetryAttemptSelect{WebhookRetryAttemptQuery: wraq}
	sbuild.label = webhookretryattempt.Label
	sbuild.flds, sbuild.scan = &wraq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WebhookRetryAttemptSelect configured with the given aggregations.
func (wraq *WebhookRetryAttemptQuery) Aggregate(fns ...AggregateFunc) *WebhookRetryAttemptSelect {
	return wraq.Select().Aggregate(fns...)
}

func (wraq *WebhookRetryAttemptQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wraq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wraq); err != nil {
				return err
			}
		}
	}
	for _, f := range wraq.ctx.Fields {
		if !webhookretryattempt.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wraq.path != nil {
		prev, err := wraq.path(ctx)
		if err != nil {
			return err
		}
		wraq.sql = prev
	}
	return nil
}

func (wraq *WebhookRetryAttemptQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WebhookRetryAttempt, error) {
	var (
		nodes = []*WebhookRetryAttempt{}
		_spec = wraq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WebhookRetryAttempt).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WebhookRetryAttempt{config: wraq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wraq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (wraq *WebhookRetryAttemptQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wraq.querySpec()
	_spec.Node.Columns = wraq.ctx.Fields
	if len(wraq.ctx.Fields) > 0 {
		_spec.Unique = wraq.ctx.Unique != nil && *wraq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wraq.driver, _spec)
}

func (wraq *WebhookRetryAttemptQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(webhookretryattempt.Table, webhookretryattempt.Columns, sqlgraph.NewFieldSpec(webhookretryattempt.FieldID, field.TypeInt))
	_spec.From = wraq.sql
	if unique := wraq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wraq.path != nil {
		_spec.Unique = true
	}
	if fields := wraq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, webhookretryattempt.FieldID)
		for i := range fields {
			if fields[i] != webhookretryattempt.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wraq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wraq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wraq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wraq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wraq *WebhookRetryAttemptQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wraq.driver.Dialect())
	t1 := builder.Table(webhookretryattempt.Table)
	columns := wraq.ctx.Fields
	if len(columns) == 0 {
		columns = webhookretryattempt.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wraq.sql != nil {
		selector = wraq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wraq.ctx.Unique != nil && *wraq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wraq.predicates {
		p(selector)
	}
	for _, p := range wraq.order {
		p(selector)
	}
	if offset := wraq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wraq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WebhookRetryAttemptGroupBy is the group-by builder for WebhookRetryAttempt entities.
type WebhookRetryAttemptGroupBy struct {
	selector
	build *WebhookRetryAttemptQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wragb *WebhookRetryAttemptGroupBy) Aggregate(fns ...AggregateFunc) *WebhookRetryAttemptGroupBy {
	wragb.fns = append(wragb.fns, fns...)
	return wragb
}

// Scan applies the selector query and scans the result into the given value.
func (wragb *WebhookRetryAttemptGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wragb.build.ctx, ent.OpQueryGroupBy)
	if err := wragb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WebhookRetryAttemptQuery, *WebhookRetryAttemptGroupBy](ctx, wragb.build, wragb, wragb.build.inters, v)
}

func (wragb *WebhookRetryAttemptGroupBy) sqlScan(ctx context.Context, root *WebhookRetryAttemptQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wragb.fns))
	for _, fn := range wragb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wragb.flds)+len(wragb.fns))
		for _, f := range *wragb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wragb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wragb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WebhookRetryAttemptSelect is the builder for selecting fields of WebhookRetryAttempt entities.
type WebhookRetryAttemptSelect struct {
	*WebhookRetryAttemptQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wras *WebhookRetryAttemptSelect) Aggregate(fns ...AggregateFunc) *WebhookRetryAttemptSelect {
	wras.fns = append(wras.fns, fns...)
	return wras
}

// Scan applies the selector query and scans the result into the given value.
func (wras *WebhookRetryAttemptSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wras.ctx, ent.OpQuerySelect)
	if err := wras.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WebhookRetryAttemptQuery, *WebhookRetryAttemptSelect](ctx, wras.WebhookRetryAttemptQuery, wras, wras.inters, v)
}

func (wras *WebhookRetryAttemptSelect) sqlScan(ctx context.Context, root *WebhookRetryAttemptQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wras.fns))
	for _, fn := range wras.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wras.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wras.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
