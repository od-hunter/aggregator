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
	"github.com/google/uuid"
	"github.com/paycrest/protocol/ent/paymentorder"
	"github.com/paycrest/protocol/ent/paymentorderrecipient"
	"github.com/paycrest/protocol/ent/predicate"
)

// PaymentOrderRecipientQuery is the builder for querying PaymentOrderRecipient entities.
type PaymentOrderRecipientQuery struct {
	config
	ctx              *QueryContext
	order            []paymentorderrecipient.OrderOption
	inters           []Interceptor
	predicates       []predicate.PaymentOrderRecipient
	withPaymentOrder *PaymentOrderQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PaymentOrderRecipientQuery builder.
func (porq *PaymentOrderRecipientQuery) Where(ps ...predicate.PaymentOrderRecipient) *PaymentOrderRecipientQuery {
	porq.predicates = append(porq.predicates, ps...)
	return porq
}

// Limit the number of records to be returned by this query.
func (porq *PaymentOrderRecipientQuery) Limit(limit int) *PaymentOrderRecipientQuery {
	porq.ctx.Limit = &limit
	return porq
}

// Offset to start from.
func (porq *PaymentOrderRecipientQuery) Offset(offset int) *PaymentOrderRecipientQuery {
	porq.ctx.Offset = &offset
	return porq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (porq *PaymentOrderRecipientQuery) Unique(unique bool) *PaymentOrderRecipientQuery {
	porq.ctx.Unique = &unique
	return porq
}

// Order specifies how the records should be ordered.
func (porq *PaymentOrderRecipientQuery) Order(o ...paymentorderrecipient.OrderOption) *PaymentOrderRecipientQuery {
	porq.order = append(porq.order, o...)
	return porq
}

// QueryPaymentOrder chains the current query on the "payment_order" edge.
func (porq *PaymentOrderRecipientQuery) QueryPaymentOrder() *PaymentOrderQuery {
	query := (&PaymentOrderClient{config: porq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := porq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := porq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(paymentorderrecipient.Table, paymentorderrecipient.FieldID, selector),
			sqlgraph.To(paymentorder.Table, paymentorder.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, paymentorderrecipient.PaymentOrderTable, paymentorderrecipient.PaymentOrderColumn),
		)
		fromU = sqlgraph.SetNeighbors(porq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PaymentOrderRecipient entity from the query.
// Returns a *NotFoundError when no PaymentOrderRecipient was found.
func (porq *PaymentOrderRecipientQuery) First(ctx context.Context) (*PaymentOrderRecipient, error) {
	nodes, err := porq.Limit(1).All(setContextOp(ctx, porq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{paymentorderrecipient.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (porq *PaymentOrderRecipientQuery) FirstX(ctx context.Context) *PaymentOrderRecipient {
	node, err := porq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PaymentOrderRecipient ID from the query.
// Returns a *NotFoundError when no PaymentOrderRecipient ID was found.
func (porq *PaymentOrderRecipientQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = porq.Limit(1).IDs(setContextOp(ctx, porq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{paymentorderrecipient.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (porq *PaymentOrderRecipientQuery) FirstIDX(ctx context.Context) int {
	id, err := porq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PaymentOrderRecipient entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PaymentOrderRecipient entity is found.
// Returns a *NotFoundError when no PaymentOrderRecipient entities are found.
func (porq *PaymentOrderRecipientQuery) Only(ctx context.Context) (*PaymentOrderRecipient, error) {
	nodes, err := porq.Limit(2).All(setContextOp(ctx, porq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{paymentorderrecipient.Label}
	default:
		return nil, &NotSingularError{paymentorderrecipient.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (porq *PaymentOrderRecipientQuery) OnlyX(ctx context.Context) *PaymentOrderRecipient {
	node, err := porq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PaymentOrderRecipient ID in the query.
// Returns a *NotSingularError when more than one PaymentOrderRecipient ID is found.
// Returns a *NotFoundError when no entities are found.
func (porq *PaymentOrderRecipientQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = porq.Limit(2).IDs(setContextOp(ctx, porq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{paymentorderrecipient.Label}
	default:
		err = &NotSingularError{paymentorderrecipient.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (porq *PaymentOrderRecipientQuery) OnlyIDX(ctx context.Context) int {
	id, err := porq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PaymentOrderRecipients.
func (porq *PaymentOrderRecipientQuery) All(ctx context.Context) ([]*PaymentOrderRecipient, error) {
	ctx = setContextOp(ctx, porq.ctx, ent.OpQueryAll)
	if err := porq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PaymentOrderRecipient, *PaymentOrderRecipientQuery]()
	return withInterceptors[[]*PaymentOrderRecipient](ctx, porq, qr, porq.inters)
}

// AllX is like All, but panics if an error occurs.
func (porq *PaymentOrderRecipientQuery) AllX(ctx context.Context) []*PaymentOrderRecipient {
	nodes, err := porq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PaymentOrderRecipient IDs.
func (porq *PaymentOrderRecipientQuery) IDs(ctx context.Context) (ids []int, err error) {
	if porq.ctx.Unique == nil && porq.path != nil {
		porq.Unique(true)
	}
	ctx = setContextOp(ctx, porq.ctx, ent.OpQueryIDs)
	if err = porq.Select(paymentorderrecipient.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (porq *PaymentOrderRecipientQuery) IDsX(ctx context.Context) []int {
	ids, err := porq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (porq *PaymentOrderRecipientQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, porq.ctx, ent.OpQueryCount)
	if err := porq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, porq, querierCount[*PaymentOrderRecipientQuery](), porq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (porq *PaymentOrderRecipientQuery) CountX(ctx context.Context) int {
	count, err := porq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (porq *PaymentOrderRecipientQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, porq.ctx, ent.OpQueryExist)
	switch _, err := porq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (porq *PaymentOrderRecipientQuery) ExistX(ctx context.Context) bool {
	exist, err := porq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PaymentOrderRecipientQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (porq *PaymentOrderRecipientQuery) Clone() *PaymentOrderRecipientQuery {
	if porq == nil {
		return nil
	}
	return &PaymentOrderRecipientQuery{
		config:           porq.config,
		ctx:              porq.ctx.Clone(),
		order:            append([]paymentorderrecipient.OrderOption{}, porq.order...),
		inters:           append([]Interceptor{}, porq.inters...),
		predicates:       append([]predicate.PaymentOrderRecipient{}, porq.predicates...),
		withPaymentOrder: porq.withPaymentOrder.Clone(),
		// clone intermediate query.
		sql:  porq.sql.Clone(),
		path: porq.path,
	}
}

// WithPaymentOrder tells the query-builder to eager-load the nodes that are connected to
// the "payment_order" edge. The optional arguments are used to configure the query builder of the edge.
func (porq *PaymentOrderRecipientQuery) WithPaymentOrder(opts ...func(*PaymentOrderQuery)) *PaymentOrderRecipientQuery {
	query := (&PaymentOrderClient{config: porq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	porq.withPaymentOrder = query
	return porq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Institution string `json:"institution,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PaymentOrderRecipient.Query().
//		GroupBy(paymentorderrecipient.FieldInstitution).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (porq *PaymentOrderRecipientQuery) GroupBy(field string, fields ...string) *PaymentOrderRecipientGroupBy {
	porq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PaymentOrderRecipientGroupBy{build: porq}
	grbuild.flds = &porq.ctx.Fields
	grbuild.label = paymentorderrecipient.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Institution string `json:"institution,omitempty"`
//	}
//
//	client.PaymentOrderRecipient.Query().
//		Select(paymentorderrecipient.FieldInstitution).
//		Scan(ctx, &v)
func (porq *PaymentOrderRecipientQuery) Select(fields ...string) *PaymentOrderRecipientSelect {
	porq.ctx.Fields = append(porq.ctx.Fields, fields...)
	sbuild := &PaymentOrderRecipientSelect{PaymentOrderRecipientQuery: porq}
	sbuild.label = paymentorderrecipient.Label
	sbuild.flds, sbuild.scan = &porq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PaymentOrderRecipientSelect configured with the given aggregations.
func (porq *PaymentOrderRecipientQuery) Aggregate(fns ...AggregateFunc) *PaymentOrderRecipientSelect {
	return porq.Select().Aggregate(fns...)
}

func (porq *PaymentOrderRecipientQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range porq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, porq); err != nil {
				return err
			}
		}
	}
	for _, f := range porq.ctx.Fields {
		if !paymentorderrecipient.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if porq.path != nil {
		prev, err := porq.path(ctx)
		if err != nil {
			return err
		}
		porq.sql = prev
	}
	return nil
}

func (porq *PaymentOrderRecipientQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PaymentOrderRecipient, error) {
	var (
		nodes       = []*PaymentOrderRecipient{}
		withFKs     = porq.withFKs
		_spec       = porq.querySpec()
		loadedTypes = [1]bool{
			porq.withPaymentOrder != nil,
		}
	)
	if porq.withPaymentOrder != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, paymentorderrecipient.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PaymentOrderRecipient).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PaymentOrderRecipient{config: porq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, porq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := porq.withPaymentOrder; query != nil {
		if err := porq.loadPaymentOrder(ctx, query, nodes, nil,
			func(n *PaymentOrderRecipient, e *PaymentOrder) { n.Edges.PaymentOrder = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (porq *PaymentOrderRecipientQuery) loadPaymentOrder(ctx context.Context, query *PaymentOrderQuery, nodes []*PaymentOrderRecipient, init func(*PaymentOrderRecipient), assign func(*PaymentOrderRecipient, *PaymentOrder)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*PaymentOrderRecipient)
	for i := range nodes {
		if nodes[i].payment_order_recipient == nil {
			continue
		}
		fk := *nodes[i].payment_order_recipient
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(paymentorder.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "payment_order_recipient" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (porq *PaymentOrderRecipientQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := porq.querySpec()
	_spec.Node.Columns = porq.ctx.Fields
	if len(porq.ctx.Fields) > 0 {
		_spec.Unique = porq.ctx.Unique != nil && *porq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, porq.driver, _spec)
}

func (porq *PaymentOrderRecipientQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(paymentorderrecipient.Table, paymentorderrecipient.Columns, sqlgraph.NewFieldSpec(paymentorderrecipient.FieldID, field.TypeInt))
	_spec.From = porq.sql
	if unique := porq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if porq.path != nil {
		_spec.Unique = true
	}
	if fields := porq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, paymentorderrecipient.FieldID)
		for i := range fields {
			if fields[i] != paymentorderrecipient.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := porq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := porq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := porq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := porq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (porq *PaymentOrderRecipientQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(porq.driver.Dialect())
	t1 := builder.Table(paymentorderrecipient.Table)
	columns := porq.ctx.Fields
	if len(columns) == 0 {
		columns = paymentorderrecipient.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if porq.sql != nil {
		selector = porq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if porq.ctx.Unique != nil && *porq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range porq.predicates {
		p(selector)
	}
	for _, p := range porq.order {
		p(selector)
	}
	if offset := porq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := porq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PaymentOrderRecipientGroupBy is the group-by builder for PaymentOrderRecipient entities.
type PaymentOrderRecipientGroupBy struct {
	selector
	build *PaymentOrderRecipientQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (porgb *PaymentOrderRecipientGroupBy) Aggregate(fns ...AggregateFunc) *PaymentOrderRecipientGroupBy {
	porgb.fns = append(porgb.fns, fns...)
	return porgb
}

// Scan applies the selector query and scans the result into the given value.
func (porgb *PaymentOrderRecipientGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, porgb.build.ctx, ent.OpQueryGroupBy)
	if err := porgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PaymentOrderRecipientQuery, *PaymentOrderRecipientGroupBy](ctx, porgb.build, porgb, porgb.build.inters, v)
}

func (porgb *PaymentOrderRecipientGroupBy) sqlScan(ctx context.Context, root *PaymentOrderRecipientQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(porgb.fns))
	for _, fn := range porgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*porgb.flds)+len(porgb.fns))
		for _, f := range *porgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*porgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := porgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PaymentOrderRecipientSelect is the builder for selecting fields of PaymentOrderRecipient entities.
type PaymentOrderRecipientSelect struct {
	*PaymentOrderRecipientQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pors *PaymentOrderRecipientSelect) Aggregate(fns ...AggregateFunc) *PaymentOrderRecipientSelect {
	pors.fns = append(pors.fns, fns...)
	return pors
}

// Scan applies the selector query and scans the result into the given value.
func (pors *PaymentOrderRecipientSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pors.ctx, ent.OpQuerySelect)
	if err := pors.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PaymentOrderRecipientQuery, *PaymentOrderRecipientSelect](ctx, pors.PaymentOrderRecipientQuery, pors, pors.inters, v)
}

func (pors *PaymentOrderRecipientSelect) sqlScan(ctx context.Context, root *PaymentOrderRecipientQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pors.fns))
	for _, fn := range pors.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pors.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pors.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
