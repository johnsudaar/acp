package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Scalingo/go-utils/mongo/document"
	"gopkg.in/mgo.v2/bson"
)

type TimerType string
type TimerDirection string

const (
	TimerCollection = "timer"

	TimerTypeClock     TimerType = "clock"
	TimerTypeCountDown TimerType = "countdown"
	TimerTypeCountUp   TimerType = "countup"
	TimerTypeExternal  TimerType = "external"

	TimerDirectionUp   TimerDirection = "up"
	TimerDirectionDown TimerDirection = "down"

	TimerActionStart = "start"
	TimerActionStop  = "stop"
	TimerActionPause = "pause"
	TimerActionReset = "reset"
)

type Timer struct {
	document.Base  `bson:",inline"`
	Name           string        `json:"name" bson:"name"`
	Type           TimerType     `json:"type" bson:"type"`
	ExternalDevice bson.ObjectId `json:"external_device" bson:"external_device,omitempty"`
	ExternalSource string        `json:"external_source" bson:"external_source,omitempty"`
	Duration       Duration      `json:"duration" bson:"duration"`
}

type TimerAction struct {
	Action string `json:"action"`
	Param  string `json:"param"`
}

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var err error
	if b[0] == '"' {
		sd := string(b[1 : len(b)-1])
		if len(sd) == 0 {
			d.Duration = time.Duration(0)
			return nil
		}
		d.Duration, err = time.ParseDuration(sd)
		if err != nil {
			return err
		}
		return nil
	}

	fmt.Println(string(b))
	if string(b) == "null" {
		d.Duration = time.Duration(0)
		return nil
	}

	var id int64
	id, err = json.Number(string(b)).Int64()
	if err != nil {
		return err
	}
	d.Duration = time.Duration(id)

	return nil
}

func (d Duration) MarshalJSON() (b []byte, err error) {
	return []byte(fmt.Sprintf(`"%s"`, d.String())), nil
}
