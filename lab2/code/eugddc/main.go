//Package eugddc stands for "EUGenius Distributed Data Collections"
// and contains the shared code of the application
package eugddc

import (
	"bufio"
	"net"

	"github.com/rs/zerolog/log"
)

//Address Defines the address of the UDP group
const Address = "239.0.0.0:8000"

type Client struct {
	Incoming chan []byte
	Outgoing chan []byte
	reader   *bufio.Reader
	writer   *bufio.Writer
}

func (client *Client) Read() {
	for {
		data := make([]byte, 1000)
		nr, err := client.reader.Read(data)
		if err != nil {
			log.Debug().Msgf("Error reading:	%v", err)
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
		Incoming: make(chan []byte),
		Outgoing: make(chan []byte),
		reader:   reader,
		writer:   writer,
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
