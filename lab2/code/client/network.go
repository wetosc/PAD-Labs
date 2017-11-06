package main

import (
	"net"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"pad.com/lab2/code/eugddc"
)

var nodes = make([]Node, 0, 6)

// Node contains info about a data node, and helps finding the maven
type Node struct {
	Nodes int
	Addr  net.Addr
}

// Step1 consist in pinging to a udp multicast group and then listening for udp unicast responses
func Step1() {
	log.Debug().Msg("Started Step 1")
	addrSender, _ := net.ResolveUDPAddr("udp", eugddc.Address)
	addrMe, _ := net.ResolveUDPAddr("udp", "localhost:9000")

	pingUDPOnce(addrMe, addrSender)
	listenUDP(addrMe)
}

// Step2 consist in finding the maven (the node with the most connections)
func Step2() Node {
	maven := nodes[0]
	for _, node := range nodes {
		if node.Nodes > maven.Nodes {
			maven = node
		}
	}
	return maven
}

// Step3 consist in connecting over TCP to the maven and reqesting all data
func Step3(addr net.Addr) {
	conn, err := net.Dial("tcp", addr.String())
	eugddc.CheckError(err, "Error creating connection")
	client := eugddc.NewClient(conn)
	client.Outgoing <- []byte("*")
	for {
		data := <-client.Incoming
		log.Debug().Msgf("Received data: %v", string(data))
	}
}

func pingUDPOnce(addr1 *net.UDPAddr, addr2 *net.UDPAddr) {
	conn, err := net.ListenUDP("udp", addr1)
	eugddc.CheckError(err, "Error creating write UDP connection")
	defer conn.Close()
	log.Info().Msgf("Pinging with 1 bit of data ...")
	_, err = conn.WriteToUDP(make([]byte, 1), addr2)
	eugddc.CheckError(err, "Error sending UDP ping")
}

func listenUDP(addr *net.UDPAddr) {
	conn, err := net.ListenUDP("udp", addr)
	eugddc.CheckError(err, "Error creating read UDP connection")
	defer conn.Close()
	data := make([]byte, 1024)
	conn.SetReadBuffer(8192)

	for {
		log.Debug().Msgf("Listening ...")
		nr, addr, err := conn.ReadFromUDP(data)
		eugddc.CheckError(err, "Error on read from UDP")
		if nr > 0 {
			data = data[:nr]
			str := string(data)
			nr, _ := strconv.Atoi(str)
			nodes = append(nodes, Node{Nodes: nr, Addr: addr})
			log.Debug().Msgf("Response from %v : %v", addr, nr)
		}
		time.Sleep(1 * time.Second)
	}
}
