package types

import (
	"github.com/graphql-go/graphql"
)

// Todo type definition.
type Person struct {
	ID        int    `db:"id" json:"id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
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
