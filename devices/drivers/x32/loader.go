package x32

import (
	"context"
	"encoding/json"

	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/params"
	"github.com/pkg/errors"
)

type X32Params struct {
	IP string `json:"ip"`
}

type x32Loader struct {
}

func NewLoader() devices.DeviceLoader {
	return x32Loader{}
}

func (x32Loader) Load(ctx context.Context, base *devices.Base, message json.RawMessage) (devices.Device, error) {
	params := X32Params{}

	err := json.Unmarshal(message, &params)
	if err != nil {
		return nil, errors.Wrap(err, "invalid payload")
	}

	var x32 X32
	x32.IP = params.IP

	x32.Base = base
	x32.log = logger.Get(ctx)
	return &x32, nil
}

func (x32Loader) Validate(message json.RawMessage) error {
	params := X32Params{}
	err := json.Unmarshal(message, &params)
	if err != nil {
		return err
	}
	return nil
}

func (x32Loader) Params() params.Params {
	return params.Params{
		"ip": params.Input{
			Type:        params.IP,
			Description: "X32 IP",
			Required:    true,
		},
	}
}
