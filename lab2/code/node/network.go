package main

import (
	"net"
	"sync"

	"github.com/rs/zerolog/log"
	"pad.com/lab2/code/eugddc"
	"pad.com/lab2/code/tcpClient"
)

type NodeState struct {
	Data       eugddc.DogDict
	SubQueries int
	State      string
}

type NodeMap struct {
	sync.Mutex
	q map[string]*NodeState
}

type ConnPool struct {
	sync.Mutex
	pool map[string]*tcpClient.Client
}

var nD = &NodeMap{q: make(map[string]*NodeState)}
var cP = &ConnPool{pool: make(map[string]*tcpClient.Client)}

func onConnect(conn net.Conn) {
	client := tcpClient.NewClient(conn)
	client.ReadAsync(onMessage)
}

func onMessage(c *tcpClient.Client, data []byte) {
	if len(data) == 0 {
		return
	}
	m, err := eugddc.NodeMessageFromJSON(data)
	eugddc.CheckError(err, "Error converting JSON  ***\n"+string(string(data))+"\n*** ")
	_, rAddr := c.Addr()
	log.Debug().Msgf("\nNew request from %v \n%v\n", rAddr, m)
	switch m.Type {
	case "REPLY":
		onREPLYMessage(m, rAddr)
	case "GET":
		onGETMessage(m)
	}
}

func onREPLYMessage(m eugddc.NodeMessage, rAddr string) {
	queryID := m.Query.ID
	nD.Lock()
	defer nD.Unlock()

	nD.q[queryID].Data[rAddr] = m.Data // Add collected data to storage
	log.Debug().Msgf("REPLY message collected data: %v,\nnecessary responses: %v", len(nD.q[queryID].Data), nD.q[queryID].SubQueries)
	if len(m.Trace) > 0 && // Check if it's not the last node in the chain
		len(nD.q[queryID].Data) == nD.q[queryID].SubQueries+1 { // Check if has all the data (+1 because of my items)
		parent, rest := m.Trace[0], m.Trace[1:]
		m.Trace = rest
		m.Data = nD.q[queryID].Data.ToSlice()
		m.Data = eugddc.Perform(m.Query, m.Data)
		newC := getClient(parent)
		newC.WriteAsync(m.ToJSON())
		log.Debug().Msgf("\nSent reply message to %v: \n%v\n", parent, m)
		delete(nD.q, queryID)
	}
}

func onGETMessage(m eugddc.NodeMessage) {
	queryID := m.Query.ID
	nD.Lock()
	defer nD.Unlock()
	if contains(m.Trace, myAddr) {
		log.Debug().Msgf("Received message from myself ???")
		return
	}
	if nD.q[queryID] != nil { // This request was here before, send an empty REPLY
		m.Type = "REPLY"
		parent, rest := m.Trace[0], m.Trace[1:]
		m.Trace = rest
		newC := getClient(parent)
		newC.WriteAsync(m.ToJSON())
		log.Debug().Msgf("\nSent empty reply message to %v: \n%v\n", parent, m)
		return
	}

	m.Trace = append([]string{myAddr}, m.Trace...) // Insert myself at begining of trace
	i := 0
	for _, node := range connections {
		if !contains(m.Trace, node) { // So it won't go back from where it came
			nC := getClient(node)
			nC.WriteAsync(m.ToJSON())
			log.Debug().Msgf("\nSent all message to %v: \n%v\n", node, m)
			i++
		}
	}
	if i == 0 { // End of line == leaf, send REPLY message with only my data
		m.Type = "REPLY"
		parent, rest := m.Trace[1], m.Trace[2:] // Remove myself (trace[0]) and get the receiver (trace[1])
		m.Trace = rest
		m.Data = append(m.Data, items...)
		m.Data = eugddc.Perform(m.Query, m.Data)
		newC := getClient(parent)
		newC.WriteAsync(m.ToJSON())
		log.Debug().Msgf("\nSent EOL message to %v: \n%v\n", parent, m)
	}
	nD.q[queryID] = &NodeState{SubQueries: i, Data: make(eugddc.DogDict)}
	nD.q[queryID].Data[myAddr] = items
}

func getClient(addr string) *tcpClient.Client {
	cP.Lock()
	defer cP.Unlock()
	cachedC := cP.pool[addr]
	if cachedC != nil && !cachedC.Closed {
		return cachedC
	}
	client := tcpClient.Connect(addr)
	cP.pool[addr] = client
	return client
}

func contains(arr []string, elem string) bool {
	for _, el := range arr {
		if el == elem {
			return true
		}
	}
	return false
}
