//Package eugddc stands for "EUGenius Distributed Data Collections"
// and contains the shared code of the application
package eugddc

import (
	"bufio"
	"io"
	"net"

	"github.com/rs/zerolog/log"
)

//MulticastAddress Defines the address of the UDP group
const MulticastAddress = "239.0.0.0:8000"

// Client is a helper for connections; allows use of channels for writing and reading
type Client struct {
	Incoming chan []byte
	Outgoing chan []byte
	reader   *bufio.Reader
	writer   *bufio.Writer
	closed   bool
}

func (client *Client) Read() {
	for {
		if client.closed {
			return
		}
		data := make([]byte, 1000)
		nr, err := client.reader.Read(data)
		CheckError(err, "Error reading")
		if err == io.EOF {
			break
		}
		if nr == 0 {
			continue
		}
		data = data[:nr]
		client.Incoming <- data
	}
}

func (client *Client) Write() {
	for data := range client.Outgoing {
		if client.closed {
			return
		}
		_, err := client.writer.Write(data)
		client.writer.Flush()
		CheckError(err, "Error writting data")
	}
}

//Listen checks if there are messages for writing or for reading
func (client *Client) Listen() {
	go client.Read()
	go client.Write()
}

// Close ends
func (client *Client) Close() {
	client.closed = true
}

//NewClient creates a new client from a connection
func NewClient(connection net.Conn) *Client {
	writer := bufio.NewWriter(connection)
	reader := bufio.NewReader(connection)

	client := &Client{
		Incoming: make(chan []byte),
		Outgoing: make(chan []byte),
		reader:   reader,
		writer:   writer,
		closed:   false,
	}

	client.Listen()

	return client
}

//CheckError checks if error is not nil and if so prints the description in logs on level INFO
func CheckError(err error, info string) {
	if err != nil {
		log.Info().Msgf("%v : %v", info, err)
	}
}
