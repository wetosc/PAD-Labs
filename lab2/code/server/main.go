package main

import (
	"flag"
	"math/rand"
	"os"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var lAddrStr string
var port = rand.Intn(10)
var nodeCounter = 1

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	parseFlags()

	log.Info().Msgf("Node started")

	lAddrStr = ":" + strconv.Itoa(9000+port)

	go Step1()
	go Step2()
	for {
	}
}

func parseFlags() {
	portPointer := flag.Int("p", port, "Set the port of the broker.")
	nodeCountPointer := flag.Int("n", nodeCounter, "Set the number of nodes connected to this one")
	logLevel := flag.String("v", "debug", "Set the verbosity level (info or debug)")
	flag.Parse()
	port = *portPointer
	nodeCounter = *nodeCountPointer
	switch *logLevel {
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
