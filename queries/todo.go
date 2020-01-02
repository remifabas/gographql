package queries

import (
	"log"
	"time"

	"github.com/gographql/types"

	"github.com/graphql-go/graphql"
)

// GetTodoQuery returns the queries available against todo type.
func GetTodoQuery() *graphql.Field {
	return &graphql.Field{

		Type: graphql.NewList(types.TodoType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var todos []types.Todo
			start := time.Now()
			// ... Implémenter la logique de base de données ici
			/* let's mock some datas */
			var oneTodo types.Todo
			oneTodo.ID = 36
			oneTodo.Title = "Assaj and Talzin"
			oneTodo.Complete = false

			todos = append(todos, oneTodo, oneTodo)
			elapsed := time.Since(start)
			log.Printf("resolve took %s", elapsed)
			return todos, nil
		},
	}
}
