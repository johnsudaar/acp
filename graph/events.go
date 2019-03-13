package graph

import (
	"context"
	"fmt"

	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/models"
	"github.com/johnsudaar/acp/utils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (d deviceGraph) SendEvent(ctx context.Context, from models.Port, name string, data interface{}) {
	log := logger.Get(ctx).WithFields(logrus.Fields{
		"event_source_device": from.DeviceID,
		"event_source_port":   from.Port,
		"event_name":          name,
	})

	log.Info("Event received")

	d.linksLock.RLock()
	defer d.linksLock.RUnlock()
	for _, port := range d.links[from] {
		log := log.WithFields(logrus.Fields{
			"event_destination_device": port.DeviceID,
			"event_destination_port":   port.Port,
		})
		log.Info("Sending event")
		dev, err := d.Get(ctx, port.DeviceID)
		if err != nil {
			log.WithError(err).Error("Fail to find device, ignoring event...")
			return
		}

		go dev.WriteEvent(ctx, port.Port, name, data)
	}
}

func (d deviceGraph) Connect(ctx context.Context, input models.Port, output models.Port) error {
	inDevice, err := d.Get(ctx, input.DeviceID)
	if err != nil {
		return errors.Wrap(err, "fail to find input device")
	}

	outDevice, err := d.Get(ctx, output.DeviceID)
	if err != nil {
		return errors.Wrap(err, "fail to find output device")
	}

	if !utils.HasString(input.Port, inDevice.InputPorts()) {
		return fmt.Errorf("Port %s not found on device %s", input.Port, input.DeviceID)
	}

	if !utils.HasString(output.Port, outDevice.OutputPorts()) {
		return fmt.Errorf("Port %s not found on device %s", output.Port, output.DeviceID)
	}

	d.linksLock.Lock()
	defer d.linksLock.Unlock()
	d.links[input] = append(d.links[input], output)
	return nil
}

func (d deviceGraph) Disconnect(ctx context.Context, input models.Port, output models.Port) {
	d.linksLock.Lock()
	defer d.linksLock.Unlock()
	ports, ok := d.links[input]
	if !ok {
		return
	}

	var newPorts []models.Port
	for _, port := range ports {
		if port != output {
			newPorts = append(newPorts, port)
		}
	}
	d.links[input] = newPorts
}
