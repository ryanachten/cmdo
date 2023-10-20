package models

type BroadcastMessage struct {
	CommandName string `json:"commandName"`
	MessageBody string `json:"messageBody"`
}

type BroadcastChannel = chan BroadcastMessage
