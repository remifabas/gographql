package types

import (
	"github.com/graphql-go/graphql"
)

// Todo type definition.
type Todo struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
}

// TodoType is the GraphQL schema for the user type.
var TodoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Todo",
	Fields: graphql.Fields{
		"id":    &graphql.Field{Type: graphql.Int},
		"Title": &graphql.Field{Type: graphql.String},
	},
})
