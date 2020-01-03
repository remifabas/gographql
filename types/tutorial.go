package types

import "github.com/graphql-go/graphql"

// Tutorial ...
type Tutorial struct {
	ID     int
	Title  string
	Author Author
}

// TutorialType is the GraphQL schema for the user type. NEW OBJECT
var TutorialType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Tutorial",
		Fields: graphql.Fields{
			"id":     &graphql.Field{Type: graphql.Int},
			"title":  &graphql.Field{Type: graphql.String},
			"author": &graphql.Field{Type: AuthorType},
		},
	},
)
