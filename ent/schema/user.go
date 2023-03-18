package schema

import (
	"context"
	"fmt"
	"regexp"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/ent/hook"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/utils"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("username").Unique(),
		field.String("email").Match(regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)).Unique(),
		field.String("first_name"),
		field.String("last_name"),
		field.Bool("is_user_confirmed").Default(false).Annotations(entgql.Skip(entgql.SkipMutationCreateInput)),
		field.String("password").Annotations(entgql.Skip(entgql.SkipType)).Sensitive(),
		field.Strings("auth_tokens").Optional().Annotations(entgql.Skip(entgql.SkipAll)),
	}
}

func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					if err := m.SetField("password", utils.HashPassword(func() ent.Value { f, _ := m.Field("password"); return f }())); err != nil {
						fmt.Printf("%v", err)
						return nil, fmt.Errorf("could hash the password")
					}
					return next.Mutate(ctx, m)
				})
			}, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne),
	}
}

// Edges of the Users.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("address_slaves", Address.Type),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
