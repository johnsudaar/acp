package models

import (
	"encoding/json"

	"github.com/Scalingo/go-utils/mongo/document"
)

const DeviceCollection = "device"

type Device struct {
	document.Base `bson:",inline"`
	Type          string            `bson:"type" json:"type"`
	Name          string            `bson:"name" json:"name"`
	Params        json.RawMessage   `bson:"params,omitempty" json:"params,omitempty"`
	DisplayOpts   DeviceDisplayOpts `bson:"display_opts,omitempty" json:"display_opts,omitempty"`
}

type DeviceDisplayOpts struct {
	Position DeviceDisplayOptsPosition `bson:"position,omitempty" json:"position,omitempty"`
}

type DeviceDisplayOptsPosition struct {
	X int `bson:"x,omitempty" json:"x"`
	Y int `bson:"y,omitempty" json:"y"`
}
