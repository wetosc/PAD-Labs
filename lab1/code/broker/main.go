package main

import (
	"flag"
	"net"
	"strconv"

	"github.com/rs/zerolog"
	"pad.com/lab1/code/eumgent"
)

// var queue = NewQueue(100)

var port = eumgent.PORT

func main() {
	room := eumgent.NewRoom()

	listener, _ := net.Listen("tcp", ":"+strconv.Itoa(port))
	for {
		conn, _ := listener.Accept()
		room.Joins <- conn
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
