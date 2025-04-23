package main
import (
	"log"
	"net/http"
	"sms-microservices/database"

	"github.com/gorilla/mux"
)

// func main() {
// 	initMongoDB()

// 	http.HandleFunc("/store", storeSmsHandler)
// 	log.Println("Storage Service running on port 8001")
// 	log.Fatal(http.ListenAndServe(":8001", nil))
// }

func main() {
	database.ConnectDB()
	r := mux.NewRouter()

	// Register Handlers
	r.HandleFunc("/store", storeSmsHandler).Methods("POST")
	r.HandleFunc("/store/allSMS", GetAllSMS).Methods("GET")

	log.Println("Storage Service running on port 8001")
	log.Fatal(http.ListenAndServe(":8001", r))
}
