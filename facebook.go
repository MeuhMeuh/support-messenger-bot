package main

import (
	"net/http"
)

func getFacebookHandler(c *Configuration) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		mode, verifyToken, challenge :=
			queryParams.Get("hub.mode"), queryParams.Get("hub.verify_token"), queryParams.Get("hub.challenge")

		if len(mode) > 0 && len(verifyToken) > 0 {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(challenge))
		} else {
			w.Write([]byte("Forbidden"))
			w.WriteHeader(http.StatusForbidden)
		}
	}
}
