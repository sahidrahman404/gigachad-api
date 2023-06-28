package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(generateKSUID),
		field.String("email"),
		field.String("username"),
		field.String("hashed_password"),
		field.String("name"),
		field.String("created_at").DefaultFunc(generateTime),
		field.Int("activated").Default(0),
		field.Int("version").Default(1),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tokens", Token.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
