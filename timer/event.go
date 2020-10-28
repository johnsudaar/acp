package timer

import "github.com/johnsudaar/acp/models"

const RealtimeUpdateChannel string = "timer"

type RealtimeUpdateEvent struct {
	Value     string                `json:"value"`
	Expired   bool                  `json:"expired"`
	Type      models.TimerType      `json:"type"`
	Direction models.TimerDirection `json:"direction"`
}
