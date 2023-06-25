package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition
type Product struct {
	ent.Schema
}

// Fields of the Entity.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
	}
}

// Edges of the Entity.
func (Product) Edges() []ent.Edge {
	return nil
}
