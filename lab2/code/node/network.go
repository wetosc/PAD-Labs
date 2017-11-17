package main

import (
	"encoding/json"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"pad.com/lab2/code/eugddc"
	"pad.com/lab2/code/tcpClient"
)

var requestID = 0
var nodeData = make(map[string]DogDict)
var querySieve = make(map[string]bool)

func onConnect(conn net.Conn) {
	client := tcpClient.NewClient(conn)
	client.ReadAsync(onMessage)
}

func onMessage(c *tcpClient.Client, data []byte) {
	if len(data) == 0 {
		return
	}
	_, remoteAddr := c.Addr()
	log.Info().Msgf("New message from %v", remoteAddr)
	if strings.Contains(remoteAddr, eugddc.MediatorAddr) {
		m, err := eugddc.MessageFromJSON(data)
		eugddc.CheckError(err, "Error converting JSON")
		onMediatorMessage(c, m)
	} else {
		m, err := NodeMessageFromJSON(data)
		eugddc.CheckError(err, "Error converting JSON")
		onNodeMessage(c, m)
	}
}

func onMediatorMessage(c *tcpClient.Client, m eugddc.Message) {
	log.Debug().Msgf("New mediator request with query: %v", m.Query)
	queryID := strconv.Itoa(requestID)
	nodeData[queryID][myAddr] = items
	requestID++
	for _, node := range connections {
		askNodeInfo(node, queryID)
	}
	time.Sleep(3 * time.Second)
	var newDogs = nodeData[queryID].toSlice()
	response := eugddc.NewMessage_Data(m.Query, newDogs)
	c.Write(eugddc.MessageToJSON(response))
	delete(nodeData, queryID)
}

func askNodeInfo(node string, queryID string) {
	c := tcpClient.Connect(node)
	m := NodeMessage{Origin: myAddr, Sender: myAddr, Query: NodeQuery{ID: queryID, Query: "*"}}
	c.Write(m.toJSON())
}

func onNodeMessage(c *tcpClient.Client, m NodeMessage) {
	log.Debug().Msgf("New node request with query: %v", m.Query)
	if m.Origin == myAddr {
		nodeData[m.Query.ID][m.Sender] = m.Data
	} else {
		_, processed := querySieve[m.Query.ID]
		if processed {
			return
		}
		for _, node := range connections {
			go func(_node string) {
				askNodeInfo(_node, m.Query.ID)
			}(node)
		}
		querySieve[m.Query.ID] = true
		newDogs := nodeData[m.Query.ID].toSlice()
		response := NodeMessage{Origin: m.Origin, Sender: myAddr, Query: m.Query, Data: newDogs}
		c.Write(response.toJSON())
	}
}

type DogDict map[string][]eugddc.Dog

func (d DogDict) toSlice() []eugddc.Dog {
	slice := make([]eugddc.Dog, 0, len(d))
	for _, value := range d {
		slice = append(slice, value...)
	}
	return slice
}

type NodeMessage struct {
	Origin string
	Sender string
	Query  NodeQuery
	Data   []eugddc.Dog
}

type NodeQuery struct {
	ID    string
	Query string
}

func NodeMessageFromJSON(data []byte) (NodeMessage, error) {
	var a NodeMessage
	err := json.Unmarshal(data, a)
	return a, err
}

func (m NodeMessage) toJSON() []byte {
	data, err := json.Marshal(m)
	eugddc.CheckError(err, "[NodeMessage] Error converting to JSON")
	return data
}
