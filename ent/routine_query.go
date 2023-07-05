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
	"github.com/sahidrahman404/gigachad-api/ent/exercise"
	"github.com/sahidrahman404/gigachad-api/ent/predicate"
	"github.com/sahidrahman404/gigachad-api/ent/routine"
	"github.com/sahidrahman404/gigachad-api/ent/routineexercise"
	"github.com/sahidrahman404/gigachad-api/ent/user"
)

// RoutineQuery is the builder for querying Routine entities.
type RoutineQuery struct {
	config
	ctx                       *QueryContext
	order                     []routine.OrderOption
	inters                    []Interceptor
	predicates                []predicate.Routine
	withExercises             *ExerciseQuery
	withUsers                 *UserQuery
	withRoutineExercises      *RoutineExerciseQuery
	modifiers                 []func(*sql.Selector)
	loadTotal                 []func(context.Context, []*Routine) error
	withNamedExercises        map[string]*ExerciseQuery
	withNamedRoutineExercises map[string]*RoutineExerciseQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RoutineQuery builder.
func (rq *RoutineQuery) Where(ps ...predicate.Routine) *RoutineQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *RoutineQuery) Limit(limit int) *RoutineQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *RoutineQuery) Offset(offset int) *RoutineQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RoutineQuery) Unique(unique bool) *RoutineQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *RoutineQuery) Order(o ...routine.OrderOption) *RoutineQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryExercises chains the current query on the "exercises" edge.
func (rq *RoutineQuery) QueryExercises() *ExerciseQuery {
	query := (&ExerciseClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(routine.Table, routine.FieldID, selector),
			sqlgraph.To(exercise.Table, exercise.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, routine.ExercisesTable, routine.ExercisesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUsers chains the current query on the "users" edge.
func (rq *RoutineQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(routine.Table, routine.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, routine.UsersTable, routine.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRoutineExercises chains the current query on the "routine_exercises" edge.
func (rq *RoutineQuery) QueryRoutineExercises() *RoutineExerciseQuery {
	query := (&RoutineExerciseClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(routine.Table, routine.FieldID, selector),
			sqlgraph.To(routineexercise.Table, routineexercise.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, routine.RoutineExercisesTable, routine.RoutineExercisesColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Routine entity from the query.
// Returns a *NotFoundError when no Routine was found.
func (rq *RoutineQuery) First(ctx context.Context) (*Routine, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{routine.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RoutineQuery) FirstX(ctx context.Context) *Routine {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Routine ID from the query.
// Returns a *NotFoundError when no Routine ID was found.
func (rq *RoutineQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{routine.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *RoutineQuery) FirstIDX(ctx context.Context) string {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Routine entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Routine entity is found.
// Returns a *NotFoundError when no Routine entities are found.
func (rq *RoutineQuery) Only(ctx context.Context) (*Routine, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{routine.Label}
	default:
		return nil, &NotSingularError{routine.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RoutineQuery) OnlyX(ctx context.Context) *Routine {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Routine ID in the query.
// Returns a *NotSingularError when more than one Routine ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *RoutineQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{routine.Label}
	default:
		err = &NotSingularError{routine.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RoutineQuery) OnlyIDX(ctx context.Context) string {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Routines.
func (rq *RoutineQuery) All(ctx context.Context) ([]*Routine, error) {
	ctx = setContextOp(ctx, rq.ctx, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Routine, *RoutineQuery]()
	return withInterceptors[[]*Routine](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *RoutineQuery) AllX(ctx context.Context) []*Routine {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Routine IDs.
func (rq *RoutineQuery) IDs(ctx context.Context) (ids []string, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, "IDs")
	if err = rq.Select(routine.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RoutineQuery) IDsX(ctx context.Context) []string {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RoutineQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*RoutineQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RoutineQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RoutineQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, "Exist")
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RoutineQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RoutineQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RoutineQuery) Clone() *RoutineQuery {
	if rq == nil {
		return nil
	}
	return &RoutineQuery{
		config:               rq.config,
		ctx:                  rq.ctx.Clone(),
		order:                append([]routine.OrderOption{}, rq.order...),
		inters:               append([]Interceptor{}, rq.inters...),
		predicates:           append([]predicate.Routine{}, rq.predicates...),
		withExercises:        rq.withExercises.Clone(),
		withUsers:            rq.withUsers.Clone(),
		withRoutineExercises: rq.withRoutineExercises.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithExercises tells the query-builder to eager-load the nodes that are connected to
// the "exercises" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RoutineQuery) WithExercises(opts ...func(*ExerciseQuery)) *RoutineQuery {
	query := (&ExerciseClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withExercises = query
	return rq
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RoutineQuery) WithUsers(opts ...func(*UserQuery)) *RoutineQuery {
	query := (&UserClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withUsers = query
	return rq
}

// WithRoutineExercises tells the query-builder to eager-load the nodes that are connected to
// the "routine_exercises" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RoutineQuery) WithRoutineExercises(opts ...func(*RoutineExerciseQuery)) *RoutineQuery {
	query := (&RoutineExerciseClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withRoutineExercises = query
	return rq
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
//	client.Routine.Query().
//		GroupBy(routine.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *RoutineQuery) GroupBy(field string, fields ...string) *RoutineGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RoutineGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = routine.Label
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
//	client.Routine.Query().
//		Select(routine.FieldName).
//		Scan(ctx, &v)
func (rq *RoutineQuery) Select(fields ...string) *RoutineSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &RoutineSelect{RoutineQuery: rq}
	sbuild.label = routine.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RoutineSelect configured with the given aggregations.
func (rq *RoutineQuery) Aggregate(fns ...AggregateFunc) *RoutineSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *RoutineQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !routine.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RoutineQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Routine, error) {
	var (
		nodes       = []*Routine{}
		_spec       = rq.querySpec()
		loadedTypes = [3]bool{
			rq.withExercises != nil,
			rq.withUsers != nil,
			rq.withRoutineExercises != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Routine).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Routine{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(rq.modifiers) > 0 {
		_spec.Modifiers = rq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withExercises; query != nil {
		if err := rq.loadExercises(ctx, query, nodes,
			func(n *Routine) { n.Edges.Exercises = []*Exercise{} },
			func(n *Routine, e *Exercise) { n.Edges.Exercises = append(n.Edges.Exercises, e) }); err != nil {
			return nil, err
		}
	}
	if query := rq.withUsers; query != nil {
		if err := rq.loadUsers(ctx, query, nodes, nil,
			func(n *Routine, e *User) { n.Edges.Users = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withRoutineExercises; query != nil {
		if err := rq.loadRoutineExercises(ctx, query, nodes,
			func(n *Routine) { n.Edges.RoutineExercises = []*RoutineExercise{} },
			func(n *Routine, e *RoutineExercise) { n.Edges.RoutineExercises = append(n.Edges.RoutineExercises, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range rq.withNamedExercises {
		if err := rq.loadExercises(ctx, query, nodes,
			func(n *Routine) { n.appendNamedExercises(name) },
			func(n *Routine, e *Exercise) { n.appendNamedExercises(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range rq.withNamedRoutineExercises {
		if err := rq.loadRoutineExercises(ctx, query, nodes,
			func(n *Routine) { n.appendNamedRoutineExercises(name) },
			func(n *Routine, e *RoutineExercise) { n.appendNamedRoutineExercises(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range rq.loadTotal {
		if err := rq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *RoutineQuery) loadExercises(ctx context.Context, query *ExerciseQuery, nodes []*Routine, init func(*Routine), assign func(*Routine, *Exercise)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*Routine)
	nids := make(map[string]map[*Routine]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(routine.ExercisesTable)
		s.Join(joinT).On(s.C(exercise.FieldID), joinT.C(routine.ExercisesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(routine.ExercisesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(routine.ExercisesPrimaryKey[0]))
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
					nids[inValue] = map[*Routine]struct{}{byID[outValue]: {}}
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
func (rq *RoutineQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*Routine, init func(*Routine), assign func(*Routine, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Routine)
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
func (rq *RoutineQuery) loadRoutineExercises(ctx context.Context, query *RoutineExerciseQuery, nodes []*Routine, init func(*Routine), assign func(*Routine, *RoutineExercise)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Routine)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(routineexercise.FieldRoutineID)
	}
	query.Where(predicate.RoutineExercise(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(routine.RoutineExercisesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.RoutineID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "routine_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (rq *RoutineQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	if len(rq.modifiers) > 0 {
		_spec.Modifiers = rq.modifiers
	}
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RoutineQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(routine.Table, routine.Columns, sqlgraph.NewFieldSpec(routine.FieldID, field.TypeString))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, routine.FieldID)
		for i := range fields {
			if fields[i] != routine.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if rq.withUsers != nil {
			_spec.Node.AddColumnOnce(routine.FieldUserID)
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RoutineQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(routine.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = routine.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedExercises tells the query-builder to eager-load the nodes that are connected to the "exercises"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (rq *RoutineQuery) WithNamedExercises(name string, opts ...func(*ExerciseQuery)) *RoutineQuery {
	query := (&ExerciseClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if rq.withNamedExercises == nil {
		rq.withNamedExercises = make(map[string]*ExerciseQuery)
	}
	rq.withNamedExercises[name] = query
	return rq
}

// WithNamedRoutineExercises tells the query-builder to eager-load the nodes that are connected to the "routine_exercises"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (rq *RoutineQuery) WithNamedRoutineExercises(name string, opts ...func(*RoutineExerciseQuery)) *RoutineQuery {
	query := (&RoutineExerciseClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if rq.withNamedRoutineExercises == nil {
		rq.withNamedRoutineExercises = make(map[string]*RoutineExerciseQuery)
	}
	rq.withNamedRoutineExercises[name] = query
	return rq
}

// RoutineGroupBy is the group-by builder for Routine entities.
type RoutineGroupBy struct {
	selector
	build *RoutineQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RoutineGroupBy) Aggregate(fns ...AggregateFunc) *RoutineGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *RoutineGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoutineQuery, *RoutineGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *RoutineGroupBy) sqlScan(ctx context.Context, root *RoutineQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RoutineSelect is the builder for selecting fields of Routine entities.
type RoutineSelect struct {
	*RoutineQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *RoutineSelect) Aggregate(fns ...AggregateFunc) *RoutineSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RoutineSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoutineQuery, *RoutineSelect](ctx, rs.RoutineQuery, rs, rs.inters, v)
}

func (rs *RoutineSelect) sqlScan(ctx context.Context, root *RoutineQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
