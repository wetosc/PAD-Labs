package main

import (
	"os"
	"time"

	"pad.com/lab2/code/tcpClient"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var myAddr string

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.TimeFieldFormat = time.StampMicro
	parseFlags()

	loadData()

	log.Info().Msgf("Node %v started", myAddr)
	log.Info().Msgf("My info: %v", items)
	log.Info().Msgf("My connections: %v", connections)

	tcpClient.StartServer(myAddr, onConnect)
}
