package tallyrecorder

import (
	"context"
	"encoding/json"

	"github.com/johnsudaar/acp/devices"
	"github.com/pkg/errors"
)

type RecorderParams struct {
	Shoot string `json:"shoot"`
}

type recorderLoader struct{}

func NewLoader() devices.DeviceLoader {
	return recorderLoader{}
}

func (recorderLoader) Load(ctx context.Context, base *devices.Base, message json.RawMessage) (devices.Device, error) {
	params := RecorderParams{}
	err := json.Unmarshal(message, &params)
	if err != nil {
		return nil, errors.Wrap(err, "invalid payload")
	}

	var recorder Recorder
	recorder.Base = base
	recorder.Shoot = params.Shoot

	return &recorder, nil
}

func (recorderLoader) Validate(message json.RawMessage) error {
	params := RecorderParams{}
	err := json.Unmarshal(message, &params)
	if err != nil {
		return err
	}
	return nil
}
