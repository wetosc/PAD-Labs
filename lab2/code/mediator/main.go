package main

import (
	"os"

	"pad.com/lab2/code/eugddc"
	"pad.com/lab2/code/tcpClient"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var port = eugddc.MediatorAddr

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	parseFlags()
	log.Info().Msgf("Mediator started on port %v", port)

	tcpClient.StartServer(port, onConnect)
}
