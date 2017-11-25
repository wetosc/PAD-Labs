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
	m := eugddc.NodeMessage{Type: "CLIENT", Trace: nil, Query: eugddc.NodeQuery{ID: "", Query: "*"}, Data: nil}
	client.Write(m.ToJSON())
	log.Debug().Msg("Sent message '*' to mediator")
}

func onMediatorMessage(c *tcpClient.Client, data []byte) {
	if len(data) == 0 {
		return
	}
	m, err := eugddc.NodeMessageFromJSON(data)
	eugddc.CheckError(err, "Error parsing mediator JSON")
	log.Debug().Msgf("Received data: %v", m.Data)
}
