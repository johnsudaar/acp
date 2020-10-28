package hyperdeck

import (
	"fmt"

	"github.com/pkg/errors"
)

/*
202 slot info:
slot id: 2
status: empty
volume name:
recording time: 0
video format: none
device serial:
device model:
firmware revision:
*/

const (
	// empty, mounting, error, mounted
	SlotStatusEmpty    = "empty"
	SlotStatusMounting = "mounting"
	SlotStatusError    = "error"
	SlotStatusMounted  = "mounted"
)

type SlotListener func(Slot)

type Slot struct {
	ID               int    `header:"slot id"`
	Status           string `header:"status"`
	VolumeName       string `header:"volume name"`
	RecordingTime    int    `header:"recording time"`
	VideoFormat      string `header:"video format"`
	DeviceSerial     string `header:"device serial"`
	DeviceModel      string `header:"device model"`
	FirmwareRevision string `header:"firmware revision"`
}

func ParseSlot(payload []byte) (Slot, error) {
	var res Slot
	err := Parse(payload, &res)
	if err != nil {
		return res, errors.Wrap(err, "fail to parse slot")
	}
	return res, nil
}

func (c *Client) SlotInfo(slot int) (Slot, error) {
	cmd := fmt.Sprintf("slot info")
	if slot != 0 {
		cmd = fmt.Sprintf("%s: slot id: %v", cmd, slot)
	}
	cmd += "\n"

	res, err := c.Send([]byte(cmd))
	if err != nil {
		return Slot{}, errors.Wrap(err, "fail to send cmd")
	}

	if IsError(res) {
		parsedError, err := ParseError(res)
		if err != nil {
			return Slot{}, errors.Wrap(err, "fail to parse error")
		} else {
			return Slot{}, parsedError
		}
	}

	slotRes, err := ParseSlot(res)
	if err != nil {
		return slotRes, errors.Wrap(err, "fail to parse slot")
	}
	return slotRes, nil
}
