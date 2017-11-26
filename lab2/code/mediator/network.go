package main

import (
	"net"
	"strconv"

	"github.com/rs/zerolog/log"
	"pad.com/lab2/code/eugddc"
	"pad.com/lab2/code/tcpClient"
)

var clients = make(map[string]*tcpClient.Client)
var pool = make(map[string]*tcpClient.Client)
var allData = make(map[string]eugddc.DogDict)
var requestID = 0

func onConnect(conn net.Conn) {
	client := tcpClient.NewClient(conn)
	client.ReadAsync(onMessage)
}

func onMessage(c *tcpClient.Client, data []byte) {
	if len(data) == 0 {
		return
	}
	m, err := eugddc.NodeMessageFromJSON(data)
	eugddc.CheckError(err, "Error converting client JSON")
	log.Info().Msgf("New message %v", m)

	switch m.Type {
	case "REPLY":
		queryID := m.Query.ID
		newC := clients[queryID]
		if newC != nil {
			_, remAddr := c.Addr()
			if allData[queryID] == nil {
				allData[queryID] = make(eugddc.DogDict)
			}
			allData[queryID][remAddr] = m.Data

			if len(allData[m.Query.ID]) == len(connections) {
				log.Debug().Msg("Received data from all connections")
				m.Data = allData[queryID].ToSlice()
				newC.Write(m.ToXML())
				delete(clients, queryID)
				delete(allData, queryID)
			}
		}
	case "CLIENT":
		newQID := strconv.Itoa(requestID)
		clients[newQID] = c
		myAddr, _ := c.Addr()
		newM := eugddc.NodeMessage{Type: "GET", Trace: []string{myAddr}, Query: eugddc.NodeQuery{ID: newQID, Query: m.Query.Query}, Data: nil}
		requestID++
		for _, node := range connections {
			nodeC := getClient(node)
			nodeC.WriteAsync(newM.ToJSON())
		}
	}
}

func getClient(addr string) *tcpClient.Client {
	cachedC := pool[addr]
	if cachedC != nil && !cachedC.Closed {
		return cachedC
	}
	client := tcpClient.Connect(addr)
	pool[addr] = client
	return client
}
