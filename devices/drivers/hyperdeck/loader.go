package hyperdeck

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/Scalingo/go-utils/logger"
	"github.com/a-contre-plongee/hyperdeck"
	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/params"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type HyperdeckParams struct {
	IP         string `json:"ip"`
	RepeaterIP string `json:"repeater_ip"`
}

type hyperdeckLoader struct{}

func NewLoader() devices.DeviceLoader {
	return hyperdeckLoader{}
}

func (hyperdeckLoader) Load(ctx context.Context, base *devices.Base, message json.RawMessage) (devices.Device, error) {
	params := HyperdeckParams{}
	err := json.Unmarshal(message, &params)
	if err != nil {
		return nil, errors.Wrap(err, "invalid payload")
	}
	log := logger.Get(ctx).WithFields(logrus.Fields{
		"type":        "hyperdeck",
		"device_name": base.Name(),
	})

	hyperdeck := &Hyperdeck{
		IP:           params.IP,
		RepeaterIP:   params.RepeaterIP,
		log:          log,
		client:       nil,
		clientLock:   &sync.RWMutex{},
		stopping:     false,
		stoppingLock: &sync.Mutex{},

		cacheLock:   &sync.Mutex{},
		clipCache:   make(map[int]hyperdeck.Clip),
		currentClip: 0,
		timecode:    hyperdeck.Timecode{},
		playing:     false,
	}
	hyperdeck.Base = base

	return hyperdeck, nil
}

func (hyperdeckLoader) Validate(message json.RawMessage) error {
	params := HyperdeckParams{}

	err := json.Unmarshal(message, &params)
	if err != nil {
		return err
	}

	if params.IP == "" {
		return fmt.Errorf("ip cannot be empty")
	}
	return nil
}

func (hyperdeckLoader) Params() params.Params {
	return params.Params{
		"ip": params.Input{
			Type:        params.String,
			Description: "Hyperdeck IP",
			Required:    true,
		},
		"repeater_ip": params.Input{
			Type:        params.String,
			Description: "Repeater Listen IP",
			Required:    false,
		},
	}
}
