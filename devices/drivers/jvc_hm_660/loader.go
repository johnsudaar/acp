package jvc

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/devices"
	"github.com/pkg/errors"
)

type CamParams struct {
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type jvcHm660Loader struct{}

func NewLoader() devices.DeviceLoader {
	return jvcHm660Loader{}
}

func (jvcHm660Loader) Load(ctx context.Context, base *devices.Base, message json.RawMessage) (devices.Device, error) {
	params := CamParams{}
	err := json.Unmarshal(message, &params)
	if err != nil {
		return nil, errors.Wrap(err, "invalid payload")
	}
	var cam JVCHM660
	cam.IP = params.IP
	cam.Port = params.Port
	cam.User = params.User
	cam.Password = params.Password
	cam.Base = base
	cam.log = logger.Get(ctx)
	cam.stoppingLock = &sync.Mutex{}
	cam.tallySync = &sync.RWMutex{}
	cam.tallyRefreshChan = make(chan bool, 1)
	return &cam, nil
}

func (jvcHm660Loader) Validate(message json.RawMessage) error {
	params := CamParams{}
	err := json.Unmarshal(message, &params)
	if err != nil {
		return err
	}

	return nil
}
