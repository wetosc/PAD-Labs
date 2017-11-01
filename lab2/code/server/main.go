package main

import (
	"net"
	"os"
	"strconv"

	"pad.com/lab2/code/eugddc"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Info().Msgf("Node started")
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
	conn2, _ := net.DialUDP("udp", nil, addr)
	str := strconv.Itoa(1)
	conn2.Write([]byte(str))
	conn2.Close()
}
