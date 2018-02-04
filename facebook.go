package main

import (
	"net/http"
)

func getFacebookHandler(c *Configuration) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		mode, verifyToken, challenge :=
			queryParams.Get("hub.mode"), queryParams.Get("hub.verify_token"), queryParams.Get("hub.challenge")

		if mode == "subscribe" && verifyToken == c.Token {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(challenge))
		} else {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Forbidden"))
		}
	}
}
