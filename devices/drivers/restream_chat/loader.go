package restream

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/params"
	"github.com/johnsudaar/acp/devices/types/chat"
	"github.com/pkg/errors"
)

type Params struct {
	Token string `json:"token"`
}

type loader struct {
}

func NewLoader() devices.DeviceLoader {
	return loader{}
}

func (loader) Load(ctx context.Context, base *devices.Base, message json.RawMessage) (devices.Device, error) {
	var params Params
	err := json.Unmarshal(message, &params)
	if err != nil {
		return nil, errors.Wrap(err, "invalid params")
	}
	return &Restream{
		Base:             base,
		token:            params.Token,
		subscriptions:    make(map[string]chat.ChatHandler),
		subscriptionLock: &sync.RWMutex{},
		connMutex:        &sync.Mutex{},
		stop:             false,
		stopMutex:        &sync.RWMutex{},
	}, nil
}

func (loader) Validate(message json.RawMessage) error {
	var params Params
	err := json.Unmarshal(message, &params)
	if err != nil {
		return err
	}

	if params.Token == "" {
		return fmt.Errorf("token cannot be empty")
	}

	return nil
}

func (loader) Params() params.Params {
	return params.Params{
		"token": params.Input{
			Type:        params.String,
			Description: "Restream chat token",
			Required:    true,
		},
	}
}
