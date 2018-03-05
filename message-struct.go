package main

// Sender defines the payload expected when a /post webhook is called by Facebook.
type Sender struct {
	ID string `json:"id"`
}

// Recipient defines a... recipient.
type Recipient struct {
	ID string `json:"id"`
}

// Message defines a message sent from the Facebook Messenger system.
type Message struct {
	Mid        string `json:"mid"`
	Text       string `json:"text"`
	QuickReply struct {
		Payload string `json:"payload"`
	} `json:"quick_reply"`
}

// ActionPayload defines the payload expected when a /post webhook is called by Facebook.
type ActionPayload struct {
	Object string `json:"object"`
	Entry  []struct {
		Messaging []struct {
			Sender    *Sender    `json:"sender"`
			Recipient *Recipient `json:"recipient"`
			Timestamp int64      `json:"timestamp"`
			Message   *Message   `json:"message"`
		} `json:"messaging"`
	} `json:"entry"`
}

// SendMessage is the tiniest representation of a message to be sent.
type SendMessage struct {
	Text string `json:"text"`
}

// BasicMessage is intended to define a message to be sent via the Messenger API.
type BasicMessage struct {
	MessagingType string       `json:"messaging_type"`
	Recipient     *Recipient   `json:"recipient"`
	Message       *SendMessage `json:"message"`
}
