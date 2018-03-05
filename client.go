package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

// MessengerClient defines a client used to communicate with the Messenger API.
type MessengerClient struct {
	PageAccessToken string
	Client          *http.Client
}

func createAPIClient(pageAccessToken string) *MessengerClient {
	return &MessengerClient{
		PageAccessToken: pageAccessToken,
		Client: &http.Client{
			Timeout: time.Second * 20,
		},
	}
}

func (client *MessengerClient) send(message []byte) (*http.Response, error) {
	return client.Client.Post(
		fmt.Sprintf("https://graph.facebook.com/v2.6/me/messages?access_token=%s", client.PageAccessToken),
		"application/json",
		bytes.NewReader(message),
	)
}
