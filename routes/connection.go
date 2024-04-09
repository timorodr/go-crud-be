package routes

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// func returns a pointer to a mongo.Client object - this function will create
// and return a reference to a MongoDB client that can interact with my DB
// Struct Client being returned
func DBinstance() *mongo.Client {
	MongoDb := "mongodb+srv://timorodr:R0dr1guez@cluster0.edwcuv3.mongodb.net/caloriesdb?retryWrites=true&w=majority"

	// Whenever we have operations with DB's we want to have a timeout so program doesnt keep waiting forever
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// attempts to connect to MongoDB server using provided context
	// context.BG implies no specific lifecycle or cancellation for this connection attempt
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDb))

	if err != nil {
		log.Fatal(err)
		return nil
	}

	fmt.Println("Connected to MongoDB")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("caloriesdb").Collection(collectionName)
	return collection
}
