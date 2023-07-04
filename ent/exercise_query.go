// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/equipment"
	"github.com/sahidrahman404/gigachad-api/ent/exercise"
	"github.com/sahidrahman404/gigachad-api/ent/exercisetype"
	"github.com/sahidrahman404/gigachad-api/ent/musclesgroup"
	"github.com/sahidrahman404/gigachad-api/ent/predicate"
	"github.com/sahidrahman404/gigachad-api/ent/routine"
	"github.com/sahidrahman404/gigachad-api/ent/routineexercise"
	"github.com/sahidrahman404/gigachad-api/ent/user"
	"github.com/sahidrahman404/gigachad-api/ent/workoutlog"
)

// ExerciseQuery is the builder for querying Exercise entities.
type ExerciseQuery struct {
	config
	ctx                  *QueryContext
	order                []exercise.OrderOption
	inters               []Interceptor
	predicates           []predicate.Exercise
	withWorkoutLogs      *WorkoutLogQuery
	withUsers            *UserQuery
	withEquipments       *EquipmentQuery
	withMusclesGroups    *MusclesGroupQuery
	withExerciseTypes    *ExerciseTypeQuery
	withRoutines         *RoutineQuery
	withRoutineExercises *RoutineExerciseQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ExerciseQuery builder.
func (eq *ExerciseQuery) Where(ps ...predicate.Exercise) *ExerciseQuery {
	eq.predicates = append(eq.predicates, ps...)
	return eq
}

// Limit the number of records to be returned by this query.
func (eq *ExerciseQuery) Limit(limit int) *ExerciseQuery {
	eq.ctx.Limit = &limit
	return eq
}

// Offset to start from.
func (eq *ExerciseQuery) Offset(offset int) *ExerciseQuery {
	eq.ctx.Offset = &offset
	return eq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eq *ExerciseQuery) Unique(unique bool) *ExerciseQuery {
	eq.ctx.Unique = &unique
	return eq
}

// Order specifies how the records should be ordered.
func (eq *ExerciseQuery) Order(o ...exercise.OrderOption) *ExerciseQuery {
	eq.order = append(eq.order, o...)
	return eq
}

// QueryWorkoutLogs chains the current query on the "workout_logs" edge.
func (eq *ExerciseQuery) QueryWorkoutLogs() *WorkoutLogQuery {
	query := (&WorkoutLogClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(exercise.Table, exercise.FieldID, selector),
			sqlgraph.To(workoutlog.Table, workoutlog.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, exercise.WorkoutLogsTable, exercise.WorkoutLogsColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUsers chains the current query on the "users" edge.
func (eq *ExerciseQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(exercise.Table, exercise.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, exercise.UsersTable, exercise.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEquipments chains the current query on the "equipments" edge.
func (eq *ExerciseQuery) QueryEquipments() *EquipmentQuery {
	query := (&EquipmentClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(exercise.Table, exercise.FieldID, selector),
			sqlgraph.To(equipment.Table, equipment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, exercise.EquipmentsTable, exercise.EquipmentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMusclesGroups chains the current query on the "muscles_groups" edge.
func (eq *ExerciseQuery) QueryMusclesGroups() *MusclesGroupQuery {
	query := (&MusclesGroupClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(exercise.Table, exercise.FieldID, selector),
			sqlgraph.To(musclesgroup.Table, musclesgroup.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, exercise.MusclesGroupsTable, exercise.MusclesGroupsColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryExerciseTypes chains the current query on the "exercise_types" edge.
func (eq *ExerciseQuery) QueryExerciseTypes() *ExerciseTypeQuery {
	query := (&ExerciseTypeClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(exercise.Table, exercise.FieldID, selector),
			sqlgraph.To(exercisetype.Table, exercisetype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, exercise.ExerciseTypesTable, exercise.ExerciseTypesColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRoutines chains the current query on the "routines" edge.
func (eq *ExerciseQuery) QueryRoutines() *RoutineQuery {
	query := (&RoutineClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(exercise.Table, exercise.FieldID, selector),
			sqlgraph.To(routine.Table, routine.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, exercise.RoutinesTable, exercise.RoutinesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRoutineExercises chains the current query on the "routine_exercises" edge.
func (eq *ExerciseQuery) QueryRoutineExercises() *RoutineExerciseQuery {
	query := (&RoutineExerciseClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(exercise.Table, exercise.FieldID, selector),
			sqlgraph.To(routineexercise.Table, routineexercise.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, exercise.RoutineExercisesTable, exercise.RoutineExercisesColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Exercise entity from the query.
// Returns a *NotFoundError when no Exercise was found.
func (eq *ExerciseQuery) First(ctx context.Context) (*Exercise, error) {
	nodes, err := eq.Limit(1).All(setContextOp(ctx, eq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{exercise.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eq *ExerciseQuery) FirstX(ctx context.Context) *Exercise {
	node, err := eq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Exercise ID from the query.
// Returns a *NotFoundError when no Exercise ID was found.
func (eq *ExerciseQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = eq.Limit(1).IDs(setContextOp(ctx, eq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{exercise.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eq *ExerciseQuery) FirstIDX(ctx context.Context) string {
	id, err := eq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Exercise entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Exercise entity is found.
// Returns a *NotFoundError when no Exercise entities are found.
func (eq *ExerciseQuery) Only(ctx context.Context) (*Exercise, error) {
	nodes, err := eq.Limit(2).All(setContextOp(ctx, eq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{exercise.Label}
	default:
		return nil, &NotSingularError{exercise.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eq *ExerciseQuery) OnlyX(ctx context.Context) *Exercise {
	node, err := eq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Exercise ID in the query.
// Returns a *NotSingularError when more than one Exercise ID is found.
// Returns a *NotFoundError when no entities are found.
func (eq *ExerciseQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = eq.Limit(2).IDs(setContextOp(ctx, eq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{exercise.Label}
	default:
		err = &NotSingularError{exercise.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eq *ExerciseQuery) OnlyIDX(ctx context.Context) string {
	id, err := eq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Exercises.
func (eq *ExerciseQuery) All(ctx context.Context) ([]*Exercise, error) {
	ctx = setContextOp(ctx, eq.ctx, "All")
	if err := eq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Exercise, *ExerciseQuery]()
	return withInterceptors[[]*Exercise](ctx, eq, qr, eq.inters)
}

// AllX is like All, but panics if an error occurs.
func (eq *ExerciseQuery) AllX(ctx context.Context) []*Exercise {
	nodes, err := eq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Exercise IDs.
func (eq *ExerciseQuery) IDs(ctx context.Context) (ids []string, err error) {
	if eq.ctx.Unique == nil && eq.path != nil {
		eq.Unique(true)
	}
	ctx = setContextOp(ctx, eq.ctx, "IDs")
	if err = eq.Select(exercise.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eq *ExerciseQuery) IDsX(ctx context.Context) []string {
	ids, err := eq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eq *ExerciseQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, eq.ctx, "Count")
	if err := eq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, eq, querierCount[*ExerciseQuery](), eq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (eq *ExerciseQuery) CountX(ctx context.Context) int {
	count, err := eq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eq *ExerciseQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, eq.ctx, "Exist")
	switch _, err := eq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (eq *ExerciseQuery) ExistX(ctx context.Context) bool {
	exist, err := eq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ExerciseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eq *ExerciseQuery) Clone() *ExerciseQuery {
	if eq == nil {
		return nil
	}
	return &ExerciseQuery{
		config:               eq.config,
		ctx:                  eq.ctx.Clone(),
		order:                append([]exercise.OrderOption{}, eq.order...),
		inters:               append([]Interceptor{}, eq.inters...),
		predicates:           append([]predicate.Exercise{}, eq.predicates...),
		withWorkoutLogs:      eq.withWorkoutLogs.Clone(),
		withUsers:            eq.withUsers.Clone(),
		withEquipments:       eq.withEquipments.Clone(),
		withMusclesGroups:    eq.withMusclesGroups.Clone(),
		withExerciseTypes:    eq.withExerciseTypes.Clone(),
		withRoutines:         eq.withRoutines.Clone(),
		withRoutineExercises: eq.withRoutineExercises.Clone(),
		// clone intermediate query.
		sql:  eq.sql.Clone(),
		path: eq.path,
	}
}

// WithWorkoutLogs tells the query-builder to eager-load the nodes that are connected to
// the "workout_logs" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *ExerciseQuery) WithWorkoutLogs(opts ...func(*WorkoutLogQuery)) *ExerciseQuery {
	query := (&WorkoutLogClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withWorkoutLogs = query
	return eq
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *ExerciseQuery) WithUsers(opts ...func(*UserQuery)) *ExerciseQuery {
	query := (&UserClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withUsers = query
	return eq
}

// WithEquipments tells the query-builder to eager-load the nodes that are connected to
// the "equipments" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *ExerciseQuery) WithEquipments(opts ...func(*EquipmentQuery)) *ExerciseQuery {
	query := (&EquipmentClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withEquipments = query
	return eq
}

// WithMusclesGroups tells the query-builder to eager-load the nodes that are connected to
// the "muscles_groups" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *ExerciseQuery) WithMusclesGroups(opts ...func(*MusclesGroupQuery)) *ExerciseQuery {
	query := (&MusclesGroupClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withMusclesGroups = query
	return eq
}

// WithExerciseTypes tells the query-builder to eager-load the nodes that are connected to
// the "exercise_types" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *ExerciseQuery) WithExerciseTypes(opts ...func(*ExerciseTypeQuery)) *ExerciseQuery {
	query := (&ExerciseTypeClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withExerciseTypes = query
	return eq
}

// WithRoutines tells the query-builder to eager-load the nodes that are connected to
// the "routines" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *ExerciseQuery) WithRoutines(opts ...func(*RoutineQuery)) *ExerciseQuery {
	query := (&RoutineClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withRoutines = query
	return eq
}

// WithRoutineExercises tells the query-builder to eager-load the nodes that are connected to
// the "routine_exercises" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *ExerciseQuery) WithRoutineExercises(opts ...func(*RoutineExerciseQuery)) *ExerciseQuery {
	query := (&RoutineExerciseClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withRoutineExercises = query
	return eq
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
//	client.Exercise.Query().
//		GroupBy(exercise.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (eq *ExerciseQuery) GroupBy(field string, fields ...string) *ExerciseGroupBy {
	eq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ExerciseGroupBy{build: eq}
	grbuild.flds = &eq.ctx.Fields
	grbuild.label = exercise.Label
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
//	client.Exercise.Query().
//		Select(exercise.FieldName).
//		Scan(ctx, &v)
func (eq *ExerciseQuery) Select(fields ...string) *ExerciseSelect {
	eq.ctx.Fields = append(eq.ctx.Fields, fields...)
	sbuild := &ExerciseSelect{ExerciseQuery: eq}
	sbuild.label = exercise.Label
	sbuild.flds, sbuild.scan = &eq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ExerciseSelect configured with the given aggregations.
func (eq *ExerciseQuery) Aggregate(fns ...AggregateFunc) *ExerciseSelect {
	return eq.Select().Aggregate(fns...)
}

func (eq *ExerciseQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range eq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, eq); err != nil {
				return err
			}
		}
	}
	for _, f := range eq.ctx.Fields {
		if !exercise.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eq.path != nil {
		prev, err := eq.path(ctx)
		if err != nil {
			return err
		}
		eq.sql = prev
	}
	return nil
}

func (eq *ExerciseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Exercise, error) {
	var (
		nodes       = []*Exercise{}
		_spec       = eq.querySpec()
		loadedTypes = [7]bool{
			eq.withWorkoutLogs != nil,
			eq.withUsers != nil,
			eq.withEquipments != nil,
			eq.withMusclesGroups != nil,
			eq.withExerciseTypes != nil,
			eq.withRoutines != nil,
			eq.withRoutineExercises != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Exercise).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Exercise{config: eq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eq.withWorkoutLogs; query != nil {
		if err := eq.loadWorkoutLogs(ctx, query, nodes,
			func(n *Exercise) { n.Edges.WorkoutLogs = []*WorkoutLog{} },
			func(n *Exercise, e *WorkoutLog) { n.Edges.WorkoutLogs = append(n.Edges.WorkoutLogs, e) }); err != nil {
			return nil, err
		}
	}
	if query := eq.withUsers; query != nil {
		if err := eq.loadUsers(ctx, query, nodes, nil,
			func(n *Exercise, e *User) { n.Edges.Users = e }); err != nil {
			return nil, err
		}
	}
	if query := eq.withEquipments; query != nil {
		if err := eq.loadEquipments(ctx, query, nodes, nil,
			func(n *Exercise, e *Equipment) { n.Edges.Equipments = e }); err != nil {
			return nil, err
		}
	}
	if query := eq.withMusclesGroups; query != nil {
		if err := eq.loadMusclesGroups(ctx, query, nodes, nil,
			func(n *Exercise, e *MusclesGroup) { n.Edges.MusclesGroups = e }); err != nil {
			return nil, err
		}
	}
	if query := eq.withExerciseTypes; query != nil {
		if err := eq.loadExerciseTypes(ctx, query, nodes, nil,
			func(n *Exercise, e *ExerciseType) { n.Edges.ExerciseTypes = e }); err != nil {
			return nil, err
		}
	}
	if query := eq.withRoutines; query != nil {
		if err := eq.loadRoutines(ctx, query, nodes,
			func(n *Exercise) { n.Edges.Routines = []*Routine{} },
			func(n *Exercise, e *Routine) { n.Edges.Routines = append(n.Edges.Routines, e) }); err != nil {
			return nil, err
		}
	}
	if query := eq.withRoutineExercises; query != nil {
		if err := eq.loadRoutineExercises(ctx, query, nodes,
			func(n *Exercise) { n.Edges.RoutineExercises = []*RoutineExercise{} },
			func(n *Exercise, e *RoutineExercise) { n.Edges.RoutineExercises = append(n.Edges.RoutineExercises, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eq *ExerciseQuery) loadWorkoutLogs(ctx context.Context, query *WorkoutLogQuery, nodes []*Exercise, init func(*Exercise), assign func(*Exercise, *WorkoutLog)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Exercise)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(workoutlog.FieldExerciseID)
	}
	query.Where(predicate.WorkoutLog(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(exercise.WorkoutLogsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ExerciseID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "exercise_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eq *ExerciseQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*Exercise, init func(*Exercise), assign func(*Exercise, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Exercise)
	for i := range nodes {
		if nodes[i].UserID == nil {
			continue
		}
		fk := *nodes[i].UserID
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
func (eq *ExerciseQuery) loadEquipments(ctx context.Context, query *EquipmentQuery, nodes []*Exercise, init func(*Exercise), assign func(*Exercise, *Equipment)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Exercise)
	for i := range nodes {
		fk := nodes[i].EquipmentID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(equipment.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "equipment_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (eq *ExerciseQuery) loadMusclesGroups(ctx context.Context, query *MusclesGroupQuery, nodes []*Exercise, init func(*Exercise), assign func(*Exercise, *MusclesGroup)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Exercise)
	for i := range nodes {
		fk := nodes[i].MusclesGroupID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(musclesgroup.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "muscles_group_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (eq *ExerciseQuery) loadExerciseTypes(ctx context.Context, query *ExerciseTypeQuery, nodes []*Exercise, init func(*Exercise), assign func(*Exercise, *ExerciseType)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Exercise)
	for i := range nodes {
		fk := nodes[i].ExerciseTypeID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(exercisetype.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "exercise_type_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (eq *ExerciseQuery) loadRoutines(ctx context.Context, query *RoutineQuery, nodes []*Exercise, init func(*Exercise), assign func(*Exercise, *Routine)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*Exercise)
	nids := make(map[string]map[*Exercise]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(exercise.RoutinesTable)
		s.Join(joinT).On(s.C(routine.FieldID), joinT.C(exercise.RoutinesPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(exercise.RoutinesPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(exercise.RoutinesPrimaryKey[1]))
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
				return append([]any{new(sql.NullString)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullString).String
				inValue := values[1].(*sql.NullString).String
				if nids[inValue] == nil {
					nids[inValue] = map[*Exercise]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Routine](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "routines" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (eq *ExerciseQuery) loadRoutineExercises(ctx context.Context, query *RoutineExerciseQuery, nodes []*Exercise, init func(*Exercise), assign func(*Exercise, *RoutineExercise)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Exercise)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(routineexercise.FieldExerciseID)
	}
	query.Where(predicate.RoutineExercise(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(exercise.RoutineExercisesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ExerciseID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "exercise_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (eq *ExerciseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eq.querySpec()
	_spec.Node.Columns = eq.ctx.Fields
	if len(eq.ctx.Fields) > 0 {
		_spec.Unique = eq.ctx.Unique != nil && *eq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, eq.driver, _spec)
}

func (eq *ExerciseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(exercise.Table, exercise.Columns, sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString))
	_spec.From = eq.sql
	if unique := eq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if eq.path != nil {
		_spec.Unique = true
	}
	if fields := eq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, exercise.FieldID)
		for i := range fields {
			if fields[i] != exercise.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if eq.withUsers != nil {
			_spec.Node.AddColumnOnce(exercise.FieldUserID)
		}
		if eq.withEquipments != nil {
			_spec.Node.AddColumnOnce(exercise.FieldEquipmentID)
		}
		if eq.withMusclesGroups != nil {
			_spec.Node.AddColumnOnce(exercise.FieldMusclesGroupID)
		}
		if eq.withExerciseTypes != nil {
			_spec.Node.AddColumnOnce(exercise.FieldExerciseTypeID)
		}
	}
	if ps := eq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eq *ExerciseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eq.driver.Dialect())
	t1 := builder.Table(exercise.Table)
	columns := eq.ctx.Fields
	if len(columns) == 0 {
		columns = exercise.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eq.sql != nil {
		selector = eq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eq.ctx.Unique != nil && *eq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range eq.predicates {
		p(selector)
	}
	for _, p := range eq.order {
		p(selector)
	}
	if offset := eq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ExerciseGroupBy is the group-by builder for Exercise entities.
type ExerciseGroupBy struct {
	selector
	build *ExerciseQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (egb *ExerciseGroupBy) Aggregate(fns ...AggregateFunc) *ExerciseGroupBy {
	egb.fns = append(egb.fns, fns...)
	return egb
}

// Scan applies the selector query and scans the result into the given value.
func (egb *ExerciseGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, egb.build.ctx, "GroupBy")
	if err := egb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExerciseQuery, *ExerciseGroupBy](ctx, egb.build, egb, egb.build.inters, v)
}

func (egb *ExerciseGroupBy) sqlScan(ctx context.Context, root *ExerciseQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(egb.fns))
	for _, fn := range egb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*egb.flds)+len(egb.fns))
		for _, f := range *egb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*egb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := egb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ExerciseSelect is the builder for selecting fields of Exercise entities.
type ExerciseSelect struct {
	*ExerciseQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (es *ExerciseSelect) Aggregate(fns ...AggregateFunc) *ExerciseSelect {
	es.fns = append(es.fns, fns...)
	return es
}

// Scan applies the selector query and scans the result into the given value.
func (es *ExerciseSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, es.ctx, "Select")
	if err := es.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExerciseQuery, *ExerciseSelect](ctx, es.ExerciseQuery, es, es.inters, v)
}

func (es *ExerciseSelect) sqlScan(ctx context.Context, root *ExerciseQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(es.fns))
	for _, fn := range es.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*es.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := es.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
