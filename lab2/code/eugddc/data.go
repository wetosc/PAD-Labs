package eugddc

import "encoding/json"
import "io/ioutil"

// Dog is the representation of data in the files
type Dog struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// DogFromJSON converts []byte to Animal struct
func DogFromJSON(data []byte) (Dog, error) {
	var a Dog
	err := json.Unmarshal(data, &a)
	return a, err
}

// DogsFromJSON converts []byte to []Animal
func DogsFromJSON(data []byte) ([]Dog, error) {
	var a []Dog
	err := json.Unmarshal(data, &a)
	return a, err
}

// Load loads Animals from 'fileName' file
func Load(fileName string) ([]Dog, error) {
	data, _ := ioutil.ReadFile(fileName)
	return DogsFromJSON(data)
}

// JSONfromDogs converts []Dog to []byte
func JSONfromDogs(items []Dog) []byte {
	data, _ := json.Marshal(items)
	return data
}
