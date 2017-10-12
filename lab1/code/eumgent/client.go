package eumgent

import (
	"bufio"
	"io"
	"net"

	"github.com/rs/zerolog/log"
)

//Client represents something like a connection, which allows you
type Client struct {
	Queue    string
	Mode     ClientMode
	Incoming chan Message
	Outgoing chan Message
	reader   *bufio.Reader
	writer   *bufio.Writer
}

func (client *Client) Read() {
	shouldBreak := false
	for {
		data := make([]byte, 1000)
		nr, err := client.reader.Read(data)

		shouldBreak = err == io.EOF

		if shouldBreak && nr <= 0 {
			log.Info().Msgf("Connexion closed:	%v", err)
			client.Mode = ERROR
			break
		}
		if err != nil {
			log.Debug().Msgf("Error reading:	%v", err)
		}
		if nr == 0 {
			continue
		}

		data = data[:nr]
		msg, err := MessageFromJSON(data)
		if err != nil {
			log.Debug().Msgf("Invalid JSON:	%v", err)
			continue
		}

		configureClient(client, msg)
		client.Incoming <- msg

		if shouldBreak {
			log.Info().Msgf("Connexion closed:	%v", err)
			client.Mode = ERROR
			break
		}
	}
}

func configureClient(client *Client, msg Message) {
	if msg.Type == PUBLISH {
		client.Mode = PUBLISH
	}
	if msg.Type == SUBSCRIBE {
		client.Mode = SUBSCRIBE
		client.Queue = msg.Queue
	}
}

func (client *Client) Write() {
	for msg := range client.Outgoing {
		data, _ := msg.ToJSON()
		_, err := client.writer.Write(data)
		client.writer.Flush()
		if err != nil {
			log.Debug().Msgf("Error writting data:	%v", err)
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
