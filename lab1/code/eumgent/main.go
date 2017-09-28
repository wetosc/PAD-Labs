// Package eumgent stands for eugenius message agent.
// It implements the shared functional of sender, receiver and message queue
package eumgent

import (
	"encoding/json"
	"net"
)

//PORT : Declares the port on which the application will listen
const PORT = 9000

//MessageType - the type of message, like post, get or other
type MessageType string

const (
	//POST means that the message will be added to the queue
	POST MessageType = "post"
	//GET means that a message will be popped from the queue
	GET MessageType = "get"
)

//Message is the base type for sent messages
type Message struct {
	Type MessageType `json:"type"`
	Info string      `json:"info"`
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

//Write writes data to connection and returns error or nil
func Write(client net.Conn, data []byte) error {
	_, err := client.Write(data)
	return err
}

// WriteString writes string to connection and return error or nil
func WriteString(client net.Conn, msg string) error {
	_, err := client.Write([]byte(msg))
	return err
}

// Read reads data from connection and returns a ([]byte, error)
func Read(client net.Conn) ([]byte, error) {
	reply := make([]byte, 100)
	length, err := client.Read(reply)
	return reply[:length], err
}

// ReadString reads data from connection and returns a (string, error)
func ReadString(client net.Conn) (string, error) {
	reply := make([]byte, 100)
	_, err := client.Read(reply)
	return string(reply), err
}
