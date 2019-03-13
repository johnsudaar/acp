package atem

import (
	"context"
	"encoding/json"
	"strconv"
	"sync"

	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/devices"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type AtemParams struct {
	IP     string `json:"ip"`      // ATEM IP
	Port   int    `json:"port"`    // ATEM Port
	MyPort int    `json:"my_port"` // Local port used for the UDP connection to the ATEM
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
	atem.MyPort = strconv.Itoa(params.MyPort)
	atem.Base = base
	atem.stoppingLock = &sync.Mutex{}

	atem.log = logger.Get(ctx).WithFields(logrus.Fields{
		"device":    atem.Name,
		"device_id": atem.ID,
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
