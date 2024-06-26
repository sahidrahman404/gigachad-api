// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
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

// WorkoutQuery is the builder for querying Workout entities.
type WorkoutQuery struct {
	config
	ctx                  *QueryContext
	order                []workout.OrderOption
	inters               []Interceptor
	predicates           []predicate.Workout
	withUsers            *UserQuery
	withExercises        *ExerciseQuery
	withWorkoutLogs      *WorkoutLogQuery
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*Workout) error
	withNamedExercises   map[string]*ExerciseQuery
	withNamedWorkoutLogs map[string]*WorkoutLogQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WorkoutQuery builder.
func (wq *WorkoutQuery) Where(ps ...predicate.Workout) *WorkoutQuery {
	wq.predicates = append(wq.predicates, ps...)
	return wq
}

// Limit the number of records to be returned by this query.
func (wq *WorkoutQuery) Limit(limit int) *WorkoutQuery {
	wq.ctx.Limit = &limit
	return wq
}

// Offset to start from.
func (wq *WorkoutQuery) Offset(offset int) *WorkoutQuery {
	wq.ctx.Offset = &offset
	return wq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wq *WorkoutQuery) Unique(unique bool) *WorkoutQuery {
	wq.ctx.Unique = &unique
	return wq
}

// Order specifies how the records should be ordered.
func (wq *WorkoutQuery) Order(o ...workout.OrderOption) *WorkoutQuery {
	wq.order = append(wq.order, o...)
	return wq
}

