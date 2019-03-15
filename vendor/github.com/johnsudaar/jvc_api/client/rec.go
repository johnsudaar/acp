package client

import "github.com/pkg/errors"

type RecControl string

const (
	RecControlStart = "Rec"
	RecControlStop  = "Stop"
)

type setRecControl struct {
	CamCtrl RecControl `json:"CamCtrl"`
}

func (c HTTPClient) SetRec(rec RecControl) error {
	_, err := c.makeRequest("SetCamCtrl", setRecControl{
		CamCtrl: rec,
	})
	if err != nil {
		return errors.Wrap(err, "fail to call HTTP API")
	}
	return nil
}
