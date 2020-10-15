package realtime

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Channel string

const (
	TallyChannel Channel = "tally"
)

type RealtimeEvent struct {
	SenderID string      `json:"sender_id"`
	Data     interface{} `json:"data"`
}

type TallyEvent struct {
	Program bool `json:"program"`
	Preview bool `json:"preview"`
}

func (r *RealtimeServer) Publish(ch Channel, payload RealtimeEvent) error {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return errors.Wrap(err, "fail to marshal event")
	}

	_, err = r.node.Publish(string(ch), payloadBytes)
	if err != nil {
		return errors.Wrap(err, "fail to publish message")
	}
	return nil
}
