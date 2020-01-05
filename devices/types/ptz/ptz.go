package ptz

import (
	"context"

	handlers "github.com/Scalingo/go-handlers"
	"github.com/johnsudaar/acp/devices/types"
	"gopkg.in/mgo.v2/bson"
)

var _ types.DeviceType = &PtzDriver{}

type PTZJoystickParams struct {
	PanSpeed   float64    `json:"pan"`
	TiltSpeed  float64    `json:"tilt"`
	ZoomSpeed  float64    `json:"zoom"`
	FocusSpeed float64    `json:"focus"`
	Buttons    PTZButtons `json:"buttons"`
}

type PTZButtons struct {
	FocusPushAuto bool `json:"focus_push_auto"`
	IrisOpen      bool `json:"iris_open"`
	IrisClose     bool `json:"iris_close"`
	Button1       bool `json:"button_1"`
	Button2       bool `json:"button_2"`
	Button3       bool `json:"button_3"`
	Button4       bool `json:"button_4"`
}

type PTZPositionParams struct {
	Pan   float64 `json:"pan"`
	Tilt  float64 `json:"tilt"`
	Zoom  float64 `json:"zoom"`
	Focus float64 `json:"focus"`
}

type Ptzable interface {
	ID() bson.ObjectId
	SendPTZJoystick(params PTZJoystickParams) error
	SetPosition(params PTZPositionParams) error
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
		"/ptz/joystick":         p.Joystick,
		"/ptz/position":         p.Position,
		"/ptz/store":            p.Store,
		"/ptz/store/{position}": p.SinglePosition,
	}
}
