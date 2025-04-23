package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type SMS struct {
	Sender  string `json:"sender"`
	Message string `json:"message"`
}

func receiveSmsHandler(w http.ResponseWriter, r *http.Request) {
	var sms SMS
	err := json.NewDecoder(r.Body).Decode(&sms)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	log.Printf("Receive SMS from %s: %s\n", sms.Sender, sms.Message)
	sendToStorage(sms)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "SMS received"})

}
