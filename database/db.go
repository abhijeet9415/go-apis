package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Global MongoDB Client
var Client *mongo.Client

// ConnectDB initializes the MongoDB connection
func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("MongoDB Connection Error:", err)
	}
	log.Println("Connected to MongoDB")
}

// GetClient returns the MongoDB Client
func GetClient() *mongo.Client {
	return Client
}

// GetCollection returns a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database("smsdb").Collection(collectionName)
}
