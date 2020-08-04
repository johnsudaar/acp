package models

import "github.com/Scalingo/go-utils/mongo/document"

const X32IntercomRooms = "x32_intercom_rooms"

type Room struct {
	document.Base `bson:",inline"`
	Name          string `json:"name" bson:"name"`
	Channel       int    `json:"channel" bson:"channel"`
	Mix           int    `json:"mix" bson:"mix"`
}
