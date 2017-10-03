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
