package tallyrecorder

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Scalingo/go-utils/mongo/document"
	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/drivers/tallyrecorder/models"
	"github.com/johnsudaar/acp/devices/types"
	"github.com/johnsudaar/acp/events"
	"github.com/sirupsen/logrus"
)

type Recorder struct {
	*devices.Base
	log   logrus.FieldLogger
	Shoot string
}

func (t *Recorder) Start() error {
	return nil
}

func (t *Recorder) Stop() error {
	return nil
}

func (*Recorder) InputPorts() []string {
	return []string{}
}

func (*Recorder) OutputPorts() []string {
	return []string{"Tally"}
}

func (*Recorder) Types() []types.Type {
	return []types.Type{}
}

func (r *Recorder) WriteEvent(ctx context.Context, toPort, name string, data interface{}) {
	if name == events.TalliesEventName {
		params, ok := data.(events.TalliesEvent)
		if !ok {
			r.log.Error("Invalid data type for tallied event")
		}

		data := models.TallyEvent{
			Time:    time.Now(),
			Shoot:   r.Shoot,
			Program: params.Program,
			Preview: params.Preview,
		}

		err := document.Create(ctx, models.TallyEventCollection, &data)
		if err != nil {
			r.log.WithError(err).Error("fail to save tally event")
		}
	}
}

func (r *Recorder) WriteRealtimeEvent(ctx context.Context, channel string, payload json.RawMessage) {
}
