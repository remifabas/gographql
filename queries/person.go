package queries

import (
	"log"
	"time"

	"github.com/gographql/types"

	"github.com/graphql-go/graphql"
)

// GetPersonQuery returns the queries available against person type.
func GetPersonQuery() *graphql.Field {
	return &graphql.Field{

		Type: graphql.NewList(types.PersonType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var persons []types.Person
			start := time.Now()
			// ... Implémenter la logique de base de données ici
			/* let's mock some datas */
			var onePerson types.Person
			onePerson.ID = 36
			onePerson.Firstname = "Rey"
			onePerson.Lastname = "Palpatine"

			persons = append(persons, onePerson, onePerson)
			elapsed := time.Since(start)
			log.Printf("resolve took %s", elapsed)
			return persons, nil
		},
	}
}
