package queries

import (
	"context"
	"log"

	"github.com/gographql/types"
	"go.mongodb.org/mongo-driver/bson"

	mongo "github.com/gographql/gateway"
)

// GetAllAuthor return all tutorials alviable
func GetAllAuthor() ([]types.Author, error) {

	client := mongo.GetMongoClient()
	collection := client.Database("graphqlApi").Collection("author")

	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	var allAuthors []types.Author
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		var result types.Author
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		allAuthors = append(allAuthors, result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	// implement database acces
	return allAuthors, nil
}
