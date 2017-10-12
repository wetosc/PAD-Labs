package eumgent

import (
	"net"
	"regexp"
	"time"

	"github.com/rs/zerolog/log"
)

//Room is a structure that allows to register publishers or subscribers
type Room struct {
	Clients  []*Client
	Joins    chan net.Conn
	Incoming chan Message
	Outgoing chan Message
}

//Deliver sends messages to all subscribers
func (room *Room) Deliver(msg Message) {
	if msg.Type != PUBLISH {
		return
	}
	msg.Type = DELIVER
	i := 0
	for _, client := range room.Clients {
		if client.Mode == SUBSCRIBE {
			match, _ := regexp.MatchString(client.Queue, msg.Queue)
			if match {
				client.Outgoing <- msg
				i++
			}
		}
	}
	if i > 0 {
		log.Info().Msgf("Delivered message: %v    to %v clients", msg, i)
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
	room.startClientsCleaner(10 * time.Second)

	return room
}

func filter(vs []*Client, f func(*Client) bool) []*Client {
	vsf := make([]*Client, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

//Each "interval" removes the clients with the mode ERROR
func (room *Room) startClientsCleaner(interval time.Duration) {
	ticker := time.NewTicker(interval)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				room.Clients = filter(room.Clients, func(client *Client) bool { return client.Mode != ERROR })
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
