package remote

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/types"
	"gopkg.in/mgo.v2/bson"
)

type Remote struct {
	*devices.Base
	RemoteID bson.ObjectId
}

func (r *Remote) InputPorts() []string {
	return []string{}
}

func (r *Remote) OutputPorts() []string {
	return []string{}
}

func (r *Remote) API() http.Handler {
	return http.NotFoundHandler()
}

func (r *Remote) Start() error {
	return nil
}

func (r *Remote) Stop() error {
	return nil
}

func (r *Remote) WriteEvent(ctx context.Context, toPort string, name string, data interface{}) {
}

func (s *Remote) WriteRealtimeEvent(ctx context.Context, channel string, payload json.RawMessage) {
}

func (r *Remote) Types() []types.Type {
	return []types.Type{}
}
