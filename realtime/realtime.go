package realtime

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Scalingo/go-utils/logger"
	"github.com/centrifugal/centrifuge"
	"github.com/johnsudaar/acp/devices"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Realtime interface {
	Publish(ch string, from string, payload interface{}) error
	Websocket(resp http.ResponseWriter, req *http.Request)
	SockJS(resp http.ResponseWriter, req *http.Request)
}

type DeviceGetter interface {
	Get(ctx context.Context, id string) (devices.Device, error)
}

type RealtimeServer struct {
	node             *centrifuge.Node
	websocketHandler http.Handler
	sockjsHandler    http.Handler
	log              logrus.FieldLogger
	getter           DeviceGetter
}

var _ Realtime = &RealtimeServer{}

func New() *RealtimeServer {
	return &RealtimeServer{}
}

func (r *RealtimeServer) Start(ctx context.Context, getter DeviceGetter) error {
	log := logger.Get(ctx).WithField("process", "websocket")
	node, err := centrifuge.New(centrifuge.DefaultConfig)
	if err != nil {
		return errors.Wrap(err, "fail to init centrifuge")
	}

	r.node = node
	r.log = log
	r.getter = getter

	node.OnConnect(func(c *centrifuge.Client) {
		transportName := c.Transport().Name()
		transportProto := c.Transport().Protocol()
		log.Infof("Client connected via %s (%s)", transportName, transportProto)
	})

	node.OnDisconnect(func(c *centrifuge.Client, e centrifuge.DisconnectEvent) {
		log.Info("Client disconnected")
	})

	node.OnSubscribe(r.onSubscribe)
	node.OnPublish(r.onPublish)

	err = node.Run()
	if err != nil {
		return errors.Wrap(err, "fail to start realtime server")
	}

	r.sockjsHandler = centrifuge.NewSockjsHandler(node, centrifuge.SockjsConfig{
		HandlerPrefix:            "/connection/sockjs",
		WebsocketReadBufferSize:  1024,
		WebsocketWriteBufferSize: 1024,
	})

	r.websocketHandler = centrifuge.NewWebsocketHandler(node, centrifuge.WebsocketConfig{})

	return nil
}

func (s *RealtimeServer) onSubscribe(c *centrifuge.Client, e centrifuge.SubscribeEvent) (centrifuge.SubscribeReply, error) {
	s.log.Infof("User subscribed on %s", e.Channel)
	return centrifuge.SubscribeReply{}, nil
}

func (s *RealtimeServer) onPublish(c *centrifuge.Client, e centrifuge.PublishEvent) (centrifuge.PublishReply, error) {
	log := s.log.WithFields(logrus.Fields{
		"channel": e.Channel,
	})
	var event UserEvent
	err := json.Unmarshal(e.Data, &event)
	if err != nil {
		log.WithError(err).Error("fail to unmarshal event")
		return centrifuge.PublishReply{}, nil
	}

	log = log.WithField("device", event.DeviceID)
	log.Info("Toto")

	ctx := logger.ToCtx(context.Background(), log)
	device, err := s.getter.Get(ctx, event.DeviceID)
	if err != nil {
		log.WithError(err).Error("device not found")
		return centrifuge.PublishReply{}, nil
	}
	device.WriteRealtimeEvent(context.Background(), e.Channel, event.Data)
	return centrifuge.PublishReply{}, nil
}
