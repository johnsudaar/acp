package client

import (
	"fmt"

	"github.com/pkg/errors"
)

type PanDirection string
type TiltDirection string
type ZoomDirection string

const (
	PanDirectionStop  PanDirection = "Stop"
	PanDirectionLeft  PanDirection = "Left"
	PanDirectionRight PanDirection = "Right"

	TiltDirectionStop TiltDirection = "Stop"
	TiltDirectionUp   TiltDirection = "Up"
	TiltDirectionDown TiltDirection = "Down"

	ZoomDirectionStop ZoomDirection = "Stop"
	ZoomDirectionWide ZoomDirection = "Wide"
	ZoomDirectionTele ZoomDirection = "Tele"

	MinJoystickSpeed = 0
	MaxJoystickSpeed = 30

	MinZoomSpeed = 0
	MaxZoomSpeed = 8
)

type JoyStickOperationParams struct {
	PanDirection  PanDirection  `json:"PanDirection"`
	PanSpeed      int           `json:"PanSpeed"`
	TiltDirection TiltDirection `json:"TiltDirection"`
	TiltSpeed     int           `json:"TiltSpeed"`
}

type ZoomParams struct {
	Direction ZoomDirection `json:"Direction"`
	Speed     int           `json:"Speed"`
}

func (c HTTPClient) JoyStickOperation(params JoyStickOperationParams) error {
	if params.PanSpeed < MinJoystickSpeed || params.PanSpeed > MaxJoystickSpeed {
		return fmt.Errorf("Invalid pan speed %v, should be between %v and %v", params.PanSpeed, MinJoystickSpeed, MaxJoystickSpeed)
	}

	if params.TiltSpeed < MinJoystickSpeed || params.TiltSpeed > MaxJoystickSpeed {
		return fmt.Errorf("Invalid tilt speed %v, should be between %v and %v", params.TiltSpeed, MinJoystickSpeed, MaxJoystickSpeed)
	}

	_, err := c.makeRequest("JoyStickOperation", params)
	if err != nil {
		return errors.Wrap(err, "fail to call HTTP API")
	}
	return nil
}

func (c HTTPClient) Zoom(params ZoomParams) error {
	if params.Speed < MinZoomSpeed || params.Speed > MaxZoomSpeed {
		return fmt.Errorf("Invalid zoom speed %v, should be between %v and %v", params.Speed, MinZoomSpeed, MaxZoomSpeed)
	}
	_, err := c.makeRequest("ZoomSwitchOperation", params)
	if err != nil {
		return errors.Wrap(err, "fail to call HTTP API")
	}
	return nil
}
