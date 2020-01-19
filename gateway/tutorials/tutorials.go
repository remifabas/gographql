package gateway

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gographql/types"

	mongo "github.com/gographql/gateway/mongo"
)

// FindAllTutorialFind all tutorial in mondb
func FindAllTutorials() []types.Tutorial {
	collection := mongo.DB.Collection("tutorial")

	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	var allTuto []types.Tutorial
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		var result types.Tutorial
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		allTuto = append(allTuto, result)
	}
	if err := cur.Err(); err != nil {
		return nil
	}

	return allTuto
}
