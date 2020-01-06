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
			authors, status := GetAllAuthor()
			return authors, status
		},
	}
}
