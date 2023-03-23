package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Drone holds the schema definition for the Drone entity.
type Drone struct {
	ent.Schema
}

// Fields of the Drone.
func (Drone) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.Float("latitude").Optional().Nillable(),
		field.Float("longtitude").Optional().Nillable(),
		field.Bool("in_warehouse").Optional().Default(true),
		field.String("plate_number").Unique(),
	}
}

// Edges of the Drone.
func (Drone) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("current_order", Order.Type).Unique(),
	}
}

func (Drone) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
