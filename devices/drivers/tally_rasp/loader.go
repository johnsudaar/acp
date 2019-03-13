package tally

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/devices"
	"github.com/pkg/errors"
)

type TallyParams struct {
	IP string `json:"ip"`
}

type raspTallyLoader struct{}

func NewLoader() devices.DeviceLoader {
	return raspTallyLoader{}
}

func (raspTallyLoader) Load(ctx context.Context, base *devices.Base, message json.RawMessage) (devices.Device, error) {
	params := TallyParams{}
	err := json.Unmarshal(message, &params)
	if err != nil {
		return nil, errors.Wrap(err, "invalid payload")
	}
	var rasp Tally
	rasp.IP = params.IP
	rasp.Base = base
	rasp.log = logger.Get(ctx)
	rasp.tallyRefreshChan = make(chan bool, 1)
	rasp.tallySync = &sync.RWMutex{}
	rasp.stoppingLock = &sync.Mutex{}
	return &rasp, nil
}

func (raspTallyLoader) Validate(message json.RawMessage) error {
	params := TallyParams{}
	err := json.Unmarshal(message, &params)
	if err != nil {
		return err
	}

	return nil
}
