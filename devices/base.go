package devices

import (
	"context"

	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/models"
	"gopkg.in/mgo.v2/bson"
)

type Base struct {
	id          bson.ObjectId
	name        string
	deviceType  string
	state       State
	eventWriter EventWriter
	realtime    RealtimeEventWriter
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

func (b *Base) PublishRealtimeEvent(ctx context.Context, ch string, data interface{}) {
	log := logger.Get(ctx)
	err := b.realtime.Publish(ch, b.ID().Hex(), data)
	if err != nil {
		log.WithError(err).Error("fail to send realtime event")
	}
}

func Import(d models.Device, writer EventWriter, realtime RealtimeEventWriter) *Base {
	return &Base{
		id:          d.ID,
		name:        d.Name,
		deviceType:  d.Type,
		state:       StateNotConnected,
		eventWriter: writer,
		realtime:    realtime,
	}
}
