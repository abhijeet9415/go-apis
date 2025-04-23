package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func sendToStorage(sms SMS) {
	fmt.Println("Send SMS to Storage")
	jsonData, _ := json.Marshal(sms)
	resp, err := http.Post("http://localhost:8001/store", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	} else {
		log.Printf("Response Status: %d %s", resp.StatusCode, resp.Status)

	}
	defer resp.Body.Close()

}