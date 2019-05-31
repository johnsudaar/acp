package devices

import (
	"context"
	"encoding/json"
	"errors"
	"sync"

	"github.com/johnsudaar/acp/devices/params"
)

var (
	ErrTypeNotFound = errors.New("Type not found")
)

type DeviceLoader interface {
	Load(context.Context, *Base, json.RawMessage) (Device, error)
	Validate(json.RawMessage) error
	Params() params.Params
}

var (
	deviceTypesLock *sync.RWMutex           = &sync.RWMutex{}
	deviceTypesMap  map[string]DeviceLoader = make(map[string]DeviceLoader)
)

func RegisterType(name string, loader DeviceLoader) {
	deviceTypesLock.Lock()
	defer deviceTypesLock.Unlock()

	deviceTypesMap[name] = loader
}

func AvailableTypes() []string {
	deviceTypesLock.RLock()
	defer deviceTypesLock.RUnlock()
	res := []string{}
	for typeName, _ := range deviceTypesMap {
		res = append(res, typeName)
	}

	return res
}

func GetLoader(name string) (DeviceLoader, error) {
	deviceTypesLock.RLock()
	defer deviceTypesLock.RUnlock()

	loader, ok := deviceTypesMap[name]
	if !ok {
		return nil, ErrTypeNotFound
	}

	return WrapLoader(loader), nil
}
