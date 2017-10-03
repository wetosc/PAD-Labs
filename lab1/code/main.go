package main

import (
	"log"
	"net"
	"strconv"
	"time"

	"pad.com/lab1/code/eumgent"
)

var queue = NewQueue(100)

func main() {
	log.Println("Hello World")
	server, err := net.Listen("tcp", ":"+strconv.Itoa(eumgent.PORT))
	if server == nil || err != nil {
		log.Panic("Error: " + err.Error())
	}

	restoredQueue := loadQueue()
	if restoredQueue != nil {
		queue = restoredQueue
	}
	log.Printf("Restored messages:\t%v", len(queue.Items))
	createQueueSaver()

	connections := connectionChannel(server)

	for {
		go handleConnection(<-connections)
	}
}

func connectionChannel(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	go func() {
		for {
			client, err := listener.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			ch <- client
		}
	}()
	return ch
}

func handleConnection(client net.Conn) {

	client.SetDeadline(time.Now().Add(time.Second * 2))

	b := make([]byte, 1024)
	length, _ := client.Read(b)
	m, err := eumgent.MessageFromJSON(b[:length])

	if err != nil {
		log.Printf("Error:\tBad JSON :: \t%v", err)
	}
	switch m.Type {
	case eumgent.GET:
		info := queue.Pop()
		if info == nil {
			log.Printf("Error:\tQueue is empty")
			break
		}
		data, _ := eumgent.Message{Type: eumgent.POST, Info: info.(string)}.ToJSON()
		client.Write(data)
		break
	case eumgent.POST:
		queue.Push(m.Info)
		data, _ := eumgent.Message{Type: eumgent.POST, Info: "OK"}.ToJSON()
		client.Write(data)
		break
	}
	log.Printf("New Message:\t%v\t%v", len(queue.Items), m.Info)
}

func createQueueSaver() {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				b, _ := queue.Serialize()
				err := SaveToFile(b)
				if err != nil {
					log.Printf("Error:\tFile not saved -- %v", err)
					break
				}
				log.Printf("Persistence:\t Save queue to file")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func loadQueue() *Queue {
	b, _ := LoadFromFile()
	queue, _ := DeSerialize(b)
	return queue
}
