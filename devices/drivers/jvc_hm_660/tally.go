package jvc

import (
	"context"

	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/types/tally"
	"github.com/johnsudaar/jvc_api/client"
)

func (j *JVCHM660) SendTally(ctx context.Context, port string, value tally.Value) {
	log := logger.Get(ctx)
	tallyStatus := client.TallyOff
	switch value {
	case tally.Program:
		tallyStatus = client.TallyProgram
	case tally.Preview:
		tallyStatus = client.TallyPreview
	}

	j.clientSync.RLock()
	client := j.client
	j.clientSync.RUnlock()
	if client == nil {
		log.Error("Client not initialized")
		return
	}

	err := client.SetStudioTally(tallyStatus)
	if err != nil {
		j.SetState(devices.StateNotConnected)
		log.WithError(err).Error("Fail to set tally")
	}
}
