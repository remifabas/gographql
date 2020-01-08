package types

import (
	"github.com/graphql-go/graphql"
)

// Status type definition.
type Status struct {
	Success string `json:"Success"`
}

// StatusType is the GraphQL schema for the user type.
var StatusType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Status",
		Fields: graphql.Fields{
			"Success": &graphql.Field{Type: graphql.String},
		},
	},
)
