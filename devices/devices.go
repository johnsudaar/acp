package devices

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/johnsudaar/acp/events"
)

type OutputPort struct {
	Device Device
	Name   string
}

func (p OutputPort) Valid() bool {
	return ValidOutput(p.Name, p.Device)
}

type Device interface {
	WriteEvent(outputName string, event events.Event) error
	Connect(inputName string, port OutputPort) error
	AvailableInputs() []string
	AvailableOutputs() []string
}

var manager *deviceManager
var managerOnce sync.Once

type initializer func(json.RawMessage) (Device, error)

type deviceManager struct {
	devices     map[string]initializer
	devicesLock sync.RWMutex
}

func get() *deviceManager {
	managerOnce.Do(func() {
		manager = &deviceManager{
			devices: make(map[string]initializer),
		}
	})
	return manager
}

func Register(name string, init initializer) {
	m := get()
	m.devicesLock.Lock()
	defer m.devicesLock.Unlock()
	m.devices[name] = init
}

func Initialize(name string, payload json.RawMessage) (Device, error) {
	m := get()
	m.devicesLock.RLock()
	defer m.devicesLock.RUnlock()
	initializer, ok := m.devices[name]
	if !ok {
		return nil, fmt.Errorf("device %s is not registered", name)
	}
	return initializer(payload)
}

func ValidInput(input string, device Device) bool {
	for _, inputName := range device.AvailableInputs() {
		if inputName == input {
			return true
		}
	}
	return false
}

func ValidOutput(output string, device Device) bool {
	for _, outputName := range device.AvailableOutputs() {
		if outputName == output {
			return true
		}
	}
	return false
}
