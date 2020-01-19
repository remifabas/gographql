package types

import (
	"time"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Tutorial ...
type Tutorial struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title     string             `json:"Title"`
	Author    primitive.ObjectID `bson:"author" json:"author,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}

// TutorialType is the GraphQL schema for the user type. NEW OBJECT
var TutorialType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Tutorial",
		Fields: graphql.Fields{
			"ID":        &graphql.Field{Type: graphql.String},
			"Title":     &graphql.Field{Type: graphql.String},
			"Author":    &graphql.Field{Type: graphql.String},
			"CreatedAt": &graphql.Field{Type: graphql.DateTime},
			"UpdatedAt": &graphql.Field{Type: graphql.DateTime},
		},
	},
)
