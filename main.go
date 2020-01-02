package main

import (
	"log"
	"net/http"

	"github.com/friendsofgo/graphiql"
	"github.com/gographql/queries"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: queries.GetRootFields(),
		}),
		/*Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootMutation",
			Fields: mutations.GetRootFields(),
		}),*/
	}

	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("Failed to create new schema, error: %v", err)
	}

	httpHandler := handler.New(&handler.Config{Schema: &schema})

	port := 8383
	http.Handle("/", httpHandler)
	log.Print("ready: listening on port :", port)

	// First argument must be same as graphql handler path
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/query")
	if err != nil {
		panic(err)
	}
	http.Handle("/graphiql", graphiqlHandler)

	http.ListenAndServe(":8383", nil)
}
