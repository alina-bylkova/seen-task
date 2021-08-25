package dto

import (
	"errors"
	"fmt"
	"testing"
)

func TestValidateEventType(t *testing.T) {
	testCases := []struct {
		eventType      *Event
		expectedResult error
		caseName       string
	}{
		{
			eventType: &Event{
				RecipientID: 1,
				VideoID:     1,
				EventType:   "LinkClicked",
			},
			expectedResult: errors.New("Invalid event type is provided"),
			caseName:       "Invalid event type provided",
		},
		{
			eventType: &Event{
				RecipientID: 1,
				VideoID:     1,
				EventType:   "LpHits",
			},
			expectedResult: nil,
			caseName:       "Valid event: LpHits",
		},
		{
			eventType: &Event{
				RecipientID: 1,
				VideoID:     1,
				EventType:   "VideoPlays",
			},
			expectedResult: nil,
			caseName:       "Valid event: VideoPlays",
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("testing %v", tc.caseName), func(t *testing.T) {
			result := tc.eventType.ValidateEventType()
			if result == nil {
				if tc.expectedResult != nil {
					t.Errorf("Expected %v, got %v", tc.expectedResult, result)
				}
			} else if result.Error() != tc.expectedResult.Error() {
				t.Errorf("Expected %v, got %v", tc.expectedResult, result)
			}
		})
	}
}
