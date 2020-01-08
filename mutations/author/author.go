package mutations

import (
	"context"

	"github.com/gographql/types"

	"github.com/graphql-go/graphql"

	mongo "github.com/gographql/gateway/mongo"
)

// GetCreateAuthorMutation creates a new user and returns it.
func GetCreateAuthorMutation() *graphql.Field {
	return &graphql.Field{
		Type:        types.AuthorType,
		Description: "Create an anthor Name and Tutorial needed",
		Args: graphql.FieldConfigArgument{
			"Name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"Tutorials": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.NewList(graphql.Int)),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			intSlice := InterfaceSliceToSlice(params.Args["Tutorials"].([]interface{}))
			author := types.Author{
				Name:      params.Args["Name"].(string),
				Tutorials: intSlice,
			}

			collection := mongo.DB.Collection("author")

			_, err := collection.InsertOne(context.TODO(), author)

			return author, err
		},
	}
}

// InitSomeAuthor ...
func InitSomeAuthor() *graphql.Field {
	return &graphql.Field{
		Type:        types.StatusType,
		Description: "Init two authors in database for test purpose",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			err := mongo.InitSomeDatas()
			var status types.Status
			if err != nil {
				status.Success = "internal error"
			} else {
				status.Success = "init done"
			}
			return status, err
		},
	}
}

// InterfaceSliceToSlice change slice interface to int
func InterfaceSliceToSlice(theInterface []interface{}) []int {
	var returnSlice []int
	for i := range theInterface {
		returnSlice = append(returnSlice, theInterface[i].(int))
	}
	return returnSlice
}
