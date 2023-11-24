package events

type EventBus struct {
	CommandOutput  CommandOutputChannel
	CommandRequest CommandRequestChannel
	CommandState   CommandStateChannel
}

// Creates event bus with initialised event channels
func CreateEventBus() EventBus {
	return EventBus{
		CommandOutput:  make(CommandOutputChannel),
		CommandRequest: make(CommandRequestChannel),
		CommandState:   make(CommandStateChannel),
	}
}

// Closed all event channels
func (bus EventBus) Close() {
	close(bus.CommandOutput)
	close(bus.CommandRequest)
	close(bus.CommandState)
}
