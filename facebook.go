package main

import (
	"fmt"
	"net/http"
)

func getFacebookHandler(c *Configuration) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		mode, verifyToken, challenge :=
			queryParams.Get("hub.mode"), queryParams.Get("hub.verify_token"), queryParams.Get("hub.challenge")

		fmt.Println(mode, verifyToken)

		if len(mode) > 0 && len(verifyToken) > 0 {
			fmt.Println("Yeah it should work")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(challenge))
		} else {
			fmt.Println("Nope sth went wrong")
			w.WriteHeader(http.StatusForbidden)
		}
	}
}
