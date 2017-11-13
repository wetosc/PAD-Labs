package main

import (
	"net"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"pad.com/lab2/code/eugddc"
)

// Step1 consist in joining a UDP multicast group and sending the number of connected nodes tho whoewer pings
func Step1() {
	// ifaces, _ := net.Interfaces()
	// fmt.Println(ifaces)
	eth0, _ := net.InterfaceByName("lo")
	// fmt.Println(eth0)
	addr, _ := net.ResolveUDPAddr("udp", eugddc.MulticastAddress)
	conn, err := net.ListenMulticastUDP("udp", eth0, addr)
	eugddc.CheckError(err, "Error joining the multicast group")
	listenUDP(conn)
}

// Step2 consist in listeng on TCP and sending data to all connections
func Step2() {
	listener, err := net.Listen("tcp", lAddrStr)
	eugddc.CheckError(err, "Error creating TCP Listener")
	log.Debug().Msgf("TCP Listener started : %v", listener)
	for {
		conn, err := listener.Accept()
		eugddc.CheckError(err, "Error accepting TCP connection")
		go handleTCPConn(conn)
	}
}

func handleTCPConn(conn net.Conn) {
	client := eugddc.NewClient(conn)
	defer conn.Close()
	defer client.Close()

LoopTCP:
	for {
		data := <-client.Incoming
		msg := string(data)
		log.Debug().Msgf("Received data request: %v", msg)
		switch msg {
		case "-*":
			client.Outgoing <- eugddc.JSONfromDogs(items)
			log.Debug().Msgf("Sent items to %v", conn.RemoteAddr().String())
			break LoopTCP
		case "*":
			collectedItems := <-collectData()
			log.Debug().Msgf("Received ALL request")
			log.Debug().Msgf("CollectedItems: %v", collectedItems)
			allItems := append(collectedItems, items...)
			client.Outgoing <- eugddc.JSONfromDogs(allItems)
			break LoopTCP
		}
	}
}

func collectData() chan []eugddc.Dog {
	outChan := make(chan []eugddc.Dog, 1000)

	for _, str := range connections {
		go func(str string) {
			conn, err := net.Dial("tcp", str)
			eugddc.CheckError(err, "Error connectiong to nodes")
			client := eugddc.NewClient(conn)
			defer conn.Close()
			defer client.Close()

			client.Outgoing <- []byte("-*")

		LoopCollect:
			for {
				select {
				case data := <-client.Incoming:
					dogs, _ := eugddc.DogsFromJSON(data)
					outChan <- dogs
				case <-time.After(time.Second * 2):
					break LoopCollect
				}
			}
		}(str)
	}
	return outChan
}

func listenUDP(conn *net.UDPConn) {
	data := make([]byte, 1000)
	conn.SetReadBuffer(8192)
	for {
		nr, addr, err := conn.ReadFromUDP(data)
		if err != nil {
			log.Info().Msgf("Error reading from UDP: %v", err)
		}
		if nr > 0 {
			data = data[:nr]
			str := string(data)
			log.Debug().Msgf("Read from UDP addr %v : 	%v", addr, str)
			sendNr(addr)
		}
	}
}

func sendNr(addr *net.UDPAddr) {
	lAddr, _ := net.ResolveUDPAddr("udp", lAddrStr)
	conn2, _ := net.DialUDP("udp", lAddr, addr)
	defer conn2.Close()
	str := strconv.Itoa(nodeCounter)
	conn2.Write([]byte(str))
}
