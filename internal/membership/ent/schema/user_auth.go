package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User_Auth holds the schema definition for the User_Auth entity.
type User_Auth struct {
	ent.Schema
}

// Fields of the User_Auth.
func (User_Auth) Fields() []ent.Field {
	return []ent.Field{
		field.String("account_number").MaxLen(100).Unique(),
		field.String("username").MaxLen(100).NotEmpty(),
		field.String("hash").MaxLen(200).NotEmpty(),
		field.Time("last_login"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the User_Auth.
func (User_Auth) Edges() []ent.Edge {
	return nil
}
