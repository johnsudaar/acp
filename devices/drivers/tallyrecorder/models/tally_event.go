package models

import (
	"time"

	"github.com/Scalingo/go-utils/mongo/document"
)

const TallyEventCollection = "tallies"

type TallyEvent struct {
	document.Base `bson:",inline"`
	Time          time.Time `json:"time" bson:"time"`
	Shoot         string    `json:"shoot" bson:"shoot"`
	Program       []string  `json:"program" bson:"program"`
	Preview       []string  `json:"preview" bson:"preview"`
}
