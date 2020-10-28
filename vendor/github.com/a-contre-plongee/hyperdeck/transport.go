package hyperdeck

import (
	"github.com/pkg/errors"
)

/*
status: stopped
speed: 0
slot id: 1
clip id: 2
single clip: false
display timecode: 00:51:40:05
timecode: 00:51:40:05
video format: 1080i50
loop: false
*/

const (
	//preview, stopped, play, forward, rewind, jog, shuttle, record
	TransportStatusPreview = "preview"
	TransportStatusStopped = "stopped"
	TransportStatusPlay    = "play"
	TransportStatusForward = "forward"
	TransportStatusRewind  = "rewind"
	TransportStatusJog     = "jog"
	TransportStatusShuttle = "shuttle"
	TransportStatusRecord  = "record"
)

type StringParser interface {
	FromString(string) error
}

type TransportListener func(Transport)

type Transport struct {
	Status          string   `header:"status"`
	Speed           int      `header:"speed"`
	SlotID          int      `header:"slot id"`
	ClipID          int      `header:"clip id"`
	SingleClip      bool     `header:"single clip"`
	DisplayTimecode Timecode `header:"display timecode"`
	Timecode        Timecode `header:"timecode"`
	VideoFormat     string   `header:"video format"`
	Loop            bool     `header:"loop"`
}

func ParseTransport(payload []byte) (Transport, error) {
	var res Transport
	err := Parse(payload, &res)
	if err != nil {
		return res, errors.Wrap(err, "fail to parse transport")
	}
	return res, nil

}

func (c *Client) TransportInfo() (Transport, error) {
	res, err := c.Send([]byte("transport info\n"))
	if err != nil {
		return Transport{}, errors.Wrap(err, "fail to send command")
	}

	if IsError(res) {
		parsedError, err := ParseError(res)
		if err != nil {
			return Transport{}, errors.Wrap(err, "fail to parse error")
		} else {
			return Transport{}, parsedError
		}
	}

	transport, err := ParseTransport(res)
	if err != nil {
		return transport, errors.Wrap(err, "fail to parse transport")
	}

	return transport, nil
}
