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
	m := eugddc.NodeMessage{Type: "CLIENT", Trace: nil, Query: eugddc.NodeQuery{ID: "", Query: "*", Params: params}, Data: nil}
	client.Write(m.ToJSON())
	log.Debug().Msg("Sent message '*' to mediator")
}

func onMediatorMessage(c *tcpClient.Client, data []byte) {
	if len(data) == 0 {
		return
	}
	if validate(data) {
		log.Info().Msgf("XML Document validated successfully")
		m, err := eugddc.NodeMessageFromXML(data)
		eugddc.CheckError(err, "Error parsing mediator XML")
		log.Debug().Msgf("Received data: %v", m.Data)
	} else {
		log.Info().Msgf("Received invalid XML response")
	}
}
