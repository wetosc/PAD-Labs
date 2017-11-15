package main

import (
	"github.com/rs/zerolog/log"
	"pad.com/lab2/code/eugddc"
	"pad.com/lab2/code/tcpClient"
)

func requestAllData() {
	client := tcpClient.TryConnectSync(eugddc.MediatorAddr)
	log.Debug().Msg("Connected to mediator")
	client.ReadAsync(onMediatorMessage)
	m := eugddc.NewMessage("*")
	client.Write(eugddc.MessageToJSON(m))
	log.Debug().Msg("Sent message '*' to mediator")
}

func onMediatorMessage(c *tcpClient.Client, data []byte) {
	if len(data) == 0 {
		return
	}
	m, err := eugddc.MessageFromJSON(data)
	eugddc.CheckError(err, "Error parsing mediator JSON")
	log.Debug().Msgf("Message: %v", m)
}
