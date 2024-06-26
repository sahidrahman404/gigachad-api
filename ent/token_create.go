// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/token"
	"github.com/sahidrahman404/gigachad-api/ent/user"
)

// TokenCreate is the builder for creating a Token entity.
type TokenCreate struct {
	config
	mutation *TokenMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetHash sets the "hash" field.
func (tc *TokenCreate) SetHash(b []byte) *TokenCreate {
	tc.mutation.SetHash(b)
	return tc
}

// SetExpiry sets the "expiry" field.
func (tc *TokenCreate) SetExpiry(s string) *TokenCreate {
	tc.mutation.SetExpiry(s)
	return tc
}

// SetScope sets the "scope" field.
func (tc *TokenCreate) SetScope(s string) *TokenCreate {
	tc.mutation.SetScope(s)
	return tc
}

// SetUserID sets the "user_id" field.
func (tc *TokenCreate) SetUserID(pk pksuid.ID) *TokenCreate {
	tc.mutation.SetUserID(pk)
	return tc
}

// SetID sets the "id" field.
func (tc *TokenCreate) SetID(pk pksuid.ID) *TokenCreate {
	tc.mutation.SetID(pk)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TokenCreate) SetNillableID(pk *pksuid.ID) *TokenCreate {
	if pk != nil {
		tc.SetID(*pk)
	}
	return tc
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (tc *TokenCreate) SetUsersID(id pksuid.ID) *TokenCreate {
	tc.mutation.SetUsersID(id)
	return tc
}

// SetUsers sets the "users" edge to the User entity.
func (tc *TokenCreate) SetUsers(u *User) *TokenCreate {
	return tc.SetUsersID(u.ID)
}

// Mutation returns the TokenMutation object of the builder.
func (tc *TokenCreate) Mutation() *TokenMutation {
	return tc.mutation
}

// Save creates the Token in the database.
func (tc *TokenCreate) Save(ctx context.Context) (*Token, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TokenCreate) SaveX(ctx context.Context) *Token {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TokenCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TokenCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TokenCreate) defaults() {
	if _, ok := tc.mutation.ID(); !ok {
		v := token.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TokenCreate) check() error {
	if _, ok := tc.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New(`ent: missing required field "Token.hash"`)}
	}
	if _, ok := tc.mutation.Expiry(); !ok {
		return &ValidationError{Name: "expiry", err: errors.New(`ent: missing required field "Token.expiry"`)}
	}
	if _, ok := tc.mutation.Scope(); !ok {
		return &ValidationError{Name: "scope", err: errors.New(`ent: missing required field "Token.scope"`)}
	}
	if _, ok := tc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Token.user_id"`)}
	}
	if _, ok := tc.mutation.UsersID(); !ok {
		return &ValidationError{Name: "users", err: errors.New(`ent: missing required edge "Token.users"`)}
	}
	return nil
}

func (tc *TokenCreate) sqlSave(ctx context.Context) (*Token, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*pksuid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TokenCreate) createSpec() (*Token, *sqlgraph.CreateSpec) {
	var (
		_node = &Token{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(token.Table, sqlgraph.NewFieldSpec(token.FieldID, field.TypeString))
	)
	_spec.OnConflict = tc.conflict
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.Hash(); ok {
		_spec.SetField(token.FieldHash, field.TypeBytes, value)
		_node.Hash = value
	}
	if value, ok := tc.mutation.Expiry(); ok {
		_spec.SetField(token.FieldExpiry, field.TypeString, value)
		_node.Expiry = value
	}
	if value, ok := tc.mutation.Scope(); ok {
		_spec.SetField(token.FieldScope, field.TypeString, value)
		_node.Scope = value
	}
	if nodes := tc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.UsersTable,
			Columns: []string{token.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Token.Create().
//		SetHash(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TokenUpsert) {
//			SetHash(v+v).
//		}).
//		Exec(ctx)
func (tc *TokenCreate) OnConflict(opts ...sql.ConflictOption) *TokenUpsertOne {
	tc.conflict = opts
	return &TokenUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TokenCreate) OnConflictColumns(columns ...string) *TokenUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TokenUpsertOne{
		create: tc,
	}
}

type (
	// TokenUpsertOne is the builder for "upsert"-ing
	//  one Token node.
	TokenUpsertOne struct {
		create *TokenCreate
	}

	// TokenUpsert is the "OnConflict" setter.
	TokenUpsert struct {
		*sql.UpdateSet
	}
)

// SetHash sets the "hash" field.
func (u *TokenUpsert) SetHash(v []byte) *TokenUpsert {
	u.Set(token.FieldHash, v)
	return u
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *TokenUpsert) UpdateHash() *TokenUpsert {
	u.SetExcluded(token.FieldHash)
	return u
}

// SetExpiry sets the "expiry" field.
func (u *TokenUpsert) SetExpiry(v string) *TokenUpsert {
	u.Set(token.FieldExpiry, v)
	return u
}

// UpdateExpiry sets the "expiry" field to the value that was provided on create.
func (u *TokenUpsert) UpdateExpiry() *TokenUpsert {
	u.SetExcluded(token.FieldExpiry)
	return u
}

// SetScope sets the "scope" field.
func (u *TokenUpsert) SetScope(v string) *TokenUpsert {
	u.Set(token.FieldScope, v)
	return u
}

// UpdateScope sets the "scope" field to the value that was provided on create.
func (u *TokenUpsert) UpdateScope() *TokenUpsert {
	u.SetExcluded(token.FieldScope)
	return u
}

// SetUserID sets the "user_id" field.
func (u *TokenUpsert) SetUserID(v pksuid.ID) *TokenUpsert {
	u.Set(token.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *TokenUpsert) UpdateUserID() *TokenUpsert {
	u.SetExcluded(token.FieldUserID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(token.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TokenUpsertOne) UpdateNewValues() *TokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(token.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Token.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TokenUpsertOne) Ignore() *TokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TokenUpsertOne) DoNothing() *TokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TokenCreate.OnConflict
// documentation for more info.
func (u *TokenUpsertOne) Update(set func(*TokenUpsert)) *TokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TokenUpsert{UpdateSet: update})
	}))
	return u
}

