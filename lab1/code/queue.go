package main

import (
	"encoding/json"
	"sync"
)

// Queue data type
type Queue struct {
	sync.Mutex
	Items []interface{} `json:"items"`
}

// NewQueue : Create Queue with capacity
func NewQueue(capacity int) *Queue {
	return &Queue{
		Items: make([]interface{}, 0, capacity),
	}
}

// Push one item to the Queue
func (q *Queue) Push(item interface{}) {
	q.Lock()
	defer q.Unlock()
	q.Items = append(q.Items, item)
}

// Pop one item from the Queue
func (q *Queue) Pop() interface{} {
	q.Lock()
	defer q.Unlock()
	if len(q.Items) == 0 {
		return nil
	}
	item := q.Items[0]
	q.Items[0] = nil
	q.Items = q.Items[1:]
	return item
}

// Serialize converts the queue in byte array
func (q *Queue) Serialize() ([]byte, error) {
	bytes, err := json.Marshal(q)
	return bytes, err
}

// DeSerialize creates the queue from byte array
func DeSerialize(b []byte) (*Queue, error) {
	var q = &Queue{}
	err := json.Unmarshal(b, q)
	return q, err
}
