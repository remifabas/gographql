package gateway

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gographql/types"

	mongo "github.com/gographql/gateway/mongo"
)

// FindAllAuthor Find all authors in mondb
func FindAllAuthor() []types.Author {
	collection := mongo.DB.Collection("author")

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
		fmt.Println(result) // output:
		if err != nil {
			log.Fatal(err)
		}
		allAuthors = append(allAuthors, result)
	}
	if err := cur.Err(); err != nil {
		return nil
	}

	return allAuthors
}
