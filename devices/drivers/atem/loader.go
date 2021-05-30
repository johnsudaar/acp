package atem

import (
	"context"
	"encoding/json"
	"strconv"
	"sync"

	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/params"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	AtemType1ME             = "atem_1_me"
	AtemTypeConstellationHD = "atem_constellation_hd"
)

type AtemParams struct {
	IP   string `json:"ip"`   // ATEM IP
	Port int    `json:"port"` // ATEM Port
	Type string `json:"type"`
}

type atemLoader struct{}

func NewLoader() devices.DeviceLoader {
	return atemLoader{}
}

func (atemLoader) Load(ctx context.Context, base *devices.Base, message json.RawMessage) (devices.Device, error) {
	params := AtemParams{}
	err := json.Unmarshal(message, &params)
	if err != nil {
		return nil, errors.Wrap(err, "invalid payload")
	}
	var atem ATEM
	atem.IP = params.IP
	atem.Port = strconv.Itoa(params.Port)
	atem.Base = base
	atem.stoppingLock = &sync.Mutex{}
	atem.atemType = params.Type
	if atem.atemType == "" {
		atem.atemType = AtemType1ME
	}

	atem.log = logger.Get(ctx).WithFields(logrus.Fields{
		"device":    atem.Name(),
		"device_id": atem.ID(),
	})
	return &atem, nil
}

func (atemLoader) Validate(message json.RawMessage) error {
	params := AtemParams{}
	err := json.Unmarshal(message, &params)
	if err != nil {
		return err
	}

	return nil
}

func (atemLoader) Params() params.Params {
	return params.Params{
		"ip": params.Input{
			Type:        params.IP,
			Description: "ATEM IP",
			Required:    true,
		},
		"port": params.Input{
			Type:        params.Number,
			Description: "ATEM Port",
			Required:    true,
			Default:     9910,
			Min:         1,
			Max:         65535,
		},
		"type": params.Input{
			Type:        params.Select,
			Description: "Atem Model",
			Required:    false,
			Default:     "atem_1_me",
			Options: []params.SelectOption{
				{
					Value: AtemType1ME,
					Name:  "Atem 1 M/E",
				}, {
					Value: AtemTypeConstellationHD,
					Name:  "Atem Constellation (HD Mode)",
				},
			},
		},
	}
}
