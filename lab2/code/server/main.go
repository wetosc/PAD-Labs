package main

import (
	"net"
	"os"
	"strconv"
	"time"

	"pad.com/lab2/code/eugddc"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var lAddrStr string

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Info().Msgf("Node started")

	lAddrStr = ":9001"

	go step1()
	go listenTCP()
	time.Sleep(10 * time.Second)
}

func step1() {
	// ifaces, _ := net.Interfaces()
	// fmt.Println(ifaces)
	eth0, _ := net.InterfaceByName("lo")
	// fmt.Println(eth0)
	addr, _ := net.ResolveUDPAddr("udp", eugddc.Address)
	conn, err := net.ListenMulticastUDP("udp", eth0, addr)
	if err != nil {
		log.Info().Msgf("Error joining the multicast group: %v", err)
	}
	listenUDP(conn)
}

func listenTCP() {
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
	for {
		data := <-client.Incoming
		log.Debug().Msgf("Received data request: %v", data)
	}
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
	str := strconv.Itoa(1)
	conn2.Write([]byte(str))

}
