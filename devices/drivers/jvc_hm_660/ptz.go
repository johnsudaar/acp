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
	if params.Buttons.Button1 {
		return cam.MoveToPreset(1)
	} else if params.Buttons.Button2 {
		return cam.MoveToPreset(2)
	} else if params.Buttons.Button3 {
		return cam.MoveToPreset(3)
	} else if params.Buttons.Button4 {
		return cam.MoveToPreset(4)
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

	focusSpeed := client.FocusStop
	if params.FocusSpeed > 0.8 {
		focusSpeed = client.FocusNear3
	} else if params.FocusSpeed > 0.5 {
		focusSpeed = client.FocusNear2
	} else if params.FocusSpeed > 0.2 {
		focusSpeed = client.FocusNear1
	} else if params.FocusSpeed < -0.8 {
		focusSpeed = client.FocusFar3
	} else if params.FocusSpeed < -0.5 {
		focusSpeed = client.FocusFar2
	} else if params.FocusSpeed < -0.2 {
		focusSpeed = client.FocusFar1
	}

	irisDirection := client.IrisStop
	if params.Buttons.IrisOpen {
		irisDirection = client.IrisOpen2
	} else if params.Buttons.IrisClose {
		irisDirection = client.IrisClose2
	}

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

	var joyErr, zoomErr, focusErr, irisErr error

	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		joyErr = cam.JoyStickOperation(joystickParams)
		wg.Done()
	}()

	go func() {
		zoomErr = cam.Zoom(zoomParams)
		wg.Done()
	}()

	go func() {
		if params.Buttons.FocusPushAuto {
			focusErr = cam.FocusPushAuto()
		} else {
			focusErr = cam.Focus(focusSpeed)
		}
		wg.Done()
	}()

	go func() {
		irisErr = cam.Iris(irisDirection)
		wg.Done()
	}()
	wg.Wait()

	if joyErr != nil {
		return errors.Wrap(joyErr, "fail to send joystick params")
	}
	if zoomErr != nil {
		return errors.Wrap(zoomErr, "fail to send zoom params")
	}
	if focusErr != nil {
		return errors.Wrap(zoomErr, "fail to send focus params")
	}
	if irisErr != nil {
		return errors.Wrap(zoomErr, "fail to send iris params")
	}

	return nil
}

func (j *JVCHM660) SetPosition(params ptz.PTZPositionParams) error {
	j.clientSync.RLock()
	cam := j.client
	j.clientSync.RUnlock()

	if cam == nil {
		return errors.New("camera not connected")
	}

	payload := client.PTControlParams{
		PanDirection:  client.PanCtrlDirectionPosition,
		PanSpeed:      client.MaxPanCtrlSpeed,
		TiltDirection: client.TiltCtrlDirectionPosition,
		TiltSpeed:     client.MaxTiltCtrlSpeed,
	}

	payload.PanPosition = int((params.Pan * client.MaxPanCtrlPosition) / 100.0)
	payload.TiltPosition = int((params.Tilt * client.MaxTiltCtrlPosition) / 100.0)
	zoom := int((params.Zoom * client.MaxZoomPosition) / 100.0)

	var wg sync.WaitGroup
	var ptErr, zoomErr error
	wg.Add(2)
	go func() {
		ptErr = cam.SetPanTilt(payload)
		wg.Done()
	}()

	go func() {
		zoomErr = cam.SetZoom(zoom)
		wg.Done()
	}()
	if ptErr != nil {
		return errors.Wrap(ptErr, "fail to send pan tilt params")
	}
	if zoomErr != nil {
		return errors.Wrap(ptErr, "fail to send zoom params")
	}

	return nil
}
