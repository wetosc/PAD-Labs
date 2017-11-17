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
	log.Info().Msgf("New message from client with query %v", m.Query)
	var allData = make([]eugddc.Dog, 0, 10)

	for _, node := range connections {

		func(_node string) {
			nodeC := tcpClient.Connect(node)
			addr1, addr2 := nodeC.Addr()
			log.Debug().Msgf("Connexion from %v to %v", addr1, addr2)
			nodeC.Write(eugddc.MessageToJSON(m))

			nodeC.ReadAsync(func(_c *tcpClient.Client, data []byte) {
				if len(data) == 0 {
					return
				}
				nodeData, err := eugddc.MessageFromJSON(data)
				eugddc.CheckError(err, "Error converting JSON")
				newDogs := nodeData.Data
				log.Info().Msgf("New message from node with %v elements", len(newDogs))
				allData = append(allData, newDogs...)
			})
		}(node)
	}
	time.Sleep(3 * time.Second)
	log.Info().Msgf("Send collected data to client: %v", allData)
	response := eugddc.MessageToJSON(eugddc.NewMessage_Data(m.Query, allData))
	c.Write(response)

}
