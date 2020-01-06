package gateway

import (
	"context"
	"fmt"
	"log"

	"github.com/gographql/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoClient() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

func InitSomeDatas() {
	client := GetMongoClient()
	collection := client.Database("graphqlApi").Collection("author")

	/* let's mock some datas */
	initialAuthors := []types.Author{
		{
			ID:        1,
			Name:      "el kabach",
			Tutorials: []int{1, 2},
		},
		{
			ID:        1,
			Name:      "djibbril",
			Tutorials: []int{1, 2},
		},
	}

	for _, author := range initialAuthors {
		insertResult, err := collection.InsertOne(context.TODO(), author)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("insert sucess : ", insertResult)
		}
	}
}
