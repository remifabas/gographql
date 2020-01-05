package queries

import (
	"github.com/gographql/types"

	"github.com/graphql-go/graphql"
)

// GetTutorialQuery returns the queries available against person type.
func GetTutorialQuery() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(types.TutorialType),
		Description: "Get Tutorial By ID",
		// We can define arguments that allow us to
		// pick specific tutorials. In this case
		// we want to be able to specify the ID of the
		// tutorial we want to retrieve
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// take in the ID argument
			id, ok := p.Args["id"].(int)
			var tutoResponse []types.Tutorial
			if ok {
				// Parse our tutorial array for the matching id
				for _, tutorial := range tutos {
					if int(tutorial.ID) == id {
						// return our tutorial
						tutoResponse = append(tutoResponse, tutorial)
						return tutoResponse, nil
					}
				}
			}
			return tutos, nil
		},
	}
}

/* let's mock some datas */
var tutos = []types.Tutorial{
	{
		ID:    1,
		Title: "title tutorial",
	},
	{
		ID:    2,
		Title: "bel kaoui",
	},
}
