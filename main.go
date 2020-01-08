package main

import (
	"log"
	"net/http"

	"github.com/friendsofgo/graphiql"
	"github.com/gographql/mutations"
	"github.com/gographql/queries"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	mongo "github.com/gographql/gateway/mongo"
)

func main() {
	// Init connexion to database and load some value in it
	mongo.OpenMongoClient()
	mongo.InitSomeDatas()
	// Create the schema configuration i.e. queries and mutations
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: queries.GetRootFields(),
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootMutation",
			Fields: mutations.GetRootFields(),
		}),
	}

	// load the schema configuration
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create new schema, error: %v", err)
	}

	// handle the schema and activate the api
	httpHandler := handler.New(&handler.Config{Schema: &schema})
	http.Handle("/graphql", httpHandler)
	log.Print("ready: listening on port :8383")

	// Handle graphiql web interface
	// First argument must be same as graphql handler path
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}
	http.Handle("/graphiql", graphiqlHandler)

	http.ListenAndServe(":8383", nil)
}
