package events

type CommandRequestType = string

const (
	CommandRequestStop  CommandRequestType = "stop"
	CommandRequestStart CommandRequestType = "start"
)

type CommandRequest struct {
	CommandName    string             `json:"commandName"`
	RequestedState CommandRequestType `json:"requestedState"`
}

type CommandRequestChannel = chan CommandRequest
