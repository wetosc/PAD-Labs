package main

import (
	"flag"
	"net"
	"os"
	"strconv"
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

	conn, err := net.Dial("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Debug().Msgf("Error connecting to the broker : %v", err)
		panic(err)
	}

	client := eumgent.NewClient(conn)

	if isSender {
		log.Info().Msgf("Client of type Sender started on port %v, with queue '%v' and info '%v'", port, queue, payload)
		startSender(client)
	} else {
		log.Info().Msgf("Client of type Subscriber started on port %v, with queue '%v'", port, queue)
		startSender(client)
	}
}

func startSender(client *eumgent.Client) {
	message := eumgent.Message{Type: eumgent.PUBLISH, Queue: queue, Info: payload}
	for {
		time.Sleep(time.Second)
		client.Outgoing <- message
	}
}

func startReceiver(client *eumgent.Client) {
	for {
		select {
		case msg := <-client.Incoming:
			log.Info().Msgf("Received message: %v", msg)
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
