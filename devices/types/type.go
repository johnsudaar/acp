package types

import (
	"context"
	"errors"

	handlers "github.com/Scalingo/go-handlers"
)

type DeviceType interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	EventSubscriptions() []string
	WriteEvent(ctx context.Context, toPort string, name string, data interface{})
	Routes() map[string]handlers.HandlerFunc
}

type Type string

const (
	TallyType    Type = "tally"
	PTZType      Type = "ptz"
	SwitcherType Type = "switcher"
)

var (
	ErrInvalidType = errors.New("Invalid device type")
)
