package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func getFacebookHandler(c *Configuration) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(time.RFC3339)
		w.Write([]byte("The time is: " + tm))
	}
}

func main() {
	configuration, err := getConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}

	r := mux.NewRouter()

	r.HandleFunc("/facebook", getFacebookHandler(&configuration)).Methods("POST").Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listening...")
	log.Fatal(srv.ListenAndServe())
}
