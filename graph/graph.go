package graph

import (
	"context"
	"sync"

	"github.com/Scalingo/go-utils/logger"
	"github.com/Scalingo/go-utils/mongo/document"
	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/models"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

var ErrNotFound = errors.New("Not found")

type Graph interface {
	All(ctx context.Context) []devices.Device
	Get(ctx context.Context, id string) (devices.Device, error)
	Add(ctx context.Context, id string) (devices.Device, error)
	Remove(ctx context.Context, id string) error

	SendEvent(ctx context.Context, from models.Port, name string, data interface{})
	Connect(ctx context.Context, input models.Port, output models.Port) error
	Disconnect(ctx context.Context, input models.Port, output models.Port)
}

func Load(ctx context.Context) (Graph, error) {
	log := logger.Get(ctx)
	graph := &deviceGraph{
		devices:     []devices.Device{},
		devicesLock: &sync.RWMutex{},

		links:     make(map[models.Port][]models.Port),
		linksLock: &sync.RWMutex{},
	}

	log.Info("Loading graph")

	var devices []models.Device
	err := document.Where(ctx, models.DeviceCollection, bson.M{}, &devices)
	if err != nil {
		return nil, errors.Wrap(err, "fail to get devices in database")
	}

	for _, d := range devices {
		log := log.WithFields(logrus.Fields{
			"id":   d.ID.Hex(),
			"name": d.Name,
			"type": d.Type,
		})
		log.Info("Loading device")

		_, err := graph.Add(ctx, d.ID.Hex())
		if err != nil {
			return nil, errors.Wrapf(err, "fail to add device (id=%s)", d.ID.Hex())
		}
	}

	var links []models.Link
	err = document.Where(ctx, models.LinkCollection, bson.M{}, &links)
	if err != nil {
		return nil, errors.Wrap(err, "fail to get links in database")
	}

	for _, link := range links {
		log := log.WithFields(logrus.Fields{
			"input_id":    link.Input.DeviceID,
			"input_port":  link.Input.Port,
			"output_id":   link.Output.DeviceID,
			"output_port": link.Output.Port,
		})
		log.Info("Loading link")

		err := graph.Connect(ctx, link.Input, link.Output)
		if err != nil {
			return nil, errors.Wrap(err, "fail to add link")
		}
	}

	// TODO: Start all devices
	return graph, nil
}
