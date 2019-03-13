package models

import (
	"encoding/json"

	"github.com/Scalingo/go-utils/mongo/document"
)

const DeviceCollection = "device"

type Device struct {
	document.Base
	Type        string          `bson:"type" json:"type"`
	Name        string          `bson:"name" json:"name"`
	Params      json.RawMessage `bson:"params,omitempty" json:"params,omitempty"`
	DisplayOpts interface{}     `bson:"display_opts,omitempty" json:"display_opts,omitempty"`
}
