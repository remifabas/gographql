package gateway

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gographql/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var err error
	// Set client options
	// when working with no docker compose
	//clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	// Connect to MongoDB
	DBCon, err := mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = DBCon.Connect(ctx)
	if err != nil {
		log.Fatal("JoWilfriedTsonga first says : ", err)
	}
	// Check the connection
	err = DBCon.Ping(ctx, nil)

	if err != nil {
		log.Fatal("JoWilfriedTsonga say : ", err)
	}
	log.Println("-> Mongo Client Connected")
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
		ID:        primitive.NewObjectID(),
		Name:      "tried something",
		Tutorials: []int{1, 2, 3},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        primitive.NewObjectID(),
		Name:      "tried something else",
		Tutorials: []int{1, 2, 3},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
