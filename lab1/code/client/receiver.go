package main

import (
	"errors"
	"io"
	"net"
	"strconv"

	"github.com/rs/zerolog/log"
	"pad.com/lab1/code/eumgent"
)

//Subscribe connects to broker trough the 'port' and subscribes to the queue 'channel'
func Subscribe(port int, channel string) (net.Conn, error) {
	subscribeMessage, _ := eumgent.Message{Type: eumgent.SUBSCRIBE, Queue: channel}.ToJSON()

	conn, err := net.Dial("tcp", ":"+strconv.Itoa(eumgent.PORT))
	if err != nil {
		log.Debug().Msgf("Error connecting to the broker : %v", err)
		return conn, err
	}

	eumgent.Write(conn, subscribeMessage)
	data, err := eumgent.Read(conn)
	if err != nil {
		log.Debug().Msgf("Error reading the response : %v", err)
	}
	response, err := eumgent.MessageFromJSON(data)

	if response.Type != eumgent.ERROR {
		log.Info().Msgf("Subscribed succesfully to queue\t%v\t%v", response.Queue, response.Info)
		return conn, err
	}

	conn.Close()
	return nil, errors.New(response.Info)
}

//Listen reads data from the connection until the connection is closed by the server
//If it receives a message, it prints it on Info level
func Listen(conn net.Conn) {
	for {
		var data = make([]byte, 0, 100)
		nr, err := conn.Read(data)
		if err != nil {
			if err == io.EOF {
				log.Info().Msgf("Connection closed by server : \t%v", err)
				break
			}
			log.Debug().Msgf("Error reading data : \t%v", err)
		}
		if nr > 0 {
			data = data[:nr]
			msg, _ := eumgent.MessageFromJSON(data)
			log.Info().Msgf("Received message: %v", msg)
		}
	}
}
