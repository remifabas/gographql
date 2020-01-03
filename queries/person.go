package queries

import (
	"github.com/gographql/types"

	"github.com/graphql-go/graphql"
)

// GetPersonQuery returns the queries available against person type.
func GetPersonQuery() *graphql.Field {
	return &graphql.Field{

		Type:        graphql.NewList(types.PersonType),
		Description: "Get Person By ID",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {

			// ... Implémenter la logique de base de données ici
			return persons, nil
		},
	}
}

/* let's mock some datas */
var persons = []types.Person{
	{
		ID:        1,
		Firstname: "Chicha Morada",
		Lastname:  "Chicha morada is a beverage originated in the Andean regions of Perú but is actually consumed at a national level (wiki)",
	},
	{
		ID:        2,
		Firstname: "mocked firstname Chicha Morada",
		Lastname:  "mocked lastname",
	},
	{
		ID:        3,
		Firstname: "GUsto SLice",
		Lastname:  "la pizzeria",
	},
}
