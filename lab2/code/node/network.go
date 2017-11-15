package main

import (
	"net"
	"time"

	"github.com/rs/zerolog/log"
	"pad.com/lab2/code/eugddc"
	"pad.com/lab2/code/tcpClient"
)

func onConnect(conn net.Conn) {
	client := tcpClient.NewClient(conn)
	client.ReadAsync(onMessage)
}

func onMessage(c *tcpClient.Client, data []byte) {
	if len(data) == 0 {
		return
	}
	m, err := eugddc.MessageFromJSON(data)
	eugddc.CheckError(err, "Error converting JSON")
	log.Debug().Msgf("New request with query: %v", m.Query)
	time.Sleep(500 * time.Millisecond)
	switch m.Query {
	case "*":
		allData := make([]eugddc.Dog, 0, 10)
		allData = append(allData, items...)
		for _, node := range connections {
			_c := tcpClient.Connect(node)
			m.Query += "-"
			_c.Write(eugddc.MessageToJSON(m))
			newData := _c.Read()
			if len(data) == 0 {
				continue
			}
			newDogs, _ := eugddc.DogsFromJSON(newData)
			allData = append(allData, newDogs...)
			log.Debug().Msgf("Received new dogs from friends: %v", newDogs)
		}

		time.Sleep(500 * time.Millisecond)
		resM := eugddc.NewMessage_Data(m.Query, allData)
		response := eugddc.MessageToJSON(resM)
		c.Write(response)
		log.Debug().Msgf("Sent response to the mediator: %v", response)

	case "*-":
		response := eugddc.JSONfromDogs(items)
		c.Write(response)
		log.Debug().Msgf("Sent response to friend: %v", items)
	}

}
