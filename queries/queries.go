package queries

import (
	"github.com/graphql-go/graphql"

	author "github.com/gographql/queries/author"
	tutorial "github.com/gographql/queries/tutorial"
)

// GetRootFields returns all the available queries.
func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"author":   author.GetAuthorQuery(),
		"tutorial": tutorial.GetTutorialQuery(),
	}
}
