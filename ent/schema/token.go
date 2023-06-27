package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Token holds the schema definition for the Token entity.
type Token struct {
	ent.Schema
}

// Fields of the Token.
func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("hash"),
		field.String("expiry"),
		field.String("scope"),
		field.String("user_id").Optional(),
	}
}

func (Token) Index() []ent.Index {
	return []ent.Index{
		index.Fields("hash"),
		index.Fields("expiry"),
		index.Fields("scope"),
	}
}

// Edges of the Token.
func (Token) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("tokens").Field("user_id").Unique(),
	}
}
