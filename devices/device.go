package devices

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/johnsudaar/acp/devices/types"
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

type RealtimeEventWriter interface {
	Publish(ch string, from string, payload interface{}) error
}

type Device interface {
	ID() bson.ObjectId
	Name() string
	Type() string
	State() State
	InputPorts() []string
	OutputPorts() []string
	API() http.Handler

	Start() error
	Stop() error

	WriteEvent(ctx context.Context, toPort string, name string, data interface{})
	WriteRealtimeEvent(ctx context.Context, channel string, payload json.RawMessage)
	Types() []types.Type
}
