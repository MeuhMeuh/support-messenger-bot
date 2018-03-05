package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	configuration, err := getConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}

	r := mux.NewRouter()

	r.HandleFunc("/webhook", subscriptionHandler(configuration)).Methods("GET")
	r.HandleFunc("/webhook", actionHandler(configuration)).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listening...")
	log.Fatal(srv.ListenAndServe())
}
