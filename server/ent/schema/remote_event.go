package schema

import (
	"edumeet/utils"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type RemoteEvent struct {
	ent.Schema
}

func (RemoteEvent) Fields() []ent.Field {
	ulid := utils.ULID{}
	return []ent.Field{
		field.String("id").DefaultFunc(ulid.GenerateUlid()).Unique(),
		field.String("url"),
	}
}

func (RemoteEvent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).Ref("remote_event").Unique(),
	}
}
