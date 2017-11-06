package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var lAddrStr string

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Info().Msgf("Node started")

	lAddrStr = ":9001"

	go Step1()
	go Step2()
	for {
	}
}
