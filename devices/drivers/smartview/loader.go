package smartview

import (
	"context"
	"encoding/json"

	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/params"
	"github.com/pkg/errors"
)

type SmartViewDuoParams struct {
	IP          string `json:"ip"`
	ScreenCount int    `json:"screen_count"`
}

type smartviewDuo struct{}

func NewLoader() devices.DeviceLoader {
	return smartviewDuo{}
}

func (smartviewDuo) Load(ctx context.Context, base *devices.Base, message json.RawMessage) (devices.Device, error) {
	params := SmartViewDuoParams{}
	err := json.Unmarshal(message, &params)
	if err != nil {
		return nil, errors.Wrap(err, "invalid payload")
	}
	var smartView SmartView
	smartView.IP = params.IP
	smartView.Base = base
	smartView.log = logger.Get(ctx)

	var outputs []string

	for i := 0; i < params.ScreenCount; i++ {
		outputs = append(outputs, "MONITOR "+string(int('A')+i))
	}

	smartView.Outputs = outputs

	return &smartView, nil
}

func (smartviewDuo) Validate(message json.RawMessage) error {
	params := SmartViewDuoParams{}
	err := json.Unmarshal(message, &params)
	if err != nil {
		return err
	}

	return nil
}

func (smartviewDuo) Params() params.Params {
	return params.Params{
		"ip": params.Input{
			Type:        params.IP,
			Description: "Screen IP",
			Required:    true,
		},
		"screen_count": params.Input{
			Type:        params.Number,
			Description: "Number of screens",
			Required:    true,
			Min:         1,
			Max:         2,
		},
	}
}
