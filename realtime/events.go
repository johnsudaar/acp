package realtime

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type RealtimeEvent struct {
	SenderID string      `json:"sender_id"`
	Data     interface{} `json:"data"`
}

type UserEvent struct {
	DeviceID string          `json:"device_id"`
	Data     json.RawMessage `json:"data"`
}

func (r *RealtimeServer) Publish(ch string, from string, data interface{}) error {
	if r.node == nil {
		return nil
	}
	event := RealtimeEvent{
		SenderID: from,
		Data:     data,
	}
	payloadBytes, err := json.Marshal(event)
	if err != nil {
		return errors.Wrap(err, "fail to marshal event")
	}

	_, err = r.node.Publish(string(ch), payloadBytes)
	if err != nil {
		return errors.Wrap(err, "fail to publish message")
	}
	return nil
}
