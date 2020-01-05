package client

import "github.com/pkg/errors"

type FocusDirection string

const (
	FocusFar1  FocusDirection = "Far1"
	FocusFar2  FocusDirection = "Far2"
	FocusFar3  FocusDirection = "Far3"
	FocusNear1 FocusDirection = "Near1"
	FocusNear2 FocusDirection = "Near2"
	FocusNear3 FocusDirection = "Near3"
	FocusStop  FocusDirection = "Stop"
)

func (c HTTPClient) Focus(direction FocusDirection) error {
	err := c.SendWebKeyEvent("Focus", string(direction))
	if err != nil {
		return errors.Wrap(err, "fail to send focus event")
	}
	return nil
}

func (c HTTPClient) FocusPushAuto() error {
	err := c.SendWebKeyEvent("Focus", "PushAuto")
	if err != nil {
		return errors.Wrap(err, "fail to send push auto event")
	}
	return nil
}
