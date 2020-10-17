package devices

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	handlers "github.com/Scalingo/go-handlers"
	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/devices/types"
	"github.com/johnsudaar/acp/devices/types/ptz"
	"github.com/johnsudaar/acp/devices/types/switcher"
	"github.com/johnsudaar/acp/devices/types/tally"
	"github.com/johnsudaar/acp/utils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

// This type is used as a proxy for the real device (when using the loader.Load method)
// it's necessarry to add the hooks needed for the device types
type DeviceWrapper struct {
	Implementation Device
	DeviceTypes    map[types.Type]types.DeviceType
	TypesLock      *sync.RWMutex
}

func Wrap(d Device) (Device, error) {
	deviceTypes := make(map[types.Type]types.DeviceType)
	for _, t := range d.Types() {
		switch t {
		case types.TallyType:
			tally, err := tally.Import(d)
			if err != nil {
				return nil, errors.Wrapf(err, "fail to import device type %s", t)
			}
			deviceTypes[t] = tally
		case types.PTZType:
			ptz, err := ptz.Import(d)
			if err != nil {
				return nil, errors.Wrapf(err, "fail to import device type %s", t)
			}
			deviceTypes[t] = ptz
		case types.SwitcherType:
			switcher, err := switcher.Import(d)
			if err != nil {
				return nil, errors.Wrapf(err, "fail to import device type %s", t)
			}
			deviceTypes[t] = switcher
		}

	}
	return &DeviceWrapper{
		Implementation: d,
		DeviceTypes:    deviceTypes,
		TypesLock:      &sync.RWMutex{},
	}, nil
}

func (d *DeviceWrapper) ID() bson.ObjectId {
	return d.Implementation.ID()
}

func (d *DeviceWrapper) Name() string {
	return d.Implementation.Name()
}

func (d *DeviceWrapper) Type() string {
	return d.Implementation.Type()
}

func (d *DeviceWrapper) State() State {
	return d.Implementation.State()
}

func (d *DeviceWrapper) InputPorts() []string {
	return d.Implementation.InputPorts()
}

func (d *DeviceWrapper) OutputPorts() []string {
	return d.Implementation.OutputPorts()
}

func (d *DeviceWrapper) API() http.Handler {
	log := logrus.New()
	log.Out = ioutil.Discard
	router := handlers.NewRouter(log)

	d.TypesLock.Lock()
	for _, t := range d.DeviceTypes {
		routes := t.Routes()
		if routes == nil {
			continue
		}

		for route, handler := range routes {
			router.HandleFunc(route, handler)
		}
	}
	d.TypesLock.Unlock()
	router.PathPrefix("/").Handler(d.Implementation.API())
	return router
}

func (d *DeviceWrapper) Start() error {
	d.TypesLock.RLock()
	defer d.TypesLock.RUnlock()

	log := logger.Default().WithField("device_name", d.Implementation.Name())

	for name, t := range d.DeviceTypes {
		ctx := logger.ToCtx(context.Background(), log.WithField("device_type", name))
		err := t.Start(ctx)
		if err != nil {
			return errors.Wrap(err, "fail to start device type")
		}
	}
	return d.Implementation.Start()
}

func (d *DeviceWrapper) Stop() error {
	d.TypesLock.RLock()
	defer d.TypesLock.RUnlock()

	log := logger.Default().WithField("device_name", d.Implementation.Name())

	for name, t := range d.DeviceTypes {
		ctx := logger.ToCtx(context.Background(), log.WithField("device_type", name))
		err := t.Stop(ctx)
		if err != nil {
			return errors.Wrap(err, "fail to start device type")
		}
	}

	return d.Implementation.Stop()
}

func (d *DeviceWrapper) WriteEvent(ctx context.Context, toPort string, name string, data interface{}) {
	d.TypesLock.RLock()
	defer d.TypesLock.RUnlock()

	for typeName, t := range d.DeviceTypes {
		if utils.HasString(name, t.EventSubscriptions()) {
			ctx := logger.ToCtx(ctx, logger.Get(ctx).WithField("device_type", typeName))
			t.WriteEvent(ctx, toPort, name, data)
		}

	}
	d.Implementation.WriteEvent(ctx, toPort, name, data)
}

func (d *DeviceWrapper) WriteRealtimeEvent(ctx context.Context, channel string, payload json.RawMessage) {
	d.TypesLock.RLock()
	defer d.TypesLock.RUnlock()

	for typeName, t := range d.DeviceTypes {
		if utils.HasString(channel, t.RealtimeEventSubscriptions()) {
			ctx := logger.ToCtx(ctx, logger.Get(ctx).WithField("device_type", typeName))
			t.WriteRealtimeEvent(ctx, channel, payload)
		}
	}

	d.Implementation.WriteRealtimeEvent(ctx, channel, payload)
}

func (d *DeviceWrapper) Types() []types.Type {
	return d.Implementation.Types()
}
