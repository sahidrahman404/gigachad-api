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
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/user"
	"github.com/sahidrahman404/gigachad-api/ent/workout"
	"github.com/sahidrahman404/gigachad-api/ent/workoutlog"
)

// WorkoutLogQuery is the builder for querying WorkoutLog entities.
type WorkoutLogQuery struct {
	config
	ctx           *QueryContext
	order         []workoutlog.OrderOption
	inters        []Interceptor
	predicates    []predicate.WorkoutLog
	withUsers     *UserQuery
	withExercises *ExerciseQuery
	withWorkouts  *WorkoutQuery
	modifiers     []func(*sql.Selector)
	loadTotal     []func(context.Context, []*WorkoutLog) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WorkoutLogQuery builder.
func (wlq *WorkoutLogQuery) Where(ps ...predicate.WorkoutLog) *WorkoutLogQuery {
	wlq.predicates = append(wlq.predicates, ps...)
	return wlq
}

// Limit the number of records to be returned by this query.
func (wlq *WorkoutLogQuery) Limit(limit int) *WorkoutLogQuery {
	wlq.ctx.Limit = &limit
	return wlq
}

// Offset to start from.
func (wlq *WorkoutLogQuery) Offset(offset int) *WorkoutLogQuery {
	wlq.ctx.Offset = &offset
	return wlq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wlq *WorkoutLogQuery) Unique(unique bool) *WorkoutLogQuery {
	wlq.ctx.Unique = &unique
	return wlq
}

// Order specifies how the records should be ordered.
func (wlq *WorkoutLogQuery) Order(o ...workoutlog.OrderOption) *WorkoutLogQuery {
	wlq.order = append(wlq.order, o...)
	return wlq
}

// QueryUsers chains the current query on the "users" edge.
func (wlq *WorkoutLogQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: wlq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workoutlog.Table, workoutlog.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workoutlog.UsersTable, workoutlog.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(wlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryExercises chains the current query on the "exercises" edge.
func (wlq *WorkoutLogQuery) QueryExercises() *ExerciseQuery {
	query := (&ExerciseClient{config: wlq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workoutlog.Table, workoutlog.FieldID, selector),
			sqlgraph.To(exercise.Table, exercise.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workoutlog.ExercisesTable, workoutlog.ExercisesColumn),
		)
		fromU = sqlgraph.SetNeighbors(wlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkouts chains the current query on the "workouts" edge.
func (wlq *WorkoutLogQuery) QueryWorkouts() *WorkoutQuery {
	query := (&WorkoutClient{config: wlq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workoutlog.Table, workoutlog.FieldID, selector),
			sqlgraph.To(workout.Table, workout.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workoutlog.WorkoutsTable, workoutlog.WorkoutsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WorkoutLog entity from the query.
// Returns a *NotFoundError when no WorkoutLog was found.
func (wlq *WorkoutLogQuery) First(ctx context.Context) (*WorkoutLog, error) {
	nodes, err := wlq.Limit(1).All(setContextOp(ctx, wlq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{workoutlog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wlq *WorkoutLogQuery) FirstX(ctx context.Context) *WorkoutLog {
	node, err := wlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WorkoutLog ID from the query.
// Returns a *NotFoundError when no WorkoutLog ID was found.
func (wlq *WorkoutLogQuery) FirstID(ctx context.Context) (id pksuid.ID, err error) {
	var ids []pksuid.ID
	if ids, err = wlq.Limit(1).IDs(setContextOp(ctx, wlq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{workoutlog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wlq *WorkoutLogQuery) FirstIDX(ctx context.Context) pksuid.ID {
	id, err := wlq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WorkoutLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WorkoutLog entity is found.
// Returns a *NotFoundError when no WorkoutLog entities are found.
func (wlq *WorkoutLogQuery) Only(ctx context.Context) (*WorkoutLog, error) {
	nodes, err := wlq.Limit(2).All(setContextOp(ctx, wlq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{workoutlog.Label}
	default:
		return nil, &NotSingularError{workoutlog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wlq *WorkoutLogQuery) OnlyX(ctx context.Context) *WorkoutLog {
	node, err := wlq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WorkoutLog ID in the query.
// Returns a *NotSingularError when more than one WorkoutLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (wlq *WorkoutLogQuery) OnlyID(ctx context.Context) (id pksuid.ID, err error) {
	var ids []pksuid.ID
	if ids, err = wlq.Limit(2).IDs(setContextOp(ctx, wlq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{workoutlog.Label}
	default:
		err = &NotSingularError{workoutlog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wlq *WorkoutLogQuery) OnlyIDX(ctx context.Context) pksuid.ID {
	id, err := wlq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WorkoutLogs.
func (wlq *WorkoutLogQuery) All(ctx context.Context) ([]*WorkoutLog, error) {
	ctx = setContextOp(ctx, wlq.ctx, "All")
	if err := wlq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WorkoutLog, *WorkoutLogQuery]()
	return withInterceptors[[]*WorkoutLog](ctx, wlq, qr, wlq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wlq *WorkoutLogQuery) AllX(ctx context.Context) []*WorkoutLog {
	nodes, err := wlq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WorkoutLog IDs.
func (wlq *WorkoutLogQuery) IDs(ctx context.Context) (ids []pksuid.ID, err error) {
	if wlq.ctx.Unique == nil && wlq.path != nil {
		wlq.Unique(true)
	}
	ctx = setContextOp(ctx, wlq.ctx, "IDs")
	if err = wlq.Select(workoutlog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wlq *WorkoutLogQuery) IDsX(ctx context.Context) []pksuid.ID {
	ids, err := wlq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wlq *WorkoutLogQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wlq.ctx, "Count")
	if err := wlq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wlq, querierCount[*WorkoutLogQuery](), wlq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wlq *WorkoutLogQuery) CountX(ctx context.Context) int {
	count, err := wlq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wlq *WorkoutLogQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wlq.ctx, "Exist")
	switch _, err := wlq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wlq *WorkoutLogQuery) ExistX(ctx context.Context) bool {
	exist, err := wlq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WorkoutLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wlq *WorkoutLogQuery) Clone() *WorkoutLogQuery {
	if wlq == nil {
		return nil
	}
	return &WorkoutLogQuery{
		config:        wlq.config,
		ctx:           wlq.ctx.Clone(),
		order:         append([]workoutlog.OrderOption{}, wlq.order...),
		inters:        append([]Interceptor{}, wlq.inters...),
		predicates:    append([]predicate.WorkoutLog{}, wlq.predicates...),
		withUsers:     wlq.withUsers.Clone(),
		withExercises: wlq.withExercises.Clone(),
		withWorkouts:  wlq.withWorkouts.Clone(),
		// clone intermediate query.
		sql:  wlq.sql.Clone(),
		path: wlq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (wlq *WorkoutLogQuery) WithUsers(opts ...func(*UserQuery)) *WorkoutLogQuery {
	query := (&UserClient{config: wlq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wlq.withUsers = query
	return wlq
}

// WithExercises tells the query-builder to eager-load the nodes that are connected to
// the "exercises" edge. The optional arguments are used to configure the query builder of the edge.
func (wlq *WorkoutLogQuery) WithExercises(opts ...func(*ExerciseQuery)) *WorkoutLogQuery {
	query := (&ExerciseClient{config: wlq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wlq.withExercises = query
	return wlq
}

// WithWorkouts tells the query-builder to eager-load the nodes that are connected to
// the "workouts" edge. The optional arguments are used to configure the query builder of the edge.
func (wlq *WorkoutLogQuery) WithWorkouts(opts ...func(*WorkoutQuery)) *WorkoutLogQuery {
	query := (&WorkoutClient{config: wlq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wlq.withWorkouts = query
	return wlq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Sets *schematype.Sets `json:"sets,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WorkoutLog.Query().
//		GroupBy(workoutlog.FieldSets).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wlq *WorkoutLogQuery) GroupBy(field string, fields ...string) *WorkoutLogGroupBy {
	wlq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WorkoutLogGroupBy{build: wlq}
	grbuild.flds = &wlq.ctx.Fields
	grbuild.label = workoutlog.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Sets *schematype.Sets `json:"sets,omitempty"`
//	}
//
//	client.WorkoutLog.Query().
//		Select(workoutlog.FieldSets).
//		Scan(ctx, &v)
func (wlq *WorkoutLogQuery) Select(fields ...string) *WorkoutLogSelect {
	wlq.ctx.Fields = append(wlq.ctx.Fields, fields...)
	sbuild := &WorkoutLogSelect{WorkoutLogQuery: wlq}
	sbuild.label = workoutlog.Label
	sbuild.flds, sbuild.scan = &wlq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WorkoutLogSelect configured with the given aggregations.
func (wlq *WorkoutLogQuery) Aggregate(fns ...AggregateFunc) *WorkoutLogSelect {
	return wlq.Select().Aggregate(fns...)
}

func (wlq *WorkoutLogQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wlq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wlq); err != nil {
				return err
			}
		}
	}
	for _, f := range wlq.ctx.Fields {
		if !workoutlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wlq.path != nil {
		prev, err := wlq.path(ctx)
		if err != nil {
			return err
		}
		wlq.sql = prev
	}
	return nil
}

func (wlq *WorkoutLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WorkoutLog, error) {
	var (
		nodes       = []*WorkoutLog{}
		_spec       = wlq.querySpec()
		loadedTypes = [3]bool{
			wlq.withUsers != nil,
			wlq.withExercises != nil,
			wlq.withWorkouts != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WorkoutLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WorkoutLog{config: wlq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(wlq.modifiers) > 0 {
		_spec.Modifiers = wlq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wlq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wlq.withUsers; query != nil {
		if err := wlq.loadUsers(ctx, query, nodes, nil,
			func(n *WorkoutLog, e *User) { n.Edges.Users = e }); err != nil {
			return nil, err
		}
	}
	if query := wlq.withExercises; query != nil {
		if err := wlq.loadExercises(ctx, query, nodes, nil,
			func(n *WorkoutLog, e *Exercise) { n.Edges.Exercises = e }); err != nil {
			return nil, err
		}
	}
	if query := wlq.withWorkouts; query != nil {
		if err := wlq.loadWorkouts(ctx, query, nodes, nil,
			func(n *WorkoutLog, e *Workout) { n.Edges.Workouts = e }); err != nil {
			return nil, err
		}
	}
	for i := range wlq.loadTotal {
		if err := wlq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wlq *WorkoutLogQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*WorkoutLog, init func(*WorkoutLog), assign func(*WorkoutLog, *User)) error {
	ids := make([]pksuid.ID, 0, len(nodes))
	nodeids := make(map[pksuid.ID][]*WorkoutLog)
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
func (wlq *WorkoutLogQuery) loadExercises(ctx context.Context, query *ExerciseQuery, nodes []*WorkoutLog, init func(*WorkoutLog), assign func(*WorkoutLog, *Exercise)) error {
	ids := make([]pksuid.ID, 0, len(nodes))
	nodeids := make(map[pksuid.ID][]*WorkoutLog)
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
func (wlq *WorkoutLogQuery) loadWorkouts(ctx context.Context, query *WorkoutQuery, nodes []*WorkoutLog, init func(*WorkoutLog), assign func(*WorkoutLog, *Workout)) error {
	ids := make([]pksuid.ID, 0, len(nodes))
	nodeids := make(map[pksuid.ID][]*WorkoutLog)
	for i := range nodes {
		fk := nodes[i].WorkoutID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(workout.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "workout_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (wlq *WorkoutLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wlq.querySpec()
	if len(wlq.modifiers) > 0 {
		_spec.Modifiers = wlq.modifiers
	}
	_spec.Node.Columns = wlq.ctx.Fields
	if len(wlq.ctx.Fields) > 0 {
		_spec.Unique = wlq.ctx.Unique != nil && *wlq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wlq.driver, _spec)
}

func (wlq *WorkoutLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(workoutlog.Table, workoutlog.Columns, sqlgraph.NewFieldSpec(workoutlog.FieldID, field.TypeString))
	_spec.From = wlq.sql
	if unique := wlq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wlq.path != nil {
		_spec.Unique = true
	}
	if fields := wlq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workoutlog.FieldID)
		for i := range fields {
			if fields[i] != workoutlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if wlq.withUsers != nil {
			_spec.Node.AddColumnOnce(workoutlog.FieldUserID)
		}
		if wlq.withExercises != nil {
			_spec.Node.AddColumnOnce(workoutlog.FieldExerciseID)
		}
		if wlq.withWorkouts != nil {
			_spec.Node.AddColumnOnce(workoutlog.FieldWorkoutID)
		}
	}
	if ps := wlq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wlq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wlq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wlq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wlq *WorkoutLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wlq.driver.Dialect())
	t1 := builder.Table(workoutlog.Table)
	columns := wlq.ctx.Fields
	if len(columns) == 0 {
		columns = workoutlog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wlq.sql != nil {
		selector = wlq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wlq.ctx.Unique != nil && *wlq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wlq.predicates {
		p(selector)
	}
	for _, p := range wlq.order {
		p(selector)
	}
	if offset := wlq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wlq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WorkoutLogGroupBy is the group-by builder for WorkoutLog entities.
type WorkoutLogGroupBy struct {
	selector
	build *WorkoutLogQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wlgb *WorkoutLogGroupBy) Aggregate(fns ...AggregateFunc) *WorkoutLogGroupBy {
	wlgb.fns = append(wlgb.fns, fns...)
	return wlgb
}

// Scan applies the selector query and scans the result into the given value.
func (wlgb *WorkoutLogGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wlgb.build.ctx, "GroupBy")
	if err := wlgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkoutLogQuery, *WorkoutLogGroupBy](ctx, wlgb.build, wlgb, wlgb.build.inters, v)
}

func (wlgb *WorkoutLogGroupBy) sqlScan(ctx context.Context, root *WorkoutLogQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wlgb.fns))
	for _, fn := range wlgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wlgb.flds)+len(wlgb.fns))
		for _, f := range *wlgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wlgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wlgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WorkoutLogSelect is the builder for selecting fields of WorkoutLog entities.
type WorkoutLogSelect struct {
	*WorkoutLogQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wls *WorkoutLogSelect) Aggregate(fns ...AggregateFunc) *WorkoutLogSelect {
	wls.fns = append(wls.fns, fns...)
	return wls
}

// Scan applies the selector query and scans the result into the given value.
func (wls *WorkoutLogSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wls.ctx, "Select")
	if err := wls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkoutLogQuery, *WorkoutLogSelect](ctx, wls.WorkoutLogQuery, wls, wls.inters, v)
}

func (wls *WorkoutLogSelect) sqlScan(ctx context.Context, root *WorkoutLogQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wls.fns))
	for _, fn := range wls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