// SetHash sets the "hash" field.
func (u *TokenUpsertOne) SetHash(v []byte) *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.SetHash(v)
	})
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *TokenUpsertOne) UpdateHash() *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateHash()
	})
}

// SetExpiry sets the "expiry" field.
func (u *TokenUpsertOne) SetExpiry(v string) *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.SetExpiry(v)
	})
}

// UpdateExpiry sets the "expiry" field to the value that was provided on create.
func (u *TokenUpsertOne) UpdateExpiry() *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateExpiry()
	})
}

// SetScope sets the "scope" field.
func (u *TokenUpsertOne) SetScope(v string) *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.SetScope(v)
	})
}

// UpdateScope sets the "scope" field to the value that was provided on create.
func (u *TokenUpsertOne) UpdateScope() *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateScope()
	})
}

// SetUserID sets the "user_id" field.
func (u *TokenUpsertOne) SetUserID(v pksuid.ID) *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *TokenUpsertOne) UpdateUserID() *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateUserID()
	})
}

// Exec executes the query.
func (u *TokenUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TokenCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TokenUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TokenUpsertOne) ID(ctx context.Context) (id pksuid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TokenUpsertOne.ID is not supported by MySQL driver. Use TokenUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TokenUpsertOne) IDX(ctx context.Context) pksuid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TokenCreateBulk is the builder for creating many Token entities in bulk.
type TokenCreateBulk struct {
	config
	err      error
	builders []*TokenCreate
	conflict []sql.ConflictOption
}

// Save creates the Token entities in the database.
func (tcb *TokenCreateBulk) Save(ctx context.Context) ([]*Token, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Token, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TokenMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TokenCreateBulk) SaveX(ctx context.Context) []*Token {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TokenCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TokenCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Token.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TokenUpsert) {
//			SetHash(v+v).
//		}).
//		Exec(ctx)
func (tcb *TokenCreateBulk) OnConflict(opts ...sql.ConflictOption) *TokenUpsertBulk {
	tcb.conflict = opts
	return &TokenUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TokenCreateBulk) OnConflictColumns(columns ...string) *TokenUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TokenUpsertBulk{
		create: tcb,
	}
}

// TokenUpsertBulk is the builder for "upsert"-ing
// a bulk of Token nodes.
type TokenUpsertBulk struct {
	create *TokenCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(token.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TokenUpsertBulk) UpdateNewValues() *TokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(token.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TokenUpsertBulk) Ignore() *TokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TokenUpsertBulk) DoNothing() *TokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TokenCreateBulk.OnConflict
// documentation for more info.
func (u *TokenUpsertBulk) Update(set func(*TokenUpsert)) *TokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TokenUpsert{UpdateSet: update})
	}))
	return u
}

// SetHash sets the "hash" field.
func (u *TokenUpsertBulk) SetHash(v []byte) *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.SetHash(v)
	})
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *TokenUpsertBulk) UpdateHash() *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateHash()
	})
}

// SetExpiry sets the "expiry" field.
func (u *TokenUpsertBulk) SetExpiry(v string) *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.SetExpiry(v)
	})
}

// UpdateExpiry sets the "expiry" field to the value that was provided on create.
func (u *TokenUpsertBulk) UpdateExpiry() *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateExpiry()
	})
}

// SetScope sets the "scope" field.
func (u *TokenUpsertBulk) SetScope(v string) *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.SetScope(v)
	})
}

// UpdateScope sets the "scope" field to the value that was provided on create.
func (u *TokenUpsertBulk) UpdateScope() *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateScope()
	})
}

// SetUserID sets the "user_id" field.
func (u *TokenUpsertBulk) SetUserID(v pksuid.ID) *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *TokenUpsertBulk) UpdateUserID() *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateUserID()
	})
}

// Exec executes the query.
func (u *TokenUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TokenCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TokenCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TokenUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
