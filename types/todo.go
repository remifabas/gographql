package types

import (
	"github.com/graphql-go/graphql"
)

// Todo type definition.
type Todo struct {
	ID        int    `db:"id" json:"id"`
	Title     string `db:"title" json:"title"`
	Complete  bool   `db:"complete" json:"complete"`
	Person_ID int    `db:"person_id" json:"person_id"`
}

// TodoType is the GraphQL schema for the user type.
var TodoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Todo",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.Int},
		"Title":     &graphql.Field{Type: graphql.String},
		"Complete":  &graphql.Field{Type: graphql.Boolean},
		"Person_ID": &graphql.Field{Type: graphql.Int},
	},
})
