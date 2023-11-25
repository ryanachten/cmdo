package events

type CommandStateChannel = chan BroadcastMessage

type CommandStateType = string

const (
	CommandStateStarted CommandStateType = "started"
	CommandStateStopped CommandStateType = "stopped"
	CommandStateFailed  CommandStateType = "failed"
)

const commandStateMessageGroup = "commandState"

func CommandStateStart(commandName string) BroadcastMessage {
	return BroadcastMessage{
		CommandName:  commandName,
		MessageGroup: commandStateMessageGroup,
		MessageType:  CommandStateStarted,
	}
}

func CommandStateStop(commandName string) BroadcastMessage {
	return BroadcastMessage{
		CommandName:  commandName,
		MessageGroup: commandStateMessageGroup,
		MessageType:  CommandStateStopped,
	}
}

func CommandStateFail(commandName string, failureReason string) BroadcastMessage {
	return BroadcastMessage{
		CommandName:  commandName,
		MessageGroup: commandStateMessageGroup,
		MessageType:  CommandStateFailed,
		MessageBody:  failureReason,
	}
}
