package remote

import (
	"context"
	"encoding/json"

	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/params"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

type Params struct {
	RemoteID bson.ObjectId `json:"remote_id"`
}

type loader struct {
}

func NewLoader() devices.DeviceLoader {
	return loader{}
}

func (l loader) Load(ctx context.Context, base *devices.Base, message json.RawMessage) (devices.Device, error) {
	params := Params{}
	err := json.Unmarshal(message, &params)
	if err != nil {
		return nil, errors.Wrap(err, "invalid payload")
	}

	var remote Remote
	remote.Base = base
	remote.RemoteID = params.RemoteID
	return &remote, nil
}

func (l loader) Validate(json.RawMessage) error {
	return nil
}

func (l loader) Params() params.Params {
	return params.Params{}
}
