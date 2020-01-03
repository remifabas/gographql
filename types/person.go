package types

import (
	"github.com/graphql-go/graphql"
)

// Person type definition.
type Person struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// PersonType is the GraphQL schema for the user type.
var PersonType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Person",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.Int},
		"Firstname": &graphql.Field{Type: graphql.String},
		"Lastname":  &graphql.Field{Type: graphql.String},
	},
})
