package main

import (
	"flag"
	"strconv"
	"strings"

	"pad.com/lab2/code/eugddc"

	"github.com/rs/zerolog"
)

var connections []string

func parseFlags() {
	logLevel := flag.String("v", "debug", "Set the verbosity level (info or debug)")
	connectionsPointer := flag.String("n", "", "Set the id of nodes from which to get data")

	flag.Parse()

	parseNodes(*connectionsPointer)

	switch *logLevel {
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

func parseNodes(str string) {
	connections = make([]string, 0, 6)
	elems := strings.Split(str, " ")
	for _, n := range elems {
		port, _ := strconv.Atoi(n)
		connections = append(connections, ":"+strconv.Itoa(eugddc.NodeBasePort+port))
	}
}
