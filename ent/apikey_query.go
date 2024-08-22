// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/paycrest/protocol/ent/apikey"
	"github.com/paycrest/protocol/ent/paymentorder"
	"github.com/paycrest/protocol/ent/predicate"
	"github.com/paycrest/protocol/ent/providerprofile"
	"github.com/paycrest/protocol/ent/senderprofile"
)

// APIKeyQuery is the builder for querying APIKey entities.
type APIKeyQuery struct {
	config
	ctx                 *QueryContext
	order               []apikey.OrderOption
	inters              []Interceptor
	predicates          []predicate.APIKey
	withSenderProfile   *SenderProfileQuery
	withProviderProfile *ProviderProfileQuery
	withPaymentOrders   *PaymentOrderQuery
	withFKs             bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the APIKeyQuery builder.
func (akq *APIKeyQuery) Where(ps ...predicate.APIKey) *APIKeyQuery {
	akq.predicates = append(akq.predicates, ps...)
	return akq
}

// Limit the number of records to be returned by this query.
func (akq *APIKeyQuery) Limit(limit int) *APIKeyQuery {
	akq.ctx.Limit = &limit
	return akq
}

// Offset to start from.
func (akq *APIKeyQuery) Offset(offset int) *APIKeyQuery {
	akq.ctx.Offset = &offset
	return akq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (akq *APIKeyQuery) Unique(unique bool) *APIKeyQuery {
	akq.ctx.Unique = &unique
	return akq
}

// Order specifies how the records should be ordered.
func (akq *APIKeyQuery) Order(o ...apikey.OrderOption) *APIKeyQuery {
	akq.order = append(akq.order, o...)
	return akq
}

// QuerySenderProfile chains the current query on the "sender_profile" edge.
func (akq *APIKeyQuery) QuerySenderProfile() *SenderProfileQuery {
	query := (&SenderProfileClient{config: akq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := akq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := akq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(apikey.Table, apikey.FieldID, selector),
			sqlgraph.To(senderprofile.Table, senderprofile.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, apikey.SenderProfileTable, apikey.SenderProfileColumn),
		)
		fromU = sqlgraph.SetNeighbors(akq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProviderProfile chains the current query on the "provider_profile" edge.
func (akq *APIKeyQuery) QueryProviderProfile() *ProviderProfileQuery {
	query := (&ProviderProfileClient{config: akq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := akq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := akq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(apikey.Table, apikey.FieldID, selector),
			sqlgraph.To(providerprofile.Table, providerprofile.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, apikey.ProviderProfileTable, apikey.ProviderProfileColumn),
		)
		fromU = sqlgraph.SetNeighbors(akq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPaymentOrders chains the current query on the "payment_orders" edge.
func (akq *APIKeyQuery) QueryPaymentOrders() *PaymentOrderQuery {
	query := (&PaymentOrderClient{config: akq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := akq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := akq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(apikey.Table, apikey.FieldID, selector),
			sqlgraph.To(paymentorder.Table, paymentorder.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, apikey.PaymentOrdersTable, apikey.PaymentOrdersColumn),
		)
		fromU = sqlgraph.SetNeighbors(akq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first APIKey entity from the query.
// Returns a *NotFoundError when no APIKey was found.
func (akq *APIKeyQuery) First(ctx context.Context) (*APIKey, error) {
	nodes, err := akq.Limit(1).All(setContextOp(ctx, akq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{apikey.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (akq *APIKeyQuery) FirstX(ctx context.Context) *APIKey {
	node, err := akq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first APIKey ID from the query.
// Returns a *NotFoundError when no APIKey ID was found.
func (akq *APIKeyQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = akq.Limit(1).IDs(setContextOp(ctx, akq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{apikey.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (akq *APIKeyQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := akq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single APIKey entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one APIKey entity is found.
// Returns a *NotFoundError when no APIKey entities are found.
func (akq *APIKeyQuery) Only(ctx context.Context) (*APIKey, error) {
	nodes, err := akq.Limit(2).All(setContextOp(ctx, akq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{apikey.Label}
	default:
		return nil, &NotSingularError{apikey.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (akq *APIKeyQuery) OnlyX(ctx context.Context) *APIKey {
	node, err := akq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only APIKey ID in the query.
// Returns a *NotSingularError when more than one APIKey ID is found.
// Returns a *NotFoundError when no entities are found.
func (akq *APIKeyQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = akq.Limit(2).IDs(setContextOp(ctx, akq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{apikey.Label}
	default:
		err = &NotSingularError{apikey.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (akq *APIKeyQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := akq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of APIKeys.
func (akq *APIKeyQuery) All(ctx context.Context) ([]*APIKey, error) {
	ctx = setContextOp(ctx, akq.ctx, ent.OpQueryAll)
	if err := akq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*APIKey, *APIKeyQuery]()
	return withInterceptors[[]*APIKey](ctx, akq, qr, akq.inters)
}

// AllX is like All, but panics if an error occurs.
func (akq *APIKeyQuery) AllX(ctx context.Context) []*APIKey {
	nodes, err := akq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of APIKey IDs.
func (akq *APIKeyQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if akq.ctx.Unique == nil && akq.path != nil {
		akq.Unique(true)
	}
	ctx = setContextOp(ctx, akq.ctx, ent.OpQueryIDs)
	if err = akq.Select(apikey.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (akq *APIKeyQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := akq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (akq *APIKeyQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, akq.ctx, ent.OpQueryCount)
	if err := akq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, akq, querierCount[*APIKeyQuery](), akq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (akq *APIKeyQuery) CountX(ctx context.Context) int {
	count, err := akq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (akq *APIKeyQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, akq.ctx, ent.OpQueryExist)
	switch _, err := akq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (akq *APIKeyQuery) ExistX(ctx context.Context) bool {
	exist, err := akq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the APIKeyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (akq *APIKeyQuery) Clone() *APIKeyQuery {
	if akq == nil {
		return nil
	}
	return &APIKeyQuery{
		config:              akq.config,
		ctx:                 akq.ctx.Clone(),
		order:               append([]apikey.OrderOption{}, akq.order...),
		inters:              append([]Interceptor{}, akq.inters...),
		predicates:          append([]predicate.APIKey{}, akq.predicates...),
		withSenderProfile:   akq.withSenderProfile.Clone(),
		withProviderProfile: akq.withProviderProfile.Clone(),
		withPaymentOrders:   akq.withPaymentOrders.Clone(),
		// clone intermediate query.
		sql:  akq.sql.Clone(),
		path: akq.path,
	}
}

// WithSenderProfile tells the query-builder to eager-load the nodes that are connected to
// the "sender_profile" edge. The optional arguments are used to configure the query builder of the edge.
func (akq *APIKeyQuery) WithSenderProfile(opts ...func(*SenderProfileQuery)) *APIKeyQuery {
	query := (&SenderProfileClient{config: akq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	akq.withSenderProfile = query
	return akq
}

// WithProviderProfile tells the query-builder to eager-load the nodes that are connected to
// the "provider_profile" edge. The optional arguments are used to configure the query builder of the edge.
func (akq *APIKeyQuery) WithProviderProfile(opts ...func(*ProviderProfileQuery)) *APIKeyQuery {
	query := (&ProviderProfileClient{config: akq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	akq.withProviderProfile = query
	return akq
}

// WithPaymentOrders tells the query-builder to eager-load the nodes that are connected to
// the "payment_orders" edge. The optional arguments are used to configure the query builder of the edge.
func (akq *APIKeyQuery) WithPaymentOrders(opts ...func(*PaymentOrderQuery)) *APIKeyQuery {
	query := (&PaymentOrderClient{config: akq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	akq.withPaymentOrders = query
	return akq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Secret string `json:"secret,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.APIKey.Query().
//		GroupBy(apikey.FieldSecret).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (akq *APIKeyQuery) GroupBy(field string, fields ...string) *APIKeyGroupBy {
	akq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &APIKeyGroupBy{build: akq}
	grbuild.flds = &akq.ctx.Fields
	grbuild.label = apikey.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Secret string `json:"secret,omitempty"`
//	}
//
//	client.APIKey.Query().
//		Select(apikey.FieldSecret).
//		Scan(ctx, &v)
func (akq *APIKeyQuery) Select(fields ...string) *APIKeySelect {
	akq.ctx.Fields = append(akq.ctx.Fields, fields...)
	sbuild := &APIKeySelect{APIKeyQuery: akq}
	sbuild.label = apikey.Label
	sbuild.flds, sbuild.scan = &akq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a APIKeySelect configured with the given aggregations.
func (akq *APIKeyQuery) Aggregate(fns ...AggregateFunc) *APIKeySelect {
	return akq.Select().Aggregate(fns...)
}

func (akq *APIKeyQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range akq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, akq); err != nil {
				return err
			}
		}
	}
	for _, f := range akq.ctx.Fields {
		if !apikey.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if akq.path != nil {
		prev, err := akq.path(ctx)
		if err != nil {
			return err
		}
		akq.sql = prev
	}
	return nil
}

func (akq *APIKeyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*APIKey, error) {
	var (
		nodes       = []*APIKey{}
		withFKs     = akq.withFKs
		_spec       = akq.querySpec()
		loadedTypes = [3]bool{
			akq.withSenderProfile != nil,
			akq.withProviderProfile != nil,
			akq.withPaymentOrders != nil,
		}
	)
	if akq.withSenderProfile != nil || akq.withProviderProfile != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, apikey.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*APIKey).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &APIKey{config: akq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, akq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := akq.withSenderProfile; query != nil {
		if err := akq.loadSenderProfile(ctx, query, nodes, nil,
			func(n *APIKey, e *SenderProfile) { n.Edges.SenderProfile = e }); err != nil {
			return nil, err
		}
	}
	if query := akq.withProviderProfile; query != nil {
		if err := akq.loadProviderProfile(ctx, query, nodes, nil,
			func(n *APIKey, e *ProviderProfile) { n.Edges.ProviderProfile = e }); err != nil {
			return nil, err
		}
	}
	if query := akq.withPaymentOrders; query != nil {
		if err := akq.loadPaymentOrders(ctx, query, nodes,
			func(n *APIKey) { n.Edges.PaymentOrders = []*PaymentOrder{} },
			func(n *APIKey, e *PaymentOrder) { n.Edges.PaymentOrders = append(n.Edges.PaymentOrders, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (akq *APIKeyQuery) loadSenderProfile(ctx context.Context, query *SenderProfileQuery, nodes []*APIKey, init func(*APIKey), assign func(*APIKey, *SenderProfile)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*APIKey)
	for i := range nodes {
		if nodes[i].sender_profile_api_key == nil {
			continue
		}
		fk := *nodes[i].sender_profile_api_key
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(senderprofile.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "sender_profile_api_key" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (akq *APIKeyQuery) loadProviderProfile(ctx context.Context, query *ProviderProfileQuery, nodes []*APIKey, init func(*APIKey), assign func(*APIKey, *ProviderProfile)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*APIKey)
	for i := range nodes {
		if nodes[i].provider_profile_api_key == nil {
			continue
		}
		fk := *nodes[i].provider_profile_api_key
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(providerprofile.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "provider_profile_api_key" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (akq *APIKeyQuery) loadPaymentOrders(ctx context.Context, query *PaymentOrderQuery, nodes []*APIKey, init func(*APIKey), assign func(*APIKey, *PaymentOrder)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*APIKey)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.PaymentOrder(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(apikey.PaymentOrdersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.api_key_payment_orders
		if fk == nil {
			return fmt.Errorf(`foreign-key "api_key_payment_orders" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "api_key_payment_orders" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (akq *APIKeyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := akq.querySpec()
	_spec.Node.Columns = akq.ctx.Fields
	if len(akq.ctx.Fields) > 0 {
		_spec.Unique = akq.ctx.Unique != nil && *akq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, akq.driver, _spec)
}

func (akq *APIKeyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(apikey.Table, apikey.Columns, sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeUUID))
	_spec.From = akq.sql
	if unique := akq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if akq.path != nil {
		_spec.Unique = true
	}
	if fields := akq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apikey.FieldID)
		for i := range fields {
			if fields[i] != apikey.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := akq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := akq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := akq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := akq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (akq *APIKeyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(akq.driver.Dialect())
	t1 := builder.Table(apikey.Table)
	columns := akq.ctx.Fields
	if len(columns) == 0 {
		columns = apikey.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if akq.sql != nil {
		selector = akq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if akq.ctx.Unique != nil && *akq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range akq.predicates {
		p(selector)
	}
	for _, p := range akq.order {
		p(selector)
	}
	if offset := akq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := akq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// APIKeyGroupBy is the group-by builder for APIKey entities.
type APIKeyGroupBy struct {
	selector
	build *APIKeyQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (akgb *APIKeyGroupBy) Aggregate(fns ...AggregateFunc) *APIKeyGroupBy {
	akgb.fns = append(akgb.fns, fns...)
	return akgb
}

// Scan applies the selector query and scans the result into the given value.
func (akgb *APIKeyGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, akgb.build.ctx, ent.OpQueryGroupBy)
	if err := akgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*APIKeyQuery, *APIKeyGroupBy](ctx, akgb.build, akgb, akgb.build.inters, v)
}

func (akgb *APIKeyGroupBy) sqlScan(ctx context.Context, root *APIKeyQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(akgb.fns))
	for _, fn := range akgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*akgb.flds)+len(akgb.fns))
		for _, f := range *akgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*akgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := akgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// APIKeySelect is the builder for selecting fields of APIKey entities.
type APIKeySelect struct {
	*APIKeyQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (aks *APIKeySelect) Aggregate(fns ...AggregateFunc) *APIKeySelect {
	aks.fns = append(aks.fns, fns...)
	return aks
}

// Scan applies the selector query and scans the result into the given value.
func (aks *APIKeySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aks.ctx, ent.OpQuerySelect)
	if err := aks.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*APIKeyQuery, *APIKeySelect](ctx, aks.APIKeyQuery, aks, aks.inters, v)
}

func (aks *APIKeySelect) sqlScan(ctx context.Context, root *APIKeyQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(aks.fns))
	for _, fn := range aks.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*aks.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aks.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
