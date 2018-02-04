package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func facebookHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC3339)
	w.Write([]byte("The time is: " + tm))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/facebook", facebookHandler)

	log.Println("Listening...")
	http.ListenAndServe(":3000", r)
}
