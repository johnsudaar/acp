package types

import (
	"context"
	"errors"
)

type DeviceType interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	EventSubscriptions() []string
	WriteEvent(ctx context.Context, toPort string, name string, data interface{})
}

type Type string

const (
	TallyType Type = "tally"
)

var (
	ErrInvalidType = errors.New("Invalid device type")
)
