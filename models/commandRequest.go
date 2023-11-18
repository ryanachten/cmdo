package models

type CommandRequestState = string

const (
	CommandRequestStop  CommandRequestState = "stop"
	CommandRequestStart CommandRequestState = "start"
)

type CommandRequest struct {
	CommandName    string              `json:"commandName"`
	RequestedState CommandRequestState `json:"requestedState"`
}

type CommandRequestChannel = chan CommandRequest
