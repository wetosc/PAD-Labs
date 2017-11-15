package main

import (
	"flag"
	"math/rand"
	"strconv"
	"strings"

	"github.com/rs/zerolog"
	"pad.com/lab2/code/eugddc"
)

var fileName = "data.json"
var items []eugddc.Dog
var connections []string

func parseFlags() {
	portPointer := flag.Int("id", rand.Intn(10), "Set the port / id of the broker.")
	fileNamePointer := flag.String("f", fileName, "Set the file from which to load the data")
	connectionsPointer := flag.String("n", "", "Set the id of nodes connected to this one")
	logLevel := flag.String("v", "debug", "Set the verbosity level (info or debug)")
	flag.Parse()
	myAddr = ":" + strconv.Itoa(9000+*portPointer)
	fileName = *fileNamePointer
	parseNodes(*connectionsPointer)
	switch *logLevel {
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

func loadData() {
	items, _ = eugddc.Load(fileName)
}

func parseNodes(str string) {
	connections = make([]string, 0, 6)
	elems := strings.Split(str, " ")
	for _, n := range elems {
		port, _ := strconv.Atoi(n)
		connections = append(connections, ":"+strconv.Itoa(9000+port))
	}
}
