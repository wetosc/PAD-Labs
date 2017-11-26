package main

import (
	"github.com/rs/zerolog/log"
	"pad.com/lab2/code/eugddc"
	"pad.com/lab2/code/tcpClient"
)

func requestAllData() {
	client := tcpClient.TryConnectSync(eugddc.MediatorAddr)
	lAddr, rAddr := client.Addr()
	log.Debug().Msgf("Connected to mediator %v -> %v", lAddr, rAddr)
	client.ReadAsync(onMediatorMessage)
	params := &eugddc.QueryParams{Action: "FILTER", Operation: "<", Param: "AGE", Value: "8"}
	m := eugddc.NodeMessage{Type: "CLIENT", Format: format, Trace: nil, Query: eugddc.NodeQuery{ID: "", Query: "*", Params: params}, Data: nil}
	client.Write(m.ToJSON())
	log.Debug().Msgf("Sent message %v to mediator", m)
}

func onMediatorMessage(c *tcpClient.Client, data []byte) {
	if len(data) == 0 {
		return
	}

	switch format {
	case "XML":
		if validateXML(data) {
			log.Info().Msgf("XML Document validated successfully")
			m, err := eugddc.NodeMessageFromXML(data)
			eugddc.CheckError(err, "Error parsing mediator XML")
			log.Debug().Msgf("Received data: %v", m.Data)
		} else {
			log.Info().Msgf("Received invalid XML response")
		}
	case "JSON":
		m, err := eugddc.NodeMessageFromJSON(data)
		eugddc.CheckError(err, "Error parsing mediator JSON")
		if validateJSON(m) {
			log.Info().Msgf("JSON Document validated successfully")
			log.Debug().Msgf("Received data: %v", m.Data)
		} else {
			log.Info().Msgf("Received invalid JSON response")
		}
	}

}
