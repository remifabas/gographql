package queries

import (
	"fmt"

	"github.com/gographql/types"

	"github.com/graphql-go/graphql"
)

// GetTodoQuery returns the queries available against user type.
func GetTodoQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.TodoType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var todos []types.Todo

			// ... Implémenter la logique de base de données ici
			fmt.Println("i'm here")
			fmt.Println("params ", params.Args)

			/* let's mock some datas */
			var oneTodo types.Todo
			oneTodo.ID = 36
			oneTodo.Title = "Assaj and Talzin"
			oneTodo.Complete = false

			todos = append(todos, oneTodo, oneTodo)

			return todos, nil
		},
	}
}