// QueryUsers chains the current query on the "users" edge.
func (wq *WorkoutQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workout.Table, workout.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workout.UsersTable, workout.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryExercises chains the current query on the "exercises" edge.
func (wq *WorkoutQuery) QueryExercises() *ExerciseQuery {
	query := (&ExerciseClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workout.Table, workout.FieldID, selector),
			sqlgraph.To(exercise.Table, exercise.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, workout.ExercisesTable, workout.ExercisesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkoutLogs chains the current query on the "workout_logs" edge.
func (wq *WorkoutQuery) QueryWorkoutLogs() *WorkoutLogQuery {
	query := (&WorkoutLogClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workout.Table, workout.FieldID, selector),
			sqlgraph.To(workoutlog.Table, workoutlog.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, workout.WorkoutLogsTable, workout.WorkoutLogsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Workout entity from the query.
// Returns a *NotFoundError when no Workout was found.
func (wq *WorkoutQuery) First(ctx context.Context) (*Workout, error) {
	nodes, err := wq.Limit(1).All(setContextOp(ctx, wq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{workout.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wq *WorkoutQuery) FirstX(ctx context.Context) *Workout {
	node, err := wq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Workout ID from the query.
// Returns a *NotFoundError when no Workout ID was found.
func (wq *WorkoutQuery) FirstID(ctx context.Context) (id pksuid.ID, err error) {
	var ids []pksuid.ID
	if ids, err = wq.Limit(1).IDs(setContextOp(ctx, wq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{workout.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wq *WorkoutQuery) FirstIDX(ctx context.Context) pksuid.ID {
	id, err := wq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Workout entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Workout entity is found.
// Returns a *NotFoundError when no Workout entities are found.
func (wq *WorkoutQuery) Only(ctx context.Context) (*Workout, error) {
	nodes, err := wq.Limit(2).All(setContextOp(ctx, wq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{workout.Label}
	default:
		return nil, &NotSingularError{workout.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wq *WorkoutQuery) OnlyX(ctx context.Context) *Workout {
	node, err := wq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Workout ID in the query.
// Returns a *NotSingularError when more than one Workout ID is found.
// Returns a *NotFoundError when no entities are found.
func (wq *WorkoutQuery) OnlyID(ctx context.Context) (id pksuid.ID, err error) {
	var ids []pksuid.ID
	if ids, err = wq.Limit(2).IDs(setContextOp(ctx, wq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{workout.Label}
	default:
		err = &NotSingularError{workout.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wq *WorkoutQuery) OnlyIDX(ctx context.Context) pksuid.ID {
	id, err := wq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Workouts.
func (wq *WorkoutQuery) All(ctx context.Context) ([]*Workout, error) {
	ctx = setContextOp(ctx, wq.ctx, "All")
	if err := wq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Workout, *WorkoutQuery]()
	return withInterceptors[[]*Workout](ctx, wq, qr, wq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wq *WorkoutQuery) AllX(ctx context.Context) []*Workout {
	nodes, err := wq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Workout IDs.
func (wq *WorkoutQuery) IDs(ctx context.Context) (ids []pksuid.ID, err error) {
	if wq.ctx.Unique == nil && wq.path != nil {
		wq.Unique(true)
	}
	ctx = setContextOp(ctx, wq.ctx, "IDs")
	if err = wq.Select(workout.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wq *WorkoutQuery) IDsX(ctx context.Context) []pksuid.ID {
	ids, err := wq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wq *WorkoutQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wq.ctx, "Count")
	if err := wq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wq, querierCount[*WorkoutQuery](), wq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wq *WorkoutQuery) CountX(ctx context.Context) int {
	count, err := wq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wq *WorkoutQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wq.ctx, "Exist")
	switch _, err := wq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wq *WorkoutQuery) ExistX(ctx context.Context) bool {
	exist, err := wq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WorkoutQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wq *WorkoutQuery) Clone() *WorkoutQuery {
	if wq == nil {
		return nil
	}
	return &WorkoutQuery{
		config:          wq.config,
		ctx:             wq.ctx.Clone(),
		order:           append([]workout.OrderOption{}, wq.order...),
		inters:          append([]Interceptor{}, wq.inters...),
		predicates:      append([]predicate.Workout{}, wq.predicates...),
		withUsers:       wq.withUsers.Clone(),
		withExercises:   wq.withExercises.Clone(),
		withWorkoutLogs: wq.withWorkoutLogs.Clone(),
		// clone intermediate query.
		sql:  wq.sql.Clone(),
		path: wq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkoutQuery) WithUsers(opts ...func(*UserQuery)) *WorkoutQuery {
	query := (&UserClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withUsers = query
	return wq
}

// WithExercises tells the query-builder to eager-load the nodes that are connected to
// the "exercises" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkoutQuery) WithExercises(opts ...func(*ExerciseQuery)) *WorkoutQuery {
	query := (&ExerciseClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withExercises = query
	return wq
}

// WithWorkoutLogs tells the query-builder to eager-load the nodes that are connected to
// the "workout_logs" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkoutQuery) WithWorkoutLogs(opts ...func(*WorkoutLogQuery)) *WorkoutQuery {
	query := (&WorkoutLogClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withWorkoutLogs = query
	return wq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Workout.Query().
//		GroupBy(workout.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wq *WorkoutQuery) GroupBy(field string, fields ...string) *WorkoutGroupBy {
	wq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WorkoutGroupBy{build: wq}
	grbuild.flds = &wq.ctx.Fields
	grbuild.label = workout.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Workout.Query().
//		Select(workout.FieldName).
//		Scan(ctx, &v)
func (wq *WorkoutQuery) Select(fields ...string) *WorkoutSelect {
	wq.ctx.Fields = append(wq.ctx.Fields, fields...)
	sbuild := &WorkoutSelect{WorkoutQuery: wq}
	sbuild.label = workout.Label
	sbuild.flds, sbuild.scan = &wq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WorkoutSelect configured with the given aggregations.
func (wq *WorkoutQuery) Aggregate(fns ...AggregateFunc) *WorkoutSelect {
	return wq.Select().Aggregate(fns...)
}

func (wq *WorkoutQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wq); err != nil {
				return err
			}
		}
	}
	for _, f := range wq.ctx.Fields {
		if !workout.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wq.path != nil {
		prev, err := wq.path(ctx)
		if err != nil {
			return err
		}
		wq.sql = prev
	}
	if workout.Policy == nil {
		return errors.New("ent: uninitialized workout.Policy (forgotten import ent/runtime?)")
	}
	if err := workout.Policy.EvalQuery(ctx, wq); err != nil {
		return err
	}
	return nil
}

func (wq *WorkoutQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Workout, error) {
	var (
		nodes       = []*Workout{}
		_spec       = wq.querySpec()
		loadedTypes = [3]bool{
			wq.withUsers != nil,
			wq.withExercises != nil,
			wq.withWorkoutLogs != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Workout).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Workout{config: wq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(wq.modifiers) > 0 {
		_spec.Modifiers = wq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wq.withUsers; query != nil {
		if err := wq.loadUsers(ctx, query, nodes, nil,
			func(n *Workout, e *User) { n.Edges.Users = e }); err != nil {
			return nil, err
		}
	}
	if query := wq.withExercises; query != nil {
		if err := wq.loadExercises(ctx, query, nodes,
			func(n *Workout) { n.Edges.Exercises = []*Exercise{} },
			func(n *Workout, e *Exercise) { n.Edges.Exercises = append(n.Edges.Exercises, e) }); err != nil {
			return nil, err
		}
	}
	if query := wq.withWorkoutLogs; query != nil {
		if err := wq.loadWorkoutLogs(ctx, query, nodes,
			func(n *Workout) { n.Edges.WorkoutLogs = []*WorkoutLog{} },
			func(n *Workout, e *WorkoutLog) { n.Edges.WorkoutLogs = append(n.Edges.WorkoutLogs, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range wq.withNamedExercises {
		if err := wq.loadExercises(ctx, query, nodes,
			func(n *Workout) { n.appendNamedExercises(name) },
			func(n *Workout, e *Exercise) { n.appendNamedExercises(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range wq.withNamedWorkoutLogs {
		if err := wq.loadWorkoutLogs(ctx, query, nodes,
			func(n *Workout) { n.appendNamedWorkoutLogs(name) },
			func(n *Workout, e *WorkoutLog) { n.appendNamedWorkoutLogs(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range wq.loadTotal {
		if err := wq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wq *WorkoutQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*Workout, init func(*Workout), assign func(*Workout, *User)) error {
	ids := make([]pksuid.ID, 0, len(nodes))
	nodeids := make(map[pksuid.ID][]*Workout)
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
func (wq *WorkoutQuery) loadExercises(ctx context.Context, query *ExerciseQuery, nodes []*Workout, init func(*Workout), assign func(*Workout, *Exercise)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[pksuid.ID]*Workout)
	nids := make(map[pksuid.ID]map[*Workout]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(workout.ExercisesTable)
		s.Join(joinT).On(s.C(exercise.FieldID), joinT.C(workout.ExercisesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(workout.ExercisesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(workout.ExercisesPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(pksuid.ID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*pksuid.ID)
				inValue := *values[1].(*pksuid.ID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Workout]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Exercise](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "exercises" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (wq *WorkoutQuery) loadWorkoutLogs(ctx context.Context, query *WorkoutLogQuery, nodes []*Workout, init func(*Workout), assign func(*Workout, *WorkoutLog)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[pksuid.ID]*Workout)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(workoutlog.FieldWorkoutID)
	}
	query.Where(predicate.WorkoutLog(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(workout.WorkoutLogsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.WorkoutID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "workout_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (wq *WorkoutQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wq.querySpec()
	if len(wq.modifiers) > 0 {
		_spec.Modifiers = wq.modifiers
	}
	_spec.Node.Columns = wq.ctx.Fields
	if len(wq.ctx.Fields) > 0 {
		_spec.Unique = wq.ctx.Unique != nil && *wq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wq.driver, _spec)
}

func (wq *WorkoutQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(workout.Table, workout.Columns, sqlgraph.NewFieldSpec(workout.FieldID, field.TypeString))
	_spec.From = wq.sql
	if unique := wq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wq.path != nil {
		_spec.Unique = true
	}
	if fields := wq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workout.FieldID)
		for i := range fields {
			if fields[i] != workout.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if wq.withUsers != nil {
			_spec.Node.AddColumnOnce(workout.FieldUserID)
		}
	}
	if ps := wq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wq *WorkoutQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wq.driver.Dialect())
	t1 := builder.Table(workout.Table)
	columns := wq.ctx.Fields
	if len(columns) == 0 {
		columns = workout.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wq.sql != nil {
		selector = wq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wq.ctx.Unique != nil && *wq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wq.predicates {
		p(selector)
	}
	for _, p := range wq.order {
		p(selector)
	}
	if offset := wq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedExercises tells the query-builder to eager-load the nodes that are connected to the "exercises"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkoutQuery) WithNamedExercises(name string, opts ...func(*ExerciseQuery)) *WorkoutQuery {
	query := (&ExerciseClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if wq.withNamedExercises == nil {
		wq.withNamedExercises = make(map[string]*ExerciseQuery)
	}
	wq.withNamedExercises[name] = query
	return wq
}

// WithNamedWorkoutLogs tells the query-builder to eager-load the nodes that are connected to the "workout_logs"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkoutQuery) WithNamedWorkoutLogs(name string, opts ...func(*WorkoutLogQuery)) *WorkoutQuery {
	query := (&WorkoutLogClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if wq.withNamedWorkoutLogs == nil {
		wq.withNamedWorkoutLogs = make(map[string]*WorkoutLogQuery)
	}
	wq.withNamedWorkoutLogs[name] = query
	return wq
}

// WorkoutGroupBy is the group-by builder for Workout entities.
type WorkoutGroupBy struct {
	selector
	build *WorkoutQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wgb *WorkoutGroupBy) Aggregate(fns ...AggregateFunc) *WorkoutGroupBy {
	wgb.fns = append(wgb.fns, fns...)
	return wgb
}

// Scan applies the selector query and scans the result into the given value.
func (wgb *WorkoutGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wgb.build.ctx, "GroupBy")
	if err := wgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkoutQuery, *WorkoutGroupBy](ctx, wgb.build, wgb, wgb.build.inters, v)
}

func (wgb *WorkoutGroupBy) sqlScan(ctx context.Context, root *WorkoutQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wgb.fns))
	for _, fn := range wgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wgb.flds)+len(wgb.fns))
		for _, f := range *wgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WorkoutSelect is the builder for selecting fields of Workout entities.
type WorkoutSelect struct {
	*WorkoutQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ws *WorkoutSelect) Aggregate(fns ...AggregateFunc) *WorkoutSelect {
	ws.fns = append(ws.fns, fns...)
	return ws
}

// Scan applies the selector query and scans the result into the given value.
func (ws *WorkoutSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ws.ctx, "Select")
	if err := ws.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkoutQuery, *WorkoutSelect](ctx, ws.WorkoutQuery, ws, ws.inters, v)
}

func (ws *WorkoutSelect) sqlScan(ctx context.Context, root *WorkoutQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ws.fns))
	for _, fn := range ws.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ws.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
