package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User_Profile holds the schema definition for the User_Profile entity.
type User_Profile struct {
	ent.Schema
}

// Fields of the User_Profile.
func (User_Profile) Fields() []ent.Field {
	return []ent.Field{
		field.String("account_number").MaxLen(100).Unique(),
		field.String("fullname").MaxLen(100).NotEmpty(),
		field.String("status").MaxLen(10).NotEmpty(),
		field.String("email").MaxLen(100).NotEmpty(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the User_Profile.
func (User_Profile) Edges() []ent.Edge {
	return nil
}
