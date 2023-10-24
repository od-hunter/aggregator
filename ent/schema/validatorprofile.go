package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ValidatorProfile holds the schema definition for the ValidatorProfile entity.
type ValidatorProfile struct {
	ent.Schema
}

// Fields of the ValidatorProfile.
func (ValidatorProfile) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("wallet_address").
			Optional(),
		field.String("host_identifier").
			Optional(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the ValidatorProfile.
func (ValidatorProfile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("validator_profile").
			Unique().
			Required().
			Immutable(),
		edge.From("validated_fulfillments", LockOrderFulfillment.Type).
			Ref("validators"),
		edge.To("api_key", APIKey.Type).
			Unique(),
	}
}
