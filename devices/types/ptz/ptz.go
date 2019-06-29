package ptz

import (
	"context"

	handlers "github.com/Scalingo/go-handlers"
	"github.com/johnsudaar/acp/devices/types"
)

var _ types.DeviceType = &PtzDriver{}

type PTZJoystickParams struct {
	PanSpeed  float64 `json:"pan"`
	TiltSpeed float64 `json:"tilt"`
	ZoomSpeed float64 `json:"zoom"`
}

type Ptzable interface {
	SendPTZJoystick(params PTZJoystickParams) error
}

type PtzDriver struct {
	device Ptzable
}

func NewPtzDriver(device Ptzable) *PtzDriver {
	return &PtzDriver{
		device: device,
	}
}

func (p *PtzDriver) Start(ctx context.Context) error {
	return nil
}

func (p *PtzDriver) Stop(ctx context.Context) error {
	return nil
}

func (p *PtzDriver) EventSubscriptions() []string {
	return []string{}
}

func (p *PtzDriver) WriteEvent(ctx context.Context, toPort string, name string, data interface{}) {
}

func (p *PtzDriver) Routes() map[string]handlers.HandlerFunc {
	return map[string]handlers.HandlerFunc{
		"/ptz/joystick": p.Joystick,
	}
}
