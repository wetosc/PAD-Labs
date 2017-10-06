package main

import (
	"net"
	"strconv"

	"github.com/rs/zerolog/log"

	"pad.com/lab1/code/eumgent"
)

//SendMessage publishes the message m to the broker using the provided port
func SendMessage(port int, m *eumgent.Message) {
	conn, err := net.Dial("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Debug().Msgf("Error connecting to the broker: %v", err.Error)
		return
	}
	defer conn.Close()

	data, err := m.ToJSON()
	if err != nil {
		log.Debug().Msgf("Error on message serialization: %v", err.Error())
	}
	err = eumgent.Write(conn, data)
	if err != nil {
		log.Debug().Msgf("Error sending the message: %v", err.Error())
	}
	log.Info().Msgf("Published a new message: %v", m)
}
