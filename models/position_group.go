package models

import "github.com/Scalingo/go-utils/mongo/document"

const (
	PositionGroupCollection = "position_group"
)

type PositionGroup struct {
	document.Base `bson:",inline"`
	Name          string `bson:"name" json:"name"`
}
