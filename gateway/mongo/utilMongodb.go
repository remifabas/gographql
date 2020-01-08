package gateway

import (
	"context"
	"fmt"
	"log"

	"github.com/gographql/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBCon and DB : client connexion mongo database and DB for database graphql
var (
	DBCon *mongo.Client
	DB    *mongo.Database
)

// OpenMongoClient open a connection to a database
func OpenMongoClient() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	var err error
	DBCon, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = DBCon.Ping(context.TODO(), nil)
	log.Println("-> Mongo Client Connected")
	if err != nil {
		log.Fatal(err)
	}
	DB = DBCon.Database("graphqlApi")
}

// InitSomeDatas init somes values in mongodb
func InitSomeDatas() error {
	collection := DB.Collection("author")
	var problem error
	for _, author := range initialAuthors {
		insertResult, err := collection.InsertOne(context.TODO(), author)
		if err != nil {
			problem = err
			log.Fatal(err)
		} else {
			fmt.Println("insert sucess : ", insertResult)
		}
	}
	return problem
}

/* let's mock some datas */
var initialAuthors = []types.Author{
	{
		ID:        1,
		Name:      "tried something",
		Tutorials: []int{1, 2, 3},
	},
	{
		ID:        1,
		Name:      "tried smething else",
		Tutorials: []int{5, 4},
	},
}
