package gateway

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gographql/types"

	authorGateway "github.com/gographql/gateway/author"
	mongo "github.com/gographql/gateway/mongo"
)

// need to find another name and place
type TutorialMongo struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title     string             `json:"Title"`
	Author    primitive.ObjectID `bson:"author" json:"author,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}

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
		var result TutorialMongo
		err := cur.Decode(&result)
		author := authorGateway.FindAuthorById(result.Author)
		if err != nil {
			log.Fatal(err)
		}
		var t types.Tutorial
		t.Author = author
		t.ID = result.ID
		t.CreatedAt = result.CreatedAt
		t.UpdatedAt = result.UpdatedAt

		allTuto = append(allTuto, t)
	}
	return allTuto
}
