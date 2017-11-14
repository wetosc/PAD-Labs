package eugddc

import (
	"io"
	"net"
)

type Client struct {
	// sync.Mutex
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	client := &Client{conn: conn}
	return client
}

func (c *Client) Write(data []byte) {
	_, err := c.conn.Write(data)
	CheckError(err, "[TCPClient] Error writing data")
}

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

func (c *Client) ReadAsync(callback func([]byte)) {
	var data []byte
	go func() {
		for {
			data = make([]byte, 1024)
			n, err := c.conn.Read(data)
			CheckError(err, "[TCPClient] Error reading data")
			callback(data[:n])
			if err == io.EOF {

			}
		}
	}()
}

func Connect(addr string) *Client {
	conn, err := net.Dial("tcp", addr)
	CheckError(err, "[TCPClient] Error creating connection")
	return NewClient(conn)
}
