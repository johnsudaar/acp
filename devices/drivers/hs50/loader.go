package hs50

import (
	"context"
	"encoding/json"

	"github.com/Scalingo/go-utils/logger"
	"github.com/a-contre-plongee/hs50/client"
	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/params"
	"github.com/pkg/errors"
)

type SwitcherParams struct {
	IP string `json:"ip"`
}

type hs50Loader struct{}

func NewLoader() devices.DeviceLoader {
	return hs50Loader{}
}

func (hs50Loader) Load(ctx context.Context, base *devices.Base, message json.RawMessage) (devices.Device, error) {
	params := SwitcherParams{}
	err := json.Unmarshal(message, &params)
	if err != nil {
		return nil, errors.Wrap(err, "invalid payload")
	}
	var switcher Switcher
	switcher.IP = params.IP
	switcher.Base = base
	switcher.log = logger.Get(ctx)
	switcher.client = client.New(switcher.IP)

	return &switcher, nil
}

func (hs50Loader) Validate(message json.RawMessage) error {
	params := SwitcherParams{}

	err := json.Unmarshal(message, &params)
	if err != nil {
		return err
	}
	return nil
}

func (hs50Loader) Params() params.Params {
	return params.Params{
		"ip": params.Input{
			Type:        params.String,
			Description: "Switcher IP",
			Required:    true,
		},
	}
}
