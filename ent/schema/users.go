package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

// Fields of the Users.
func (Users) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.New()).Unique(),
		field.String("username").Unique(),
		field.String("email").Match(regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)).Unique(),
		field.String("first_name"),
		field.String("last_name"),
		field.Bool("is_user_confirmed").Default(false),
		field.String("password"),
	}
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return nil
}
