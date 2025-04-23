package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction Transaction
	_ = json.NewDecoder(r.Body).Decode(&transaction)

	transaction.ID = primitive.NewObjectID()
	transaction.Timestamp = primitive.NewDateTimeFromTime(time.Now())

	// INsert to DB
	_, err := transactionDB.InsertOne(context.TODO(), transaction)
	if err != nil {
		// fatal.error(err)
		http.Error(w, "Failed To Save Transaction", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Transaction saved successfully",
	})
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	var filters bson.M = bson.M{}

	//queries
	query := r.URL.Query()

	userID := query.Get("user_id")
	if userID != "" {
		filters["user_id"] = userID
	}

	if tType := query.Get("type"); tType != "" {
		filters["type"] = tType // credit/debit
	}

	// get all transaction detail

	var transactions []Transaction
	cursor, err := transactionDB.Find(context.TODO(), filters)
	if err != nil {
		http.Error(w, "Failed to fetch transactions", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var transaction Transaction
		if err := cursor.Decode(&transaction); err == nil {
			transaction.Timestamp = primitive.NewDateTimeFromTime(transaction.Timestamp.Time())
			transactions = append(transactions, transaction)
		}
	}

	// Handle cases when no transactions are found
	if len(transactions) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "No transactions found"})
		return
	}

	// Return transactions
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}
