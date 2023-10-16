package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SenderProfile holds the schema definition for the SenderProfile entity.
type SenderProfile struct {
	ent.Schema
}

// Fields of the SenderProfile.
func (SenderProfile) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("webhook_url").Optional(),
		field.Strings("domain_whitelist").
			Default([]string{}),
	}
}

// Edges of the SenderProfile.
func (SenderProfile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("sender_profile").
			Unique().
			Required().
			Immutable(),
	}
}
