package main

import (
	"log"
	"net"
	"strconv"
	"time"

	"pad.com/lab1/eumgent"
)

var message = eumgent.Message{Type: eumgent.GET, Info: "Give messages!!"}

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
			response, err := eumgent.Read(conn)
			if err != nil {
				log.Println(err)
			} else {
				obj, _ := eumgent.MessageFromJSON(response)
				log.Printf("Response:\t%v", obj.Info)
			}
			conn.Close()
		}()
	}
}
