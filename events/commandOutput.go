package events

type CommandOutputType = string

type CommandOutputChannel = chan BroadcastMessage

const commandOutputMessageGroup = "commandOutput"

func CommandOutputInformation(commandName string, messageBody string) BroadcastMessage {
	return BroadcastMessage{
		MessageType:  "information",
		MessageGroup: commandOutputMessageGroup,
		CommandName:  commandName,
		MessageBody:  messageBody,
	}
}

func CommandOutputError(commandName string, messageBody string) BroadcastMessage {
	return BroadcastMessage{
		MessageType:  "information",
		MessageGroup: commandOutputMessageGroup,
		CommandName:  commandName,
		MessageBody:  messageBody,
	}
}
