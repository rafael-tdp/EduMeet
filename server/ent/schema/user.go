package schema

import (
	"edumeet/utils"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	ulid := utils.ULID{}
	return []ent.Field{
		field.String("id").DefaultFunc(ulid.GenerateUlid()).Unique(),
		field.String("email").Unique(),
		field.String("username"),
		field.String("lastname"),
		field.String("firstname"),
		field.String("password").Sensitive(),
		field.Time("birthDate"),
		field.String("bio").Optional().Nillable(),
		field.String("picture").Optional().Nillable(),
		field.Bool("activated").Default(false),
		field.Int("reportNumber").Default(0),
		field.Float("lng").Optional().Nillable(),
		field.Float("lat").Optional().Nillable(),
		field.Time("created_at").Default(time.Now),
		field.String("code").Optional().Nillable(),
		field.Enum("role").Values("SUPER ADMIN", "ADMIN", "USER").Default("USER"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("badges", Badge.Type),
		edge.To("subjects", Subject.Type),
		edge.To("events", Event.Type),
		edge.To("messages", Message.Type),
		edge.To("reports", Reporting.Type),
		edge.To("participants", Participant.Type),
	}
}
