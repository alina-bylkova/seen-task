package dto

import (
	"errors"
	"log"
)

const (
	LpHits     = "LpHits"
	VideoPlays = "VideoPlays"
)

// Event is a data transfer object used in REST
type Event struct {
	RecipientID int64  `json:"recipient_id"`
	VideoID     int64  `json:"video_id"`
	EventType   string `json:"event_type"`
}

// ValidateEventType is a method used to validate event type
func (e *Event) ValidateEventType() error {
	if e.EventType == LpHits || e.EventType == VideoPlays {
		return nil
	}
	log.Printf("Invalid event type is provided: %s", e.EventType)
	return errors.New("Invalid event type is provided")
}
