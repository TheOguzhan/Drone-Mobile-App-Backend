package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("qr_code").Default(uuid.NewString()).Annotations(entgql.Skip(entgql.SkipMutationCreateInput)),
		field.Time("date_of_the_order").Default(time.Now).Annotations(entgql.Skip(entgql.SkipMutationCreateInput)),
		field.Bool("completed").Default(false).Annotations(entgql.Skip(entgql.SkipMutationCreateInput)),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier_drone", Drone.Type).
			Ref("current_order").Unique().Required().Annotations(entgql.Skip(entgql.SkipMutationCreateInput)),
		edge.From("user_order", User.Type).
			Ref("user_orders").Unique().Required().Annotations(entgql.Skip(entgql.SkipMutationCreateInput)),
		edge.From("order_warehouse", Warehouse.Type).
			Ref("warehouse_order").Unique().Required(),
		edge.From("order_address", Address.Type).
			Ref("address_order").Unique().Required(),
		edge.From("order_product", Product.Type).
			Ref("product_order").Unique().Required(),
	}
}

func (Order) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
