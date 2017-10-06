package main

import (
	"flag"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"pad.com/lab1/code/eumgent"
)

var port = eumgent.PORT
var isSender = true
var payload = "Lorem Ipsum"
var queue = "default"

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	parseFlags()
	isSender = false
	if isSender {
		log.Info().Msgf("Client of type Sender started on port %v, with queue '%v' and info '%v'", port, queue, payload)
		message := eumgent.Message{Type: eumgent.PUBLISH, Queue: queue, Info: payload}
		for {
			time.Sleep(time.Second)
			SendMessage(port, &message)
		}
	} else {
		log.Info().Msgf("Client of type Subscriber started on port %v, with queue '%v'", port, queue)
		conn, err := Subscribe(port, queue)
		if err == nil {
			Listen(conn)
		} else {
			log.Debug().Msgf("The client will close now.\t%v", err)
		}
	}
}

func parseFlags() {
	clientType := flag.String("type", "sender", "Set the type of client, sender or receiver")
	info := flag.String("info", "", "Set the payload for published messages")
	portPointer := flag.Int("port", port, "Set the port of the broker.")
	logLevel := flag.String("v", "debug", "Set the verbosity level (info or debug)")
	queuePointer := flag.String("queue", queue, "Set the queue on which the client will work")
	flag.Parse()
	port = *portPointer
	switch *clientType {
	case "sender":
		isSender = true
	case "receiver":
		isSender = false
	default:
		panic("Wrong type flag")
	}
	if len(*info) > 0 {
		payload = *info
	}
	switch *logLevel {
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	queue = *queuePointer
}
