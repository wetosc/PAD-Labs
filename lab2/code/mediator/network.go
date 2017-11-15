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
	var allData = make([]eugddc.Dog, 0, 10)

	for _, node := range connections {

		func(_node string) {
			nodeC := tcpClient.Connect(node)
			nodeC.Write(eugddc.MessageToJSON(m))

			nodeC.ReadAsync(func(_c *tcpClient.Client, data []byte) {
				if len(data) == 0 {
					return
				}
				nodeData, err := eugddc.MessageFromJSON(data)
				newDogs := nodeData.Data
				eugddc.CheckError(err, "Error converting JSON")
				allData = append(allData, newDogs...)
			})
		}(node)
	}
	time.Sleep(3 * time.Second)
	log.Info().Msgf("Send collected data to client: %v", allData)
	response := eugddc.MessageToJSON(eugddc.NewMessage_Data(m.Query, allData))
	c.Write(response)

}
