package queries

import (
	"github.com/gographql/types"

	"github.com/graphql-go/graphql"
)

// GetAuthorQuery returns the queries available against person type.
func GetAuthorQuery() *graphql.Field {
	return &graphql.Field{

		Type:        graphql.NewList(types.AuthorType),
		Description: "Get Author By ID",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {

			// ... Implémenter la logique de base de données ici
			return authors, nil
		},
	}
}

/* let's mock some datas */
var authors = []types.Author{
	{
		ID:        1,
		Name:      "rl kabach",
		Tutorials: []int{1, 2},
	},
	{
		ID:        1,
		Name:      "rl kabach",
		Tutorials: []int{1, 2},
	},
}
