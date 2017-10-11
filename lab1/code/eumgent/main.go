// Package eumgent stands for eugenius message agent.
// It implements the shared functional of the client and the message broker
// Usually, you may not need it
package eumgent

import (
	"bufio"
	"encoding/json"
	"io"
	"net"

	"github.com/rs/zerolog/log"
)

//PORT : Declares the port on which the application will listen
const PORT = 9000

//MessageType - the type of message, like post, get or other
type MessageType string

//ClientType - the type of client: PUBLISHER or SUBSCRIBER
type ClientType string

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

//Client represents something like a connection, which allows you
type Client struct {
	Queue    string
	Mode     ClientType
	Incoming chan Message
	Outgoing chan Message
	reader   *bufio.Reader
	writer   *bufio.Writer
}

func (client *Client) Read() {
	for {
		line, _, err := client.reader.ReadLine()
		if err != nil {
			log.Debug().Msgf("Error reading:\t%v", err)
			if err == io.EOF {
				log.Info().Msgf("Connexion closed:\t%v", err)
				break
			}
		}
		msg, err := MessageFromJSON(line)
		if err != nil {
			log.Debug().Msgf("Invalid JSON:\t%v", err)
		}
		if msg.Type == PUBLISH {
			client.Mode = PUBLISH
		}
		if msg.Type == SUBSCRIBE {
			client.Mode = SUBSCRIBE
			client.Queue = msg.Queue
		}
		client.Incoming <- msg
	}
}

func (client *Client) Write() {
	for msg := range client.Outgoing {
		data, _ := msg.ToJSON()
		_, err := client.writer.Write(data)
		client.writer.Flush()
		if err != nil {
			log.Debug().Msgf("Error writting data:\t%v", err)
		}
	}
}

//Listen checks if there are messages for writing or for reading
func (client *Client) Listen() {
	go client.Read()
	go client.Write()
}

//NewClient creates a new client from a connection
func NewClient(connection net.Conn) *Client {
	writer := bufio.NewWriter(connection)
	reader := bufio.NewReader(connection)

	client := &Client{
		Incoming: make(chan Message),
		Outgoing: make(chan Message),
		reader:   reader,
		writer:   writer,
	}

	client.Listen()

	return client
}

//Room is a structure that allows to register publishers or subscribers
type Room struct {
	Clients  []*Client
	Joins    chan net.Conn
	Incoming chan Message
	Outgoing chan Message
}

//Deliver sends messages to all subscribers
func (room *Room) Deliver(msg Message) {
	log.Info().Msgf("Delivering message: %v", msg)
	for _, client := range room.Clients {
		if client.Mode == SUBSCRIBE && client.Queue == msg.Queue {
			client.Outgoing <- msg
		}
	}
}

//Join adds a new publisher
func (room *Room) Join(conn net.Conn) {
	log.Info().Msgf("New client connected: %v", conn)
	client := NewClient(conn)
	room.Clients = append(room.Clients, client)
	go func() {
		for {
			room.Incoming <- <-client.Incoming
		}
	}()
}

//Listen waits for messages
func (room *Room) Listen() {
	go func() {
		for {
			select {
			case data := <-room.Incoming:
				room.Deliver(data)
			case conn := <-room.Joins:
				room.Join(conn)
			}
		}
	}()
}

//NewRoom creates a new room
func NewRoom() *Room {
	room := &Room{
		Clients:  make([]*Client, 0),
		Joins:    make(chan net.Conn),
		Incoming: make(chan Message),
		Outgoing: make(chan Message),
	}

	room.Listen()

	return room
}
