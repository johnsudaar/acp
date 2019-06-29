package jvc

import (
	"sync"

	"github.com/johnsudaar/acp/devices/types/ptz"
	"github.com/johnsudaar/jvc_api/client"
	"github.com/pkg/errors"
)

var EPSILON float64 = 0.00000001

func (j *JVCHM660) SendPTZJoystick(params ptz.PTZJoystickParams) error {
	j.clientSync.RLock()
	cam := j.client
	j.clientSync.RUnlock()

	if cam == nil {
		return errors.New("camera not connected")
	}

	panDirection := client.PanDirectionStop
	if params.PanSpeed < -EPSILON {
		panDirection = client.PanDirectionLeft
		params.PanSpeed = -params.PanSpeed
	} else if params.PanSpeed > EPSILON {
		panDirection = client.PanDirectionRight
	}
	panSpeed := int(params.PanSpeed*(client.MaxJoystickSpeed-client.MinJoystickSpeed) + client.MinJoystickSpeed)

	tiltDirection := client.TiltDirectionStop
	if params.TiltSpeed < -EPSILON {
		tiltDirection = client.TiltDirectionDown
		params.TiltSpeed = -params.TiltSpeed
	} else if params.TiltSpeed > EPSILON {
		tiltDirection = client.TiltDirectionUp
	}
	tiltSpeed := int(params.TiltSpeed*(client.MaxJoystickSpeed-client.MinJoystickSpeed) + client.MinJoystickSpeed)

	zoomDirection := client.ZoomDirectionStop
	if params.ZoomSpeed < -EPSILON {
		zoomDirection = client.ZoomDirectionWide
		params.ZoomSpeed = -params.ZoomSpeed
	} else if params.ZoomSpeed > EPSILON {
		zoomDirection = client.ZoomDirectionTele
	}
	zoomSpeed := int(params.ZoomSpeed*(client.MaxZoomSpeed-client.MinZoomSpeed) + client.MinZoomSpeed)

	//log := logrus.New()

	//log.WithFields(logrus.Fields{
	//	"zoom_speed":     zoomSpeed,
	//	"zoom_direction": zoomDirection,
	//	"pan_speed":      panSpeed,
	//	"pan_direction":  panDirection,
	//	"tilt_speed":     tiltSpeed,
	//	"tilt_direction": tiltDirection,
	//}).Info("Applying changes !")

	joystickParams := client.JoyStickOperationParams{
		PanDirection:  panDirection,
		PanSpeed:      panSpeed,
		TiltDirection: tiltDirection,
		TiltSpeed:     tiltSpeed,
	}

	zoomParams := client.ZoomParams{
		Speed:     zoomSpeed,
		Direction: zoomDirection,
	}

	var joyErr error
	var zoomErr error

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		joyErr = cam.JoyStickOperation(joystickParams)
		wg.Done()
	}()

	go func() {
		zoomErr = cam.Zoom(zoomParams)
		wg.Done()
	}()
	wg.Wait()

	if joyErr != nil {
		return errors.Wrap(joyErr, "fail to send joystick params")
	}

	if zoomErr != nil {
		return errors.Wrap(zoomErr, "fail to send zoom params")
	}

	return nil
}
