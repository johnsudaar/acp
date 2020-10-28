package types

import (
	"context"
	"encoding/json"
	"errors"

	handlers "github.com/Scalingo/go-handlers"
)

type DeviceType interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	EventSubscriptions() []string
	RealtimeEventSubscriptions() []string
	WriteEvent(ctx context.Context, toPort string, name string, data interface{})
	WriteRealtimeEvent(ctx context.Context, channel string, payload json.RawMessage)
	Routes() map[string]handlers.HandlerFunc
}

type Type string

const (
	TallyType    Type = "tally"
	PTZType      Type = "ptz"
	SwitcherType Type = "switcher"
	TimerType    Type = "timer"
)

var (
	ErrInvalidType = errors.New("Invalid device type")
)
