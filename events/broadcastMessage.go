package events

// Base type for messages broadcasted using WebSockets
type BroadcastMessage struct {
	CommandName  string `json:"commandName"`
	MessageGroup string `json:"messageGroup"`
	MessageType  string `json:"messageType"`
	MessageBody  string `json:"messageBody"`
}
