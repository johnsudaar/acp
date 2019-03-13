package devices

import (
	"context"

	"github.com/johnsudaar/acp/models"
	"gopkg.in/mgo.v2/bson"
)

type State string

const (
	StateNotConnected State = "Not connected"
	StateConnected    State = "Connected"
)

type EventWriter interface {
	SendEvent(ctx context.Context, from models.Port, name string, data interface{})
}

type Device interface {
	ID() bson.ObjectId
	Name() string
	Type() string
	State() State
	InputPorts() []string
	OutputPorts() []string

	Start() error
	Stop() error

	WriteEvent(ctx context.Context, toPort string, name string, data interface{})
}
