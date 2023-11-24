package events

type CommandStateChannel = chan BroadcastMessage

const commandStateMessageGroup = "commandState"

func CommandStateStart(commandName string) BroadcastMessage {
	return BroadcastMessage{
		CommandName:  commandName,
		MessageGroup: commandStateMessageGroup,
		MessageType:  "started",
	}
}

func CommandStateStop(commandName string) BroadcastMessage {
	return BroadcastMessage{
		CommandName:  commandName,
		MessageGroup: commandStateMessageGroup,
		MessageType:  "stopped",
	}
}

func CommandStateFail(commandName string, failureReason string) BroadcastMessage {
	return BroadcastMessage{
		CommandName:  commandName,
		MessageGroup: commandStateMessageGroup,
		MessageType:  "failed",
		MessageBody:  failureReason,
	}
}
