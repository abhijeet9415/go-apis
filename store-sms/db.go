package main
import (
	"sms-microservices/database" // Import centralized database package
	"go.mongodb.org/mongo-driver/mongo"
)

// GetStoreCollection returns the "messages" collection from MongoDB
func GetStoreCollection() *mongo.Collection {
	return database.GetCollection("messages")
}
