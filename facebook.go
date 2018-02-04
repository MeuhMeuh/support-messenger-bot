package main

import (
	"net/http"
)

func getFacebookHandler(c *Configuration) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Token is: " + c.Token))
	}
}
