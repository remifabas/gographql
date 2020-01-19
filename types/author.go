package types

import (
	"time"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Author type definition.
type Author struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name      string             `json:"Name"`
	Tutorials []int              `json:"Tutorials"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}

// AuthorType is the GraphQL schema for the user type.
var AuthorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"ID":        &graphql.Field{Type: graphql.String},
			"Name":      &graphql.Field{Type: graphql.String},
			"Tutorials": &graphql.Field{Type: graphql.NewList(graphql.Int)},
			"CreatedAt": &graphql.Field{Type: graphql.DateTime},
			"UpdatedAt": &graphql.Field{Type: graphql.DateTime},
		},
	},
)
