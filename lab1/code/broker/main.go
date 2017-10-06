package main

import (
	"flag"
	"net"
	"os"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"pad.com/lab1/code/eumgent"
)

//Subscriber stores info about a subscriber, to be able to send him messages
type Subscriber struct {
	Conn  net.Conn
	Queue string
}

var queue = NewQueue(100)
var subscribers = NewQueue(100)
var port = eumgent.PORT

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	parseFlags()

	log.Info().Msg("Message broker started")

	server, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Debug().Msgf("Unable to launch server: %v", err)
	}
	connections := connectionChannel(server)

	for {
		go handleConnection(<-connections)
	}
}

func parseFlags() {
	portPointer := flag.Int("port", port, "Set the port of the broker.")
	logLevel := flag.String("v", "debug", "Set the verbosity level (info or debug)")
	flag.Parse()
	port = *portPointer

	switch *logLevel {
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

func connectionChannel(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	go func() {
		for {
			client, err := listener.Accept()
			if err != nil {
				log.Debug().Msgf("Error accepting the connection : %v", err)
				continue
			}
			ch <- client
		}
	}()
	return ch
}

func handleConnection(conn net.Conn) {

	// client.SetDeadline(time.Now().Add(time.Second * 2))

	b, err := eumgent.Read(conn)
	if err != nil {
		log.Debug().Msgf("Unable to read from connection: %v", err)
		return
	}
	m, err := eumgent.MessageFromJSON(b)
	if err != nil {
		log.Debug().Msgf("Bad data, unable to parse JSON : %v", err)
		return
	}

	switch m.Type {
	case eumgent.PUBLISH:

		log.Info().Msgf("Message received: %v", m)

		data, _ := eumgent.Message{Type: eumgent.RESPONSE, Queue: m.Queue}.ToJSON()
		eumgent.Write(conn, data)

		filtered := subscribers.Filter(func(item interface{}) bool {
			log.Debug().Msgf("Filtering: %v", item)
			return item.(Subscriber).Queue == m.Queue
		})

		newMData, _ := eumgent.Message{Type: eumgent.DELIVER, Queue: m.Queue, Info: m.Info}.ToJSON()

		for _, subscriber := range filtered {

			go func(item Subscriber) {

				err := eumgent.Write(item.Conn, newMData)
				if err != nil {
					log.Debug().Msgf("Error delivering message : %v", err)
					log.Info().Msgf("Unable to deliver message, the subscriber will be removed")
					// subscribers.Delete(item)
				}

			}(subscriber.(Subscriber))

		}
		conn.Close()

	case eumgent.SUBSCRIBE:
		newClient := Subscriber{Queue: m.Queue, Conn: conn}
		subscribers.Push(newClient)
		data, _ := eumgent.Message{Type: eumgent.RESPONSE, Queue: m.Queue}.ToJSON()
		eumgent.Write(conn, data)
		log.Info().Msgf("Subscriber connected on queue '%v', conn: %v", newClient.Queue, conn)
	}
}
