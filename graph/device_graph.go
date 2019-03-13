package graph

import (
	"context"
	"sync"

	"github.com/Scalingo/go-utils/mongo/document"
	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/models"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

var (
	ErrInvalidID = errors.New("invalid device id")
)

type deviceGraph struct {
	devicesLock *sync.RWMutex
	devices     []devices.Device

	linksLock *sync.RWMutex
	links     map[models.Port][]models.Port
}

func (d deviceGraph) All(ctx context.Context) []devices.Device {
	d.devicesLock.RLock()
	defer d.devicesLock.RUnlock()
	return d.devices
}

func (d deviceGraph) Get(ctx context.Context, id string) (devices.Device, error) {
	d.devicesLock.RLock()
	defer d.devicesLock.RUnlock()
	for _, dev := range d.devices {
		if dev.ID().Hex() == id {
			return dev, nil
		}
	}

	return nil, ErrNotFound
}

// Add a device to the device graph
// it will fetch it in the database, load it, add it to the graph and start it
func (d *deviceGraph) Add(ctx context.Context, id string) (devices.Device, error) {
	// Check if the id is a valid ID
	if !bson.IsObjectIdHex(id) {
		return nil, ErrInvalidID
	}

	deviceID := bson.ObjectIdHex(id)

	deviceModel := models.Device{}

	// Find the device in the mongo database
	err := document.Find(ctx, models.DeviceCollection, deviceID, &deviceModel)
	if err != nil {
		return nil, errors.Wrap(err, "fail to find device in database")
	}

	// Import the generic device
	genericDevice := devices.Import(deviceModel, d)

	// Find the correct loader for this device
	loader, err := devices.GetLoader(deviceModel.Type)
	if err != nil {
		return nil, errors.Wrap(err, "fail to get loader")
	}

	// Load and initialize the device driver
	device, err := loader.Load(ctx, genericDevice, deviceModel.Params)
	if err != nil {
		return nil, errors.Wrap(err, "fail to start device")
	}

	// Add the device into the graph
	d.devicesLock.Lock()
	d.devices = append(d.devices, device)
	d.devicesLock.Unlock()

	err = device.Start()
	if err != nil {
		return nil, errors.Wrap(err, "fail to start device")
	}

	return device, nil
}

func (d *deviceGraph) Remove(ctx context.Context, id string) error {
	// TODO: Remove all device links
	d.devicesLock.Lock()
	defer d.devicesLock.Unlock()
	var devices []devices.Device
	for _, device := range d.devices {
		if device.ID().Hex() == id {
			device.Stop()
		} else {
			devices = append(devices, device)
		}
	}
	d.devices = devices
	return nil
}
