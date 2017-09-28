package main

import (
	"log"
	"net"
	"strconv"
	"time"

	"pad.com/lab1/eumgent"
)

var message = eumgent.Message{Type: eumgent.POST, Info: "Some info"}

func main() {
	for {
		time.Sleep(time.Second)
		go func() {
			conn, err := net.Dial("tcp", ":"+strconv.Itoa(eumgent.PORT))
			if err != nil {
				log.Fatal(err)
			}
			data, _ := message.ToJSON()
			eumgent.Write(conn, data)
			response, err := eumgent.ReadString(conn)
			if err != nil {
				log.Println(err)
			} else {
				log.Printf("Response:\t%v", response)
			}
			conn.Close()
		}()
	}
}
