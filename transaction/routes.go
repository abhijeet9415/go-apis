package main

import "github.com/gorilla/mux"

func SetupTransaction(r *mux.Router) {
	r.HandleFunc("/transactions", SaveTransaction).Methods("POST")
	r.HandleFunc("/transactions", GetTransaction).Methods("GET")

}
