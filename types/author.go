package types

import (
	"github.com/graphql-go/graphql"
)

// Author type definition.
type Author struct {
	ID        int    `json:"id"`
	Name      string `json:"Name"`
	Tutorials []int  `json:"Tutorials"`
}

// AuthorType is the GraphQL schema for the user type.
var AuthorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"ID":        &graphql.Field{Type: graphql.Int},
			"Name":      &graphql.Field{Type: graphql.String},
			"Tutorials": &graphql.Field{Type: graphql.NewList(graphql.Int)},
		},
	},
)
