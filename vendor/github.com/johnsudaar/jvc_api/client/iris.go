package client

import "github.com/pkg/errors"

type IrisDirection string

const (
	IrisOpen1  IrisDirection = "Open1"
	IrisOpen2  IrisDirection = "Open2"
	IrisOpen3  IrisDirection = "Open3"
	IrisClose1 IrisDirection = "Close1"
	IrisClose2 IrisDirection = "Close2"
	IrisClose3 IrisDirection = "Close3"
	IrisStop   IrisDirection = "Stop"
)

func (c HTTPClient) Iris(direction IrisDirection) error {
	err := c.SendWebKeyEvent("Iris", string(direction))
	if err != nil {
		return errors.Wrap(err, "fail to send iris event")
	}
	return nil
}
