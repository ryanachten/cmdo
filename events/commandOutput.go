package events

type CommandOutputType = string

type CommandOutputChannel = chan BroadcastMessage

func CommandOutputInformation(commandName string, messageBody string) BroadcastMessage {
	return BroadcastMessage{
		MessageType: "information",
		CommandName: commandName,
		MessageBody: messageBody,
	}
}

func CommandOutputError(commandName string, messageBody string) BroadcastMessage {
	return BroadcastMessage{
		MessageType: "information",
		CommandName: commandName,
		MessageBody: messageBody,
	}
}
