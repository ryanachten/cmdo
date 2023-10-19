package models

type BroadcastMessage struct {
	CommandName string
	MessageBody string
}

type BroadcastChannel = chan BroadcastMessage
