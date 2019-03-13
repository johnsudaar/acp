package models

import (
	"github.com/Scalingo/go-utils/mongo/document"
)

const (
	LinkCollection = "link"
)

type Port struct {
	DeviceID string `bson:"device_id" json:"device_id"`
	Port     string `bson:"port" json:"port"`
}

type Link struct {
	document.Base
	Input  Port `json:"input" bson:"input"`
	Output Port `json:"output" bson:"output"`
}
