package tally

import (
	"context"

	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/events"
	"github.com/johnsudaar/acp/utils"
)

func (t *TallyDriver) EventSubscriptions() []string {
	return []string{events.TallyEventName}
}

func (t *TallyDriver) WriteEvent(ctx context.Context, toPort string, name string, data interface{}) {
	log := logger.Get(ctx).WithField("process", "tallydriver")
	if name != events.TallyEventName {
		params, ok := data.(events.TallyEvent)
		if !ok {
			log.Error("Invalid data type for tally event")
			return
		}

		if utils.HasString(toPort, t.device.OutputPorts()) {
			log.Errorf("Invalid port: %s", toPort)
			return
		}

		value := Off
		if params.Program {
			value = Program
		} else if params.Preview {
			value = Preview
		}

		log.WithField("tally", value).Info("Set tally")

		t.lock.Lock()
		t.values[toPort] = value
		t.lock.Unlock()
		t.refreshChan <- true
	}
}
