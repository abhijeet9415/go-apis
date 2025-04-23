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
	InitUserDB() // Initialize user collection

	// Create Router
	r := mux.NewRouter()
	SetupRoutes(r)

	log.Println("User Service running on port 8002...")
	log.Fatal(http.ListenAndServe(":8002", r))
}
