package main

import (
	"flag"

	"github.com/rs/zerolog"
)

var format string

func parseFlags() {
	logLevel := flag.String("v", "debug", "Set the verbosity level (info or debug)")
	formatPointer := flag.String("f", "", "Set the format in which to get data")

	flag.Parse()

	format = *formatPointer

	switch *formatPointer {
	case "xml", "XML":
		format = "XML"
	case "json", "JSON":
		format = "JSON"
	default:
		format = "JSON"
	}
	switch *logLevel {
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
