package queries

import (
	"github.com/graphql-go/graphql"

	tutorial "github.com/gographql/queries/tutorial"
)

// GetRootFields returns all the available queries.
func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"person":   GetPersonQuery(),
		"author":   GetAuthorQuery(),
		"tutorial": tutorial.GetTutorialQuery(),
	}
}
