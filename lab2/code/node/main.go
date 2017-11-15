package main

import (
	"os"

	"pad.com/lab2/code/tcpClient"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var myAddr string

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	parseFlags()

	loadData()

	log.Info().Msgf("Node %v started", myAddr)
	log.Info().Msgf("My info: %v", items)

	tcpClient.StartServer(myAddr, onConnect)
}
