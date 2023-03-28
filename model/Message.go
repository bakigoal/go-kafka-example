package model

import "encoding/json"

type Predecessor struct {
	MdmCode string `json:"mdmCode,omitempty"`
	Id      string `json:"id,omitempty"`
}

type Message struct {
	MdmCode      string        `json:"mdmCode,omitempty"`
	Predecessors []Predecessor `json:"predecessors"`
}

func (m Message) ToByteArray() []byte {
	b, err := json.Marshal(m)
	if err != nil {
		panic(any(err))
	}
	return b
}

func FromByteArray(b []byte) Message {
	message := Message{}
	err := json.Unmarshal(b, &message)
	if err != nil {
		panic(any(err))
	}
	return message
}
