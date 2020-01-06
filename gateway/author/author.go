package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gographql/types"
)

func FindAllAuthor(c *mongo.Collection) string {
	// create a value into which the result can be decoded
	var authors []types.Author

	err := c.Find(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

	out, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}
