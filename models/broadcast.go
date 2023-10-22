package models

type MessageType = string

const (
	InformationMessage MessageType = "information"
	ErrorMessage       MessageType = "error"
)

type BroadcastMessage struct {
	CommandName string      `json:"commandName"`
	MessageType MessageType `json:"messageType"`
	MessageBody string      `json:"messageBody"`
}

type BroadcastChannel = chan BroadcastMessage
