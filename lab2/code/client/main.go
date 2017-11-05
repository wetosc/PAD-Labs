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

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Info().Msgf("Client started")

	go step1()
	time.Sleep(10 * time.Second)
}

func step1() {
	log.Info().Msg("Started Step 1")
	addrSender, _ := net.ResolveUDPAddr("udp", eugddc.Address)
	addrMe, _ := net.ResolveUDPAddr("udp", "localhost:9000")

	pingUDPOnce(addrMe, addrSender)
	listenUDP(addrMe)
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
			log.Debug().Msgf("Response from %v : %v", addr, nr)
		}
		time.Sleep(1 * time.Second)
	}
}
