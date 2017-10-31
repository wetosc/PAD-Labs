package main

import (
	"encoding/binary"
	"net"
	"os"
	"time"

	"pad.com/lab2/code/eugddc"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Info().Msgf("Client started")
	addrSender, _ := net.ResolveUDPAddr("udp", eugddc.Address)
	addrMe, _ := net.ResolveUDPAddr("udp", "localhost:9000")

	pingUDPOnce(addrMe, addrSender)
	listenUDP(addrMe)
}

func pingUDPOnce(addr1 *net.UDPAddr, addr2 *net.UDPAddr) {
	conn, err := net.DialUDP("udp", addr1, addr2)
	eugddc.CheckError(err, "Error creating write UDP connection")
	defer conn.Close()
	log.Info().Msgf("Pinging with 1 bit of data ...")
	_, err = conn.Write(make([]byte, 1))
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
			nr := binary.BigEndian.Uint16(data)
			log.Debug().Msgf("Response from %v : %v", addr, nr)
		}
		time.Sleep(1 * time.Second)
	}
}

// func pingUDP(conn *net.UDPConn) {
// 	data := make([]byte, 1000)
// 	conn.SetReadBuffer(8192)
// 	for {
// 		_, err := conn.Write([]byte("hello, world\n"))
// 		if err != nil {
// 			log.Info().Msgf("Error writting data: %v", err)
// 		}
// 		time.Sleep(1 * time.Second)
// 		nr, addr, err := conn.ReadFromUDP(data)
// 		if nr > 0 {
// 			data = data[:nr]
// 			str := string(data)
// 			log.Debug().Msgf("Read from UDP addr %v : 	%v", addr, str)
// 		}
// 		time.Sleep(1 * time.Second)
// 	}
// }
