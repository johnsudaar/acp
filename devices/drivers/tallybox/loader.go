package tallybox

import (
	"context"
	"encoding/json"

	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/params"
)

func NewLoader() devices.DeviceLoader {
	return tallyBoxLoader{}
}

type tallyBoxLoader struct{}

func (t tallyBoxLoader) Load(ctx context.Context, base *devices.Base, message json.RawMessage) (devices.Device, error) {
	var tallybox tallybox
	tallybox.Base = base
	return &tallybox, nil
}

func (t tallyBoxLoader) Validate(json.RawMessage) error {
	return nil
}

func (t tallyBoxLoader) Params() params.Params {
	return params.Params{}
}
