package hs50

import (
	"context"
	"encoding/json"

	"github.com/a-contre-plongee/hs50/client"
	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/types"
	"github.com/johnsudaar/acp/devices/types/switcher"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var _ switcher.Switchable = &Switcher{}

type Switcher struct {
	*devices.Base
	IP     string
	log    logrus.FieldLogger
	client client.Client
}

func (s *Switcher) Start() error {
	return nil
}

func (s *Switcher) Stop() error {
	return nil
}

func (s *Switcher) WriteEvent(ctx context.Context, toPort string, name string, data interface{}) {
}

func (s *Switcher) Types() []types.Type {
	return []types.Type{types.SwitcherType}
}

func (s *Switcher) InputPorts() []string {
	return []string{
		"Input 1",
		"Input 2",
		"Input 3",
		"Input 4",
		"Input 5",
	}
}

func (s *Switcher) OutputPorts() []string {
	return []string{
		"PGM",
		"MV",
	}
}

func (s *Switcher) Switch(output, input string) error {
	if output == "MV" {
		return errors.New("Cannot switch the MV output")
	}

	var inputValue client.Input

	switch input {
	case "Input 1":
		inputValue = client.Input1
	case "Input 2":
		inputValue = client.Input2
	case "Input 3":
		inputValue = client.Input3
	case "Input 4":
		inputValue = client.Input4
	case "Input 5":
		inputValue = client.Input5
	default:
		return errors.New("invalid input: " + input)
	}

	err := s.client.SwitchBus(client.BusPgm, inputValue)
	if err != nil {
		return errors.Wrap(err, "fail to switch bus")
	}
	return nil
}

func (s *Switcher) WriteRealtimeEvent(ctx context.Context, channel string, payload json.RawMessage) {
}
