// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/exercise"
	"github.com/sahidrahman404/gigachad-api/ent/predicate"
	"github.com/sahidrahman404/gigachad-api/ent/routine"
	"github.com/sahidrahman404/gigachad-api/ent/routineexercise"
	"github.com/sahidrahman404/gigachad-api/ent/user"
)

// RoutineExerciseQuery is the builder for querying RoutineExercise entities.
type RoutineExerciseQuery struct {
	config
	ctx           *QueryContext
	order         []routineexercise.OrderOption
	inters        []Interceptor
	predicates    []predicate.RoutineExercise
	withRoutines  *RoutineQuery
	withExercises *ExerciseQuery
	withUsers     *UserQuery
	modifiers     []func(*sql.Selector)
	loadTotal     []func(context.Context, []*RoutineExercise) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RoutineExerciseQuery builder.
func (req *RoutineExerciseQuery) Where(ps ...predicate.RoutineExercise) *RoutineExerciseQuery {
	req.predicates = append(req.predicates, ps...)
	return req
}

// Limit the number of records to be returned by this query.
func (req *RoutineExerciseQuery) Limit(limit int) *RoutineExerciseQuery {
	req.ctx.Limit = &limit
	return req
}

// Offset to start from.
func (req *RoutineExerciseQuery) Offset(offset int) *RoutineExerciseQuery {
	req.ctx.Offset = &offset
	return req
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (req *RoutineExerciseQuery) Unique(unique bool) *RoutineExerciseQuery {
	req.ctx.Unique = &unique
	return req
}

// Order specifies how the records should be ordered.
func (req *RoutineExerciseQuery) Order(o ...routineexercise.OrderOption) *RoutineExerciseQuery {
	req.order = append(req.order, o...)
	return req
}

// QueryRoutines chains the current query on the "routines" edge.
func (req *RoutineExerciseQuery) QueryRoutines() *RoutineQuery {
	query := (&RoutineClient{config: req.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := req.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := req.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(routineexercise.Table, routineexercise.FieldID, selector),
			sqlgraph.To(routine.Table, routine.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, routineexercise.RoutinesTable, routineexercise.RoutinesColumn),
		)
		fromU = sqlgraph.SetNeighbors(req.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryExercises chains the current query on the "exercises" edge.
func (req *RoutineExerciseQuery) QueryExercises() *ExerciseQuery {
	query := (&ExerciseClient{config: req.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := req.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := req.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(routineexercise.Table, routineexercise.FieldID, selector),
			sqlgraph.To(exercise.Table, exercise.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, routineexercise.ExercisesTable, routineexercise.ExercisesColumn),
		)
		fromU = sqlgraph.SetNeighbors(req.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUsers chains the current query on the "users" edge.
func (req *RoutineExerciseQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: req.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := req.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := req.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(routineexercise.Table, routineexercise.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, routineexercise.UsersTable, routineexercise.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(req.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RoutineExercise entity from the query.
// Returns a *NotFoundError when no RoutineExercise was found.
func (req *RoutineExerciseQuery) First(ctx context.Context) (*RoutineExercise, error) {
	nodes, err := req.Limit(1).All(setContextOp(ctx, req.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{routineexercise.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (req *RoutineExerciseQuery) FirstX(ctx context.Context) *RoutineExercise {
	node, err := req.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RoutineExercise ID from the query.
// Returns a *NotFoundError when no RoutineExercise ID was found.
func (req *RoutineExerciseQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = req.Limit(1).IDs(setContextOp(ctx, req.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{routineexercise.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (req *RoutineExerciseQuery) FirstIDX(ctx context.Context) string {
	id, err := req.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RoutineExercise entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RoutineExercise entity is found.
// Returns a *NotFoundError when no RoutineExercise entities are found.
func (req *RoutineExerciseQuery) Only(ctx context.Context) (*RoutineExercise, error) {
	nodes, err := req.Limit(2).All(setContextOp(ctx, req.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{routineexercise.Label}
	default:
		return nil, &NotSingularError{routineexercise.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (req *RoutineExerciseQuery) OnlyX(ctx context.Context) *RoutineExercise {
	node, err := req.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RoutineExercise ID in the query.
// Returns a *NotSingularError when more than one RoutineExercise ID is found.
// Returns a *NotFoundError when no entities are found.
func (req *RoutineExerciseQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = req.Limit(2).IDs(setContextOp(ctx, req.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{routineexercise.Label}
	default:
		err = &NotSingularError{routineexercise.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (req *RoutineExerciseQuery) OnlyIDX(ctx context.Context) string {
	id, err := req.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RoutineExercises.
func (req *RoutineExerciseQuery) All(ctx context.Context) ([]*RoutineExercise, error) {
	ctx = setContextOp(ctx, req.ctx, "All")
	if err := req.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RoutineExercise, *RoutineExerciseQuery]()
	return withInterceptors[[]*RoutineExercise](ctx, req, qr, req.inters)
}

// AllX is like All, but panics if an error occurs.
func (req *RoutineExerciseQuery) AllX(ctx context.Context) []*RoutineExercise {
	nodes, err := req.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RoutineExercise IDs.
func (req *RoutineExerciseQuery) IDs(ctx context.Context) (ids []string, err error) {
	if req.ctx.Unique == nil && req.path != nil {
		req.Unique(true)
	}
	ctx = setContextOp(ctx, req.ctx, "IDs")
	if err = req.Select(routineexercise.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (req *RoutineExerciseQuery) IDsX(ctx context.Context) []string {
	ids, err := req.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (req *RoutineExerciseQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, req.ctx, "Count")
	if err := req.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, req, querierCount[*RoutineExerciseQuery](), req.inters)
}

// CountX is like Count, but panics if an error occurs.
func (req *RoutineExerciseQuery) CountX(ctx context.Context) int {
	count, err := req.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (req *RoutineExerciseQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, req.ctx, "Exist")
	switch _, err := req.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (req *RoutineExerciseQuery) ExistX(ctx context.Context) bool {
	exist, err := req.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RoutineExerciseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (req *RoutineExerciseQuery) Clone() *RoutineExerciseQuery {
	if req == nil {
		return nil
	}
	return &RoutineExerciseQuery{
		config:        req.config,
		ctx:           req.ctx.Clone(),
		order:         append([]routineexercise.OrderOption{}, req.order...),
		inters:        append([]Interceptor{}, req.inters...),
		predicates:    append([]predicate.RoutineExercise{}, req.predicates...),
		withRoutines:  req.withRoutines.Clone(),
		withExercises: req.withExercises.Clone(),
		withUsers:     req.withUsers.Clone(),
		// clone intermediate query.
		sql:  req.sql.Clone(),
		path: req.path,
	}
}

// WithRoutines tells the query-builder to eager-load the nodes that are connected to
// the "routines" edge. The optional arguments are used to configure the query builder of the edge.
func (req *RoutineExerciseQuery) WithRoutines(opts ...func(*RoutineQuery)) *RoutineExerciseQuery {
	query := (&RoutineClient{config: req.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	req.withRoutines = query
	return req
}

// WithExercises tells the query-builder to eager-load the nodes that are connected to
// the "exercises" edge. The optional arguments are used to configure the query builder of the edge.
func (req *RoutineExerciseQuery) WithExercises(opts ...func(*ExerciseQuery)) *RoutineExerciseQuery {
	query := (&ExerciseClient{config: req.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	req.withExercises = query
	return req
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (req *RoutineExerciseQuery) WithUsers(opts ...func(*UserQuery)) *RoutineExerciseQuery {
	query := (&UserClient{config: req.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	req.withUsers = query
	return req
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		RestTimer int `json:"rest_timer,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RoutineExercise.Query().
//		GroupBy(routineexercise.FieldRestTimer).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (req *RoutineExerciseQuery) GroupBy(field string, fields ...string) *RoutineExerciseGroupBy {
	req.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RoutineExerciseGroupBy{build: req}
	grbuild.flds = &req.ctx.Fields
	grbuild.label = routineexercise.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		RestTimer int `json:"rest_timer,omitempty"`
//	}
//
//	client.RoutineExercise.Query().
//		Select(routineexercise.FieldRestTimer).
//		Scan(ctx, &v)
func (req *RoutineExerciseQuery) Select(fields ...string) *RoutineExerciseSelect {
	req.ctx.Fields = append(req.ctx.Fields, fields...)
	sbuild := &RoutineExerciseSelect{RoutineExerciseQuery: req}
	sbuild.label = routineexercise.Label
	sbuild.flds, sbuild.scan = &req.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RoutineExerciseSelect configured with the given aggregations.
func (req *RoutineExerciseQuery) Aggregate(fns ...AggregateFunc) *RoutineExerciseSelect {
	return req.Select().Aggregate(fns...)
}

func (req *RoutineExerciseQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range req.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, req); err != nil {
				return err
			}
		}
	}
	for _, f := range req.ctx.Fields {
		if !routineexercise.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if req.path != nil {
		prev, err := req.path(ctx)
		if err != nil {
			return err
		}
		req.sql = prev
	}
	return nil
}

func (req *RoutineExerciseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RoutineExercise, error) {
	var (
		nodes       = []*RoutineExercise{}
		_spec       = req.querySpec()
		loadedTypes = [3]bool{
			req.withRoutines != nil,
			req.withExercises != nil,
			req.withUsers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RoutineExercise).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RoutineExercise{config: req.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(req.modifiers) > 0 {
		_spec.Modifiers = req.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, req.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := req.withRoutines; query != nil {
		if err := req.loadRoutines(ctx, query, nodes, nil,
			func(n *RoutineExercise, e *Routine) { n.Edges.Routines = e }); err != nil {
			return nil, err
		}
	}
	if query := req.withExercises; query != nil {
		if err := req.loadExercises(ctx, query, nodes, nil,
			func(n *RoutineExercise, e *Exercise) { n.Edges.Exercises = e }); err != nil {
			return nil, err
		}
	}
	if query := req.withUsers; query != nil {
		if err := req.loadUsers(ctx, query, nodes, nil,
			func(n *RoutineExercise, e *User) { n.Edges.Users = e }); err != nil {
			return nil, err
		}
	}
	for i := range req.loadTotal {
		if err := req.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (req *RoutineExerciseQuery) loadRoutines(ctx context.Context, query *RoutineQuery, nodes []*RoutineExercise, init func(*RoutineExercise), assign func(*RoutineExercise, *Routine)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*RoutineExercise)
	for i := range nodes {
		fk := nodes[i].RoutineID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(routine.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "routine_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (req *RoutineExerciseQuery) loadExercises(ctx context.Context, query *ExerciseQuery, nodes []*RoutineExercise, init func(*RoutineExercise), assign func(*RoutineExercise, *Exercise)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*RoutineExercise)
	for i := range nodes {
		fk := nodes[i].ExerciseID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(exercise.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "exercise_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (req *RoutineExerciseQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*RoutineExercise, init func(*RoutineExercise), assign func(*RoutineExercise, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*RoutineExercise)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (req *RoutineExerciseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := req.querySpec()
	if len(req.modifiers) > 0 {
		_spec.Modifiers = req.modifiers
	}
	_spec.Node.Columns = req.ctx.Fields
	if len(req.ctx.Fields) > 0 {
		_spec.Unique = req.ctx.Unique != nil && *req.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, req.driver, _spec)
}

func (req *RoutineExerciseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(routineexercise.Table, routineexercise.Columns, sqlgraph.NewFieldSpec(routineexercise.FieldID, field.TypeString))
	_spec.From = req.sql
	if unique := req.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if req.path != nil {
		_spec.Unique = true
	}
	if fields := req.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, routineexercise.FieldID)
		for i := range fields {
			if fields[i] != routineexercise.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if req.withRoutines != nil {
			_spec.Node.AddColumnOnce(routineexercise.FieldRoutineID)
		}
		if req.withExercises != nil {
			_spec.Node.AddColumnOnce(routineexercise.FieldExerciseID)
		}
		if req.withUsers != nil {
			_spec.Node.AddColumnOnce(routineexercise.FieldUserID)
		}
	}
	if ps := req.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := req.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := req.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := req.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (req *RoutineExerciseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(req.driver.Dialect())
	t1 := builder.Table(routineexercise.Table)
	columns := req.ctx.Fields
	if len(columns) == 0 {
		columns = routineexercise.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if req.sql != nil {
		selector = req.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if req.ctx.Unique != nil && *req.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range req.predicates {
		p(selector)
	}
	for _, p := range req.order {
		p(selector)
	}
	if offset := req.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := req.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RoutineExerciseGroupBy is the group-by builder for RoutineExercise entities.
type RoutineExerciseGroupBy struct {
	selector
	build *RoutineExerciseQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (regb *RoutineExerciseGroupBy) Aggregate(fns ...AggregateFunc) *RoutineExerciseGroupBy {
	regb.fns = append(regb.fns, fns...)
	return regb
}

// Scan applies the selector query and scans the result into the given value.
func (regb *RoutineExerciseGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, regb.build.ctx, "GroupBy")
	if err := regb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoutineExerciseQuery, *RoutineExerciseGroupBy](ctx, regb.build, regb, regb.build.inters, v)
}

func (regb *RoutineExerciseGroupBy) sqlScan(ctx context.Context, root *RoutineExerciseQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(regb.fns))
	for _, fn := range regb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*regb.flds)+len(regb.fns))
		for _, f := range *regb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*regb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := regb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RoutineExerciseSelect is the builder for selecting fields of RoutineExercise entities.
type RoutineExerciseSelect struct {
	*RoutineExerciseQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (res *RoutineExerciseSelect) Aggregate(fns ...AggregateFunc) *RoutineExerciseSelect {
	res.fns = append(res.fns, fns...)
	return res
}

// Scan applies the selector query and scans the result into the given value.
func (res *RoutineExerciseSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, res.ctx, "Select")
	if err := res.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoutineExerciseQuery, *RoutineExerciseSelect](ctx, res.RoutineExerciseQuery, res, res.inters, v)
}

func (res *RoutineExerciseSelect) sqlScan(ctx context.Context, root *RoutineExerciseQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(res.fns))
	for _, fn := range res.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*res.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := res.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
