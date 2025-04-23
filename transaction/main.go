package main

import (
	"log"
	"net/http"
	"sms-microservices/database"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize MongoDB connection
	database.ConnectDB()
	InitTransactionDB() // Initialize transaction collection

	// Create Router
	r := mux.NewRouter()
	SetupTransaction(r)

	log.Println("Transaction Service running on port 8003...")
	log.Fatal(http.ListenAndServe(":8003", r))
}
