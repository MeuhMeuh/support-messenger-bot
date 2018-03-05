package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func subscriptionHandler(c *Configuration) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		mode, verifyToken, challenge :=
			queryParams.Get("hub.mode"), queryParams.Get("hub.verify_token"), queryParams.Get("hub.challenge")

		if mode == "subscribe" && verifyToken == c.Token {
			log.Println("200 OK")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(challenge))
		} else {
			log.Println("403 FORBIDDEN")
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		}
	}
}

func actionHandler(c *Configuration) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewReader(body))
		if r.Header.Get("X-Hub-Signature") == "" || !isSignatureMatching(c.AppSecret, body, r.Header.Get("X-Hub-Signature")[5:]) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		decoder := json.NewDecoder(r.Body)
		var t ActionPayload
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		log.Println(t)

		if t.Object == "page" {
			log.Println(t.Entry)
		}

		w.Write([]byte("OK POST"))
	}
}
