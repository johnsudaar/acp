package jvc

import (
	"context"
	"time"

	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/events"
	"github.com/johnsudaar/jvc_api/client"
)

func (j *JVCHM660) tallyLoop() {
	log := j.log.WithField("source", "tallyloop")
	for {
		// Mutex protection
		j.stoppingLock.Lock()
		stopping := j.stopping
		j.stoppingLock.Unlock()
		if stopping {
			return
		}

		j.tallySync.RLock()
		tallyStatus := j.tallyStatus
		j.tallySync.RUnlock()

		if j.client != nil && tallyStatus != "" {
			log.WithField("status", tallyStatus).Info("Sending status")

			err := j.client.SetStudioTally(tallyStatus)
			if err != nil {
				log.WithError(err).Error("fail to set tally")
				j.SetState(devices.StateNotConnected)
			}
		}

		t := time.NewTimer(3 * time.Second)
		// Wait for tally refresh or 3 seconds
		select {
		case <-t.C:
		case <-j.tallyRefreshChan:
		}
	}
}

func (j *JVCHM660) WriteEvent(ctx context.Context, toPort string, name string, data interface{}) {
	if name == events.TallyEventName {
		params, ok := data.(events.TallyEvent)
		if !ok {
			j.log.Error("Invalid data type for tally event")
		}

		tallyValue := client.TallyOff
		if params.Program {
			tallyValue = client.TallyProgram
		} else if params.Preview {
			tallyValue = client.TallyPreview
		}

		j.tallySync.Lock()
		j.tallyStatus = tallyValue
		j.tallySync.Unlock()

		j.tallyRefreshChan <- true
	}
}
