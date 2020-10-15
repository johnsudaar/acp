package realtime

import (
	"context"
	"net/http"

	"github.com/Scalingo/go-utils/logger"
	"github.com/centrifugal/centrifuge"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Realtime interface {
	Publish(ch Channel, payload RealtimeEvent) error
	Websocket(resp http.ResponseWriter, req *http.Request)
	SockJS(resp http.ResponseWriter, req *http.Request)
}

type RealtimeServer struct {
	node             *centrifuge.Node
	websocketHandler http.Handler
	sockjsHandler    http.Handler
	log              logrus.FieldLogger
}

var _ Realtime = &RealtimeServer{}

func Start(ctx context.Context) (*RealtimeServer, error) {
	log := logger.Get(ctx).WithField("process", "websocket")
	node, err := centrifuge.New(centrifuge.DefaultConfig)
	if err != nil {
		return nil, errors.Wrap(err, "fail to init centrifuge")
	}

	server := RealtimeServer{
		node: node,
		log:  log,
	}

	node.OnConnect(func(c *centrifuge.Client) {
		transportName := c.Transport().Name()
		transportProto := c.Transport().Protocol()
		log.Infof("Client connected via %s (%s)", transportName, transportProto)
	})

	node.OnDisconnect(func(c *centrifuge.Client, e centrifuge.DisconnectEvent) {
		log.Info("Client disconnected")
	})

	node.OnSubscribe(server.onSubscribe)

	err = node.Run()
	if err != nil {
		return nil, errors.Wrap(err, "fail to start realtime server")
	}

	server.sockjsHandler = centrifuge.NewSockjsHandler(node, centrifuge.SockjsConfig{
		HandlerPrefix:            "/connection/sockjs",
		WebsocketReadBufferSize:  1024,
		WebsocketWriteBufferSize: 1024,
	})

	server.websocketHandler = centrifuge.NewWebsocketHandler(node, centrifuge.WebsocketConfig{})

	return &server, nil
}

func (s *RealtimeServer) onSubscribe(c *centrifuge.Client, e centrifuge.SubscribeEvent) (centrifuge.SubscribeReply, error) {
	s.log.Infof("User subscribed on %s", e.Channel)
	return centrifuge.SubscribeReply{}, nil
}
