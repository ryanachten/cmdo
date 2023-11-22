package events

type EventBus struct {
	CommandOutput  CommandOutputChannel
	CommandRequest CommandRequestChannel
}

// Creates event bus with initialised event channels
func CreateEventBus() EventBus {
	return EventBus{
		CommandOutput:  make(CommandOutputChannel),
		CommandRequest: make(CommandRequestChannel),
	}
}

// Closed all event channels
func (bus EventBus) Close() {
	close(bus.CommandOutput)
	close(bus.CommandRequest)
}
