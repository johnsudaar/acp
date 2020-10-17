package tallybox

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/types"
	rpio "github.com/stianeikeland/go-rpio"
)

type tallybox struct {
	*devices.Base
}

func (t *tallybox) InputPorts() []string {
	return []string{}
}

func (t *tallybox) OutputPorts() []string {
	return []string{
		"Input_1",
		"Input_2",
		"Input_3",
		"Input_4",
	}
}

func (t *tallybox) API() http.Handler {
	return http.NotFoundHandler()
}

func (t *tallybox) Start() error {
	err := rpio.Open()
	if err != nil {
		return err
	}

	t.initGPIO()

	return nil
}

func (t *tallybox) Stop() error {
	return rpio.Close()
}

func (t *tallybox) WriteEvent(ctx context.Context, toPort string, name string, data interface{}) {
}

func (t *tallybox) Types() []types.Type {
	return []types.Type{types.TallyType}
}

func (t *tallybox) WriteRealtimeEvent(ctx context.Context, channel string, payload json.RawMessage) {
}
