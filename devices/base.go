package devices

import (
	"context"

	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/models"
	"github.com/johnsudaar/acp/realtime"
	"gopkg.in/mgo.v2/bson"
)

type Base struct {
	id          bson.ObjectId
	name        string
	deviceType  string
	state       State
	eventWriter EventWriter
	realtime    realtime.Realtime
}

func (b Base) ID() bson.ObjectId {
	return b.id
}

func (b Base) Name() string {
	return b.name
}

func (b Base) Type() string {
	return b.deviceType
}

func (b Base) State() State {
	return b.state
}

func (b *Base) SetState(state State) {
	b.state = state
}

func (b *Base) SendEvent(ctx context.Context, from string, name string, data interface{}) {
	b.eventWriter.SendEvent(ctx, models.Port{
		DeviceID: b.ID().Hex(),
		Port:     from,
	}, name, data)
}

func (b *Base) PublishRealtimeEvent(ctx context.Context, ch realtime.Channel, data interface{}) {
	log := logger.Get(ctx)
	event := realtime.RealtimeEvent{
		SenderID: b.ID().Hex(),
		Data:     data,
	}
	err := b.realtime.Publish(ch, event)
	if err != nil {
		log.WithError(err).Error("fail to send realtime event")
	}
}

func Import(d models.Device, writer EventWriter, realtime realtime.Realtime) *Base {
	return &Base{
		id:          d.ID,
		name:        d.Name,
		deviceType:  d.Type,
		state:       StateNotConnected,
		eventWriter: writer,
		realtime:    realtime,
	}
}
