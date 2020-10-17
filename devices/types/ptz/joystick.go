package ptz

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func (p *PtzDriver) JoystickAction(body []byte) error {
	var payload PTZJoystickParams
	err := json.Unmarshal(body, &payload)
	if err != nil {
		return nil
	}
	if payload.PanSpeed > 1 {
		payload.PanSpeed = 1
	}
	if payload.PanSpeed < -1 {
		payload.PanSpeed = -1
	}
	if payload.TiltSpeed > 1 {
		payload.TiltSpeed = 1
	}
	if payload.TiltSpeed < -1 {
		payload.TiltSpeed = -1
	}
	if payload.ZoomSpeed > 1 {
		payload.ZoomSpeed = 1
	}
	if payload.ZoomSpeed < -1 {
		payload.ZoomSpeed = -1
	}

	err = p.device.SendPTZJoystick(payload)
	if err != nil {
		return errors.Wrap(err, "fail to send PTZ")
	}
	return nil
}
