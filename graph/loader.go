package graph

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/johnsudaar/acp/devices"
	"github.com/pkg/errors"
)

type GraphJSON struct {
	Devices []DeviceJSON `json:"devices"`
	Links   []LinkJSON   `json:"links"`
}

type DeviceJSON struct {
	Device     string          `json:"device"`
	ID         string          `json:"id"`
	Attributes json.RawMessage `json:"attributes"`
}

type LinkJSON struct {
	From DevicePort `json:"from"`
	To   DevicePort `json:"to"`
}

func (g GraphJSON) ToGraph() (Graph, error) {
	graph := Graph{
		Devices: make(map[string]Device),
	}

	for _, d := range g.Devices {
		if _, ok := graph.Devices[d.ID]; ok {
			return graph, fmt.Errorf("duplicate ID: %s", d.ID)
		}

		realDevice, err := devices.Initialize(d.Device, d.Attributes)
		if err != nil {
			return graph, errors.Wrapf(err, "fail to intiailize %s", d.ID)
		}
		graph.Devices[d.ID] = Device{
			Type:   d.Device,
			Device: realDevice,
		}
	}

	for _, l := range g.Links {
		deviceFrom, ok := graph.Devices[l.From.ID]
		if !ok {
			return graph, fmt.Errorf("device %s not found", l.From.ID)
		}
		deviceTo, ok := graph.Devices[l.To.ID]
		if !ok {
			return graph, fmt.Errorf("device %s not found", l.From.ID)
		}

		err := deviceTo.Device.Connect(l.To.Port, devices.OutputPort{
			Device: deviceFrom.Device,
			Name:   l.From.Port,
		})
		if err != nil {
			return graph, errors.Wrapf(err, "fail to connect %s %s to %s %s", l.From.ID, l.From.Port, l.To.ID, l.To.Port)
		}
	}

	return graph, nil
}

func Load(path string) (Graph, error) {
	file, err := os.Open(path)
	if err != nil {
		return Graph{}, errors.Wrap(err, "fail to open file")
	}

	var tempGraph GraphJSON

	err = json.NewDecoder(file).Decode(&tempGraph)
	if err != nil {
		return Graph{}, errors.Wrap(err, "fail to read graph definition")
	}

	graph, err := tempGraph.ToGraph()
	if err != nil {
		return Graph{}, errors.Wrap(err, "invalid graph definition")
	}
	return graph, nil
}
