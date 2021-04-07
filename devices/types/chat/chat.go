package chat

import (
	"context"
	"encoding/json"

	"github.com/Scalingo/go-handlers"
	"github.com/johnsudaar/acp/devices/types"
)

type ChatMessage struct {
	ID        string `json:"id"`
	Timestamp int64  `json:"timestamp,omitempty"`
	From      string `json:"from,omitempty"`
	Channel   string `json:"channel,omitempty"`
	Message   string `json:"message,omitempty"`
}

type ChatHandler func(ctx context.Context, message ChatMessage)

type Chattable interface {
	ChatSubscribe(ctx context.Context, c ChatHandler) string
	ChatUnsubscribe(ctx context.Context, id string)
}

var _ types.DeviceType = &ChatDriver{}

type ChatDriver struct {
	device Chattable
}

func NewChatDriver(device Chattable) *ChatDriver {
	return &ChatDriver{
		device: device,
	}
}

func (d *ChatDriver) Start(ctx context.Context) error {
	return nil
}

func (d *ChatDriver) Stop(ctx context.Context) error {
	return nil
}

func (d *ChatDriver) EventSubscriptions() []string {
	return []string{}
}

func (d *ChatDriver) RealtimeEventSubscriptions() []string {
	return []string{}
}

func (d *ChatDriver) WriteEvent(ctx context.Context, toPort, name string, data interface{}) {
}

func (d *ChatDriver) WriteRealtimeEvent(ctx context.Context, channel string, payload json.RawMessage) {
}

func (d *ChatDriver) Routes() map[string]handlers.HandlerFunc {
	return map[string]handlers.HandlerFunc{}
}
