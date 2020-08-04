package x32

import (
	"context"

	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/types"
	"github.com/sirupsen/logrus"
)

type X32 struct {
	*devices.Base
	IP  string
	log logrus.FieldLogger
}

func (x *X32) Start() error {
	return nil
}

func (x *X32) Stop() error {
	return nil
}

func (x *X32) WriteEvent(ctx context.Context, toPort string, name string, data interface{}) {
}

func (x *X32) Types() []types.Type {
	return []types.Type{}
}

func (x *X32) InputPorts() []string {
	return []string{}
}

func (x *X32) OutputPorts() []string {
	return []string{}
}
