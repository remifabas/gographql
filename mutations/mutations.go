package mutations

import (
	"github.com/graphql-go/graphql"

	author "github.com/gographql/mutations/author"
)

// GetRootFields returns all the available mutations.
func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"createAuthor": author.GetCreateAuthorMutation(),
	}
}
