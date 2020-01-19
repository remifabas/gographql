package gateway

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

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

// InsertAuthor insert in mongodb some authors
func InsertAuthor(author []types.Author) ([]types.Author, error) {

	collection := mongo.DB.Collection("author")

	for _, a := range author {
		_, err := collection.InsertOne(context.TODO(), a)
		if err != nil {
			log.Fatal(err)
			return nil, nil
		}
	}

	return author, nil
}

// FindAuthorById ...
func FindAuthorById(oid primitive.ObjectID) types.Author {
	var theAuthor types.Author
	c := mongo.DB.Collection("author")
	err := c.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&theAuthor)

	if err != nil {
		log.Fatal(err)
	}
	return theAuthor

}
