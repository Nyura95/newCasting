package nyubroadcasting

import "time"

// ExternalCommunication ...
type ExternalCommunication struct {
	Type      string
	Message   string
	CreatedAt time.Time
}
