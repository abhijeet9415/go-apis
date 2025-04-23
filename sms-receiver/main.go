package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/sms", receiveSmsHandler)
	fmt.Println("Sms Receiver Service Running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
