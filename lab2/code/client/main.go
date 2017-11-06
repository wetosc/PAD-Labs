package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Info().Msgf("Client started")

	go step1()
	time.Sleep(10 * time.Second)
	if len(nodes) == 0 {
		panic("No nodes detected, something is wrong")
	}
	maven := step2()
	log.Debug().Msgf("Maven: %v", maven.Addr)
	step3(maven.Addr)
}
