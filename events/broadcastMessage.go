package events

// Base type for messages broadcasted using WebSockets
type BroadcastMessage struct {
	CommandCategory string `json:"commandCategory"`
	CommandName     string `json:"commandName"`
	MessageType     string `json:"messageType"`
	MessageBody     string `json:"messageBody"`
}
