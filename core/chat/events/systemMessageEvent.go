package events

// SystemMessageEvent is a message displayed in chat on behalf of the server.
type SystemMessageEvent struct {
	Event
	MessageEvent
}

// GetBroadcastPayload will return the object to send to all chat users.
func (e *SystemMessageEvent) GetBroadcastPayload(serverName string) EventPayload {
	return EventPayload{
		"id":        e.ID,
		"timestamp": e.Timestamp,
		"body":      e.Body,
		"type":      SystemMessageSent,
		"user": EventPayload{
			"displayName": serverName,
		},
	}
}

// GetMessageType will return the event type for this message.
func (e *SystemMessageEvent) GetMessageType() EventType {
	return SystemMessageSent
}
