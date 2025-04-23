package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type SMS struct {
	Sender  string `json:"sender"`
	Message string `json:"message"`
}

func storeSmsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("store SMS handler Started")
	var sms SMS
	if err := json.NewDecoder(r.Body).Decode(&sms); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	collection := GetStoreCollection()
	_, err := collection.InsertOne(context.TODO(), bson.M{
		"sender": sms.Sender,
		 "message": sms.Message,
		})
	if err != nil {
		http.Error(w, "Failed to store message", http.StatusInternalServerError)
		return
	}

	log.Println("Message stored successfully in MongoDB")
	w.WriteHeader(http.StatusCreated)
}
