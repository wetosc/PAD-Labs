// Package eumgent stands for eugenius message agent.
// It implements the shared functional of the client and the message broker
// Usually, you may not need it
package eumgent

import (
	"encoding/json"
)

//PORT : Declares the port on which the application will listen
const PORT = 9000

//MessageType - the type of message, like post, get or other
type MessageType string

//ClientMode - the type of client: PUBLISHER / SUBSCRIBER / Error
type ClientMode string

const (
	// PUBLISH : Client sends a message to the broker
	PUBLISH = "PUBLISH"

	// DELIVER : Broker sends a message to the subscriber
	DELIVER = "DELIVER"

	// SUBSCRIBE : Subscribes a client to the queue "queue"
	SUBSCRIBE = "SUBSCRIBE"

	// RESPONSE : Notifies that a transaction completed succesfully. Usually can be ignored
	RESPONSE = "RESPONSE"

	// ERROR : Signifies an error in the request, like trying to subscribe to an unexistent queue.
	ERROR = "ERROR"
)

//Message is the base type for sent messages
type Message struct {
	Type  MessageType `json:"type"`
	Queue string      `json:"queue"`
	Info  string      `json:"info"`
}

//ToJSON creates a json object from a message
func (m Message) ToJSON() ([]byte, error) {
	return json.Marshal(&m)
}

//MessageFromJSON creates a message from json data
func MessageFromJSON(data []byte) (Message, error) {
	var m Message
	err := json.Unmarshal(data, &m)
	return m, err
}
