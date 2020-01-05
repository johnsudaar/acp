package client

import (
	"fmt"

	"github.com/pkg/errors"
)

type PanDirection string
type TiltDirection string
type ZoomDirection string
type PanCtrlDirection string
type TiltCtrlDirection string

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

	PanCtrlDirectionStop     PanCtrlDirection = "Stop"
	PanCtrlDirectionLeft     PanCtrlDirection = "Left"
	PanCtrlDirectionRight    PanCtrlDirection = "Right"
	PanCtrlDirectionHome     PanCtrlDirection = "Home"
	PanCtrlDirectionPosition PanCtrlDirection = "Position"

	TiltCtrlDirectionStop     TiltCtrlDirection = "Stop"
	TiltCtrlDirectionUp       TiltCtrlDirection = "Up"
	TiltCtrlDirectionDown     TiltCtrlDirection = "Down"
	TiltCtrlDirectionHome     TiltCtrlDirection = "Home"
	TiltCtrlDirectionPosition TiltCtrlDirection = "Position"

	MinPanCtrlPosition = 0
	MaxPanCtrlPosition = 35080

	MinTiltCtrlPosition = 0
	MaxTiltCtrlPosition = 12080

	MinPanCtrlSpeed = 0
	MaxPanCtrlSpeed = 30

	MinTiltCtrlSpeed = 0
	MaxTiltCtrlSpeed = 30

	MinZoomPosition = 0
	MaxZoomPosition = 499
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

type PTZPresetParams struct {
	Operation string `json:"Operation"`
	No        int    `json:"No"`
}

type PTControlParams struct {
	PanDirection  PanCtrlDirection  `json:"PanDirection"`
	PanPosition   int               `json:"PanPosition"`
	PanSpeed      int               `json:"PanSpeed"`
	TiltDirection TiltCtrlDirection `json:"TiltDirection"`
	TiltPosition  int               `json:"TiltPosition"`
	TiltSpeed     int               `json:"TiltSpeed"`
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

func (c HTTPClient) MoveToPreset(preset int) error {
	_, err := c.makeRequest("SetPTZPreset", PTZPresetParams{
		Operation: "Move",
		No:        preset,
	})
	if err != nil {
		return errors.Wrap(err, "fail to move to preset")
	}
	return nil
}

func (c HTTPClient) SetPanTilt(params PTControlParams) error {
	if params.PanSpeed < MinPanCtrlSpeed || params.PanSpeed > MaxPanCtrlSpeed {
		return fmt.Errorf("Invalid pan speed %v, should be between %v and %v", params.PanSpeed, MinPanCtrlSpeed, MaxPanCtrlSpeed)
	}
	if params.PanPosition < MinPanCtrlPosition || params.PanPosition > MaxPanCtrlPosition {
		return fmt.Errorf("Invalid pan speed %v, should be between %v and %v", params.PanPosition, MinPanCtrlPosition, MaxPanCtrlPosition)
	}

	if params.TiltSpeed < MinTiltCtrlSpeed || params.TiltSpeed > MaxTiltCtrlSpeed {
		return fmt.Errorf("Invalid tilt speed %v, should be between %v and %v", params.TiltSpeed, MinTiltCtrlSpeed, MaxTiltCtrlSpeed)
	}
	if params.TiltPosition < MinTiltCtrlPosition || params.TiltPosition > MaxTiltCtrlPosition {
		return fmt.Errorf("Invalid tilt speed %v, should be between %v and %v", params.TiltPosition, MinTiltCtrlPosition, MaxTiltCtrlPosition)
	}

	_, err := c.makeRequest("SetPTCtrl", params)
	if err != nil {
		return errors.Wrap(err, "fail to call HTTP API")
	}
	return nil
}

func (c HTTPClient) SetZoom(position int) error {
	if position < MinZoomPosition || position > MaxZoomPosition {
		return fmt.Errorf("Invalid zoom position %v, should be between %v and %v", position, MinZoomPosition, MaxZoomPosition)
	}

	_, err := c.makeRequest("SetZoomCtrl", map[string]int{
		"Position": position,
	})
	if err != nil {
		return errors.Wrap(err, "fail to call HTTP API")
	}

	return nil
}
