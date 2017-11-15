//Package eugddc stands for "EUGenius Distributed Data Collections"
// and contains the shared code of the application
package eugddc

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

const MediatorAddr = ":9876"
const ClientAddr = ":9000"
const NodeBasePort = 9000

//CheckError checks if error is not nil and if so prints the description in logs on level INFO
func CheckError(err error, info string) {
	if err != nil {
		log.Info().Msgf("%v : %v", info, err)
	}
}

type Message struct {
	Query string
	Data  []Dog
}

func NewMessage(query string) Message {
	return Message{Query: query}
}

func NewMessage_Data(query string, dogs []Dog) Message {
	return Message{Query: query, Data: dogs}
}

func MessageFromJSON(data []byte) (Message, error) {
	var a Message
	err := json.Unmarshal(data, &a)
	return a, err
}

func MessageToJSON(a Message) []byte {
	data, err := json.Marshal(a)
	CheckError(err, "[EUGDDC] Error converting data to []byte")
	return data
}
