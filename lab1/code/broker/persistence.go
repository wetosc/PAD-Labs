package main

import (
	"io/ioutil"
)

const filePath = "dat"

// SaveToFile saves b []byte to the standard file
func SaveToFile(b []byte) error {
	err := ioutil.WriteFile(filePath, b, 0644)
	return err
}

// LoadFromFile reads the standard file and return []byte
func LoadFromFile() ([]byte, error) {
	b, err := ioutil.ReadFile(filePath)
	return b, err
}


// func createQueueSaver() {
// 	ticker := time.NewTicker(5 * time.Second)
// 	quit := make(chan struct{})
// 	go func() {
// 		for {
// 			select {
// 			case <-ticker.C:
// 				b, _ := queue.Serialize()
// 				err := SaveToFile(b)
// 				if err != nil {
// 					log.Printf("Error:\tFile not saved -- %v", err)
// 					break
// 				}
// 				log.Printf("Persistence:\t Save queue to file")
// 			case <-quit:
// 				ticker.Stop()
// 				return
// 			}
// 		}
// 	}()
// }

// func loadQueue() *Queue {
// 	b, _ := LoadFromFile()
// 	queue, _ := DeSerialize(b)
// 	return queue
// }