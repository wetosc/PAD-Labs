//Package eugddc stands for "EUGenius Distributed Data Collections"
// and contains the shared code of the application
package eugddc

import (
	"encoding/json"
	"encoding/xml"

	"github.com/rs/zerolog/log"
)

const MediatorAddr = ":9876"
const ClientAddr = ":9000"
const NodeBasePort = 15100

//CheckError checks if error is not nil and if so prints the description in logs on level INFO
func CheckError(err error, info string) {
	if err != nil {
		log.Info().Msgf("%v : %v", info, err)
	}
}

type NodeMessage struct {
	Type  string
	Trace []string
	Query NodeQuery
	Data  []Dog
}

type NodeQuery struct {
	ID     string
	Query  string
	Params *QueryParams
}

type QueryParams struct {
	// ACTION = FILTER
	// PARAM = NAME, AGE
	// OPERATION =  < > =
	// VALUE = string or int
	Action    string
	Param     string
	Operation string
	Value     string
}

func NodeMessageFromJSON(data []byte) (NodeMessage, error) {
	var a NodeMessage
	err := json.Unmarshal(data, &a)
	return a, err
}

func NodeMessageFromXML(data []byte) (NodeMessage, error) {
	var a NodeMessage
	err := xml.Unmarshal(data, &a)
	return a, err
}

func (m NodeMessage) ToJSON() []byte {
	data, err := json.Marshal(m)
	CheckError(err, "[NodeMessage] Error converting to JSON")
	return data
}

func (m NodeMessage) ToXML() []byte {
	data, err := xml.Marshal(m)
	CheckError(err, "[NodeMessage] Error converting to XML")
	return data
}
