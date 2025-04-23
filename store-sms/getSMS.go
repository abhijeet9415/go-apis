package main


import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)
//getting all messages 
func GetAllSMS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching all SMS messages...")

	w.Header().Set("Content-Type", "application/json")

	// Get collection
	collection := GetStoreCollection()

	// Fetch all messages
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, "Failed to retrieve messages", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var messages []SMS

	// Decode messages
	for cursor.Next(context.TODO()) {
		var sms SMS
		if err := cursor.Decode(&sms); err != nil {
			http.Error(w, "Error decoding message", http.StatusInternalServerError)
			return
		}
		messages = append(messages, sms)
	}

	// Check if no messages found
	if len(messages) == 0 {
		http.Error(w, "No messages found", http.StatusNotFound)
		return
	}

	// Send response
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
