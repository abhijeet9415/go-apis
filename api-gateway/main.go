package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/sms", func(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "http://sms-receiver-service:8000/api/sms", http.StatusTemporaryRedirect)
    })

    log.Println("API Gateway running on port 8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
