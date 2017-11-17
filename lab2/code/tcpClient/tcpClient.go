package tcpClient

import (
	"io"
	"net"
	"time"

	"github.com/rs/zerolog/log"
)

type Client struct {
	conn net.Conn
}

func (c *Client) Addr() (string, string) {
	return c.conn.LocalAddr().String(), c.conn.RemoteAddr().String()
}

func NewClient(conn net.Conn) *Client {
	client := &Client{conn: conn}
	return client
}

func (c *Client) Write(data []byte) {
	_, err := c.conn.Write(data)
	CheckError(err, "[TCPClient] Error writing data")
}

// Read reads sync from client and returns []byte.
// If an error occurs, the data may be empty, so you have to check for yourself.
func (c *Client) Read() []byte {
	data := make([]byte, 1024)
	n, err := c.conn.Read(data)
	CheckError(err, "[TCPClient] Error reading data")
	return data[:n]
}

func (c *Client) WriteAsync(data []byte) {
	go func() {
		c.Write(data)
	}()
}

func (c *Client) ReadAsync(callback func(*Client, []byte)) {
	var data []byte
	go func() {
		for {
			data = make([]byte, 1024)
			n, err := c.conn.Read(data)
			callback(c, data[:n])
			if err == io.EOF {
				CheckError(err, "[TCPClient] The connection closed")
				break
			} else {
				CheckError(err, "[TCPClient] Error reading data")
			}
		}
	}()
}

func Connect(addr string) *Client {
	conn, err := net.Dial("tcp", addr)
	CheckError(err, "[TCPClient] Error creating connection")
	return NewClient(conn)
}

func TryConnectSync(addr string) *Client {
	conn, err := net.Dial("tcp", addr)
	for err != nil {
		conn, err = net.Dial("tcp", addr)
		CheckError(err, "[TCPClient] Error creating connection")
		time.Sleep(1000 * time.Millisecond)
	}
	return NewClient(conn)
}

func StartServer(addr string, callback func(net.Conn)) {
	listener, err := net.Listen("tcp", addr)
	CheckError(err, "[TCPClient] Error creating the listener")
	for {
		conn, err := listener.Accept()
		CheckError(err, "[TCPClient] Error accepting connection")
		callback(conn)
	}
}

func StartServerAsync(addr string, callback func(net.Conn)) {
	listener, err := net.Listen("tcp", addr)
	CheckError(err, "[TCPClient] Error creating the listener")
	go func() {
		for {
			conn, err := listener.Accept()
			CheckError(err, "[TCPClient] Error accepting connection")
			callback(conn)
		}
	}()
}

//CheckError checks if error is not nil and if so prints the description in logs on level INFO
func CheckError(err error, info string) {
	if err != nil {
		log.Debug().Msgf("%v : %v", info, err)
	}
}
