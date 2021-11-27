package models

import (
	"github.com/Scalingo/go-utils/mongo/document"
	"gopkg.in/mgo.v2/bson"
)

const (
	PtzPositionCollection = "ptz"
)

type PtzPosition struct {
	document.Base   `bson:",inline"`
	DeviceID        bson.ObjectId  `json:"device_id" bson:"device_id"`
	Name            string         `json:"name" bson:"name"`
	Pan             float64        `json:"pan" bson:"pan"`
	Tilt            float64        `json:"tilt" bson:"tilt"`
	Zoom            float64        `json:"zoom" bson:"zoom"`
	Focus           float64        `json:"focus" bson:"focus"`
	PositionGroupID *bson.ObjectId `json:"position_group_id" bson:"position_group_id"`
}
